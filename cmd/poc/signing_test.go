package main

import (
	"regexp"
	"strings"
	"test-zkp/cmd/poc/api/model"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	TESTMESSAGE   = `Vosvor Exchange rocks!`
	GOODSIGNATURE = `0xe2bbe850d6dee98081fd99110aa32f4f8a6cbe6a6498bd2d17fef1d29e59cc44728b95504b55957a616286fc8099a08544d380e2d392f3792d0bf3d61987190601`
	PUBKEY        = `0x6FBc600C44DC1Cd1321F893bb1F04AA004C3158F`
)

func TestSignerSignMessage(t *testing.T) {

	// Check if API key matches the one in our environment
	envkey = model.ZKPKEY
	require.NotEmpty(t, envkey)

	// Set test message
	testmessage := TESTMESSAGE

	signing := runTestSIGNING(t, "signing", "--smilo", "--msg="+testmessage, "--zkpkey="+envkey)

	signing.WaitExit()

	re1, _ := regexp.Compile(".*value=.*") // look for value=... (here we find our hash value)
	s1 := re1.FindString(signing.StderrText())

	re2, _ := regexp.Compile(".*Signature=.*") // look for Signature=... (here we find our hash value)
	s2 := re2.FindString(signing.StderrText())

	starthash := strings.Index(s1, "value=") // find location of value= in this line
	response1 := s1[starthash+6:]            // capture the string after =

	// describe the expected hash value, which should match the hash we found in the output
	expectedhash := "0x0926ffcabd20504ba5035b9e6ea3201f1d48610aff761c5489be0718a0002b5c"
	require.Contains(t, response1, expectedhash)

	startsig := strings.Index(s2, "Signature=") // find location of value= in this line
	response2 := s2[startsig+10:]               // capture the string after =

	// describe the expected hash value, which should match the hash we found in the output
	expectedsig := GOODSIGNATURE
	require.Contains(t, response2, expectedsig)

	expectedMessages := []string{
		`signerAction completed`,
	}

	for _, m := range expectedMessages {
		require.Contains(t, signing.StderrText(), m)
	}
}

func TestSignerVerifySignatureMessage(t *testing.T) {

	// Set test message
	testmessage := TESTMESSAGE
	testsignature := GOODSIGNATURE
	pubkey := PUBKEY

	signing := runTestSIGNING(t, "signing", "--verifysig", "--msg="+testmessage, "--pubkey="+pubkey, "--signature="+testsignature, "--zkpkey="+envkey)

	signing.WaitExit()

	expectedMessages := []string{
		`Signature correct!`,
		`signerAction completed`,
	}

	for _, m := range expectedMessages {
		require.Contains(t, signing.StderrText(), m)
	}
}

func TestSignerVerifySignatureMessage_WrongSignature(t *testing.T) {

	// Set test message
	testmessage := TESTMESSAGE
	testsignature := `0xe2bbe850d6dee98081fd99110aa32f4f8a6cbe6a6498bd2d17fef1d29e59cc44728b95504b55957a616286fc8099a08544d380e2d392f3792d0bf3d61987190001`
	pubkey := PUBKEY

	signing := runTestSIGNING(t, "signing", "--verifysig", "--msg="+testmessage, "--pubkey="+pubkey, "--signature="+testsignature, "--zkpkey="+envkey)

	signing.WaitExit()

	expectedMessages := []string{
		`Signature NOT correct!`,
		`signerAction completed`,
	}

	for _, m := range expectedMessages {
		require.Contains(t, signing.StderrText(), m)
	}
}

func TestSignerVerifySignatureMessage_WrongPublicKey(t *testing.T) {

	// Set test message
	testmessage := TESTMESSAGE
	testsignature := GOODSIGNATURE
	pubkey := `0x6FBc600C44DC1Cd1321F893bb1F04AA004C30000`

	signing := runTestSIGNING(t, "signing", "--verifysig", "--msg="+testmessage, "--pubkey="+pubkey, "--signature="+testsignature, "--zkpkey="+envkey)

	signing.WaitExit()

	expectedMessages := []string{
		`Signature NOT correct!`,
		`signerAction completed`,
	}

	for _, m := range expectedMessages {
		require.Contains(t, signing.StderrText(), m)
	}
}

func TestSignerSignFile(t *testing.T) {

	testfile := `testdocument.pdf`

	signing := runTestSIGNING(t, "signing", "--smilo", "--filesign", "--file="+testfile, "--zkpkey="+envkey)
	signing.WaitExit()

	re1, _ := regexp.Compile(".*hash=.*") // look for hash=... (here we find our hash value)
	s1 := re1.FindString(signing.StderrText())

	re2, _ := regexp.Compile(".*Signature=.*") // look for Signature=... (here we find our hash value)
	s2 := re2.FindString(signing.StderrText())

	starthash := strings.Index(s1, "hash=") // find location of hash= in this line
	response1 := s1[starthash+5:]           // capture the string after =

	// describe the expected hash value, which should match the hash we found in the output
	expectedhash := "0x687aab4ecbd96d9f196544db1d4eeafe14f193d8cbe8381de63ba873097f2e1e"
	require.Contains(t, response1, expectedhash)

	startsig := strings.Index(s2, "Signature=") // find location of value= in this line
	response2 := s2[startsig+10:]               // capture the string after =

	// describe the expected hash value, which should match the hash we found in the output
	expectedsig := "0x30fe2768c7365d370c8a1556a9d96dca07d0ab380cfac2a041099746ba8ffc936db557cdb310d1cf0d94e0bc587079a95b3e8b812225cd59e22789163222ab8700"
	require.Contains(t, response2, expectedsig)

	expectedMessages := []string{
		`signerAction completed`,
	}

	for _, m := range expectedMessages {
		require.Contains(t, signing.StderrText(), m)
	}
}

func TestSignerVerifySignatureFile(t *testing.T) {

	// Set test file and signature and pubkey vars
	testfile := `testdocument.pdf`
	testsignature := `0x30fe2768c7365d370c8a1556a9d96dca07d0ab380cfac2a041099746ba8ffc936db557cdb310d1cf0d94e0bc587079a95b3e8b812225cd59e22789163222ab8700`
	pubkey := PUBKEY

	signing := runTestSIGNING(t, "signing", "--verifyfilesig", "--file="+testfile, "--pubkey="+pubkey, "--signature="+testsignature, "--zkpkey="+envkey)

	signing.WaitExit()

	expectedMessages := []string{
		`Signature correct!`,
		`signerAction completed`,
	}

	for _, m := range expectedMessages {
		require.Contains(t, signing.StderrText(), m)
	}
}
