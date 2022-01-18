package sha3

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"math"
	"math/big"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"test-zkp/cmd/poc/api/model"
	deployer "test-zkp/cmd/poc/sha3/deployer"
	"time"

	"golang.org/x/crypto/sha3"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	log "github.com/sirupsen/logrus"
)

type ORDERTRADE struct {
	Id string                 `json:"_id"` // just need to be able to capture the _id from the parsed json Txn
	X  map[string]interface{} `json:"-"`   // te remainder of the ORDERTRADE struct we can ignore
}

var (
	ZKPKEY   = model.ZKPKEY
	MNEMONIC = model.MNEMONIC
	HDPATH   = model.HDPATH
)

const (
	chainID                  = 5777
	chainPORT                = 7545
	smiloNET                 = 20080914
	smiloPORT                = 442
	smiloURL                 = "https://api.smilo.foundation"
	apiKeyLen                = 32
	trxJson                  = `{"id": "1234567","recipients": ["1000000","111111","999999"]}`
	AbigenVersion            = "github.com/ethereum/go-ethereum/cmd/abigen@v1.10.3"
	filePath                 = "cmd/poc/signing/files/"
	testFilePath             = "signing/files/"
	TxnTypeOrder             = "ORDER"
	TxnTypeOtcOrder          = "OTCORDER"
	TxnTypeTrade             = "TRADE"
	OrderStateCreated        = "CREATED"
	OrderStateExpired        = "EXPIRED"
	OrderStateTraded         = "TRADED"
	TradeStateCreated        = "CREATED"
	TradeStateSignatures     = "SIGNATURES"
	TradeStateInstructions   = "INSTRUCTIONS"
	TradeStateShippingAdvice = "SHIPPINGADVICE"
	TradeStareDocuments      = "DOCUMENTS"
	TradeStateCompletion     = "COMPLETION"
	TradeStateIntention      = "INTENTION"
	TradeStateDeclaration    = "DECLARATION"
	TradeStateClosed         = "CLOSED"
)

var (
	wrapperGo  = "sha3/deployer/wrapper.go"
	loggerSol  = "registry-contract/contracts/Logger.sol"
	basePath   = "sha3/deployer/"
	abiBinPath = "bin/abigen"
	testPKey   = "" // will contain the 1st private key from Ganache (provided through ENV)
	smiloPKey  = "" // will contain private key from smilo wallet (provided by HD-Wallet, through ENV)
)

var workingDir = ""

func init() {
	var err error
	workingDir, err = os.Getwd()
	if err != nil {
		panic(err)
	}
	if strings.Contains(workingDir, "server") {
		workingDir = workingDir + "/../../../cmd/poc"
	} else if !strings.Contains(workingDir, "/cmd") {
		workingDir = workingDir + "/cmd/poc"
	}
	loggerSol = filepath.Join(workingDir, "../..", loggerSol)
	wrapperGo = filepath.Join(workingDir, wrapperGo)
	basePath = filepath.Join(workingDir, basePath)
	abiBinPath = filepath.Join(workingDir, "../..", abiBinPath)
}

// GOBIN environment variable.
func GOBIN() string {
	if os.Getenv("GOBIN") == "" {
		log.Fatal("GOBIN not set")
	}

	return os.Getenv("GOBIN")
}

// Gettypename helper function to get typename from a typevalue
func Gettypename(tt uint8) (string, error) {

	var strType string
	var err error

	switch tt {
	case 0:
		strType = TxnTypeOrder
	case 1:
		strType = TxnTypeOtcOrder
	case 2:
		strType = TxnTypeTrade
	default:
		log.Error("unknown transaction type")
		err = errors.New("unknown transaction type")
	}
	if err != nil {
		return "", err
	}

	return strType, nil
}

