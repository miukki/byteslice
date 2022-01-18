package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBytesliceInit(t *testing.T) {
	zkp := runTestZKP(t, "byteslice", "--init")

	zkp.WaitExit()

	expectedMessages := []string{
		`bytesliceAction completed`,
	}

	for _, m := range expectedMessages {
		require.Contains(t, zkp.StderrText(), m)
	}

}
