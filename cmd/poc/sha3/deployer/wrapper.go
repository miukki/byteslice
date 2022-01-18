// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethereum

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

// LoggerTStateLog is an auto generated low-level Go binding around an user-defined struct.
type LoggerTStateLog struct {
	Tt        uint8
	St        uint8
	Timestamp *big.Int
}

// LoggerABI is the input ABI used to generate the binding from.
const LoggerABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_identifier\",\"type\":\"bytes32\"},{\"internalType\":\"enumLogger.TType\",\"name\":\"_tp\",\"type\":\"uint8\"}],\"name\":\"getHash\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_identifier\",\"type\":\"bytes32\"},{\"internalType\":\"enumLogger.TType\",\"name\":\"_tp\",\"type\":\"uint8\"}],\"name\":\"getHistoricStates\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"tt\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"st\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"internalType\":\"structLogger.TStateLog[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_identifier\",\"type\":\"bytes32\"},{\"internalType\":\"enumLogger.TType\",\"name\":\"_tp\",\"type\":\"uint8\"}],\"name\":\"getLog\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"enumLogger.TType\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"hash\",\"type\":\"string\"},{\"internalType\":\"enumLogger.TType\",\"name\":\"tp\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"st\",\"type\":\"uint8\"}],\"name\":\"hashLog\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// LoggerFuncSigs maps the 4-byte function signature to its string representation.
var LoggerFuncSigs = map[string]string{
	"db72ae36": "getHash(bytes32,uint8)",
	"6989b836": "getHistoricStates(bytes32,uint8)",
	"c3b4204c": "getLog(bytes32,uint8)",
	"4aabdfc1": "hashLog(bytes32,string,uint8,uint8)",
}