// Getstatename helper function to get statename from a value for a particular type
func Getstatename(tt uint8, st uint8) (string, error) {

	// Convert State to string type
	var strState string
	var err error

	// First deal with Order states, then Trade states
	if tt <= 1 {
		switch st {
		case 0:
			strState = OrderStateCreated
		case 1:
			strState = OrderStateExpired
		case 2:
			strState = OrderStateTraded
		default:
			log.Error("unknown order state")
			err = errors.New("unknown order state")
		}
	} else { // Now deal with trade states
		switch st {
		case 0:
			strState = TradeStateCreated
		case 1:
			strState = TradeStateSignatures
		case 2:
			strState = TradeStateInstructions
		case 3:
			strState = TradeStateShippingAdvice
		case 4:
			strState = TradeStareDocuments
		case 5:
			strState = TradeStateCompletion
		case 6:
			strState = TradeStateIntention
		case 7:
			strState = TradeStateDeclaration
		case 8:
			strState = TradeStateClosed
		default:
			log.Error("unknown trade state")
			err = errors.New("unknown trade state")
		}
	}
	if err != nil {
		return "", err
	}

	return strState, nil
}

// GetTypeVal helper function will return the type value for the type-string
func GetTypeVal(ttype string) (uint8, error) {

	var TypeVal uint8
	var err error

	// if transaction type is order or otc order, set order state
	switch ttype {
	case TxnTypeOrder:
		TypeVal = 0
	case TxnTypeOtcOrder:
		TypeVal = 1
	case TxnTypeTrade:
		TypeVal = 2
	default:
		log.Error("unknown type used")
		err = errors.New("unknown Type used")
	}

	return TypeVal, err
}

// GetStateVal helper function will return the state value for the state-string + type value given
func GetStateVal(state string, tt uint8) (uint8, error) {

	var StateVal uint8
	var err error

	// if transaction type is order or otc order, set order state
	if tt <= 1 {
		switch state {
		case OrderStateCreated:
			StateVal = 0
		case OrderStateExpired:
			StateVal = 1
		case OrderStateTraded:
			StateVal = 2
		default:
			log.Error("unknown order state used")
			err = errors.New("unknown orderState used")
		}
	} else { // if not order or otcorder, then this must be a trade so mapping trade state now to uint8 value
		switch state {
		case TradeStateCreated:
			StateVal = 0
		case TradeStateSignatures:
			StateVal = 1
		case TradeStateInstructions:
			StateVal = 2
		case TradeStateShippingAdvice:
			StateVal = 3
		case TradeStareDocuments:
			StateVal = 4
		case TradeStateCompletion:
			StateVal = 5
		case TradeStateIntention:
			StateVal = 6
		case TradeStateDeclaration:
			StateVal = 7
		case TradeStateClosed:
			StateVal = 8
		default:
			log.Error("unknown trade state used")
			err = errors.New("unknown tradeState used")
		}
	}
	if err != nil {
		return uint8(0), err
	}

	return StateVal, nil
}

//nolint:gosec
func ensureAbigen() error {

	// Make sure abigen is downloaded and available.
	argsGet := []string{"get", AbigenVersion}
	cmd := exec.Command(filepath.Join(runtime.GOROOT(), "bin", "go"), argsGet...)

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Errorf("could not list pkgs: %v\n%s", err, string(out))
		return err
	}

	cmd = exec.Command(filepath.Join(GOBIN(), "abigen"))
	cmd.Args = append(cmd.Args, "--version")

	log.Info("Checking abigen version ...", strings.Join(cmd.Args, " \\\n"))
	cmd.Stderr, cmd.Stdout = os.Stderr, os.Stdout

	if err := cmd.Run(); err != nil {
		log.Error("Error: Could not Checking abigen version ... ", "error: ", err, ", cmd: ", cmd)
		return err
	}

	return nil
}

