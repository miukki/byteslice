package signing

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"strings"
	"test-zkp/cmd/poc/api/model"

	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"

	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	log "github.com/sirupsen/logrus"
)

type OtcTXN struct {
	Id         string
	Recipients []string
}

var (
	testPKey  = "" // will contain the 1st private key from Ganache (provided through ENV)
	smiloPKey = "" // will contain private key from smilo wallet (provided by HD-Wallet, through ENV)
)

var (
	ZKPKEY   = model.ZKPKEY
	MNEMONIC = model.MNEMONIC
	HDPATH   = model.HDPATH
)

const (
	filePath     = "cmd/poc/signing/files/"
	testFilePath = "signing/files/"
)

// VerifyFileSignature will check if the signature provided is signed by the entity belonging to the publickey provided
func VerifyFileSignature(filename string, apikey string, publickey string, signature string) (bool, error) {

	// start setting the signatureOK to false, for we have not tested if it was fine
	signatureOK := false

	// Check if API key matches the one in our environment
	if ZKPKEY != apikey {
		// if no API key found than bail out
		log.Error("no legal test zkp api key used")
		return false, errors.New("no legal test zkp api key used")
	}

	// We only continue with a valid API key

	// var to hold (truncated) signature
	var sig string

	// strip 0x from signature if this is how this values start
	if signature[:2] == "0x" || signature[:2] == "0X" {
		sig = signature[2:]
	} else {
		sig = signature
	}

	sigdata, err := hex.DecodeString(sig)
	if err != nil {
		return false, err
	}

	if len(sigdata) != 65 {
		return false, errors.New("signature should be 65 bytes long")
	}

	// Adjust signature to make it compatible with ether.js
	if sigdata[64] == 27 || sigdata[64] == 28 {
		sigdata[64] -= 27
	}

	// Give feedback on what we are about to do
	log.WithField("Processing", filePath+filename).Info("test the file that will be processed")

	// for now we have a hardcoded file, read the entire file in a byte array
	filebytes, err := os.ReadFile(filePath + filename)
	if err != nil { // If not found in the normal path, try the test path
		filebytes, err = os.ReadFile(testFilePath + filename)
		if err != nil {
			return false, err
		}
	}

	// Now proceed verifying the filebytes with signature given the public key
	signatureOK, err = VerifyFileBytesSignature(filebytes, ZKPKEY, publickey, signature)
	if err != nil {
		return false, err
	}

	return signatureOK, nil
}

// VerifyFileBytesSignature will check if the signature provided is signed by the entity belonging to the publickey provided
func VerifyFileBytesSignature(filebytes []byte, apikey string, publickey string, signature string) (bool, error) {

	// start setting the signatureOK to false, for we have not tested if it was fine
	signatureOK := false

	// Check if API key matches the one in our environment
	if ZKPKEY != apikey {
		// if no API key found than bail out
		log.Error("no legal test zkp api key used")
		return false, errors.New("no legal test zkp api key used")
	}

	// We only continue with a valid API key

	// var to hold (truncated) signature
	var sig string

	// strip 0x from signature if this is how this values start
	if signature[:2] == "0x" || signature[:2] == "0X" {
		sig = signature[2:]
	} else {
		sig = signature
	}

	sigdata, err := hex.DecodeString(sig)
	if err != nil {
		return false, err
	}

	if len(sigdata) != 65 {
		return false, errors.New("signature should be 65 bytes long")
	}

	// Adjust signature to make it compatible with ether.js
	if sigdata[64] == 27 || sigdata[64] == 28 {
		sigdata[64] -= 27
	}

	// Calculate hash from file and add ethereum header
	filehash := prefixHash(filebytes)

	// log hash of message
	log.WithField("Hash value", filehash.String()).Info("hash of the file")

	pubkeyfromsig, err := crypto.SigToPub(filehash.Bytes(), sigdata)
	if err != nil {
		return false, err
	}

	// create public key address from public key
	pubkeyaddr := crypto.PubkeyToAddress(*pubkeyfromsig)
	log.WithField("public key from signature", pubkeyaddr).Info("Test public key address from signature")

	// Only continue if the address is the same as the one that we are testing
	if strings.Compare(strings.ToLower(pubkeyaddr.String()), strings.ToLower(publickey)) == 0 {

		// sign message and using the private key (smilo or ganache) then capture signature
		sigPublicKey, err := crypto.Ecrecover(filehash.Bytes(), sigdata)
		if err != nil {
			return false, err
		}

		// create []byte from *ecdsa.PublicKey value
		pubkeydata := crypto.FromECDSAPub(pubkeyfromsig)

		// Assert that signature found matches the one that we expected
		signatureOK = bytes.Equal(sigPublicKey, pubkeydata)
	}

	if signatureOK {
		log.Info("Signature correct!")
	} else {
		log.Info("Signature NOT correct!")
	}

	return signatureOK, nil
}

