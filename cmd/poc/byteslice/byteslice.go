package byteslice

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"io"
	"math"
	"math/big"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"test-zkp/cmd/poc/byteslice/circuit"

	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/consensys/gnark-crypto/ecc/bn254/fp"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr"

	log "github.com/sirupsen/logrus"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark-crypto/hash"
	"github.com/consensys/gnark/backend"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"
)

type OtcTXN struct {
	Id         string
	Recipients []string
}

const (
	chainID       = 5777
	chainPORT     = 7545
	testSeedVal   = "TestOTC"
	trxJson       = `{"id": "1234567","recipients": ["1000000","111111","999999"]}`
	deployPKey    = "49af4bf2b88e38698be51535d73daed1c7cb5a5d2439c86a28d78e82dbc18b73" // private key from Ganache
	AbigenVersion = "github.com/ethereum/go-ethereum/cmd/abigen@v1.10.3"
)

var (
	r1csPath    = "byteslice/circuit/mimc.r1cs"
	pkPath      = "byteslice/circuit/mimc.pk"
	vkPath      = "byteslice/circuit/mimc.vk"
	verifierSol = "registry-contract/contracts/Verifier.sol"
	wrapperGo   = "byteslice/circuit/wrapper.go"
	basePath    = "byteslice/circuit/"
	abiBinPath  = "bin/abigen"
)

var workingDir = ""

func init() {
	var err error
	workingDir, err = os.Getwd()
	if err != nil {
		panic(err)
	}
	if !strings.Contains(workingDir, "/cmd") {
		workingDir = workingDir + "/cmd/poc"
	}
	r1csPath = filepath.Join(workingDir, r1csPath)
	pkPath = filepath.Join(workingDir, pkPath)
	vkPath = filepath.Join(workingDir, vkPath)
	verifierSol = filepath.Join(workingDir, "../..", verifierSol)
	wrapperGo = filepath.Join(workingDir, wrapperGo)
	basePath = filepath.Join(workingDir, basePath)
	abiBinPath = filepath.Join(workingDir, "../..", abiBinPath)
}

func Run() error {

	// check that init was performed
	if _, err := os.Stat(r1csPath); os.IsNotExist(err) {
		log.Fatal("please run with --init flag first to serialize circuit, keys and solidity contract")
		return err
	}

	err := execute()

	if err != nil {
		log.WithError(err).Error("Failed to setup ADB data")
		return err
	}

	log.Info("Finished setting up ADB data")
	return nil
}