func deploySolidity(smilo bool) (*deployer.Logger, string, error) {

	// run abigen to generate go wrapper
	if _, err := os.Stat(filepath.Join(GOBIN(), "abigen")); os.IsNotExist(err) {
		err = ensureAbigen()
		if err != nil {
			return nil, "", err
		}
	}

	cmd := exec.Command(abiBinPath, "--sol", loggerSol, "--pkg", "ethereum", "--out", wrapperGo)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Error("Error: Could not execute abigen ... ", "error: ", err, ", cmd: ", cmd)
		return nil, "", err
	}

	var auth *bind.TransactOpts
	var client *ethclient.Client
	var err error

	if smilo {
		// prepare for write and setup authentication for account with deployPKey, let function estimate gas
		auth, client, err = prepareForWrite(smiloPKey, smiloNET, smiloPORT, smilo, 500000)
	} else {
		// prepare for write and setup authentication for account with deployPKey, let function estimate gas
		auth, client, err = prepareForWrite(testPKey, chainID, chainPORT, smilo, 0)
	}
	if err != nil {
		return nil, "", err
	}

	// now deploy the Logger contract
	address, tx, loggerContract, err := deployer.DeployLogger(auth, client)
	if err != nil {
		return nil, "", err
	}

	// log address and transaction hash for deploying this contract
	log.WithField("address", address.Hex()).Info("contract deployed at")
	log.WithField("hash", tx.Hash().Hex()).Info("hash of deploy transaction")

	// convert balance to ETH
	txcost := new(big.Float)
	txcost.SetString(tx.Cost().String())
	ethCosts := new(big.Float).Quo(txcost, big.NewFloat(math.Pow10(18)))

	log.WithField("cost", ethCosts.String()).Info("cost to deploy in ETH")

	// return pointer to contract
	return loggerContract, address.Hex(), nil
}

func Run() error {
	err := execute()
	if err != nil {
		log.WithError(err).Error("Failed to setup ADB data")
		return err
	}

	log.Info("Finished setting up ADB data")
	return nil
}

// CreateApiKey to create a random base64 API-Key
func CreateApiKey() (string, error) {

	// Create a slice to hold the API key with specified Length = apiKeyLen
	b := make([]byte, apiKeyLen)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	// Encode the random values captured in b in a base64 string value.
	apikey := base64.URLEncoding.EncodeToString(b)

	// Return API key to caller
	return apikey, nil
}

// prepareForWrite helper function to prepare for a smart contract write call and get a client for the contract
func prepareForWrite(acc string, chainid int, chainport int, smilo bool, gas int64) (*bind.TransactOpts, *ethclient.Client, error) {

	// build server URL
	var server string
	if smilo {
		server = smiloURL
	} else {
		server = strings.Join([]string{"http://localhost:", strconv.Itoa(chainport)}, "")
	}

	// log URL of RPC server
	log.WithField("URL", server).Info("test URL of RPC server")

	client, err := ethclient.Dial(server)
	if err != nil {
		return nil, nil, err
	}

	// get private key (parameter)
	privateKey, err := crypto.HexToECDSA(acc)
	if err != nil {
		return nil, nil, err
	}

	// get deployer address from public key
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	callAddr := crypto.PubkeyToAddress(*publicKeyECDSA)
	account := common.HexToAddress(callAddr.String())
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		return nil, nil, err
	}

	// convert balance to ETH
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))

	// log balance of the account used to deploy the contract
	log.WithField("balance", ethValue).Info("testing balance of account used")

	nonce, err := client.PendingNonceAt(context.Background(), account)
	if err != nil {
		return nil, nil, err
	}

	// log nonce
	log.WithField("nonce", nonce).Info("testing the nonce")

	// initialize gasPrice to be gas, the parameter used
	gasPrice := big.NewInt(gas)

	// if gas parameter was set to zero, estimate the required gas
	if gas == 0 {
		gasPrice, err = client.SuggestGasPrice(context.Background())
		if err != nil {
			return nil, nil, err
		}
	}

	// setup authentication
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(int64(chainid)))
	if err != nil {
		return nil, nil, err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	auth.GasLimit = 3000000    // in units
	auth.GasPrice = gasPrice

	// return euthentication and client
	return auth, client, nil
}

