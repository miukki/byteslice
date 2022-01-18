// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package circuit

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// PairingABI is the input ABI used to generate the binding from.
const PairingABI = "[]"

// PairingBin is the compiled bytecode used for deploying new contracts.
var PairingBin = "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220a0432c0762809bc1e24590ff97852fc22aa4476d6d519fc34b3b04466d6e889a64736f6c63430008040033"

// DeployPairing deploys a new Ethereum contract, binding an instance of Pairing to it.
func DeployPairing(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Pairing, error) {
	parsed, err := abi.JSON(strings.NewReader(PairingABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PairingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Pairing{PairingCaller: PairingCaller{contract: contract}, PairingTransactor: PairingTransactor{contract: contract}, PairingFilterer: PairingFilterer{contract: contract}}, nil
}

// Pairing is an auto generated Go binding around an Ethereum contract.
type Pairing struct {
	PairingCaller     // Read-only binding to the contract
	PairingTransactor // Write-only binding to the contract
	PairingFilterer   // Log filterer for contract events
}

// PairingCaller is an auto generated read-only Go binding around an Ethereum contract.
type PairingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PairingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PairingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PairingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PairingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PairingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PairingSession struct {
	Contract     *Pairing          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PairingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PairingCallerSession struct {
	Contract *PairingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// PairingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PairingTransactorSession struct {
	Contract     *PairingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// PairingRaw is an auto generated low-level Go binding around an Ethereum contract.
type PairingRaw struct {
	Contract *Pairing // Generic contract binding to access the raw methods on
}

// PairingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PairingCallerRaw struct {
	Contract *PairingCaller // Generic read-only contract binding to access the raw methods on
}

// PairingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PairingTransactorRaw struct {
	Contract *PairingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPairing creates a new instance of Pairing, bound to a specific deployed contract.
func NewPairing(address common.Address, backend bind.ContractBackend) (*Pairing, error) {
	contract, err := bindPairing(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pairing{PairingCaller: PairingCaller{contract: contract}, PairingTransactor: PairingTransactor{contract: contract}, PairingFilterer: PairingFilterer{contract: contract}}, nil
}

// NewPairingCaller creates a new read-only instance of Pairing, bound to a specific deployed contract.
func NewPairingCaller(address common.Address, caller bind.ContractCaller) (*PairingCaller, error) {
	contract, err := bindPairing(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PairingCaller{contract: contract}, nil
}

// NewPairingTransactor creates a new write-only instance of Pairing, bound to a specific deployed contract.
func NewPairingTransactor(address common.Address, transactor bind.ContractTransactor) (*PairingTransactor, error) {
	contract, err := bindPairing(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PairingTransactor{contract: contract}, nil
}

// NewPairingFilterer creates a new log filterer instance of Pairing, bound to a specific deployed contract.
func NewPairingFilterer(address common.Address, filterer bind.ContractFilterer) (*PairingFilterer, error) {
	contract, err := bindPairing(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PairingFilterer{contract: contract}, nil
}

// bindPairing binds a generic wrapper to an already deployed contract.
func bindPairing(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PairingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pairing *PairingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pairing.Contract.PairingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pairing *PairingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pairing.Contract.PairingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pairing *PairingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pairing.Contract.PairingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pairing *PairingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pairing.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pairing *PairingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pairing.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pairing *PairingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pairing.Contract.contract.Transact(opts, method, params...)
}

// VerifierABI is the input ABI used to generate the binding from.
const VerifierABI = "[{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[1]\",\"name\":\"input\",\"type\":\"uint256[1]\"}],\"name\":\"verifyProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"r\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// VerifierFuncSigs maps the 4-byte function signature to its string representation.
var VerifierFuncSigs = map[string]string{
	"43753b4d": "verifyProof(uint256[2],uint256[2][2],uint256[2],uint256[1])",
}

// VerifierBin is the compiled bytecode used for deploying new contracts.
var VerifierBin = "0x608060405234801561001057600080fd5b50611219806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c806343753b4d14610030575b600080fd5b61004361003e366004610fd2565b610057565b604051901515815260200160405180910390f35b6000610061610e0a565b6040805180820182528751815260208089015181830152908352815160808101835287515181840190815288518301516060830152815282518084018452888301805151825251830151818401528183015283820152815180830183528651815286820151918101919091529082015260006100db610586565b6040805180820190915260008082526020820152835151919250906000805160206111c4833981519152116101575760405162461bcd60e51b815260206004820152601760248201527f76657269666965722d61582d6774652d7072696d652d7100000000000000000060448201526064015b60405180910390fd5b8251602001516000805160206111c4833981519152116101b95760405162461bcd60e51b815260206004820152601760248201527f76657269666965722d61592d6774652d7072696d652d71000000000000000000604482015260640161014e565b602083015151516000805160206111c48339815191521161021c5760405162461bcd60e51b815260206004820152601860248201527f76657269666965722d6258302d6774652d7072696d652d710000000000000000604482015260640161014e565b6020838101510151516000805160206111c4833981519152116102815760405162461bcd60e51b815260206004820152601860248201527f76657269666965722d6259302d6774652d7072696d652d710000000000000000604482015260640161014e565b6020838101515101516000805160206111c4833981519152116102e65760405162461bcd60e51b815260206004820152601860248201527f76657269666965722d6258312d6774652d7072696d652d710000000000000000604482015260640161014e565b60208381015181015101516000805160206111c48339815191521161034d5760405162461bcd60e51b815260206004820152601860248201527f76657269666965722d6259312d6774652d7072696d652d710000000000000000604482015260640161014e565b6040830151516000805160206111c4833981519152116103af5760405162461bcd60e51b815260206004820152601760248201527f76657269666965722d63582d6774652d7072696d652d71000000000000000000604482015260640161014e565b6000805160206111c4833981519152836040015160200151106104145760405162461bcd60e51b815260206004820152601760248201527f76657269666965722d63592d6774652d7072696d652d71000000000000000000604482015260640161014e565b60005b6001811015610532577f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000186826001811061046157634e487b7160e01b600052603260045260246000fd5b6020020151106104b35760405162461bcd60e51b815260206004820152601f60248201527f76657269666965722d6774652d736e61726b2d7363616c61722d6669656c6400604482015260640161014e565b61051e8261051985608001518460016104cc9190611124565b600281106104ea57634e487b7160e01b600052603260045260246000fd5b602002015189856001811061050f57634e487b7160e01b600052603260045260246000fd5b60200201516108a6565b610942565b91508061052a81611172565b915050610417565b50608082015151610544908290610942565b905061057a61055684600001516109da565b84602001518460000151856020015185876040015189604001518960600151610a70565b98975050505050505050565b61058e610e5b565b6040805180820182527f127b8e496ad01562fc353c1f3cae13b2d7a75923ec43530f7ed8fcd9adc0771c81527f2c7e476a0cbb1d5fda67e61e9008b2cb6ea5486c55f096f0120802f98b86691b6020808301919091529083528151608080820184527f291cd73b89af987f6c337b19783e4dd0b209f2b55e65a6c77e408c162a7f876a8285019081527f0e42898e90db971b92a4822378461e23ed29c289520b3983906856f877306800606080850191909152908352845180860186527f20a948113bde7521bcb9ecea89069cfe875599a2cc77352a3fb97b779f016fba81527f2ecb61af163cdd828513fc42f194d5fd94051034350971dbbc9f234ad3c6ea74818601528385015285840192909252835180820185527f080b010e09ea1a29a9b5edd9e630047567d1bff4206660cc59f0fb69a67e9ff58186019081527f1eccf9311a44dee9fa9f2d17138882da02ab147dd3e806db602320208d5bcef3828501528152845180860186527f0611ed8c9714afd3a772e467b627be49cca0a868714ab561e0aa13d3b482906e81527f034ffe7bb937245501d8ec744b10b7ae698577dbf0f09025cda057299ea9f073818601528185015285850152835180820185527f2ebc37963436c2ae3d895c5527d5f01198f4a784c373f5eda4443eab013359518186019081527f0d4ce86316a5c6fc10eadc5c0761f7ea3cfe8180ee1a8271bddd15c3f7489154828501528152845180860186527f29ad2318394104f3114c07d89aaadfe789c7f6d02b274db265149ca5808fd09e81527edddbd673fcf58c81879d377a0dc513cdc43cc8c03af560b5134d25549ba72b818601528185015291850191909152825180840184527f1086507bf6fd2a8a06f05b74a98388a8b647afc6b3eb7bd3160618ad48bd398281527f1e1082c54fd3b8e7367cf2e788e820ae8168e810d060ebd5dff9cefcc235dc0b8184015290840180519190915282518084019093527f28422f37902dc6ebc8e3f37ab5abab54599cf5fd059769e8b55f35be779c0b7583527f1be67a1497cd27c9a571af8b7761fb1939657615ead8ceac54bf26aa700e34058383015251015290565b60408051808201909152600080825260208201526108c2610eac565b835181526020808501519082015260408101839052600060608360808460076107d05a03fa90508080156108f5576108f7565bfe5b508061093a5760405162461bcd60e51b81526020600482015260126024820152711c185a5c9a5b99cb5b5d5b0b59985a5b195960721b604482015260640161014e565b505092915050565b604080518082019091526000808252602082015261095e610eca565b8351815260208085015181830152835160408301528301516060808301919091526000908360c08460066107d05a03fa90508080156108f557508061093a5760405162461bcd60e51b81526020600482015260126024820152711c185a5c9a5b99cb5859190b59985a5b195960721b604482015260640161014e565b604080518082019091526000808252602082015281511580156109ff57506020820151155b15610a1d575050604080518082019091526000808252602082015290565b6040518060400160405280836000015181526020016000805160206111c48339815191528460200151610a50919061118d565b610a68906000805160206111c483398151915261115b565b905292915050565b60408051608080820183528a825260208083018a90528284018890526060808401879052845192830185528b83528282018a9052828501889052820185905283516018808252610320820190955260009491859190839082016103008036833701905050905060005b6004811015610d88576000610aef82600661113c565b9050858260048110610b1157634e487b7160e01b600052603260045260246000fd5b60200201515183610b23836000611124565b81518110610b4157634e487b7160e01b600052603260045260246000fd5b602002602001018181525050858260048110610b6d57634e487b7160e01b600052603260045260246000fd5b60200201516020015183826001610b849190611124565b81518110610ba257634e487b7160e01b600052603260045260246000fd5b602002602001018181525050848260048110610bce57634e487b7160e01b600052603260045260246000fd5b6020020151515183610be1836002611124565b81518110610bff57634e487b7160e01b600052603260045260246000fd5b602002602001018181525050848260048110610c2b57634e487b7160e01b600052603260045260246000fd5b6020020151516001602002015183610c44836003611124565b81518110610c6257634e487b7160e01b600052603260045260246000fd5b602002602001018181525050848260048110610c8e57634e487b7160e01b600052603260045260246000fd5b602002015160200151600060028110610cb757634e487b7160e01b600052603260045260246000fd5b602002015183610cc8836004611124565b81518110610ce657634e487b7160e01b600052603260045260246000fd5b602002602001018181525050848260048110610d1257634e487b7160e01b600052603260045260246000fd5b602002015160200151600160028110610d3b57634e487b7160e01b600052603260045260246000fd5b602002015183610d4c836005611124565b81518110610d6a57634e487b7160e01b600052603260045260246000fd5b60209081029190910101525080610d8081611172565b915050610ad9565b50610d91610ee8565b6000602082602086026020860160086107d05a03fa90508080156108f5575080610df55760405162461bcd60e51b81526020600482015260156024820152741c185a5c9a5b99cb5bdc18dbd9194b59985a5b1959605a1b604482015260640161014e565b505115159d9c50505050505050505050505050565b6040805160a081019091526000606082018181526080830191909152815260208101610e34610f06565b8152602001610e56604051806040016040528060008152602001600081525090565b905290565b6040805160e08101909152600060a0820181815260c0830191909152815260208101610e85610f06565b8152602001610e92610f06565b8152602001610e9f610f06565b8152602001610e56610f26565b60405180606001604052806003906020820280368337509192915050565b60405180608001604052806004906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b6040518060400160405280610f19610f5f565b8152602001610e56610f5f565b60405180604001604052806002905b6040805180820190915260008082526020820152815260200190600190039081610f355790505090565b60405180604001604052806002906020820280368337509192915050565b600082601f830112610f8d578081fd5b610f956110bc565b808385604086011115610fa6578384fd5b835b6002811015610fc7578135845260209384019390910190600101610fa8565b509095945050505050565b600080600080610120808688031215610fe9578485fd5b610ff38787610f7d565b9450604087605f880112611005578485fd5b61100d6110bc565b8082890160c08a018b811115611021578889fd5b885b600281101561104a576110368d84610f7d565b855260209094019391850191600101611023565b508298506110588c82610f7d565b975050505050508661011f87011261106e578182fd5b6110766110f3565b80610100880189848a01111561108a578485fd5b8493505b60018410156110ae5780358352600193909301926020928301920161108e565b509598949750929550505050565b6040805190810167ffffffffffffffff811182821017156110ed57634e487b7160e01b600052604160045260246000fd5b60405290565b6040516020810167ffffffffffffffff811182821017156110ed57634e487b7160e01b600052604160045260246000fd5b60008219821115611137576111376111ad565b500190565b6000816000190483118215151615611156576111566111ad565b500290565b60008282101561116d5761116d6111ad565b500390565b6000600019821415611186576111866111ad565b5060010190565b6000826111a857634e487b7160e01b81526012600452602481fd5b500690565b634e487b7160e01b600052601160045260246000fdfe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47a2646970667358221220bbf8656ae82b154f6d721e232945e0ec898ade9a9a413a96bcff3cbe5a48354d64736f6c63430008040033"

// DeployVerifier deploys a new Ethereum contract, binding an instance of Verifier to it.
func DeployVerifier(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Verifier, error) {
	parsed, err := abi.JSON(strings.NewReader(VerifierABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VerifierBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Verifier{VerifierCaller: VerifierCaller{contract: contract}, VerifierTransactor: VerifierTransactor{contract: contract}, VerifierFilterer: VerifierFilterer{contract: contract}}, nil
}

// Verifier is an auto generated Go binding around an Ethereum contract.
type Verifier struct {
	VerifierCaller     // Read-only binding to the contract
	VerifierTransactor // Write-only binding to the contract
	VerifierFilterer   // Log filterer for contract events
}

// VerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type VerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VerifierSession struct {
	Contract     *Verifier         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VerifierCallerSession struct {
	Contract *VerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// VerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VerifierTransactorSession struct {
	Contract     *VerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// VerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type VerifierRaw struct {
	Contract *Verifier // Generic contract binding to access the raw methods on
}

// VerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VerifierCallerRaw struct {
	Contract *VerifierCaller // Generic read-only contract binding to access the raw methods on
}

// VerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VerifierTransactorRaw struct {
	Contract *VerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVerifier creates a new instance of Verifier, bound to a specific deployed contract.
func NewVerifier(address common.Address, backend bind.ContractBackend) (*Verifier, error) {
	contract, err := bindVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Verifier{VerifierCaller: VerifierCaller{contract: contract}, VerifierTransactor: VerifierTransactor{contract: contract}, VerifierFilterer: VerifierFilterer{contract: contract}}, nil
}

// NewVerifierCaller creates a new read-only instance of Verifier, bound to a specific deployed contract.
func NewVerifierCaller(address common.Address, caller bind.ContractCaller) (*VerifierCaller, error) {
	contract, err := bindVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VerifierCaller{contract: contract}, nil
}

// NewVerifierTransactor creates a new write-only instance of Verifier, bound to a specific deployed contract.
func NewVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*VerifierTransactor, error) {
	contract, err := bindVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VerifierTransactor{contract: contract}, nil
}

// NewVerifierFilterer creates a new log filterer instance of Verifier, bound to a specific deployed contract.
func NewVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*VerifierFilterer, error) {
	contract, err := bindVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VerifierFilterer{contract: contract}, nil
}

// bindVerifier binds a generic wrapper to an already deployed contract.
func bindVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VerifierABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Verifier *VerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Verifier.Contract.VerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Verifier *VerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Verifier.Contract.VerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Verifier *VerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Verifier.Contract.VerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Verifier *VerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Verifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Verifier *VerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Verifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Verifier *VerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Verifier.Contract.contract.Transact(opts, method, params...)
}

// VerifyProof is a free data retrieval call binding the contract method 0x43753b4d.
//
// Solidity: function verifyProof(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[1] input) view returns(bool r)
func (_Verifier *VerifierCaller) VerifyProof(opts *bind.CallOpts, a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input [1]*big.Int) (bool, error) {
	var out []interface{}
	err := _Verifier.contract.Call(opts, &out, "verifyProof", a, b, c, input)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyProof is a free data retrieval call binding the contract method 0x43753b4d.
//
// Solidity: function verifyProof(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[1] input) view returns(bool r)
func (_Verifier *VerifierSession) VerifyProof(a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input [1]*big.Int) (bool, error) {
	return _Verifier.Contract.VerifyProof(&_Verifier.CallOpts, a, b, c, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0x43753b4d.
//
// Solidity: function verifyProof(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[1] input) view returns(bool r)
func (_Verifier *VerifierCallerSession) VerifyProof(a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, input [1]*big.Int) (bool, error) {
	return _Verifier.Contract.VerifyProof(&_Verifier.CallOpts, a, b, c, input)
}