func Chainverify() error {

	log.Info("Generate proof and verify on-chain")

	// check that init was performed
	if _, err := os.Stat(r1csPath); os.IsNotExist(err) {
		log.Fatal("please run with --init flag first to serialize circuit, keys and solidity contract")
		return err
	}

	// deploy smart contract
	verifierContract, err := deploySolidity()
	if err != nil {
		return err
	}

	// read R1CS, proving key and verifying keys
	r1cs := groth16.NewCS(ecc.BN254)
	pk := groth16.NewProvingKey(ecc.BN254)
	vk := groth16.NewVerifyingKey(ecc.BN254)
	err = deserialize(r1cs, r1csPath)
	if err != nil {
		return err
	}
	err = deserialize(pk, pkPath)
	if err != nil {
		return err
	}
	err = deserialize(vk, vkPath)
	if err != nil {
		return err
	}

	var witness circuit.Circuit

	// set recipient for who we will generate the proof
	//  ["1000000","111111","999999"]
	proofFor := "999999"

	// log balance of the account used to deploy the contract
	log.WithField("recipient", proofFor).Info("generating proof for")

	//frontend-prover:::Create private witness with sum ([]byte of hash) from trxJson. Prover knows secret--A,B,C and public--HASH
	sum, err := createWitness(trxJson, proofFor, &witness)
	if err != nil {
		return err
	}

	//setup proof for backend-verifier
	proof, err := groth16.Prove(r1cs, pk, &witness)
	if err != nil {
		return err
	}

	//backend-verifier:: Create public witness. Verifier knows public--HASH and  generated proof groth16.Proof
	var publicWitness circuit.Circuit
	publicWitness.Hash.Assign(sum)

	err = groth16.Verify(proof, vk, &publicWitness)
	if err != nil {
		return err
	}
	log.Info(`tested Prove/Verify successfully locally, now we can deploy as next step, proof is ready to deploy`)

	// solidity contract inputs
	// a, b and c are the 3 ecc points in the proof we feed to the pairing
	// they are stored in the same order in the golang data structure
	// each coordinate is a field element, of size fp.Bytes bytes
	var (
		a     [2]*big.Int
		b     [2][2]*big.Int
		c     [2]*big.Int
		input [1]*big.Int
	)

	// get proof bytes
	var buf bytes.Buffer
	_, err = proof.WriteRawTo(&buf)
	if err != nil {
		return err
	}
	proofBytes := buf.Bytes()

	// proof.Ar, proof.Bs, proof.Krs
	const fpSize = fp.Bytes
	a[0] = new(big.Int).SetBytes(proofBytes[fpSize*0 : fpSize*1])
	a[1] = new(big.Int).SetBytes(proofBytes[fpSize*1 : fpSize*2])
	b[0][0] = new(big.Int).SetBytes(proofBytes[fpSize*2 : fpSize*3])
	b[0][1] = new(big.Int).SetBytes(proofBytes[fpSize*3 : fpSize*4])
	b[1][0] = new(big.Int).SetBytes(proofBytes[fpSize*4 : fpSize*5])
	b[1][1] = new(big.Int).SetBytes(proofBytes[fpSize*5 : fpSize*6])
	c[0] = new(big.Int).SetBytes(proofBytes[fpSize*6 : fpSize*7])
	c[1] = new(big.Int).SetBytes(proofBytes[fpSize*7 : fpSize*8])

	// (correct) public witness
	input[0] = new(big.Int).SetBytes(sum)

	// (wrong) public witness
	//input[0] = new(big.Int).SetUint64(42)
	// call the contract

	log.Info("proof", a, b, c, input)
	res, err := verifierContract.VerifyProof(nil, a, b, c, input)
	if err != nil {
		return err
	}

	if !res {
		log.Fatal("calling the verifier on chain didn't succeed, but should have")
	}
	log.Info("successfully verified proof on-chain")

	return nil
}

func Debug() error {

	log.Info("Debug byteslice circuit")

	var witness circuit.Circuit

	// compiles our circuit into a R1CS
	log.Info("compiling circuit")
	r1cs, err := frontend.Compile(ecc.BN254, backend.GROTH16, &witness)
	if err != nil {
		return err
	}

	//setup pk, vk as part of algorithms
	log.Info("running groth16.Setup")
	pk, vk, err := groth16.Setup(r1cs)
	if err != nil {
		return err
	}

	// build witness for party with ID="1000000"
	//  ["1000000","111111","999999"]
	proofFor := "999999"

	sum, err := createWitness(trxJson, proofFor, &witness)
	if err != nil {
		return err
	}

	//setup proof
	proof, err := groth16.Prove(r1cs, pk, &witness)
	if err != nil {
		return err
	}

	//backend-verifier
	//Assign HASH=Y by verifier (server)
	var publicWitness circuit.Circuit
	publicWitness.Hash.Assign(sum)

	err = groth16.Verify(proof, vk, &publicWitness)
	if err != nil {
		return err
	}

	log.Info("successfully verified proof off-chain")

	return nil

}

func Init() error {

	log.Info("Initializing circuit and generate solidity")

	var mimcCircuit circuit.Circuit

	// compiles our circuit into a R1CS
	log.Info("compiling circuit")
	r1cs, err := frontend.Compile(ecc.BN254, backend.GROTH16, &mimcCircuit)
	if err != nil {
		return err
	}

	//setup pk, vk as part of algorithms
	log.Info("running groth16.Setup")
	pk, vk, err := groth16.Setup(r1cs)
	if err != nil {
		return err
	}

	// serialize R1CS, proving & verifying key
	log.WithField("r1csPath", r1csPath).Info("serialize R1CS (circuit)")
	err = serialize(r1cs, r1csPath)
	if err != nil {
		return err
	}

	log.WithField("pkPath", pkPath).Info("serialize proving key")
	err = serialize(pk, pkPath)
	if err != nil {
		return err
	}

	log.WithField("vkPath", vkPath).Info("serialize verifying key")
	err = serialize(vk, vkPath)
	if err != nil {
		return err
	}

	// export verifying key to solidity
	log.WithField("verifierSol", verifierSol).Info("export solidity verifier")

	f, err := os.Create(verifierSol)
	if err != nil {
		return err
	}
	err = vk.ExportSolidity(f)
	if err != nil {
		return err
	}

	// run abigen to generate go wrapper
	if _, err = os.Stat(filepath.Join(GOBIN(), "abigen")); os.IsNotExist(err) {
		err = ensureAbigen()
		if err != nil {
			return err
		}
	}

	cmd := exec.Command(abiBinPath, "--sol", verifierSol, "--pkg", "circuit", "--out", wrapperGo)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err = cmd.Run(); err != nil {
		log.Error("Error: Could not execute abigen ... ", "error: ", err, ", cmd: ", cmd)
		return err
	}
	return nil
}

