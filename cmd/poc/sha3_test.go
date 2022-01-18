package main

import (
	"regexp"
	"strings"
	"test-zkp/cmd/poc/api/model"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

const (
	TESTFILE = `testdocument.pdf`
)

// global var with the address of the smart contract
var addr string

// where we temporarily hold the api key
var envkey string

func TestSha3(t *testing.T) {
	zkp := runTestZKP(t, "sha3")

	zkp.WaitExit()

	expectedMessages := []string{
		`sha3Action completed`,
	}

	for _, m := range expectedMessages {
		require.Contains(t, zkp.StderrText(), m)
	}
}

func TestSha3HashText(t *testing.T) {

	// Check if API key matches the one in our environment
	envkey = model.ZKPKEY
	require.NotEmpty(t, envkey)

	teststr := `Test`
	hashtype := `NEWLEGACYKECCAK256`

	zkp := runTestZKP(t, "sha3", "--texthash", "--txt="+teststr, "--hashtype="+hashtype, "--zkpkey="+envkey)

	zkp.WaitExit()

	re, _ := regexp.Compile(".*hash=0x[0-9a-fA-F]+.*") // look for Sha3=0x... (hash signature)
	s := re.FindString(zkp.StderrText())
	start := strings.Index(s, "=") // find '=' in string s
	response := s[start:]          // capture everything after '=' until the end of the line
	if len(response) > 66 {
		response = response[len(response)-66:] // capture sha3 value in response
	}

	// Now we require the hash to be exactlty this
	require.EqualValues(t, "0x3f1e3e8564db8b21b9db081be0cb8f2ead919a78240c0bd2c084e762018f0e78", response)

	expectedMessages := []string{
		`sha3Action completed`,
	}

	for _, m := range expectedMessages {
		require.Contains(t, zkp.StderrText(), m)
	}
}

func TestSha3HashFile(t *testing.T) {

	// Check if API key matches the one in our environment
	envkey = model.ZKPKEY
	require.NotEmpty(t, envkey)

	testfile := TESTFILE
	hashtype := `NEWLEGACYKECCAK256`

	zkp := runTestZKP(t, "sha3", "--filehash", "--file="+testfile, "--hashtype="+hashtype, "--zkpkey="+envkey)

	zkp.WaitExit()

	re, _ := regexp.Compile(".*hash=0x[0-9a-fA-F]+.*") // look for Sha3=0x... (hash signature)
	s := re.FindString(zkp.StderrText())
	start := strings.Index(s, "=") // find '=' in string s
	response := s[start:]          // capture everything after '=' until the end of the line
	if len(response) > 66 {
		response = response[len(response)-66:] // capture sha3 value in response
	}

	// Now we require the hash to be exactlty this
	require.EqualValues(t, "0xeb1f8e1ea395313426bbdfbf2060409e07d9136c0fb86c54c4f4e3c85411e361", response)

	expectedMessages := []string{
		`sha3Action completed`,
	}

	for _, m := range expectedMessages {
		require.Contains(t, zkp.StderrText(), m)
	}
}

func TestSha3Deploy(t *testing.T) {

	// Check if API key matches the one in our environment
	envkey = model.ZKPKEY
	require.NotEmpty(t, envkey)

	mnemonic := model.MNEMONIC
	require.NotEmpty(t, mnemonic)

	zkp := runTestZKP(t, "sha3", "--smilo", "--deploy", "--zkpkey="+envkey)

	zkp.WaitExit()

	expectedMessages := []string{
		`sha3Action completed`,
	}

	for _, m := range expectedMessages {
		require.Contains(t, zkp.StderrText(), m)
	}

	re, _ := regexp.Compile(".*address=0x[0-9a-fA-F]+")
	s := re.FindString(zkp.StderrText())
	addr = s[len(s)-42:]
}

func TestSha3LogHash(t *testing.T) {

	// first wait for 25 secs to let the block be mined
	time.Sleep(25 * time.Second)

	// order json
	ojson := `{"_id": "trades/1234567","recipients": ["1000000","111111","999999"]}`

	// call loghash function
	zkp := runTestZKP(t, "sha3", "--smilo", "--loghash", "--address="+addr, "--txn="+ojson, "--type=TRADE", "--state=COMPLETION", "--zkpkey="+envkey)

	zkp.WaitExit()

	expectedMessages := []string{
		`sha3Action completed`,
	}

	for _, m := range expectedMessages {
		require.Contains(t, zkp.StderrText(), m)
	}
}

func TestSha3LogTradeStateOutOfOrder(t *testing.T) {

	// first wait for 25 secs to let the block be mined
	time.Sleep(25 * time.Second)

	// order json
	ojson := `{"_id": "trades/1","recipients": ["1000000","111111","999999"]}`

	// call loghash function
	zkp1 := runTestZKP(t, "sha3", "--smilo", "--loghash", "--address="+addr, "--txn="+ojson, "--type=TRADE", "--state=CREATED", "--zkpkey="+envkey)
	zkp1.WaitExit()

	// call loghash function
	zkp2 := runTestZKP(t, "sha3", "--smilo", "--loghash", "--address="+addr, "--txn="+ojson, "--type=TRADE", "--state=CLOSED", "--zkpkey="+envkey)
	zkp2.WaitExit()

	// first wait for 5 secs to let the log instruction settle
	time.Sleep(5 * time.Second)

	// call loghash function
	zkp3 := runTestZKP(t, "sha3", "--smilo", "--loghash", "--address="+addr, "--txn="+ojson, "--type=TRADE", "--state=SIGNATURES", "--zkpkey="+envkey)
	zkp3.WaitExit()

	zkp4 := runTestZKP(t, "sha3", "-smilo", "--gethistory", "--address="+addr, "--id=1", "--type=TRADE", "--zkpkey="+envkey)
	zkp4.WaitExit()

	re, _ := regexp.Compile(".*states=.*") // look for States=... (historical states signature)
	s := re.FindString(zkp4.StderrText())

	start := strings.Index(s, "states=") // find '=' in string s
	response := s[start:]                // capture everything after '=' until the end of the line

	notExpectedMessage := "SIGNATURES" // we do not expect to find this string in response

	// Now require that the Historical states does not include SIGNATURES in the response
	// for we are not allowed to log a state that is lower than the last logged state which in this test
	// case is the last state CLOSED.
	require.NotContains(t, response, notExpectedMessage)
}

func TestSha3LogOrderStateOutOfOrder(t *testing.T) {

	// first wait for 25 secs to let the block be mined
	time.Sleep(25 * time.Second)

	// order json
	ojson := `{"_id": "orders/2","recipients": ["1000000","111111","999999"]}`

	// call loghash function
	zkp := runTestZKP(t, "sha3", "--smilo", "--loghash", "--address="+addr, "--txn="+ojson, "--type=ORDER", "--state=CREATED", "--zkpkey="+envkey)
	zkp.WaitExit()

	// call loghash function
	zkp1 := runTestZKP(t, "sha3", "--smilo", "--loghash", "--address="+addr, "--txn="+ojson, "--type=ORDER", "--state=TRADED", "--zkpkey="+envkey)
	zkp1.WaitExit()

	// first wait for 5 secs to let the log instruction settle
	time.Sleep(5 * time.Second)

	// call loghash function
	zkp2 := runTestZKP(t, "sha3", "--smilo", "--loghash", "--address="+addr, "--txn="+ojson, "--type=ORDER", "--state=EXPIRED", "--zkpkey="+envkey)
	zkp2.WaitExit()

	zkp3 := runTestZKP(t, "sha3", "-smilo", "--gethistory", "--address="+addr, "--id=2", "--type=ORDER", "--zkpkey="+envkey)
	zkp3.WaitExit()

	re, _ := regexp.Compile(".*states=.*") // look for States=... (historical states signature)
	s := re.FindString(zkp3.StderrText())

	start := strings.Index(s, "states=") // find '=' in string s
	response := s[start:]                // capture everything after '=' until the end of the line

	notExpectedMessage := "EXPIRED" // we do not expect to find this string in response

	// Now require that the Historical states does not include EXPIRED in response
	// for we are not allowed to log a state that is lower than the last logged state which in this test
	// case is the last state TRADED.
	require.NotContains(t, response, notExpectedMessage)
}

func TestSha3GetLog(t *testing.T) {

	// first wait for 25 secs to let the block be mined
	time.Sleep(25 * time.Second)

	zkp := runTestZKP(t, "sha3", "-smilo", "--getlog", "--address="+addr, "--id=1234567", "--type=TRADE", "--zkpkey="+envkey)
	zkp.WaitExit()

	re, _ := regexp.Compile(".*Sha3=0x[0-9a-fA-F]+.*") // look for Sha3=0x... (hash signature)
	s := re.FindString(zkp.StderrText())
	start := strings.Index(s, "=") // find '=' in string s
	response := s[start:]          // capture everything after '=' until the end of the line
	if len(response) > 66 {
		response = response[len(response)-66:] // capture sha3 value in response
	}

	// Now we require the hash to be exactlty this
	require.EqualValues(t, "0x7324821d6206da064bdf8c587076cf1e5c92f2120ba4f13d0a7b7fb433706e9a", response)

	re, _ = regexp.Compile(".*Type=.*") // look for Type=
	s = re.FindString(zkp.StderrText())
	start = strings.Index(s, "=") // find '=' in string s
	response = s[start:]          // capture everything '=' until the end of the line
	if len(response) > 1 {
		response = response[len(response)-1:] // capture type (single digit) value in response
	}

	// Now we require the type to be exactlty this
	require.EqualValues(t, "2", response) // 2 --> TYPE="TRADE"

	re, _ = regexp.Compile(".*State=.*") // look for State=
	s = re.FindString(zkp.StderrText())
	start = strings.Index(s, "=") // find '=' in string s
	response = s[start:]          // capture everything '=' until the end of the line
	if len(response) > 1 {
		response = response[len(response)-1:] // capture state (single digit) value in response
	}

	// Now we require the state to be exactlty this
	require.EqualValues(t, "5", response) // 5 --> STATE="COMPLETION"
}

func TestSha3GetHash(t *testing.T) {

	zkp := runTestZKP(t, "sha3", "-smilo", "--gethash", "--address="+addr, "--id=1234567", "--type=TRADE", "--zkpkey="+envkey)
	zkp.WaitExit()

	re, _ := regexp.Compile(".*Sha3=0x[0-9a-fA-F]+.*") // look for Sha3=0x... (hash signature)
	s := re.FindString(zkp.StderrText())
	start := strings.Index(s, "=") // find '=' in string s
	response := s[start:]          // capture everything after '=' until the end of the line
	if len(response) > 66 {
		response = response[len(response)-66:] // capture sha3 value in response
	}

	// Now we require the hash to be exactlty this
	require.EqualValues(t, "0x7324821d6206da064bdf8c587076cf1e5c92f2120ba4f13d0a7b7fb433706e9a", response)
}

func TestSha3GetHistory(t *testing.T) {

	zkp := runTestZKP(t, "sha3", "-smilo", "--gethistory", "--address="+addr, "--id=1234567", "--type=TRADE", "--zkpkey="+envkey)
	zkp.WaitExit()

	re, _ := regexp.Compile(".*states=.*") // look for States=... (historical states signature)
	s := re.FindString(zkp.StderrText())
	start := strings.Index(s, "=") // find '=' in string s
	response := s[start:]          // capture everything after '=' until the end of the line

	// One of the states included should be COMPLETION
	require.Contains(t, response, "COMPLETION")
}

func TestSha3VerifyLatest(t *testing.T) {

	// order json
	ojson := `{"_id": "orders/1","recipients": ["1000000","111111","999999"]}`

	// call loghash function
	zkp := runTestZKP(t, "sha3", "--smilo", "--loghash", "--address="+addr, "--txn="+ojson, "--type=ORDER", "--state=CREATED", "--zkpkey="+envkey)
	zkp.WaitExit()

	// first wait for 5 secs to let the log instruction settle
	time.Sleep(5 * time.Second)

	zkp1 := runTestZKP(t, "sha3", "-smilo", "--verifylatest", "--address="+addr, "--id=1", "--type=ORDER", "--txn="+ojson, "--zkpkey="+envkey)
	zkp1.WaitExit()

	re, _ := regexp.Compile(".*Result=.*") // look for States=... (historical states signature)
	s := re.FindString(zkp1.StderrText())

	start := strings.Index(s, "Result=") // find '=' in string s
	response := s[start:]                // capture everything after '=' until the end of the line

	start = strings.Index(response, "=") // find '=' in string s
	response = response[start+1:]        // capture everything '=' until the end of the line

	// Now we require the hash to be exactlty this
	require.EqualValues(t, "true", response)
}