// LoggerBin is the compiled bytecode used for deploying new contracts.
var LoggerBin = "0x608060405234801561001057600080fd5b50600680546001600160a01b031916331790556115af806100326000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80634aabdfc1146100515780636989b83614610066578063c3b4204c1461008f578063db72ae36146100b2575b600080fd5b61006461005f3660046112b6565b6100d2565b005b61007961007436600461128a565b610a69565b60405161008691906113df565b60405180910390f35b6100a261009d36600461128a565b610c84565b6040516100869493929190611458565b6100c56100c036600461128a565b61105d565b604051610086919061143e565b6006546001600160a01b031633146100e957600080fd5b60008260028111156100fd576100fd611537565b148061011a5750600182600281111561011857610118611537565b145b806101365750600282600281111561013457610134611537565b145b61015b5760405162461bcd60e51b8152600401610152906114a9565b60405180910390fd5b600082600281111561016f5761016f611537565b148015610189575060008481526003602052604090205415155b1561027657600084815260036020526040812080546101aa906001906114d7565b815481106101ba576101ba61154d565b6000918252602091829020604080516060810182526002909302909101805460ff8082168552610100909104811694840185905260019091015491830191909152909250831610156102705760405162461bcd60e51b815260206004820152603960248201527f546865206e6577206f72646572207374617465206d757374206265206869676860448201527832b91037b91039b0b6b2903a3432b710383932bb34b7bab99760391b6064820152608401610152565b506104a6565b600182600281111561028a5761028a611537565b1480156102a4575060008481526004602052604090205415155b1561038f57600084815260046020526040812080546102c5906001906114d7565b815481106102d5576102d561154d565b6000918252602091829020604080516060810182526002909302909101805460ff8082168552610100909104811694840185905260019091015491830191909152909250831610156102705760405162461bcd60e51b815260206004820152603d60248201527f546865206e6577204f54432d6f72646572207374617465206d7573742062652060448201527f686967686572206f722073616d65207468656e2070726576696f75732e0000006064820152608401610152565b60028260028111156103a3576103a3611537565b1480156103bd575060008481526005602052604090205415155b156104a657600084815260056020526040812080546103de906001906114d7565b815481106103ee576103ee61154d565b6000918252602091829020604080516060810182526002909302909101805460ff8082168552610100909104811694840185905260019091015491830191909152909250831610156104a45760405162461bcd60e51b815260206004820152603960248201527f546865206e6577207472616465207374617465206d757374206265206869676860448201527832b91037b91039b0b6b2903a3432b710383932bb34b7bab99760391b6064820152608401610152565b505b4260028360028111156104bb576104bb611537565b14156106b95760088260ff1611156105155760405162461bcd60e51b815260206004820152601b60248201527f5472616465207374617465206973206e6f7420636f72726563742e00000000006044820152606401610152565b60006040518060a0016040528086815260200185600281111561053a5761053a611537565b8152602001600081526020018460ff16600881111561055b5761055b611537565b600881111561056c5761056c611537565b8152602090810184905260008881526002825260409020825180519394508493919261059d928492909101906111cc565b50602082015160018083018054909160ff19909116908360028111156105c5576105c5611537565b0217905550604082015160018201805461ff0019166101008360028111156105ef576105ef611537565b0217905550606082015160018201805462ff000019166201000083600881111561061b5761061b611537565b02179055506080820151816002015590505060056000878152602001908152602001600020604051806060016040528086600281111561065d5761065d611537565b60ff90811682528681166020808401919091526040928301879052845460018181018755600096875295829020855160029092020180549286015184166101000261ffff19909316919093161717815591015191015550610a62565b60018360028111156106cd576106cd611537565b141561086f5760028260ff1611156107275760405162461bcd60e51b815260206004820152601f60248201527f4f54432d6f72646572207374617465206973206e6f7420636f72726563742e006044820152606401610152565b60006040518060a0016040528086815260200185600281111561074c5761074c611537565b81526020018460ff16600281111561076657610766611537565b600281111561077757610777611537565b81526020016000815260209081018490526000888152600182526040902082518051939450849391926107af928492909101906111cc565b50602082015160018083018054909160ff19909116908360028111156107d7576107d7611537565b0217905550604082015160018201805461ff00191661010083600281111561080157610801611537565b0217905550606082015160018201805462ff000019166201000083600881111561082d5761082d611537565b02179055506080820151816002015590505060046000878152602001908152602001600020604051806060016040528086600281111561065d5761065d611537565b60028260ff1611156108c35760405162461bcd60e51b815260206004820152601b60248201527f4f72646572207374617465206973206e6f7420636f72726563742e00000000006044820152606401610152565b60006040518060a001604052808681526020018560028111156108e8576108e8611537565b81526020018460ff16600281111561090257610902611537565b600281111561091357610913611537565b8152602001600081526020908101849052600088815280825260409020825180519394508493919261094a928492909101906111cc565b50602082015160018083018054909160ff199091169083600281111561097257610972611537565b0217905550604082015160018201805461ff00191661010083600281111561099c5761099c611537565b0217905550606082015160018201805462ff00001916620100008360088111156109c8576109c8611537565b021790555060808201518160020155905050600360008781526020019081526020016000206040518060600160405280866002811115610a0a57610a0a611537565b60ff90811682528681166020808401919091526040928301879052845460018181018755600096875295829020855160029092020180549286015184166101000261ffff199093169190931617178155910151910155505b5050505050565b60606000826002811115610a7f57610a7f611537565b1480610a9c57506001826002811115610a9a57610a9a611537565b145b80610ab857506002826002811115610ab657610ab6611537565b145b610ad45760405162461bcd60e51b8152600401610152906114a9565b6000826002811115610ae857610ae8611537565b1415610b7457600083815260036020908152604080832080548251818502810185019093528083529193909284015b82821015610b695760008481526020908190206040805160608101825260028602909201805460ff80821685526101009091041683850152600190810154918301919091529083529092019101610b17565b505050509050610c7e565b6001826002811115610b8857610b88611537565b1415610c0657600083815260046020908152604080832080548251818502810185019093528083529193909284018215610b695760008481526020908190206040805160608101825260028602909201805460ff80821685526101009091041683850152600190810154918301919091529083529092019101610b17565b600083815260056020908152604080832080548251818502810185019093528083529193909284018215610b695760008481526020908190206040805160608101825260028602909201805460ff80821685526101009091041683850152600190810154918301919091529083529092019101610b17565b92915050565b60606000808080856002811115610c9d57610c9d611537565b1480610cba57506001856002811115610cb857610cb8611537565b145b80610cd657506002856002811115610cd457610cd4611537565b145b610cf25760405162461bcd60e51b8152600401610152906114a9565b6000856002811115610d0657610d06611537565b1415610e9957600086815260208190526040808220815160a08101909252805482908290610d33906114fc565b80601f0160208091040260200160405190810160405280929190818152602001828054610d5f906114fc565b8015610dac5780601f10610d8157610100808354040283529160200191610dac565b820191906000526020600020905b815481529060010190602001808311610d8f57829003601f168201915b5050509183525050600182015460209091019060ff166002811115610dd357610dd3611537565b6002811115610de457610de4611537565b81526020016001820160019054906101000a900460ff166002811115610e0c57610e0c611537565b6002811115610e1d57610e1d611537565b81526020016001820160029054906101000a900460ff166008811115610e4557610e45611537565b6008811115610e5657610e56611537565b815260200160028201548152505090508060000151816020015182604001516002811115610e8657610e86611537565b8360800151945094509450945050611054565b6001856002811115610ead57610ead611537565b1415610eda57600086815260016020526040808220815160a08101909252805482908290610d33906114fc565b600086815260026020526040808220815160a08101909252805482908290610f01906114fc565b80601f0160208091040260200160405190810160405280929190818152602001828054610f2d906114fc565b8015610f7a5780601f10610f4f57610100808354040283529160200191610f7a565b820191906000526020600020905b815481529060010190602001808311610f5d57829003601f168201915b5050509183525050600182015460209091019060ff166002811115610fa157610fa1611537565b6002811115610fb257610fb2611537565b81526020016001820160019054906101000a900460ff166002811115610fda57610fda611537565b6002811115610feb57610feb611537565b81526020016001820160029054906101000a900460ff16600881111561101357611013611537565b600881111561102457611024611537565b815260200160028201548152505090508060000151816020015182606001516008811115610e8657610e86611537565b92959194509250565b6060600082600281111561107357611073611537565b14806110905750600182600281111561108e5761108e611537565b145b806110ac575060028260028111156110aa576110aa611537565b145b6110c85760405162461bcd60e51b8152600401610152906114a9565b60008260028111156110dc576110dc611537565b141561118057600083815260208190526040902080546110fb906114fc565b80601f0160208091040260200160405190810160405280929190818152602001828054611127906114fc565b80156111745780601f1061114957610100808354040283529160200191611174565b820191906000526020600020905b81548152906001019060200180831161115757829003601f168201915b50505050509050610c7e565b600182600281111561119457611194611537565b14156111b357600083815260016020526040902080546110fb906114fc565b600083815260026020526040902080546110fb906114fc565b8280546111d8906114fc565b90600052602060002090601f0160209004810192826111fa5760008555611240565b82601f1061121357805160ff1916838001178555611240565b82800160010185558215611240579182015b82811115611240578251825591602001919060010190611225565b5061124c929150611250565b5090565b5b8082111561124c5760008155600101611251565b80356003811061127457600080fd5b919050565b803560ff8116811461127457600080fd5b6000806040838503121561129d57600080fd5b823591506112ad60208401611265565b90509250929050565b600080600080608085870312156112cc57600080fd5b84359350602085013567ffffffffffffffff808211156112eb57600080fd5b818701915087601f8301126112ff57600080fd5b81358181111561131157611311611563565b604051601f8201601f19908116603f0116810190838211818310171561133957611339611563565b816040528281528a602084870101111561135257600080fd5b82602086016020830137600060208483010152809750505050505061137960408601611265565b915061138760608601611279565b905092959194509250565b6000815180845260005b818110156113b85760208185018101518683018201520161139c565b818111156113ca576000602083870101525b50601f01601f19169290920160200192915050565b602080825282518282018190526000919060409081850190868401855b82811015611431578151805160ff908116865287820151168786015285015185850152606090930192908501906001016113fc565b5091979650505050505050565b6020815260006114516020830184611392565b9392505050565b60808152600061146b6080830187611392565b90506003851061148b57634e487b7160e01b600052602160045260246000fd5b84602083015260ff8416604083015282606083015295945050505050565b6020808252601490820152732a3cb8329034b9903737ba1031b7b93932b1ba1760611b604082015260600190565b6000828210156114f757634e487b7160e01b600052601160045260246000fd5b500390565b600181811c9082168061151057607f821691505b6020821081141561153157634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052602160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052604160045260246000fdfea26469706673582212206a3c50f3a8af65029cd1b7abbbc7b39305cb4e4457581e68e18f98c024fdb70e64736f6c63430008060033"