// GOBIN environment variable.
func GOBIN() string {
	if os.Getenv("GOBIN") == "" {
		log.Fatal("GOBIN not set")
	}

	return os.Getenv("GOBIN")
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

// serialize gnark object to given file
func serialize(gnarkObject io.WriterTo, fileName string) error {
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}

	_, err = gnarkObject.WriteTo(f)
	if err != nil {
		return err
	}

	return nil
}

// deserialize gnark object from given file
func deserialize(gnarkObject io.ReaderFrom, fileName string) error {

	f, err := os.Open(filepath.Join(basePath, filepath.Base(fileName)))
	if err != nil {
		return err
	}

	_, err = gnarkObject.ReadFrom(f)
	if err != nil {
		return err
	}

	return nil
}

func deploySolidity() (*circuit.Verifier, error) {

	// build server URL
	server := strings.Join([]string{"http://localhost:", strconv.Itoa(chainPORT)}, "")

	// log URL of RPC server
	log.WithField("URL", server).Info("test URL of RPC server")

	client, err := ethclient.Dial(server)
	if err != nil {
		return nil, err
	}

	privateKey, err := crypto.HexToECDSA(deployPKey)
	if err != nil {
		return nil, err
	}

	// get deployer address from public key
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	deployAddr := crypto.PubkeyToAddress(*publicKeyECDSA)
	account := common.HexToAddress(deployAddr.String())
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		return nil, err
	}

	// convert balance to ETH
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))

	// log balance of the account used to deploy the contract
	log.WithField("balance", ethValue).Info("testing balance of account used to deploy")

	// deploy verifier contract
	log.Println("deploying verifier contract on chain")

	nonce, err := client.PendingNonceAt(context.Background(), account)
	if err != nil {
		return nil, err
	}

	// log nonce
	log.WithField("nonce", nonce).Info("testing the nonce")

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	// setup authentication
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(chainID))
	if err != nil {
		return nil, err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	auth.GasLimit = 3000000    // in units
	auth.GasPrice = gasPrice

	//Use this when deploying to simulatedBackend
	//
	//_, _, verifierContract, err := circuit.DeployVerifier(auth, simulatedBackend)
	address, tx, verifierContract, err := circuit.DeployVerifier(auth, client)
	if err != nil {
		return nil, err
	}

	// log address and transaction hash for deploying this contract
	log.WithField("address", address.Hex()).Info("contract deployed at")
	log.WithField("hash", tx.Hash().Hex()).Info("hash of deploy transaction")

	// convert balance to ETH
	txcost := new(big.Float)
	txcost.SetString(tx.Cost().String())
	ethCosts := new(big.Float).Quo(txcost, big.NewFloat(math.Pow10(18)))

	log.WithField("cost", ethCosts.String()).Info("cost to deploy in ETH")

	//simulatedBackend.Commit()
	return verifierContract, nil
}