// getInstance helper function to prepare for a smart contract read-only call and get an instance of the contract
func getInstance(smilo bool, id string, addr string) (*deployer.Logger, [32]byte, error) {

	// depending on the smilo flag set server url
	var server string
	if smilo {
		server = smiloURL
	} else {
		// build server URL
		server = strings.Join([]string{"http://localhost:", strconv.Itoa(chainPORT)}, "")
	}

	// log URL of RPC server
	log.WithField("URL", server).Info("test URL of RPC server")

	client, err := ethclient.Dial(server)
	if err != nil {
		return nil, [32]byte{}, err
	}

	// Convert OderID from sample struct to [32]byte array
	txIdBytes := []byte(id)
	var txId [32]byte
	copy(txId[:], txIdBytes)

	//Log the result, and Ethereum compatible Keccak256 sha3 value
	log.WithField("ID 32Bytes val ", "0x"+common.Bytes2Hex(txIdBytes)+"00000000000000000000000000000000000000000000000000000000000000").Info("From ID =" + id)

	log.WithField("address", addr).Info("Contract at address")
	address := common.HexToAddress(addr)
	instance, err := deployer.NewLogger(address, client)
	if err != nil {
		return nil, [32]byte{}, err
	}

	return instance, txId, nil
}

// VerifyLatest function to verify a hash of a json record with the lastest
func VerifyLatest(smilo bool, id string, ttype string, txjson string, addr string, apikey string) (bool, error) {

	// Check if API key matches the one in our environment
	if ZKPKEY != apikey {
		// if no API key found than bail out
		log.Error("no legal test zkp api key used")
		return false, errors.New("no legal test zkp api key used")
	}

	log.WithField("json", txjson).Info("Check if this json matches the latest state")

	// Get instance to contract
	instance, txId, err := getInstance(smilo, id, addr)
	if err != nil {
		return false, err
	}

	typeVal, err := GetTypeVal(ttype)
	if err != nil {
		return false, err
	}

	// Call the contract at the given address to retrieve the hash of the latest state of order/trade with id
	log.Println(">>>> Call GetHash()")
	txHash, err := instance.GetHash(nil, txId, typeVal)
	if err != nil {
		return false, err
	}

	// Self generate a sha3 from the json string for which the claim is that it is consistent with the latest state
	myHash, err := Calchash([]byte(txjson), "NEWLEGACYKECCAK256")
	if err != nil {
		return false, err
	}

	// Compare sha3 hash given with the one found on the blockchain result is true for an exact match
	result := strings.Compare(txHash, common.BytesToHash(myHash).String()) == 0

	// Write result to log
	log.WithField("Result", result).Info("Testing hash of json with latest state")

	return result, nil
}

// Getlog function to retrieve all log details of an order or trade with id
func Getlog(smilo bool, id string, ttype string, addr string, apikey string) (string, string, string, error) {

	// Check if API key matches the one in our environment
	if ZKPKEY != apikey {
		// if no API key found than bail out
		log.Error("no legal test zkp api key used")
		return "", "", "", errors.New("no legal test zkp api key used")
	}

	// Get instance to contract
	instance, txId, err := getInstance(smilo, id, addr)
	if err != nil {
		return "", "", "", err
	}

	//Get Type val from type
	typeVal, err := GetTypeVal(ttype)
	if err != nil {
		return "", "", "", err
	}

	// call the contract at the given address
	log.Println(">>>> Call GetLog()")
	txHash, txType, txState, txTime, err := instance.GetLog(nil, txId, typeVal)
	if err != nil {
		return "", "", "", err
	}

	// convert unix time from BigInt to Int64
	txTime64 := txTime.Int64()

	// Convert date-time to string
	t := time.Unix(txTime64, 0)
	strDate := t.Format(time.UnixDate)

	// log what we got back from the LogList call
	log.WithField("Sha3", txHash).Info("transaction hash value")
	log.WithField("Type", txType).Info("transaction type")
	log.WithField("State", txState).Info("transaction state")
	log.WithField("Time", strDate).Info("block time")

	// Get type name from txType value
	strType, err := Gettypename(txType)
	if err != nil {
		return "", "", "", err
	}

	// Get state name from txState value given a type value
	strState, err := Getstatename(txType, txState)
	if err != nil {
		return "", "", "", err
	}

	return txHash, strType, strState, nil
}

