# ZKP Gnark

# Objective ot this POC
1. Evaluate the mechanics of ZKP using the gnark kit in a local golang context
    * __Business flow tested: Proof to be a recipient of an OTC trade__
    * __Setup and release 1 constraint system__
    * __Create public- and private witness__
    * __Run function in golang to proof the recipient is part of a given OTC identified by transactionID__
2. Demonstrate implementation and deployment of an Ethereum smart contract to verify on-chain
    * __generate verify function and deploy to chain__
    * __setup proof using private key and private witness__
    * __convert proof and public witness (hash of member slice) to solidity data-types__
    * __call the on-chain version of verify function__
3. Rebuild POC in Dusk/Plonk
4. (optional) test other networks/alternatives.
    * this starts with the identification of suitable alternatives given the experiences with gnark and dusk.

# How to:

## Circuit POC
```
go run cmd/poc/main.go circuit
```


# ABI generate 
```
solc --abi ./registry-contract/contracts/Verifier.sol 
```

# Demo Verifier.sol
## 1. Setup your go envs:
export GO111MODULE=on
export GOPATH=
export GOROOT=
export GOBIN=

## 2. Run:
```
go mod download
go mod tidy
```

## 3. Run byteslice to generate Contract:

```
make byteslice_init
make byteslice_chain
```

## 4. Contract should be regenerated (check git status):
```
➜  zkp git:(master) ✗ git status
        modified:   cmd/poc/byteslice/circuit/wrapper.go
        modified:   registry-contract/contracts/Verifier.sol

```

## 4. Prepare frontend environement:
### 4.1 Run steps for truffle
```
cd ./registry-contract
truffle compile

cd ./registry-contract 
truffle migrate

```

## 5. Run tests with truffle [optional]

```
#tests located in registry-contract/test/
cd ./registry-contract 
truffle test

```


## 6. Run React Client
```
yarn --cwd client install
yarn --cwd client start

```


