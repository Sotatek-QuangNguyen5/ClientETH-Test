// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package exchangecoin

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ExchangecoinMetaData contains all meta data concerning the Exchangecoin contract.
var ExchangecoinMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"creator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_messageHash\",\"type\":\"bytes32\"}],\"name\":\"getEthSignedMessageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_message\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"getMessageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"receiverBalance\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_ethSignedMessageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"recoverSigner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"splitSignature\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_message\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"verify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610c7a806100606000396000f3fe60806040526004361061007b5760003560e01c8063a35f8a401161004e578063a35f8a401461011d578063a7bb58031461015a578063abe5026a14610199578063fa540801146101d65761007b565b806302d05d3f146100805780630b5fb5e2146100ab57806312065fe0146100b557806397aba7f9146100e0575b600080fd5b34801561008c57600080fd5b50610095610213565b6040516100a29190610426565b60405180910390f35b6100b3610237565b005b3480156100c157600080fd5b506100ca610239565b6040516100d7919061045a565b60405180910390f35b3480156100ec57600080fd5b5061010760048036038101906101029190610605565b610241565b6040516101149190610426565b60405180910390f35b34801561012957600080fd5b50610144600480360381019061013f919061075a565b6102b0565b60405161015191906107ec565b60405180910390f35b34801561016657600080fd5b50610181600480360381019061017c9190610807565b6102e9565b6040516101909392919061086c565b60405180910390f35b3480156101a557600080fd5b506101c060048036038101906101bb91906108a3565b610351565b6040516101cd9190610983565b60405180910390f35b3480156101e257600080fd5b506101fd60048036038101906101f8919061099e565b6103b5565b60405161020a91906107ec565b60405180910390f35b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b565b600047905090565b600080600080610250856102e9565b9250925092506001868285856040516000815260200160405260405161027994939291906109cb565b6020604051602081039080840390855afa15801561029b573d6000803e3d6000fd5b50505060206040510351935050505092915050565b6000848484846040516020016102c99493929190610aea565b604051602081830303815290604052805190602001209050949350505050565b60008060006041845114610332576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161032990610b91565b60405180910390fd5b6020840151925060408401519150606084015160001a90509193909250565b600080610360878787876102b0565b9050600061036d826103b5565b90508873ffffffffffffffffffffffffffffffffffffffff166103908286610241565b73ffffffffffffffffffffffffffffffffffffffff1614925050509695505050505050565b6000816040516020016103c89190610c1e565b604051602081830303815290604052805190602001209050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610410826103e5565b9050919050565b61042081610405565b82525050565b600060208201905061043b6000830184610417565b92915050565b6000819050919050565b61045481610441565b82525050565b600060208201905061046f600083018461044b565b92915050565b6000604051905090565b600080fd5b600080fd5b6000819050919050565b61049c81610489565b81146104a757600080fd5b50565b6000813590506104b981610493565b92915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610512826104c9565b810181811067ffffffffffffffff82111715610531576105306104da565b5b80604052505050565b6000610544610475565b90506105508282610509565b919050565b600067ffffffffffffffff8211156105705761056f6104da565b5b610579826104c9565b9050602081019050919050565b82818337600083830152505050565b60006105a86105a384610555565b61053a565b9050828152602081018484840111156105c4576105c36104c4565b5b6105cf848285610586565b509392505050565b600082601f8301126105ec576105eb6104bf565b5b81356105fc848260208601610595565b91505092915050565b6000806040838503121561061c5761061b61047f565b5b600061062a858286016104aa565b925050602083013567ffffffffffffffff81111561064b5761064a610484565b5b610657858286016105d7565b9150509250929050565b61066a81610405565b811461067557600080fd5b50565b60008135905061068781610661565b92915050565b61069681610441565b81146106a157600080fd5b50565b6000813590506106b38161068d565b92915050565b600067ffffffffffffffff8211156106d4576106d36104da565b5b6106dd826104c9565b9050602081019050919050565b60006106fd6106f8846106b9565b61053a565b905082815260208101848484011115610719576107186104c4565b5b610724848285610586565b509392505050565b600082601f830112610741576107406104bf565b5b81356107518482602086016106ea565b91505092915050565b600080600080608085870312156107745761077361047f565b5b600061078287828801610678565b9450506020610793878288016106a4565b935050604085013567ffffffffffffffff8111156107b4576107b3610484565b5b6107c08782880161072c565b92505060606107d1878288016106a4565b91505092959194509250565b6107e681610489565b82525050565b600060208201905061080160008301846107dd565b92915050565b60006020828403121561081d5761081c61047f565b5b600082013567ffffffffffffffff81111561083b5761083a610484565b5b610847848285016105d7565b91505092915050565b600060ff82169050919050565b61086681610850565b82525050565b600060608201905061088160008301866107dd565b61088e60208301856107dd565b61089b604083018461085d565b949350505050565b60008060008060008060c087890312156108c0576108bf61047f565b5b60006108ce89828a01610678565b96505060206108df89828a01610678565b95505060406108f089828a016106a4565b945050606087013567ffffffffffffffff81111561091157610910610484565b5b61091d89828a0161072c565b935050608061092e89828a016106a4565b92505060a087013567ffffffffffffffff81111561094f5761094e610484565b5b61095b89828a016105d7565b9150509295509295509295565b60008115159050919050565b61097d81610968565b82525050565b60006020820190506109986000830184610974565b92915050565b6000602082840312156109b4576109b361047f565b5b60006109c2848285016104aa565b91505092915050565b60006080820190506109e060008301876107dd565b6109ed602083018661085d565b6109fa60408301856107dd565b610a0760608301846107dd565b95945050505050565b60008160601b9050919050565b6000610a2882610a10565b9050919050565b6000610a3a82610a1d565b9050919050565b610a52610a4d82610405565b610a2f565b82525050565b6000819050919050565b610a73610a6e82610441565b610a58565b82525050565b600081519050919050565b600081905092915050565b60005b83811015610aad578082015181840152602081019050610a92565b60008484015250505050565b6000610ac482610a79565b610ace8185610a84565b9350610ade818560208601610a8f565b80840191505092915050565b6000610af68287610a41565b601482019150610b068286610a62565b602082019150610b168285610ab9565b9150610b228284610a62565b60208201915081905095945050505050565b600082825260208201905092915050565b7f696e76616c6964207369676e6174757265206c656e6774680000000000000000600082015250565b6000610b7b601883610b34565b9150610b8682610b45565b602082019050919050565b60006020820190508181036000830152610baa81610b6e565b9050919050565b7f6a6f6b6572000000000000000000000000000000000000000000000000000000600082015250565b6000610be7600583610a84565b9150610bf282610bb1565b600582019050919050565b6000819050919050565b610c18610c1382610489565b610bfd565b82525050565b6000610c2982610bda565b9150610c358284610c07565b6020820191508190509291505056fea26469706673582212208fcc9ccb690350c42d0fdc36549923d51b10e4aea607e6bafb4bafd73dcc3bd264736f6c63430008110033",
}