// Gethash function to just retrieve the sha3 value of an order or trade with id
func Gethash(smilo bool, id string, ttype string, addr string, apikey string) (string, error) {

	// Check if API key matches the one in our environment
	if ZKPKEY != apikey {
		// if no API key found than bail out
		log.Error("no legal test zkp api key used")
		return "", errors.New("no legal test zkp api key used")
	}

	// We only continue with a valid API key

	// Get instance to contract
	instance, txId, err := getInstance(smilo, id, addr)
	if err != nil {
		return "", err
	}

	typeVal, err := GetTypeVal(ttype)
	if err != nil {
		return "", err
	}

	// call the contract at the given address
	log.Println(">>>> Call GetHash()")
	txHash, err := instance.GetHash(nil, txId, typeVal)
	if err != nil {
		return "", err
	}

	// log what we got back from the LogList call
	log.WithField("Sha3", txHash).Info("transaction hash value")

	return txHash, nil
}

// Gethistory to retrieve the intrinsic state log of an order or trade with id
func Gethistory(smilo bool, id string, ttype string, addr string, apikey string) (string, string, error) {

	// Check if API key matches the one in our environment
	if ZKPKEY != apikey {
		// if no API key found than bail out
		log.Error("no legal test zkp api key used")
		return "", "", errors.New("no legal test zkp api key used")
	}

	// We only continue with a valid API key

	// Get instance to contract
	instance, txId, err := getInstance(smilo, id, addr)
	if err != nil {
		return "", "", err
	}

	typeVal, err := GetTypeVal(ttype)
	if err != nil {
		return "", "", err
	}

	// call the contract at the given address
	log.Println(">>>> Call GetHistoricalStates()")
	txHistStates, err := instance.GetHistoricStates(nil, txId, typeVal)
	if err != nil {
		return "", "", err
	}

	// initialize JSON string
	txJSON := `[`

	// some vars needed in the process
	txType := uint8(0)
	var txState string

	// reformat array in JSON array string
	for i, stateval := range txHistStates {

		// Write statevalue to log
		log.WithField("StateValue", i).Info(stateval)

		if txType == 0 {
			txType = stateval.Tt
		} else {
			if txType != stateval.Tt {
				log.Error("inconsistent log")
				return "", "", errors.New("inconsistent log")
			}
			txJSON = txJSON + `,`
		}

		txJSON = txJSON + `{"state":"`
		txState, err = Getstatename(stateval.Tt, stateval.St)
		if err != nil {
			return "", "", err
		}
		txJSON = txJSON + txState + `",{"time":"`

		uxTime := time.Unix(stateval.Timestamp.Int64(), 0)
		txTime := uxTime.Format(time.RFC3339)

		txJSON = txJSON + txTime + `"}`
	}

	txJSON = txJSON + `]`

	strType, err := Gettypename(txType)
	if err != nil {
		return "", "", nil
	}

	log.WithField("type", strType).Info("Testing the type of this log item.")
	log.WithField("json states", txJSON).Info("Testing the JSON states array")

	return strType, txJSON, nil
}

// CalcFilehash calculate hash for text
func CalcFilehash(filename string, hashtype string, apikey string) (string, error) {

	// Check if API key matches the one in our environment
	if ZKPKEY != apikey {
		// if no API key found than bail out
		log.Error("no legal test zkp api key used")
		return "", errors.New("no legal test zkp api key used")
	}

	// We only continue with a valid API key

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
	hashvalue, err := CalcFileByteshash(filebytes, hashtype, apikey)
	if err != nil {
		return "", err
	}

	return hashvalue, nil
}

// CalcFileByteshash calculate hash for text
func CalcFileByteshash(filebytes []byte, hashtype string, apikey string) (string, error) {

	// Check if API key matches the one in our environment
	if ZKPKEY != apikey {
		// if no API key found than bail out
		log.Error("no legal test zkp api key used")
		return "", errors.New("no legal test zkp api key used")
	}

	// We only continue with a valid API key

	// Calculte hash value in bytes
	bytehash, err := Calchash(filebytes, hashtype)
	if err != nil {
		return "", nil
	}

	// Convert bytes to hash and make that a string
	hashvalue := common.BytesToHash(bytehash).String()

	log.WithField("hash", hashvalue).Info("\"" + hashtype + "\" hash value logged")

	return hashvalue, nil
}