// VerifySignature will check if the signature provided is signed by the entity belonging to the publickey provided
func VerifySignature(message string, apikey string, publickey string, signature string) (bool, error) {

	// start setting the signatureOK to false, for we have not tested if it was fine
	signatureOK := false

	// Check if API key matches the one in our environment
	if ZKPKEY != apikey {
		// if no API key found than bail out
		log.Error("no legal test zkp api key used")
		return false, errors.New("no legal test zkp api key used")
	}

	// We only continue with a valid API key

	// var to hold (truncated) signature
	var sig string

	// strip 0x from signature if this is how this values start
	if signature[:2] == "0x" || signature[:2] == "0X" {
		sig = signature[2:]
	} else {
		sig = signature
	}

	sigdata, err := hex.DecodeString(sig)
	if err != nil {
		return false, err
	}

	if len(sigdata) != 65 {
		return false, errors.New("signature should be 65 bytes long")
	}

	// Adjust signature to make it compatible with ether.js
	if sigdata[64] == 27 || sigdata[64] == 28 {
		sigdata[64] -= 27
	}

	// Calculate hash of messageBytes and add ethereum header, this is the hash that we sign
	ethHashResult := prefixHash([]byte(message))

	// log hash of message
	log.WithField("Hash value", ethHashResult.String()).Info("hash of the message")

	pubkeyfromsig, err := crypto.SigToPub(ethHashResult.Bytes(), sigdata)
	if err != nil {
		return false, err
	}

	// create public key address from public key
	pubkeyaddr := crypto.PubkeyToAddress(*pubkeyfromsig)
	log.WithField("public key from signature", pubkeyaddr).Info("Test public key address from signature")

	// Only continue if the address is the same as the one that we are testing
	if strings.Compare(strings.ToLower(pubkeyaddr.String()), strings.ToLower(publickey)) == 0 {

		// sign message and using the private key (smilo or ganache) then capture signature
		sigPublicKey, err := crypto.Ecrecover(ethHashResult.Bytes(), sigdata)
		if err != nil {
			return false, err
		}

		// create []byte from *ecdsa.PublicKey value
		pubkeydata := crypto.FromECDSAPub(pubkeyfromsig)

		// Assert that signature found matches the one that we expected
		signatureOK = bytes.Equal(sigPublicKey, pubkeydata)
	}

	if signatureOK {
		log.Info("Signature correct!")
	} else {
		log.Info("Signature NOT correct!")
	}

	return signatureOK, nil
}

// FileSigner to sign the hash of a given file
func FileSigner(smilo bool, filename string, apikey string) (string, error) {
	var err error
	var pk *ecdsa.PrivateKey

	// Check if API key matches the one in our environment
	if ZKPKEY != apikey {
		// if no API key found than bail out
		log.Error("no legal test zkp api key used")
		return "", errors.New("no legal test zkp api key used")
	}

	// We only continue with a valid API key

	// use Smilo HD-wallet if the smilo switch was set
	if smilo {

		err = readKeyFromHdWallet()
		if err != nil {
			return "", err
		}

		// get private key (parameter)
		pk, err = crypto.HexToECDSA(smiloPKey)
		if err != nil {
			return "", err
		}

	} else {

		// Try to fetch Ganache pk from the environment
		testPKey = model.GanacheKey
		if testPKey == "" {
			log.Error("could not find Ganache key in env")
			return "", errors.New("could not find Ganache key in env")
		}

		// get private key (parameter)
		pk, err = crypto.HexToECDSA(testPKey)
		if err != nil {
			return "", err
		}
	}

	// get deployer address from public key
	publicKey := pk.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	// Get the address (publickey) as a hexvalue
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	// log public key
	log.WithField("public key", address).Info("Public key of the key that put the signature")

	// Give feedback on what we are about to do
	log.WithField("Processing", filePath+filename).Info("test the file that will be processed")

	// for now we have a hardcoded file, read the entire file in a byte array
	filebytes, err := os.ReadFile(filePath + filename)
	if err != nil { // If not found in the normal path, try the test path
		filebytes, err = os.ReadFile(testFilePath + filename)
		if err != nil {
			return "", err
		}
	}

	// Proceed signing the fileBytes
	signature, err := FileBytesSigner(smilo, filebytes, apikey)
	if err != nil {
		return "", err
	}

	return signature, nil
}