// function to build public witness and hash from Json order string
func createWitness(ordstr string, party string, witness *circuit.Circuit) ([]byte, error) {

	//.//set witness
	//.var witness circuit.Circuit

	//Store Json in struct and print struct values
	var order OtcTXN
	err := json.Unmarshal([]byte(ordstr), &order)
	if err != nil {
		return nil, err
	}

	log.
		WithField("tx.Id", order.Id).
		WithField("tx.Recipients", order.Recipients).
		Info("Going to generate witness circuit")

	// stop processing when there are not exactly 3 receipients
	// ToDo: In a next refactoring step we need to redesign to accept 2+ recipients
	if len(order.Recipients) != 3 {

		log.Error("Error: Not exactly 3 recipients found in order ... ", "Receipients: ", order.Recipients)
		return nil, err
	}

	// First sort the array of recipients
	sort.Strings(order.Recipients)

	// Set seed to fixed string value
	fullHashFn := hash.MIMC_BN254.New(testSeedVal)

	// reset hash functions
	fullHashFn.Reset()

	memberFound := false
	var a, b, c, orderID fr.Element

	// write orderID in r1cs as private witness, and add orderSlice to Hash
	orderID.SetBytes([]byte(order.Id))
	witness.OrderID.Assign(orderID)
	x := orderID.Bytes()
	orderSlice := x[:]
	_, err = fullHashFn.Write(orderSlice)
	if err != nil {
		return nil, err
	}

	// Initialize sliceID
	sliceID := 0

	// go through the sorted list of recipients and write private witnesses, this means slices are written in the
	// same order for all combinations that we seek.
	for _, v := range order.Recipients {

		// set sliceID to first slice
		sliceID++

		// if party matches the id found in the list of recipients, indicate that party is one of the recipients
		if strings.Compare(party, v) == 0 {
			memberFound = true
		}

		log.Info("Writing secret witness ByteSliceRecipient" + strconv.Itoa(sliceID) + " for: \"" + v + "\"")
		b.SetBytes([]byte(v))
		if sliceID == 1 {
			witness.ByteSliceRecipient1.Assign(b)

			// now write slice1
			a.SetBytes([]byte(v))
			x1 := a.Bytes()
			slice1 := x1[:]
			_, err = fullHashFn.Write(slice1)
			if err != nil {
				return nil, err
			}

		} else if sliceID == 2 {
			witness.ByteSliceRecipient2.Assign(b)

			// now write slice2
			b.SetBytes([]byte(v))
			x2 := b.Bytes()
			slice2 := x2[:]
			_, err = fullHashFn.Write(slice2)
			if err != nil {
				return nil, err
			}
		} else {
			witness.ByteSliceRecipient3.Assign(b)

			// now write slice3
			c.SetBytes([]byte(v))
			x3 := c.Bytes()
			slice3 := x3[:]
			_, err = fullHashFn.Write(slice3)
			if err != nil {
				return nil, err
			}
		}
	}

	// assign full sum to witness (this is our public witness)
	sum := fullHashFn.Sum(nil)
	witness.Hash.Assign(sum)

	if !memberFound {
		log.Error("Error: Cannot generate proof for non-member of order ... ", "Party: ", party)
		return nil, errors.New("not a member, cannot generate proof")
	}

	return sum, nil
}

func execute() error {

	log.Info("Generate proof and verify off-chain")

	// check that init was performed
	if _, err := os.Stat(r1csPath); os.IsNotExist(err) {
		log.Fatal("please run with --init flag first to serialize circuit, keys and solidity contract")
		return err
	}

	// read R1CS, proving key and verifying keys
	r1cs := groth16.NewCS(ecc.BN254)
	pk := groth16.NewProvingKey(ecc.BN254)
	vk := groth16.NewVerifyingKey(ecc.BN254)
	err := deserialize(r1cs, r1csPath)
	if err != nil {
		return err
	}
	err = deserialize(pk, pkPath)
	if err != nil {
		return err
	}
	err = deserialize(vk, vkPath)
	if err != nil {
		return err
	}

	// Generate public witness and sum ([]byte of hash) from trxJson
	var witness circuit.Circuit

	// set party for who we generate proof
	//  ["1000000","111111","999999"]
	proofFor := "1000000"

	// build witness for party with ID="1000000"
	sum, err := createWitness(trxJson, proofFor, &witness)
	if err != nil {
		return err
	}

	//setup proof
	proof, err := groth16.Prove(r1cs, pk, &witness)
	if err != nil {
		return err
	}

	//backend-verifier
	//Assign HASH=Y by verifier (server)
	var publicWitness circuit.Circuit
	publicWitness.Hash.Assign(sum)

	err = groth16.Verify(proof, vk, &publicWitness)
	if err != nil {
		return err
	}

	log.Info("successfully verified proof off-chain")

	return nil
}