// CalcTexthash calculate hash for text
func CalcTexthash(instr string, hashtype string, apikey string) (string, error) {

	// Check if API key matches the one in our environment
	if ZKPKEY != apikey {
		// if no API key found than bail out
		log.Error("no legal test zkp api key used")
		return "", errors.New("no legal test zkp api key used")
	}

	// We only continue with a valid API key

	// Calculte hash value in bytes
	bytehash, err := Calchash([]byte(instr), hashtype)
	if err != nil {
		return "", nil
	}

	// Convert bytes to hash and make that a string
	hashvalue := common.BytesToHash(bytehash).String()

	log.WithField("hash", hashvalue).Info("\"" + hashtype + "\" hash value logged")

	return hashvalue, nil
}

// Loghash log transaction order, otcorder or trade id, hash, type and state to blockchain
func Loghash(smilo bool, txjsn string, ttype string, state string, addr string, apikey string) (string, string, error) {

	var err error
	var auth *bind.TransactOpts
	var client *ethclient.Client

	log.WithField("JSON", txjsn).Info("Transaction as we received it")
	log.WithField("TYPE", ttype).Info("Type of transaction")
	log.WithField("STATE", state).Info("State of transaction")

	// Check if API key matches the one in our environment
	if ZKPKEY != apikey {
		// if no API key found than bail out
		log.Error("no legal test zkp api key used")
		return "", "", errors.New("no legal test zkp api key used")
	}

	// We only continue with a valid API key

	if smilo {
		// read keys from Smilo HdWallet
		err = readKeyFromHdWallet()
		if err != nil {
			return "", "", err
		}

		// prepare for write and setup authentication for account with deployPKey, let function estimate gas
		auth, client, err = prepareForWrite(smiloPKey, smiloNET, smiloPORT, smilo, 500000)
	} else {
		testPKey = model.GanacheKey
		if testPKey == "" {
			log.Error("could not find ganache key in environment")
			return "", "", errors.New("could not find ganache key in environment")
		}
		// prepare for write and setup authentication for account with deployPKey, let function estimate gas
		auth, client, err = prepareForWrite(testPKey, chainID, chainPORT, smilo, 0)
	}
	if err != nil {
		return "", "", err
	}

	// -- Prepare parameters and call logHash smart-contract

	// Store mocked order Json in struct and print struct values
	var tx ORDERTRADE
	err = json.Unmarshal([]byte(txjsn), &tx)
	if err != nil {
		return "", "", err
	}
	log.
		WithField("tx.Id", tx.Id).
		Debug("will calculate hash")

	s := strings.Split(tx.Id, "/")
	if len(s) != 2 { // "_id must have a '/' and therefore the split should result in s having 2 elements
		return "", "", err
	}

	id := s[1]

	// -- check state so that we avoid calling the log function for same or "TRADED" and "CLOSED" states
	h, gottype, gotstate, err := Getlog(smilo, id, ttype, addr, apikey)
	if err != nil {
		return "", "", err
	}

	// Return with error when we try to log the same state
	if gottype == ttype && h != "" && strings.Compare(gotstate, state) == 0 {
		log.Error("not possible to log the same state")
		return "", "", errors.New("not possible to log the same state")
	}

	// Calculate hash of Json string
	ethHashResult, err := Calchash([]byte(txjsn), "NEWLEGACYKECCAK256")
	if err != nil {
		return "", "", err
	}

	// Convert OderID from sample struct to [32]byte array
	txIdBytes := []byte(id)
	var txId [32]byte
	copy(txId[:], txIdBytes)

	// Convert typename in uint8 value
	transType, err := GetTypeVal(ttype)
	if err != nil {
		return "", "", err
	}

	// order state in uint8 value
	var orderState uint8

	orderState, err = GetStateVal(state, transType)
	if err != nil {
		return "", "", err
	}

	gotstateval, err := GetStateVal(gotstate, transType)
	if err != nil {
		return "", "", err
	}

	// Return with error when we try to log a state before the last logged state
	if orderState < gotstateval {
		log.Error("cannot log a state before the last logged state")
		return "", "", errors.New("cannot log a state before the last logged state")
	}

	//Log the result, and Ethereum compatible Keccak256 sha3 value
	log.WithField("ID 32Bytes val ", "0x"+common.Bytes2Hex(txIdBytes)+"00000000000000000000000000000000000000000000000000000000000000").Info("From ID =" + id)

	hash := common.BytesToHash(ethHashResult).String()
	log.WithField("SHA3", hash).Info("From " + txjsn)

	log.WithField("State", orderState).Info("State = " + state)
	log.WithField("TransactionType", transType).Info("Type = " + ttype)

	log.WithField("address", addr).Info("Contract at address")
	address := common.HexToAddress(addr)
	instance, err := deployer.NewLogger(address, client)
	if err != nil {
		return "", "", err
	}

	// call the contract at the given address
	log.Println(">>>> Call Hashlog()")
	txResult, err := instance.HashLog(auth, txId, common.BytesToHash(ethHashResult).String(), transType, orderState)
	if err != nil {
		return "", "", err
	}

	// log hash of transaction
	log.WithField("hash", txResult.Hash().Hex()).Info("hash of logHash transaction")

	// convert balance to ETH
	txcost := new(big.Float)
	txcost.SetString(txResult.Cost().String())

	ethCosts := new(big.Float).Quo(txcost, big.NewFloat(math.Pow10(18)))

	log.WithField("cost", ethCosts.String()).Info("cost to call logHash function in ETH")

	return hash, id, nil
}

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