// FileBytesSigner to sign the hash of a given file
func FileBytesSigner(smilo bool, filebytes []byte, apikey string) (string, error) {
	var err error
	var pk *ecdsa.PrivateKey

	// Check if API key matches the one in our environment
	if ZKPKEY != apikey {
		// if no API key found than bail out
		log.Error("no legal test zkp api key used")
		return "", errors.New("no legal test zkp api key used")
	}

	// We only continue with a valid API key

	// use Smilo HD-wallet if the smilo switch was set
	if smilo {

		err = readKeyFromHdWallet()
		if err != nil {
			return "", err
		}

		// get private key (parameter)
		pk, err = crypto.HexToECDSA(smiloPKey)
		if err != nil {
			return "", err
		}

	} else {

		// Try to fetch Ganache pk from the environment
		testPKey = model.GanacheKey
		if testPKey == "" {
			log.Error("could not find Ganache key in env")
			return "", errors.New("could not find Ganache key in env")
		}

		// get private key (parameter)
		pk, err = crypto.HexToECDSA(testPKey)
		if err != nil {
			return "", err
		}

		// get deployer address from public key
		publicKey := pk.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		}

		// Get the address (publickey) as a hexvalue
		address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

		// log public key
		log.WithField("ganache key", address).Info("Public key used in Ganache")
	}

	// Calculate hash of fole and add ethereum prefix
	filehash := prefixHash(filebytes)

	// log hash of file
	log.WithField("file hash", filehash.String()).Info("test hash of file")

	// sign message and using the private key (smilo or ganache) then capture signature
	signatureBytes, err := crypto.Sign(filehash.Bytes(), pk)
	if err != nil {
		return "", err
	}

	signature := hexutil.Encode(signatureBytes)

	// log signarure on message
	log.WithField("Signature", signature).Info("signature on the file")

	return signature, nil
}

// Signer the default entrypoint which sets up the keys
func Signer(smilo bool, message string, apikey string) (string, error) {
	var err error
	var pk *ecdsa.PrivateKey

	// Check if API key matches the one in our environment
	if ZKPKEY != apikey {
		// if no API key found than bail out
		log.Error("no legal test zkp api key used")
		return "", errors.New("no legal test zkp api key used")
	}

	// We only continue with a valid API key

	// use Smilo HD-wallet if the smilo switch was set
	if smilo {

		err = readKeyFromHdWallet()
		if err != nil {
			return "", err
		}

		// get private key (parameter)
		pk, err = crypto.HexToECDSA(smiloPKey)
		if err != nil {
			return "", err
		}

	} else {

		// Try to fetch Ganache pk from the environment
		testPKey = model.GanacheKey
		if testPKey == "" {
			log.Error("could not find Ganache key in env")
			return "", errors.New("could not find Ganache key in env")
		}

		// get private key (parameter)
		pk, err = crypto.HexToECDSA(testPKey)
		if err != nil {
			return "", err
		}

		// get deployer address from public key
		publicKey := pk.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		}

		// Get the address (publickey) as a hexvalue
		address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

		// log public key
		log.WithField("ganache key", address).Info("Public key used in Ganache")
	}

	//Calculate hash of messageBytes with ethereum prefix, this is the hash that we sign
	ethHashResult := prefixHash([]byte(message))

	// log hash of message
	log.WithField("Hash value", ethHashResult.String()).Info("hash of the message")

	// sign message and using the private key (smilo or ganache) then capture signature
	signatureBytes, err := crypto.Sign(ethHashResult.Bytes(), pk)
	if err != nil {
		return "", err
	}

	signature := hexutil.Encode(signatureBytes)

	// log signature on message
	log.WithField("Signature", signature).Info("signature of the message")

	return signature, nil
}

// adds a prefix to the message before its hashed
func prefixHash(data []byte) common.Hash {
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(data), data)
	return crypto.Keccak256Hash([]byte(msg))
}

// readKeyFromHdWallet helper function to read the keys from environment; either Smilo HD-Wallet or Ganache key
func readKeyFromHdWallet() error {

	// Setup wallet from mnemonic string
	wallet, err := hdwallet.NewFromMnemonic(MNEMONIC)
	if err != nil {
		return err
	}

	// Set HD wallet path and extract account details
	path := hdwallet.MustParseDerivationPath(HDPATH)
	account, err := wallet.Derive(path, false)
	if err != nil {
		return err
	}

	pkey, err := wallet.PrivateKeyHex(account)
	if err != nil {
		return err
	}

	// log public key
	log.WithField("smilo key", account.Address.Hex()).Info("Public key used on Smilo")

	// set smiloKey
	smiloPKey = pkey

	return nil
}
