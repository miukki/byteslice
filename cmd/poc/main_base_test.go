package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"sync"
	"testing"
	"text/template"
	"time"

	"github.com/docker/docker/pkg/reexec"
	"github.com/stretchr/testify/require"
)

type TestExecCommand struct {
	*testing.T
	Func    template.FuncMap
	Data    interface{}
	Cleanup func()
	cmd     *exec.Cmd
	stdout  *bufio.Reader
	stdin   io.WriteCloser
	stderr  *testlog
	Err     error
}

type testzkp struct {
	*TestExecCommand
}

type testsigning struct {
	*TestExecCommand
}

type testlog struct {
	t   *testing.T
	mu  sync.Mutex
	buf bytes.Buffer
}

func init() {
	reexec.Register("zkp-test", func() {
		app := prepare()
		if err := app.Run(os.Args); err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		os.Exit(0)
	})

	reexec.Register("signing-test", func() {
		app := prepare()
		if err := app.Run(os.Args); err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		os.Exit(0)
	})
}

func (tl *testlog) Write(b []byte) (n int, err error) {
	lines := bytes.Split(b, []byte("\n"))
	for _, line := range lines {
		if len(line) > 0 {
			tl.t.Logf("stderr: %s", line)
		}
	}
	tl.mu.Lock()
	defer tl.mu.Unlock()
	tl.buf.Write(b)
	return len(b), err
}

func (tt *TestExecCommand) Run(name string, args ...string) {
	tt.stderr = &testlog{t: tt.T}
	tt.cmd = &exec.Cmd{
		Path:   reexec.Self(),
		Args:   append([]string{name}, args...),
		Stderr: tt.stderr,
	}
	//tt.cmd = exec.Command(reexec.Self(), append([]string{name}, args...)...)
	//tt.cmd.Stderr = tt.stderr
	stdout, err := tt.cmd.StdoutPipe()
	require.Nil(tt, err)
	tt.stdout = bufio.NewReader(stdout)
	if tt.stdin, err = tt.cmd.StdinPipe(); err != nil {
		require.Nil(tt, err)
	}
	if err := tt.cmd.Start(); err != nil {
		require.Nil(tt, err)
	}
}

func (tt *TestExecCommand) ExpectExit() {
	var output []byte
	tt.withKillTimeout(func() {
		output, _ = ioutil.ReadAll(tt.stdout)
	})
	tt.WaitExit()
	if tt.Cleanup != nil {
		tt.Cleanup()
	}
	if len(output) > 0 {
		tt.Errorf("stdout unmatched:\n%s", output)
	}
}

func (tt *TestExecCommand) ExpectExitWith(want string) {
	var output []byte
	tt.withKillTimeout(func() {
		output, _ = ioutil.ReadAll(tt.stdout)
	})
	tt.WaitExit()
	if tt.Cleanup != nil {
		tt.Cleanup()
	}
	bufStr := string(output)

	if len(want) > 0 {
		require.NotEmpty(tt, bufStr)
	}
	if len(output) > 0 {
		require.Contains(tt, bufStr, want, "stdout unmatched, want: %s: found: \n%s", want, bufStr)
	}
}

func (tt *TestExecCommand) WaitExit() {
	tt.Err = tt.cmd.Wait()
}

func (tt *TestExecCommand) withKillTimeout(fn func()) {
	timeout := time.AfterFunc(60*time.Second, func() {
		tt.Log("process timeout, killing")
		tt.Kill()
	})
	defer timeout.Stop()
	fn()
}

func (tt *TestExecCommand) Kill() {
	_ = tt.cmd.Process.Kill()
	if tt.Cleanup != nil {
		tt.Cleanup()
	}
}

func (tt *TestExecCommand) Expect(tplsource string) {
	tpl := template.Must(template.New("").Funcs(tt.Func).Parse(tplsource))
	wantbuf := new(bytes.Buffer)
	require.Nil(tt, tpl.Execute(wantbuf, tt.Data))

	want := bytes.TrimPrefix(wantbuf.Bytes(), []byte("\n"))
	tt.matchExactOutput(want)

	tt.Logf("stdout matched:\n%s", want)
}

func (tt *TestExecCommand) matchExactOutput(want []byte) {
	buf := make([]byte, len(want))
	n := 0
	tt.withKillTimeout(func() { n, _ = io.ReadFull(tt.stdout, buf) })
	buf = buf[:n]
	if n < len(want) || !bytes.Equal(buf, want) {
		buf = append(buf, make([]byte, tt.stdout.Buffered())...)
		_, _ = tt.stdout.Read(buf[n:])
		require.Equal(tt, want, buf)
	}
}

func (tt *TestExecCommand) StderrText() string {
	tt.stderr.mu.Lock()
	defer tt.stderr.mu.Unlock()
	return tt.stderr.buf.String()
}

func newTestCommand(t *testing.T, data interface{}) *TestExecCommand {
	return &TestExecCommand{T: t, Data: data}
}

func runTestZKP(t *testing.T, args ...string) *testzkp {
	tt := &testzkp{}
	tt.TestExecCommand = newTestCommand(t, tt)
	tt.Run("zkp-test", args...)
	return tt
}

func runTestSIGNING(t *testing.T, args ...string) *testsigning {
	tt := &testsigning{}
	tt.TestExecCommand = newTestCommand(t, tt)
	tt.Run("signing-test", args...)
	return tt
}

func TestMain(m *testing.M) {
	if reexec.Init() {
		return
	}
	os.Exit(m.Run())
}

func TestHelpCommand(t *testing.T) {
	zkp := runTestZKP(t, "--help")
	zkp.ExpectExitWith("help")
}

func TestInvalidCommand(t *testing.T) {
	zkp := runTestZKP(t, "--doritos")

	zkp.ExpectExitWith("doritos")
}