func Deploy(smilo bool, apikey string) (string, error) {

	// Check if API key matches the one in our environment
	if ZKPKEY != apikey {
		// if no API key found than bail out
		log.Error("no legal test zkp api key used")
		return "", errors.New("no legal test zkp api key used")
	}

	// We only continue with a valid API key

	// use Smilo HD-wallet if the smilo switch was set
	if smilo {

		err := readKeyFromHdWallet()
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
	}

	// deploy smart contract for audit logging of order- and transaction states
	log.Info(">>>>> Deploying audit logger smartcontract to Ethereum <<<<<")
	_, address, err := deploySolidity(smilo)
	if err != nil {
		return "", err
	}

	return address, nil
}

func execute() error {

	//Store mocked order Json in struct and print struct values
	var tx ORDERTRADE
	err := json.Unmarshal([]byte(trxJson), &tx)
	if err != nil {
		return err
	}
	log.
		WithField("tx.Id", tx.Id).
		Debug("will calculate hash")

	//Calculate hash of Json string
	ethHashResult, err := Calchash([]byte(trxJson), "NEWLEGACYKECCAK256")
	if err != nil {
		return err
	}

	//Log the result, and Ethereum compatible Keccak256 sha3 value
	log.WithField("SHA3", common.BytesToHash(ethHashResult)).Info("From " + trxJson)

	return nil
}

// Calchash log transaction order, otcorder or trade id, hash, type and state to blockchain
func Calchash(bytes []byte, htype string) ([]byte, error) {

	var err error
	var bs []byte

	switch htype {
	case "KECCAK256":
		bs = crypto.Keccak256Hash(bytes).Bytes()
	case "NEWLEGACYKECCAK256":
		h := sha3.NewLegacyKeccak256()
		_, err = h.Write(bytes)
		if err == nil {
			bs = h.Sum(nil)
		}
	case "NEW256":
		h := sha3.New256()
		_, err = h.Write(bytes)
		if err == nil {
			bs = h.Sum(nil)
		}
	default:
		log.Error("non supported hash algorithm")
		err = errors.New("non supported hash algorithm")
	}
	if err != nil {
		return []byte{}, err
	}

	return bs, nil
}