// DeployLogger deploys a new Ethereum contract, binding an instance of Logger to it.
func DeployLogger(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Logger, error) {
	parsed, err := abi.JSON(strings.NewReader(LoggerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(LoggerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Logger{LoggerCaller: LoggerCaller{contract: contract}, LoggerTransactor: LoggerTransactor{contract: contract}, LoggerFilterer: LoggerFilterer{contract: contract}}, nil
}

// Logger is an auto generated Go binding around an Ethereum contract.
type Logger struct {
	LoggerCaller     // Read-only binding to the contract
	LoggerTransactor // Write-only binding to the contract
	LoggerFilterer   // Log filterer for contract events
}

// LoggerCaller is an auto generated read-only Go binding around an Ethereum contract.
type LoggerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LoggerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LoggerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LoggerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LoggerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LoggerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LoggerSession struct {
	Contract     *Logger           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LoggerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LoggerCallerSession struct {
	Contract *LoggerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// LoggerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LoggerTransactorSession struct {
	Contract     *LoggerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LoggerRaw is an auto generated low-level Go binding around an Ethereum contract.
type LoggerRaw struct {
	Contract *Logger // Generic contract binding to access the raw methods on
}

// LoggerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LoggerCallerRaw struct {
	Contract *LoggerCaller // Generic read-only contract binding to access the raw methods on
}

// LoggerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LoggerTransactorRaw struct {
	Contract *LoggerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLogger creates a new instance of Logger, bound to a specific deployed contract.
func NewLogger(address common.Address, backend bind.ContractBackend) (*Logger, error) {
	contract, err := bindLogger(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Logger{LoggerCaller: LoggerCaller{contract: contract}, LoggerTransactor: LoggerTransactor{contract: contract}, LoggerFilterer: LoggerFilterer{contract: contract}}, nil
}

// NewLoggerCaller creates a new read-only instance of Logger, bound to a specific deployed contract.
func NewLoggerCaller(address common.Address, caller bind.ContractCaller) (*LoggerCaller, error) {
	contract, err := bindLogger(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LoggerCaller{contract: contract}, nil
}

// NewLoggerTransactor creates a new write-only instance of Logger, bound to a specific deployed contract.
func NewLoggerTransactor(address common.Address, transactor bind.ContractTransactor) (*LoggerTransactor, error) {
	contract, err := bindLogger(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LoggerTransactor{contract: contract}, nil
}

// NewLoggerFilterer creates a new log filterer instance of Logger, bound to a specific deployed contract.
func NewLoggerFilterer(address common.Address, filterer bind.ContractFilterer) (*LoggerFilterer, error) {
	contract, err := bindLogger(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LoggerFilterer{contract: contract}, nil
}

// bindLogger binds a generic wrapper to an already deployed contract.
func bindLogger(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LoggerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Logger *LoggerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Logger.Contract.LoggerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Logger *LoggerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Logger.Contract.LoggerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Logger *LoggerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Logger.Contract.LoggerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Logger *LoggerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Logger.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Logger *LoggerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Logger.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Logger *LoggerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Logger.Contract.contract.Transact(opts, method, params...)
}

// GetHash is a free data retrieval call binding the contract method 0xdb72ae36.
//
// Solidity: function getHash(bytes32 _identifier, uint8 _tp) view returns(string)
func (_Logger *LoggerCaller) GetHash(opts *bind.CallOpts, _identifier [32]byte, _tp uint8) (string, error) {
	var out []interface{}
	err := _Logger.contract.Call(opts, &out, "getHash", _identifier, _tp)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetHash is a free data retrieval call binding the contract method 0xdb72ae36.
//
// Solidity: function getHash(bytes32 _identifier, uint8 _tp) view returns(string)
func (_Logger *LoggerSession) GetHash(_identifier [32]byte, _tp uint8) (string, error) {
	return _Logger.Contract.GetHash(&_Logger.CallOpts, _identifier, _tp)
}

// GetHash is a free data retrieval call binding the contract method 0xdb72ae36.
//
// Solidity: function getHash(bytes32 _identifier, uint8 _tp) view returns(string)
func (_Logger *LoggerCallerSession) GetHash(_identifier [32]byte, _tp uint8) (string, error) {
	return _Logger.Contract.GetHash(&_Logger.CallOpts, _identifier, _tp)
}

// GetHistoricStates is a free data retrieval call binding the contract method 0x6989b836.
//
// Solidity: function getHistoricStates(bytes32 _identifier, uint8 _tp) view returns((uint8,uint8,uint256)[])
func (_Logger *LoggerCaller) GetHistoricStates(opts *bind.CallOpts, _identifier [32]byte, _tp uint8) ([]LoggerTStateLog, error) {
	var out []interface{}
	err := _Logger.contract.Call(opts, &out, "getHistoricStates", _identifier, _tp)

	if err != nil {
		return *new([]LoggerTStateLog), err
	}

	out0 := *abi.ConvertType(out[0], new([]LoggerTStateLog)).(*[]LoggerTStateLog)

	return out0, err

}

// GetHistoricStates is a free data retrieval call binding the contract method 0x6989b836.
//
// Solidity: function getHistoricStates(bytes32 _identifier, uint8 _tp) view returns((uint8,uint8,uint256)[])
func (_Logger *LoggerSession) GetHistoricStates(_identifier [32]byte, _tp uint8) ([]LoggerTStateLog, error) {
	return _Logger.Contract.GetHistoricStates(&_Logger.CallOpts, _identifier, _tp)
}

// GetHistoricStates is a free data retrieval call binding the contract method 0x6989b836.
//
// Solidity: function getHistoricStates(bytes32 _identifier, uint8 _tp) view returns((uint8,uint8,uint256)[])
func (_Logger *LoggerCallerSession) GetHistoricStates(_identifier [32]byte, _tp uint8) ([]LoggerTStateLog, error) {
	return _Logger.Contract.GetHistoricStates(&_Logger.CallOpts, _identifier, _tp)
}

// GetLog is a free data retrieval call binding the contract method 0xc3b4204c.
//
// Solidity: function getLog(bytes32 _identifier, uint8 _tp) view returns(string, uint8, uint8, uint256)
func (_Logger *LoggerCaller) GetLog(opts *bind.CallOpts, _identifier [32]byte, _tp uint8) (string, uint8, uint8, *big.Int, error) {
	var out []interface{}
	err := _Logger.contract.Call(opts, &out, "getLog", _identifier, _tp)

	if err != nil {
		return *new(string), *new(uint8), *new(uint8), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(uint8)).(*uint8)
	out2 := *abi.ConvertType(out[2], new(uint8)).(*uint8)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, err

}

// GetLog is a free data retrieval call binding the contract method 0xc3b4204c.
//
// Solidity: function getLog(bytes32 _identifier, uint8 _tp) view returns(string, uint8, uint8, uint256)
func (_Logger *LoggerSession) GetLog(_identifier [32]byte, _tp uint8) (string, uint8, uint8, *big.Int, error) {
	return _Logger.Contract.GetLog(&_Logger.CallOpts, _identifier, _tp)
}

// GetLog is a free data retrieval call binding the contract method 0xc3b4204c.
//
// Solidity: function getLog(bytes32 _identifier, uint8 _tp) view returns(string, uint8, uint8, uint256)
func (_Logger *LoggerCallerSession) GetLog(_identifier [32]byte, _tp uint8) (string, uint8, uint8, *big.Int, error) {
	return _Logger.Contract.GetLog(&_Logger.CallOpts, _identifier, _tp)
}

// HashLog is a paid mutator transaction binding the contract method 0x4aabdfc1.
//
// Solidity: function hashLog(bytes32 id, string hash, uint8 tp, uint8 st) returns()
func (_Logger *LoggerTransactor) HashLog(opts *bind.TransactOpts, id [32]byte, hash string, tp uint8, st uint8) (*types.Transaction, error) {
	return _Logger.contract.Transact(opts, "hashLog", id, hash, tp, st)
}

// HashLog is a paid mutator transaction binding the contract method 0x4aabdfc1.
//
// Solidity: function hashLog(bytes32 id, string hash, uint8 tp, uint8 st) returns()
func (_Logger *LoggerSession) HashLog(id [32]byte, hash string, tp uint8, st uint8) (*types.Transaction, error) {
	return _Logger.Contract.HashLog(&_Logger.TransactOpts, id, hash, tp, st)
}

// HashLog is a paid mutator transaction binding the contract method 0x4aabdfc1.
//
// Solidity: function hashLog(bytes32 id, string hash, uint8 tp, uint8 st) returns()
func (_Logger *LoggerTransactorSession) HashLog(id [32]byte, hash string, tp uint8, st uint8) (*types.Transaction, error) {
	return _Logger.Contract.HashLog(&_Logger.TransactOpts, id, hash, tp, st)
}