// ExchangecoinABI is the input ABI used to generate the binding from.
// Deprecated: Use ExchangecoinMetaData.ABI instead.
var ExchangecoinABI = ExchangecoinMetaData.ABI

// ExchangecoinBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ExchangecoinMetaData.Bin instead.
var ExchangecoinBin = ExchangecoinMetaData.Bin

// DeployExchangecoin deploys a new Ethereum contract, binding an instance of Exchangecoin to it.
func DeployExchangecoin(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Exchangecoin, error) {
	parsed, err := ExchangecoinMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ExchangecoinBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Exchangecoin{ExchangecoinCaller: ExchangecoinCaller{contract: contract}, ExchangecoinTransactor: ExchangecoinTransactor{contract: contract}, ExchangecoinFilterer: ExchangecoinFilterer{contract: contract}}, nil
}

// Exchangecoin is an auto generated Go binding around an Ethereum contract.
type Exchangecoin struct {
	ExchangecoinCaller     // Read-only binding to the contract
	ExchangecoinTransactor // Write-only binding to the contract
	ExchangecoinFilterer   // Log filterer for contract events
}

// ExchangecoinCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExchangecoinCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangecoinTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExchangecoinTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangecoinFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExchangecoinFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangecoinSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExchangecoinSession struct {
	Contract     *Exchangecoin     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExchangecoinCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExchangecoinCallerSession struct {
	Contract *ExchangecoinCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ExchangecoinTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExchangecoinTransactorSession struct {
	Contract     *ExchangecoinTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ExchangecoinRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExchangecoinRaw struct {
	Contract *Exchangecoin // Generic contract binding to access the raw methods on
}

// ExchangecoinCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExchangecoinCallerRaw struct {
	Contract *ExchangecoinCaller // Generic read-only contract binding to access the raw methods on
}

// ExchangecoinTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExchangecoinTransactorRaw struct {
	Contract *ExchangecoinTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExchangecoin creates a new instance of Exchangecoin, bound to a specific deployed contract.
func NewExchangecoin(address common.Address, backend bind.ContractBackend) (*Exchangecoin, error) {
	contract, err := bindExchangecoin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Exchangecoin{ExchangecoinCaller: ExchangecoinCaller{contract: contract}, ExchangecoinTransactor: ExchangecoinTransactor{contract: contract}, ExchangecoinFilterer: ExchangecoinFilterer{contract: contract}}, nil
}

// NewExchangecoinCaller creates a new read-only instance of Exchangecoin, bound to a specific deployed contract.
func NewExchangecoinCaller(address common.Address, caller bind.ContractCaller) (*ExchangecoinCaller, error) {
	contract, err := bindExchangecoin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExchangecoinCaller{contract: contract}, nil
}

// NewExchangecoinTransactor creates a new write-only instance of Exchangecoin, bound to a specific deployed contract.
func NewExchangecoinTransactor(address common.Address, transactor bind.ContractTransactor) (*ExchangecoinTransactor, error) {
	contract, err := bindExchangecoin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExchangecoinTransactor{contract: contract}, nil
}

// NewExchangecoinFilterer creates a new log filterer instance of Exchangecoin, bound to a specific deployed contract.
func NewExchangecoinFilterer(address common.Address, filterer bind.ContractFilterer) (*ExchangecoinFilterer, error) {
	contract, err := bindExchangecoin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExchangecoinFilterer{contract: contract}, nil
}

// bindExchangecoin binds a generic wrapper to an already deployed contract.
func bindExchangecoin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ExchangecoinABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Exchangecoin *ExchangecoinRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Exchangecoin.Contract.ExchangecoinCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Exchangecoin *ExchangecoinRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Exchangecoin.Contract.ExchangecoinTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Exchangecoin *ExchangecoinRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Exchangecoin.Contract.ExchangecoinTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Exchangecoin *ExchangecoinCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Exchangecoin.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Exchangecoin *ExchangecoinTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Exchangecoin.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Exchangecoin *ExchangecoinTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Exchangecoin.Contract.contract.Transact(opts, method, params...)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() view returns(address)
func (_Exchangecoin *ExchangecoinCaller) Creator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Exchangecoin.contract.Call(opts, &out, "creator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() view returns(address)
func (_Exchangecoin *ExchangecoinSession) Creator() (common.Address, error) {
	return _Exchangecoin.Contract.Creator(&_Exchangecoin.CallOpts)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() view returns(address)
func (_Exchangecoin *ExchangecoinCallerSession) Creator() (common.Address, error) {
	return _Exchangecoin.Contract.Creator(&_Exchangecoin.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Exchangecoin *ExchangecoinCaller) GetBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Exchangecoin.contract.Call(opts, &out, "getBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Exchangecoin *ExchangecoinSession) GetBalance() (*big.Int, error) {
	return _Exchangecoin.Contract.GetBalance(&_Exchangecoin.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Exchangecoin *ExchangecoinCallerSession) GetBalance() (*big.Int, error) {
	return _Exchangecoin.Contract.GetBalance(&_Exchangecoin.CallOpts)
}

// GetEthSignedMessageHash is a free data retrieval call binding the contract method 0xfa540801.
//
// Solidity: function getEthSignedMessageHash(bytes32 _messageHash) pure returns(bytes32)
func (_Exchangecoin *ExchangecoinCaller) GetEthSignedMessageHash(opts *bind.CallOpts, _messageHash [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Exchangecoin.contract.Call(opts, &out, "getEthSignedMessageHash", _messageHash)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetEthSignedMessageHash is a free data retrieval call binding the contract method 0xfa540801.
//
// Solidity: function getEthSignedMessageHash(bytes32 _messageHash) pure returns(bytes32)
func (_Exchangecoin *ExchangecoinSession) GetEthSignedMessageHash(_messageHash [32]byte) ([32]byte, error) {
	return _Exchangecoin.Contract.GetEthSignedMessageHash(&_Exchangecoin.CallOpts, _messageHash)
}

// GetEthSignedMessageHash is a free data retrieval call binding the contract method 0xfa540801.
//
// Solidity: function getEthSignedMessageHash(bytes32 _messageHash) pure returns(bytes32)
func (_Exchangecoin *ExchangecoinCallerSession) GetEthSignedMessageHash(_messageHash [32]byte) ([32]byte, error) {
	return _Exchangecoin.Contract.GetEthSignedMessageHash(&_Exchangecoin.CallOpts, _messageHash)
}

// GetMessageHash is a free data retrieval call binding the contract method 0xa35f8a40.
//
// Solidity: function getMessageHash(address _to, uint256 _amount, string _message, uint256 _nonce) pure returns(bytes32)
func (_Exchangecoin *ExchangecoinCaller) GetMessageHash(opts *bind.CallOpts, _to common.Address, _amount *big.Int, _message string, _nonce *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Exchangecoin.contract.Call(opts, &out, "getMessageHash", _to, _amount, _message, _nonce)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetMessageHash is a free data retrieval call binding the contract method 0xa35f8a40.
//
// Solidity: function getMessageHash(address _to, uint256 _amount, string _message, uint256 _nonce) pure returns(bytes32)
func (_Exchangecoin *ExchangecoinSession) GetMessageHash(_to common.Address, _amount *big.Int, _message string, _nonce *big.Int) ([32]byte, error) {
	return _Exchangecoin.Contract.GetMessageHash(&_Exchangecoin.CallOpts, _to, _amount, _message, _nonce)
}

// GetMessageHash is a free data retrieval call binding the contract method 0xa35f8a40.
//
// Solidity: function getMessageHash(address _to, uint256 _amount, string _message, uint256 _nonce) pure returns(bytes32)
func (_Exchangecoin *ExchangecoinCallerSession) GetMessageHash(_to common.Address, _amount *big.Int, _message string, _nonce *big.Int) ([32]byte, error) {
	return _Exchangecoin.Contract.GetMessageHash(&_Exchangecoin.CallOpts, _to, _amount, _message, _nonce)
}

// RecoverSigner is a free data retrieval call binding the contract method 0x97aba7f9.
//
// Solidity: function recoverSigner(bytes32 _ethSignedMessageHash, bytes _signature) pure returns(address)
func (_Exchangecoin *ExchangecoinCaller) RecoverSigner(opts *bind.CallOpts, _ethSignedMessageHash [32]byte, _signature []byte) (common.Address, error) {
	var out []interface{}
	err := _Exchangecoin.contract.Call(opts, &out, "recoverSigner", _ethSignedMessageHash, _signature)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RecoverSigner is a free data retrieval call binding the contract method 0x97aba7f9.
//
// Solidity: function recoverSigner(bytes32 _ethSignedMessageHash, bytes _signature) pure returns(address)
func (_Exchangecoin *ExchangecoinSession) RecoverSigner(_ethSignedMessageHash [32]byte, _signature []byte) (common.Address, error) {
	return _Exchangecoin.Contract.RecoverSigner(&_Exchangecoin.CallOpts, _ethSignedMessageHash, _signature)
}

// RecoverSigner is a free data retrieval call binding the contract method 0x97aba7f9.
//
// Solidity: function recoverSigner(bytes32 _ethSignedMessageHash, bytes _signature) pure returns(address)
func (_Exchangecoin *ExchangecoinCallerSession) RecoverSigner(_ethSignedMessageHash [32]byte, _signature []byte) (common.Address, error) {
	return _Exchangecoin.Contract.RecoverSigner(&_Exchangecoin.CallOpts, _ethSignedMessageHash, _signature)
}

// SplitSignature is a free data retrieval call binding the contract method 0xa7bb5803.
//
// Solidity: function splitSignature(bytes sig) pure returns(bytes32 r, bytes32 s, uint8 v)
func (_Exchangecoin *ExchangecoinCaller) SplitSignature(opts *bind.CallOpts, sig []byte) (struct {
	R [32]byte
	S [32]byte
	V uint8
}, error) {
	var out []interface{}
	err := _Exchangecoin.contract.Call(opts, &out, "splitSignature", sig)

	outstruct := new(struct {
		R [32]byte
		S [32]byte
		V uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.R = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.S = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.V = *abi.ConvertType(out[2], new(uint8)).(*uint8)

	return *outstruct, err

}

// SplitSignature is a free data retrieval call binding the contract method 0xa7bb5803.
//
// Solidity: function splitSignature(bytes sig) pure returns(bytes32 r, bytes32 s, uint8 v)
func (_Exchangecoin *ExchangecoinSession) SplitSignature(sig []byte) (struct {
	R [32]byte
	S [32]byte
	V uint8
}, error) {
	return _Exchangecoin.Contract.SplitSignature(&_Exchangecoin.CallOpts, sig)
}

// SplitSignature is a free data retrieval call binding the contract method 0xa7bb5803.
//
// Solidity: function splitSignature(bytes sig) pure returns(bytes32 r, bytes32 s, uint8 v)
func (_Exchangecoin *ExchangecoinCallerSession) SplitSignature(sig []byte) (struct {
	R [32]byte
	S [32]byte
	V uint8
}, error) {
	return _Exchangecoin.Contract.SplitSignature(&_Exchangecoin.CallOpts, sig)
}

// Verify is a free data retrieval call binding the contract method 0xabe5026a.
//
// Solidity: function verify(address _signer, address _to, uint256 _amount, string _message, uint256 _nonce, bytes signature) pure returns(bool)
func (_Exchangecoin *ExchangecoinCaller) Verify(opts *bind.CallOpts, _signer common.Address, _to common.Address, _amount *big.Int, _message string, _nonce *big.Int, signature []byte) (bool, error) {
	var out []interface{}
	err := _Exchangecoin.contract.Call(opts, &out, "verify", _signer, _to, _amount, _message, _nonce, signature)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Verify is a free data retrieval call binding the contract method 0xabe5026a.
//
// Solidity: function verify(address _signer, address _to, uint256 _amount, string _message, uint256 _nonce, bytes signature) pure returns(bool)
func (_Exchangecoin *ExchangecoinSession) Verify(_signer common.Address, _to common.Address, _amount *big.Int, _message string, _nonce *big.Int, signature []byte) (bool, error) {
	return _Exchangecoin.Contract.Verify(&_Exchangecoin.CallOpts, _signer, _to, _amount, _message, _nonce, signature)
}

// Verify is a free data retrieval call binding the contract method 0xabe5026a.
//
// Solidity: function verify(address _signer, address _to, uint256 _amount, string _message, uint256 _nonce, bytes signature) pure returns(bool)
func (_Exchangecoin *ExchangecoinCallerSession) Verify(_signer common.Address, _to common.Address, _amount *big.Int, _message string, _nonce *big.Int, signature []byte) (bool, error) {
	return _Exchangecoin.Contract.Verify(&_Exchangecoin.CallOpts, _signer, _to, _amount, _message, _nonce, signature)
}

// ReceiverBalance is a paid mutator transaction binding the contract method 0x0b5fb5e2.
//
// Solidity: function receiverBalance() payable returns()
func (_Exchangecoin *ExchangecoinTransactor) ReceiverBalance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Exchangecoin.contract.Transact(opts, "receiverBalance")
}

// ReceiverBalance is a paid mutator transaction binding the contract method 0x0b5fb5e2.
//
// Solidity: function receiverBalance() payable returns()
func (_Exchangecoin *ExchangecoinSession) ReceiverBalance() (*types.Transaction, error) {
	return _Exchangecoin.Contract.ReceiverBalance(&_Exchangecoin.TransactOpts)
}

// ReceiverBalance is a paid mutator transaction binding the contract method 0x0b5fb5e2.
//
// Solidity: function receiverBalance() payable returns()
func (_Exchangecoin *ExchangecoinTransactorSession) ReceiverBalance() (*types.Transaction, error) {
	return _Exchangecoin.Contract.ReceiverBalance(&_Exchangecoin.TransactOpts)
}
