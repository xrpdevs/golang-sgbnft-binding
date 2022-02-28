// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package SGBFTSO_ERC721

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

// SGBFERC721MetaData contains all meta data concerning the SGBFERC721 contract.
var SGBFERC721MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"PermanentURI\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_od\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"authorDividend\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_od\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"ownerDividend\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_report\",\"type\":\"string\"}],\"name\":\"setupStatus\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"_documentation_read_first\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultRoyaltyPercentage\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"doPause\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"doUnPause\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"freeMintingAllowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isLimitedCollection\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"itemsLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenOwner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"royaltyInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"royaltyAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"setAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"baseURI_\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"_omitContractAddr\",\"type\":\"bool\"}],\"name\":\"setBaseURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"value\",\"type\":\"uint24\"}],\"name\":\"setDefaultRoyalties\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"setLimitedEdition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"setMintingFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dest\",\"type\":\"address\"}],\"name\":\"setMintingFeeDestination\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"setMintingFeeForToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"value\",\"type\":\"uint24\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"setRoyaltiesForToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"collectionOwner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"baseURI\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"mintCost\",\"type\":\"uint256\"}],\"name\":\"setup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokensOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// SGBFERC721ABI is the input ABI used to generate the binding from.
// Deprecated: Use SGBFERC721MetaData.ABI instead.
var SGBFERC721ABI = SGBFERC721MetaData.ABI

// SGBFERC721 is an auto generated Go binding around an Ethereum contract.
type SGBFERC721 struct {
	SGBFERC721Caller     // Read-only binding to the contract
	SGBFERC721Transactor // Write-only binding to the contract
	SGBFERC721Filterer   // Log filterer for contract events
}

// SGBFERC721Caller is an auto generated read-only Go binding around an Ethereum contract.
type SGBFERC721Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SGBFERC721Transactor is an auto generated write-only Go binding around an Ethereum contract.
type SGBFERC721Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SGBFERC721Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SGBFERC721Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SGBFERC721Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SGBFERC721Session struct {
	Contract     *SGBFERC721       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SGBFERC721CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SGBFERC721CallerSession struct {
	Contract *SGBFERC721Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// SGBFERC721TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SGBFERC721TransactorSession struct {
	Contract     *SGBFERC721Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// SGBFERC721Raw is an auto generated low-level Go binding around an Ethereum contract.
type SGBFERC721Raw struct {
	Contract *SGBFERC721 // Generic contract binding to access the raw methods on
}

// SGBFERC721CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SGBFERC721CallerRaw struct {
	Contract *SGBFERC721Caller // Generic read-only contract binding to access the raw methods on
}

// SGBFERC721TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SGBFERC721TransactorRaw struct {
	Contract *SGBFERC721Transactor // Generic write-only contract binding to access the raw methods on
}

// NewSGBFERC721 creates a new instance of SGBFERC721, bound to a specific deployed contract.
func NewSGBFERC721(address common.Address, backend bind.ContractBackend) (*SGBFERC721, error) {
	contract, err := bindSGBFERC721(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SGBFERC721{SGBFERC721Caller: SGBFERC721Caller{contract: contract}, SGBFERC721Transactor: SGBFERC721Transactor{contract: contract}, SGBFERC721Filterer: SGBFERC721Filterer{contract: contract}}, nil
}

// NewSGBFERC721Caller creates a new read-only instance of SGBFERC721, bound to a specific deployed contract.
func NewSGBFERC721Caller(address common.Address, caller bind.ContractCaller) (*SGBFERC721Caller, error) {
	contract, err := bindSGBFERC721(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SGBFERC721Caller{contract: contract}, nil
}

// NewSGBFERC721Transactor creates a new write-only instance of SGBFERC721, bound to a specific deployed contract.
func NewSGBFERC721Transactor(address common.Address, transactor bind.ContractTransactor) (*SGBFERC721Transactor, error) {
	contract, err := bindSGBFERC721(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SGBFERC721Transactor{contract: contract}, nil
}

// NewSGBFERC721Filterer creates a new log filterer instance of SGBFERC721, bound to a specific deployed contract.
func NewSGBFERC721Filterer(address common.Address, filterer bind.ContractFilterer) (*SGBFERC721Filterer, error) {
	contract, err := bindSGBFERC721(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SGBFERC721Filterer{contract: contract}, nil
}

// bindSGBFERC721 binds a generic wrapper to an already deployed contract.
func bindSGBFERC721(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SGBFERC721ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SGBFERC721 *SGBFERC721Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SGBFERC721.Contract.SGBFERC721Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SGBFERC721 *SGBFERC721Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SGBFERC721.Contract.SGBFERC721Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SGBFERC721 *SGBFERC721Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SGBFERC721.Contract.SGBFERC721Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SGBFERC721 *SGBFERC721CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SGBFERC721.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SGBFERC721 *SGBFERC721TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SGBFERC721.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SGBFERC721 *SGBFERC721TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SGBFERC721.Contract.contract.Transact(opts, method, params...)
}

// DocumentationReadFirst is a free data retrieval call binding the contract method 0xb638a1c3.
//
// Solidity: function _documentation_read_first() view returns(string)
func (_SGBFERC721 *SGBFERC721Caller) DocumentationReadFirst(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SGBFERC721.contract.Call(opts, &out, "_documentation_read_first")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// DocumentationReadFirst is a free data retrieval call binding the contract method 0xb638a1c3.
//
// Solidity: function _documentation_read_first() view returns(string)
func (_SGBFERC721 *SGBFERC721Session) DocumentationReadFirst() (string, error) {
	return _SGBFERC721.Contract.DocumentationReadFirst(&_SGBFERC721.CallOpts)
}

// DocumentationReadFirst is a free data retrieval call binding the contract method 0xb638a1c3.
//
// Solidity: function _documentation_read_first() view returns(string)
func (_SGBFERC721 *SGBFERC721CallerSession) DocumentationReadFirst() (string, error) {
	return _SGBFERC721.Contract.DocumentationReadFirst(&_SGBFERC721.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_SGBFERC721 *SGBFERC721Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SGBFERC721.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_SGBFERC721 *SGBFERC721Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _SGBFERC721.Contract.BalanceOf(&_SGBFERC721.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_SGBFERC721 *SGBFERC721CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _SGBFERC721.Contract.BalanceOf(&_SGBFERC721.CallOpts, owner)
}

// DefaultRoyaltyPercentage is a free data retrieval call binding the contract method 0x139d8325.
//
// Solidity: function defaultRoyaltyPercentage() view returns(uint24)
func (_SGBFERC721 *SGBFERC721Caller) DefaultRoyaltyPercentage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SGBFERC721.contract.Call(opts, &out, "defaultRoyaltyPercentage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DefaultRoyaltyPercentage is a free data retrieval call binding the contract method 0x139d8325.
//
// Solidity: function defaultRoyaltyPercentage() view returns(uint24)
func (_SGBFERC721 *SGBFERC721Session) DefaultRoyaltyPercentage() (*big.Int, error) {
	return _SGBFERC721.Contract.DefaultRoyaltyPercentage(&_SGBFERC721.CallOpts)
}

// DefaultRoyaltyPercentage is a free data retrieval call binding the contract method 0x139d8325.
//
// Solidity: function defaultRoyaltyPercentage() view returns(uint24)
func (_SGBFERC721 *SGBFERC721CallerSession) DefaultRoyaltyPercentage() (*big.Int, error) {
	return _SGBFERC721.Contract.DefaultRoyaltyPercentage(&_SGBFERC721.CallOpts)
}

// FreeMintingAllowance is a free data retrieval call binding the contract method 0xe259f2f3.
//
// Solidity: function freeMintingAllowance() view returns(uint256)
func (_SGBFERC721 *SGBFERC721Caller) FreeMintingAllowance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SGBFERC721.contract.Call(opts, &out, "freeMintingAllowance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FreeMintingAllowance is a free data retrieval call binding the contract method 0xe259f2f3.
//
// Solidity: function freeMintingAllowance() view returns(uint256)
func (_SGBFERC721 *SGBFERC721Session) FreeMintingAllowance() (*big.Int, error) {
	return _SGBFERC721.Contract.FreeMintingAllowance(&_SGBFERC721.CallOpts)
}

// FreeMintingAllowance is a free data retrieval call binding the contract method 0xe259f2f3.
//
// Solidity: function freeMintingAllowance() view returns(uint256)
func (_SGBFERC721 *SGBFERC721CallerSession) FreeMintingAllowance() (*big.Int, error) {
	return _SGBFERC721.Contract.FreeMintingAllowance(&_SGBFERC721.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_SGBFERC721 *SGBFERC721Caller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _SGBFERC721.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_SGBFERC721 *SGBFERC721Session) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _SGBFERC721.Contract.GetApproved(&_SGBFERC721.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_SGBFERC721 *SGBFERC721CallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _SGBFERC721.Contract.GetApproved(&_SGBFERC721.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_SGBFERC721 *SGBFERC721Caller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _SGBFERC721.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_SGBFERC721 *SGBFERC721Session) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _SGBFERC721.Contract.IsApprovedForAll(&_SGBFERC721.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_SGBFERC721 *SGBFERC721CallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _SGBFERC721.Contract.IsApprovedForAll(&_SGBFERC721.CallOpts, owner, operator)
}

// IsLimitedCollection is a free data retrieval call binding the contract method 0x274bbb64.
//
// Solidity: function isLimitedCollection() view returns(bool)
func (_SGBFERC721 *SGBFERC721Caller) IsLimitedCollection(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SGBFERC721.contract.Call(opts, &out, "isLimitedCollection")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLimitedCollection is a free data retrieval call binding the contract method 0x274bbb64.
//
// Solidity: function isLimitedCollection() view returns(bool)
func (_SGBFERC721 *SGBFERC721Session) IsLimitedCollection() (bool, error) {
	return _SGBFERC721.Contract.IsLimitedCollection(&_SGBFERC721.CallOpts)
}

// IsLimitedCollection is a free data retrieval call binding the contract method 0x274bbb64.
//
// Solidity: function isLimitedCollection() view returns(bool)
func (_SGBFERC721 *SGBFERC721CallerSession) IsLimitedCollection() (bool, error) {
	return _SGBFERC721.Contract.IsLimitedCollection(&_SGBFERC721.CallOpts)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_SGBFERC721 *SGBFERC721Caller) IsPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SGBFERC721.contract.Call(opts, &out, "isPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_SGBFERC721 *SGBFERC721Session) IsPaused() (bool, error) {
	return _SGBFERC721.Contract.IsPaused(&_SGBFERC721.CallOpts)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_SGBFERC721 *SGBFERC721CallerSession) IsPaused() (bool, error) {
	return _SGBFERC721.Contract.IsPaused(&_SGBFERC721.CallOpts)
}

// ItemsLimit is a free data retrieval call binding the contract method 0x6fded1fa.
//
// Solidity: function itemsLimit() view returns(uint256)
func (_SGBFERC721 *SGBFERC721Caller) ItemsLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SGBFERC721.contract.Call(opts, &out, "itemsLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ItemsLimit is a free data retrieval call binding the contract method 0x6fded1fa.
//
// Solidity: function itemsLimit() view returns(uint256)
func (_SGBFERC721 *SGBFERC721Session) ItemsLimit() (*big.Int, error) {
	return _SGBFERC721.Contract.ItemsLimit(&_SGBFERC721.CallOpts)
}

// ItemsLimit is a free data retrieval call binding the contract method 0x6fded1fa.
//
// Solidity: function itemsLimit() view returns(uint256)
func (_SGBFERC721 *SGBFERC721CallerSession) ItemsLimit() (*big.Int, error) {
	return _SGBFERC721.Contract.ItemsLimit(&_SGBFERC721.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SGBFERC721 *SGBFERC721Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SGBFERC721.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SGBFERC721 *SGBFERC721Session) Name() (string, error) {
	return _SGBFERC721.Contract.Name(&_SGBFERC721.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SGBFERC721 *SGBFERC721CallerSession) Name() (string, error) {
	return _SGBFERC721.Contract.Name(&_SGBFERC721.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SGBFERC721 *SGBFERC721Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SGBFERC721.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SGBFERC721 *SGBFERC721Session) Owner() (common.Address, error) {
	return _SGBFERC721.Contract.Owner(&_SGBFERC721.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SGBFERC721 *SGBFERC721CallerSession) Owner() (common.Address, error) {
	return _SGBFERC721.Contract.Owner(&_SGBFERC721.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_SGBFERC721 *SGBFERC721Caller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _SGBFERC721.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_SGBFERC721 *SGBFERC721Session) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _SGBFERC721.Contract.OwnerOf(&_SGBFERC721.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_SGBFERC721 *SGBFERC721CallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _SGBFERC721.Contract.OwnerOf(&_SGBFERC721.CallOpts, tokenId)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_SGBFERC721 *SGBFERC721Caller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SGBFERC721.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_SGBFERC721 *SGBFERC721Session) Paused() (bool, error) {
	return _SGBFERC721.Contract.Paused(&_SGBFERC721.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_SGBFERC721 *SGBFERC721CallerSession) Paused() (bool, error) {
	return _SGBFERC721.Contract.Paused(&_SGBFERC721.CallOpts)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 value) view returns(address receiver, uint256 royaltyAmount)
func (_SGBFERC721 *SGBFERC721Caller) RoyaltyInfo(opts *bind.CallOpts, tokenId *big.Int, value *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	var out []interface{}
	err := _SGBFERC721.contract.Call(opts, &out, "royaltyInfo", tokenId, value)

	outstruct := new(struct {
		Receiver      common.Address
		RoyaltyAmount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Receiver = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.RoyaltyAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 value) view returns(address receiver, uint256 royaltyAmount)
func (_SGBFERC721 *SGBFERC721Session) RoyaltyInfo(tokenId *big.Int, value *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	return _SGBFERC721.Contract.RoyaltyInfo(&_SGBFERC721.CallOpts, tokenId, value)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 value) view returns(address receiver, uint256 royaltyAmount)
func (_SGBFERC721 *SGBFERC721CallerSession) RoyaltyInfo(tokenId *big.Int, value *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	return _SGBFERC721.Contract.RoyaltyInfo(&_SGBFERC721.CallOpts, tokenId, value)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SGBFERC721 *SGBFERC721Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _SGBFERC721.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SGBFERC721 *SGBFERC721Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SGBFERC721.Contract.SupportsInterface(&_SGBFERC721.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SGBFERC721 *SGBFERC721CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SGBFERC721.Contract.SupportsInterface(&_SGBFERC721.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SGBFERC721 *SGBFERC721Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SGBFERC721.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SGBFERC721 *SGBFERC721Session) Symbol() (string, error) {
	return _SGBFERC721.Contract.Symbol(&_SGBFERC721.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SGBFERC721 *SGBFERC721CallerSession) Symbol() (string, error) {
	return _SGBFERC721.Contract.Symbol(&_SGBFERC721.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_SGBFERC721 *SGBFERC721Caller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SGBFERC721.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_SGBFERC721 *SGBFERC721Session) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _SGBFERC721.Contract.TokenByIndex(&_SGBFERC721.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_SGBFERC721 *SGBFERC721CallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _SGBFERC721.Contract.TokenByIndex(&_SGBFERC721.CallOpts, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_SGBFERC721 *SGBFERC721Caller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _SGBFERC721.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_SGBFERC721 *SGBFERC721Session) TokenURI(tokenId *big.Int) (string, error) {
	return _SGBFERC721.Contract.TokenURI(&_SGBFERC721.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_SGBFERC721 *SGBFERC721CallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _SGBFERC721.Contract.TokenURI(&_SGBFERC721.CallOpts, tokenId)
}

// TokensOfOwnerByIndex is a free data retrieval call binding the contract method 0x4707f44f.
//
// Solidity: function tokensOfOwnerByIndex(address owner, uint256 index) view returns(uint256 tokenId)
func (_SGBFERC721 *SGBFERC721Caller) TokensOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SGBFERC721.contract.Call(opts, &out, "tokensOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokensOfOwnerByIndex is a free data retrieval call binding the contract method 0x4707f44f.
//
// Solidity: function tokensOfOwnerByIndex(address owner, uint256 index) view returns(uint256 tokenId)
func (_SGBFERC721 *SGBFERC721Session) TokensOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _SGBFERC721.Contract.TokensOfOwnerByIndex(&_SGBFERC721.CallOpts, owner, index)
}

// TokensOfOwnerByIndex is a free data retrieval call binding the contract method 0x4707f44f.
//
// Solidity: function tokensOfOwnerByIndex(address owner, uint256 index) view returns(uint256 tokenId)
func (_SGBFERC721 *SGBFERC721CallerSession) TokensOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _SGBFERC721.Contract.TokensOfOwnerByIndex(&_SGBFERC721.CallOpts, owner, index)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SGBFERC721 *SGBFERC721Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SGBFERC721.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SGBFERC721 *SGBFERC721Session) TotalSupply() (*big.Int, error) {
	return _SGBFERC721.Contract.TotalSupply(&_SGBFERC721.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SGBFERC721 *SGBFERC721CallerSession) TotalSupply() (*big.Int, error) {
	return _SGBFERC721.Contract.TotalSupply(&_SGBFERC721.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_SGBFERC721 *SGBFERC721Transactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SGBFERC721.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_SGBFERC721 *SGBFERC721Session) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SGBFERC721.Contract.Approve(&_SGBFERC721.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_SGBFERC721 *SGBFERC721TransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SGBFERC721.Contract.Approve(&_SGBFERC721.TransactOpts, to, tokenId)
}

// DoPause is a paid mutator transaction binding the contract method 0x67d0661d.
//
// Solidity: function doPause() returns(bool)
func (_SGBFERC721 *SGBFERC721Transactor) DoPause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SGBFERC721.contract.Transact(opts, "doPause")
}

// DoPause is a paid mutator transaction binding the contract method 0x67d0661d.
//
// Solidity: function doPause() returns(bool)
func (_SGBFERC721 *SGBFERC721Session) DoPause() (*types.Transaction, error) {
	return _SGBFERC721.Contract.DoPause(&_SGBFERC721.TransactOpts)
}

// DoPause is a paid mutator transaction binding the contract method 0x67d0661d.
//
// Solidity: function doPause() returns(bool)
func (_SGBFERC721 *SGBFERC721TransactorSession) DoPause() (*types.Transaction, error) {
	return _SGBFERC721.Contract.DoPause(&_SGBFERC721.TransactOpts)
}

// DoUnPause is a paid mutator transaction binding the contract method 0x6f37f59c.
//
// Solidity: function doUnPause() returns(bool)
func (_SGBFERC721 *SGBFERC721Transactor) DoUnPause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SGBFERC721.contract.Transact(opts, "doUnPause")
}

// DoUnPause is a paid mutator transaction binding the contract method 0x6f37f59c.
//
// Solidity: function doUnPause() returns(bool)
func (_SGBFERC721 *SGBFERC721Session) DoUnPause() (*types.Transaction, error) {
	return _SGBFERC721.Contract.DoUnPause(&_SGBFERC721.TransactOpts)
}

// DoUnPause is a paid mutator transaction binding the contract method 0x6f37f59c.
//
// Solidity: function doUnPause() returns(bool)
func (_SGBFERC721 *SGBFERC721TransactorSession) DoUnPause() (*types.Transaction, error) {
	return _SGBFERC721.Contract.DoUnPause(&_SGBFERC721.TransactOpts)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address tokenOwner, uint256 tokenID) payable returns()
func (_SGBFERC721 *SGBFERC721Transactor) Mint(opts *bind.TransactOpts, tokenOwner common.Address, tokenID *big.Int) (*types.Transaction, error) {
	return _SGBFERC721.contract.Transact(opts, "mint", tokenOwner, tokenID)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address tokenOwner, uint256 tokenID) payable returns()
func (_SGBFERC721 *SGBFERC721Session) Mint(tokenOwner common.Address, tokenID *big.Int) (*types.Transaction, error) {
	return _SGBFERC721.Contract.Mint(&_SGBFERC721.TransactOpts, tokenOwner, tokenID)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address tokenOwner, uint256 tokenID) payable returns()
func (_SGBFERC721 *SGBFERC721TransactorSession) Mint(tokenOwner common.Address, tokenID *big.Int) (*types.Transaction, error) {
	return _SGBFERC721.Contract.Mint(&_SGBFERC721.TransactOpts, tokenOwner, tokenID)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SGBFERC721 *SGBFERC721Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SGBFERC721.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SGBFERC721 *SGBFERC721Session) RenounceOwnership() (*types.Transaction, error) {
	return _SGBFERC721.Contract.RenounceOwnership(&_SGBFERC721.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SGBFERC721 *SGBFERC721TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SGBFERC721.Contract.RenounceOwnership(&_SGBFERC721.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_SGBFERC721 *SGBFERC721Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SGBFERC721.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_SGBFERC721 *SGBFERC721Session) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SGBFERC721.Contract.SafeTransferFrom(&_SGBFERC721.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_SGBFERC721 *SGBFERC721TransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SGBFERC721.Contract.SafeTransferFrom(&_SGBFERC721.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_SGBFERC721 *SGBFERC721Transactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _SGBFERC721.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_SGBFERC721 *SGBFERC721Session) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _SGBFERC721.Contract.SafeTransferFrom0(&_SGBFERC721.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_SGBFERC721 *SGBFERC721TransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _SGBFERC721.Contract.SafeTransferFrom0(&_SGBFERC721.TransactOpts, from, to, tokenId, _data)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address admin) returns()
func (_SGBFERC721 *SGBFERC721Transactor) SetAdmin(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _SGBFERC721.contract.Transact(opts, "setAdmin", admin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address admin) returns()
func (_SGBFERC721 *SGBFERC721Session) SetAdmin(admin common.Address) (*types.Transaction, error) {
	return _SGBFERC721.Contract.SetAdmin(&_SGBFERC721.TransactOpts, admin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address admin) returns()
func (_SGBFERC721 *SGBFERC721TransactorSession) SetAdmin(admin common.Address) (*types.Transaction, error) {
	return _SGBFERC721.Contract.SetAdmin(&_SGBFERC721.TransactOpts, admin)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_SGBFERC721 *SGBFERC721Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _SGBFERC721.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_SGBFERC721 *SGBFERC721Session) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _SGBFERC721.Contract.SetApprovalForAll(&_SGBFERC721.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_SGBFERC721 *SGBFERC721TransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _SGBFERC721.Contract.SetApprovalForAll(&_SGBFERC721.TransactOpts, operator, approved)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0xb64b21ca.
//
// Solidity: function setBaseURI(string baseURI_, bool _omitContractAddr) returns()
func (_SGBFERC721 *SGBFERC721Transactor) SetBaseURI(opts *bind.TransactOpts, baseURI_ string, _omitContractAddr bool) (*types.Transaction, error) {
	return _SGBFERC721.contract.Transact(opts, "setBaseURI", baseURI_, _omitContractAddr)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0xb64b21ca.
//
// Solidity: function setBaseURI(string baseURI_, bool _omitContractAddr) returns()
func (_SGBFERC721 *SGBFERC721Session) SetBaseURI(baseURI_ string, _omitContractAddr bool) (*types.Transaction, error) {
	return _SGBFERC721.Contract.SetBaseURI(&_SGBFERC721.TransactOpts, baseURI_, _omitContractAddr)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0xb64b21ca.
//
// Solidity: function setBaseURI(string baseURI_, bool _omitContractAddr) returns()
func (_SGBFERC721 *SGBFERC721TransactorSession) SetBaseURI(baseURI_ string, _omitContractAddr bool) (*types.Transaction, error) {
	return _SGBFERC721.Contract.SetBaseURI(&_SGBFERC721.TransactOpts, baseURI_, _omitContractAddr)
}

// SetDefaultRoyalties is a paid mutator transaction binding the contract method 0xf71c9056.
//
// Solidity: function setDefaultRoyalties(address beneficiary, uint24 value) returns()
func (_SGBFERC721 *SGBFERC721Transactor) SetDefaultRoyalties(opts *bind.TransactOpts, beneficiary common.Address, value *big.Int) (*types.Transaction, error) {
	return _SGBFERC721.contract.Transact(opts, "setDefaultRoyalties", beneficiary, value)
}

// SetDefaultRoyalties is a paid mutator transaction binding the contract method 0xf71c9056.
//
// Solidity: function setDefaultRoyalties(address beneficiary, uint24 value) returns()
func (_SGBFERC721 *SGBFERC721Session) SetDefaultRoyalties(beneficiary common.Address, value *big.Int) (*types.Transaction, error) {
	return _SGBFERC721.Contract.SetDefaultRoyalties(&_SGBFERC721.TransactOpts, beneficiary, value)
}

// SetDefaultRoyalties is a paid mutator transaction binding the contract method 0xf71c9056.
//
// Solidity: function setDefaultRoyalties(address beneficiary, uint24 value) returns()
func (_SGBFERC721 *SGBFERC721TransactorSession) SetDefaultRoyalties(beneficiary common.Address, value *big.Int) (*types.Transaction, error) {
	return _SGBFERC721.Contract.SetDefaultRoyalties(&_SGBFERC721.TransactOpts, beneficiary, value)
}

// SetLimitedEdition is a paid mutator transaction binding the contract method 0x8f99fa1b.
//
// Solidity: function setLimitedEdition(uint256 limit) returns()
func (_SGBFERC721 *SGBFERC721Transactor) SetLimitedEdition(opts *bind.TransactOpts, limit *big.Int) (*types.Transaction, error) {
	return _SGBFERC721.contract.Transact(opts, "setLimitedEdition", limit)
}

// SetLimitedEdition is a paid mutator transaction binding the contract method 0x8f99fa1b.
//
// Solidity: function setLimitedEdition(uint256 limit) returns()
func (_SGBFERC721 *SGBFERC721Session) SetLimitedEdition(limit *big.Int) (*types.Transaction, error) {
	return _SGBFERC721.Contract.SetLimitedEdition(&_SGBFERC721.TransactOpts, limit)
}

// SetLimitedEdition is a paid mutator transaction binding the contract method 0x8f99fa1b.
//
// Solidity: function setLimitedEdition(uint256 limit) returns()
func (_SGBFERC721 *SGBFERC721TransactorSession) SetLimitedEdition(limit *big.Int) (*types.Transaction, error) {
	return _SGBFERC721.Contract.SetLimitedEdition(&_SGBFERC721.TransactOpts, limit)
}

// SetMintingFee is a paid mutator transaction binding the contract method 0x238a4709.
//
// Solidity: function setMintingFee(uint256 amount) returns()
func (_SGBFERC721 *SGBFERC721Transactor) SetMintingFee(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _SGBFERC721.contract.Transact(opts, "setMintingFee", amount)
}

// SetMintingFee is a paid mutator transaction binding the contract method 0x238a4709.
//
// Solidity: function setMintingFee(uint256 amount) returns()
func (_SGBFERC721 *SGBFERC721Session) SetMintingFee(amount *big.Int) (*types.Transaction, error) {
	return _SGBFERC721.Contract.SetMintingFee(&_SGBFERC721.TransactOpts, amount)
}

// SetMintingFee is a paid mutator transaction binding the contract method 0x238a4709.
//
// Solidity: function setMintingFee(uint256 amount) returns()
func (_SGBFERC721 *SGBFERC721TransactorSession) SetMintingFee(amount *big.Int) (*types.Transaction, error) {
	return _SGBFERC721.Contract.SetMintingFee(&_SGBFERC721.TransactOpts, amount)
}

// SetMintingFeeDestination is a paid mutator transaction binding the contract method 0x85c56e96.
//
// Solidity: function setMintingFeeDestination(address dest) returns()
func (_SGBFERC721 *SGBFERC721Transactor) SetMintingFeeDestination(opts *bind.TransactOpts, dest common.Address) (*types.Transaction, error) {
	return _SGBFERC721.contract.Transact(opts, "setMintingFeeDestination", dest)
}

// SetMintingFeeDestination is a paid mutator transaction binding the contract method 0x85c56e96.
//
// Solidity: function setMintingFeeDestination(address dest) returns()
func (_SGBFERC721 *SGBFERC721Session) SetMintingFeeDestination(dest common.Address) (*types.Transaction, error) {
	return _SGBFERC721.Contract.SetMintingFeeDestination(&_SGBFERC721.TransactOpts, dest)
}

// SetMintingFeeDestination is a paid mutator transaction binding the contract method 0x85c56e96.
//
// Solidity: function setMintingFeeDestination(address dest) returns()
func (_SGBFERC721 *SGBFERC721TransactorSession) SetMintingFeeDestination(dest common.Address) (*types.Transaction, error) {
	return _SGBFERC721.Contract.SetMintingFeeDestination(&_SGBFERC721.TransactOpts, dest)
}

// SetMintingFeeForToken is a paid mutator transaction binding the contract method 0x3a96896f.
//
// Solidity: function setMintingFeeForToken(uint256 amount, uint256 tokenId) returns()
func (_SGBFERC721 *SGBFERC721Transactor) SetMintingFeeForToken(opts *bind.TransactOpts, amount *big.Int, tokenId *big.Int) (*types.Transaction, error) {
	return _SGBFERC721.contract.Transact(opts, "setMintingFeeForToken", amount, tokenId)
}

// SetMintingFeeForToken is a paid mutator transaction binding the contract method 0x3a96896f.
//
// Solidity: function setMintingFeeForToken(uint256 amount, uint256 tokenId) returns()
func (_SGBFERC721 *SGBFERC721Session) SetMintingFeeForToken(amount *big.Int, tokenId *big.Int) (*types.Transaction, error) {
	return _SGBFERC721.Contract.SetMintingFeeForToken(&_SGBFERC721.TransactOpts, amount, tokenId)
}

// SetMintingFeeForToken is a paid mutator transaction binding the contract method 0x3a96896f.
//
// Solidity: function setMintingFeeForToken(uint256 amount, uint256 tokenId) returns()
func (_SGBFERC721 *SGBFERC721TransactorSession) SetMintingFeeForToken(amount *big.Int, tokenId *big.Int) (*types.Transaction, error) {
	return _SGBFERC721.Contract.SetMintingFeeForToken(&_SGBFERC721.TransactOpts, amount, tokenId)
}

// SetRoyaltiesForToken is a paid mutator transaction binding the contract method 0xf34cf497.
//
// Solidity: function setRoyaltiesForToken(address beneficiary, uint24 value, uint256 tokenId) returns()
func (_SGBFERC721 *SGBFERC721Transactor) SetRoyaltiesForToken(opts *bind.TransactOpts, beneficiary common.Address, value *big.Int, tokenId *big.Int) (*types.Transaction, error) {
	return _SGBFERC721.contract.Transact(opts, "setRoyaltiesForToken", beneficiary, value, tokenId)
}

// SetRoyaltiesForToken is a paid mutator transaction binding the contract method 0xf34cf497.
//
// Solidity: function setRoyaltiesForToken(address beneficiary, uint24 value, uint256 tokenId) returns()
func (_SGBFERC721 *SGBFERC721Session) SetRoyaltiesForToken(beneficiary common.Address, value *big.Int, tokenId *big.Int) (*types.Transaction, error) {
	return _SGBFERC721.Contract.SetRoyaltiesForToken(&_SGBFERC721.TransactOpts, beneficiary, value, tokenId)
}

// SetRoyaltiesForToken is a paid mutator transaction binding the contract method 0xf34cf497.
//
// Solidity: function setRoyaltiesForToken(address beneficiary, uint24 value, uint256 tokenId) returns()
func (_SGBFERC721 *SGBFERC721TransactorSession) SetRoyaltiesForToken(beneficiary common.Address, value *big.Int, tokenId *big.Int) (*types.Transaction, error) {
	return _SGBFERC721.Contract.SetRoyaltiesForToken(&_SGBFERC721.TransactOpts, beneficiary, value, tokenId)
}

// Setup is a paid mutator transaction binding the contract method 0x3b140730.
//
// Solidity: function setup(string name, string symbol, address collectionOwner, string baseURI, uint256 mintCost) returns()
func (_SGBFERC721 *SGBFERC721Transactor) Setup(opts *bind.TransactOpts, name string, symbol string, collectionOwner common.Address, baseURI string, mintCost *big.Int) (*types.Transaction, error) {
	return _SGBFERC721.contract.Transact(opts, "setup", name, symbol, collectionOwner, baseURI, mintCost)
}

// Setup is a paid mutator transaction binding the contract method 0x3b140730.
//
// Solidity: function setup(string name, string symbol, address collectionOwner, string baseURI, uint256 mintCost) returns()
func (_SGBFERC721 *SGBFERC721Session) Setup(name string, symbol string, collectionOwner common.Address, baseURI string, mintCost *big.Int) (*types.Transaction, error) {
	return _SGBFERC721.Contract.Setup(&_SGBFERC721.TransactOpts, name, symbol, collectionOwner, baseURI, mintCost)
}

// Setup is a paid mutator transaction binding the contract method 0x3b140730.
//
// Solidity: function setup(string name, string symbol, address collectionOwner, string baseURI, uint256 mintCost) returns()
func (_SGBFERC721 *SGBFERC721TransactorSession) Setup(name string, symbol string, collectionOwner common.Address, baseURI string, mintCost *big.Int) (*types.Transaction, error) {
	return _SGBFERC721.Contract.Setup(&_SGBFERC721.TransactOpts, name, symbol, collectionOwner, baseURI, mintCost)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_SGBFERC721 *SGBFERC721Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SGBFERC721.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_SGBFERC721 *SGBFERC721Session) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SGBFERC721.Contract.TransferFrom(&_SGBFERC721.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_SGBFERC721 *SGBFERC721TransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SGBFERC721.Contract.TransferFrom(&_SGBFERC721.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SGBFERC721 *SGBFERC721Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SGBFERC721.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SGBFERC721 *SGBFERC721Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SGBFERC721.Contract.TransferOwnership(&_SGBFERC721.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SGBFERC721 *SGBFERC721TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SGBFERC721.Contract.TransferOwnership(&_SGBFERC721.TransactOpts, newOwner)
}

// SGBFERC721ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the SGBFERC721 contract.
type SGBFERC721ApprovalIterator struct {
	Event *SGBFERC721Approval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SGBFERC721ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SGBFERC721Approval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SGBFERC721Approval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SGBFERC721ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SGBFERC721ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SGBFERC721Approval represents a Approval event raised by the SGBFERC721 contract.
type SGBFERC721Approval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_SGBFERC721 *SGBFERC721Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*SGBFERC721ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _SGBFERC721.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &SGBFERC721ApprovalIterator{contract: _SGBFERC721.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_SGBFERC721 *SGBFERC721Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *SGBFERC721Approval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _SGBFERC721.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SGBFERC721Approval)
				if err := _SGBFERC721.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_SGBFERC721 *SGBFERC721Filterer) ParseApproval(log types.Log) (*SGBFERC721Approval, error) {
	event := new(SGBFERC721Approval)
	if err := _SGBFERC721.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SGBFERC721ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the SGBFERC721 contract.
type SGBFERC721ApprovalForAllIterator struct {
	Event *SGBFERC721ApprovalForAll // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SGBFERC721ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SGBFERC721ApprovalForAll)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SGBFERC721ApprovalForAll)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SGBFERC721ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SGBFERC721ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SGBFERC721ApprovalForAll represents a ApprovalForAll event raised by the SGBFERC721 contract.
type SGBFERC721ApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_SGBFERC721 *SGBFERC721Filterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*SGBFERC721ApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _SGBFERC721.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &SGBFERC721ApprovalForAllIterator{contract: _SGBFERC721.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_SGBFERC721 *SGBFERC721Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *SGBFERC721ApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _SGBFERC721.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SGBFERC721ApprovalForAll)
				if err := _SGBFERC721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_SGBFERC721 *SGBFERC721Filterer) ParseApprovalForAll(log types.Log) (*SGBFERC721ApprovalForAll, error) {
	event := new(SGBFERC721ApprovalForAll)
	if err := _SGBFERC721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SGBFERC721OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SGBFERC721 contract.
type SGBFERC721OwnershipTransferredIterator struct {
	Event *SGBFERC721OwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SGBFERC721OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SGBFERC721OwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SGBFERC721OwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SGBFERC721OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SGBFERC721OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SGBFERC721OwnershipTransferred represents a OwnershipTransferred event raised by the SGBFERC721 contract.
type SGBFERC721OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SGBFERC721 *SGBFERC721Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SGBFERC721OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SGBFERC721.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SGBFERC721OwnershipTransferredIterator{contract: _SGBFERC721.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SGBFERC721 *SGBFERC721Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SGBFERC721OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SGBFERC721.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SGBFERC721OwnershipTransferred)
				if err := _SGBFERC721.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SGBFERC721 *SGBFERC721Filterer) ParseOwnershipTransferred(log types.Log) (*SGBFERC721OwnershipTransferred, error) {
	event := new(SGBFERC721OwnershipTransferred)
	if err := _SGBFERC721.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SGBFERC721PausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the SGBFERC721 contract.
type SGBFERC721PausedIterator struct {
	Event *SGBFERC721Paused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SGBFERC721PausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SGBFERC721Paused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SGBFERC721Paused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SGBFERC721PausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SGBFERC721PausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SGBFERC721Paused represents a Paused event raised by the SGBFERC721 contract.
type SGBFERC721Paused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_SGBFERC721 *SGBFERC721Filterer) FilterPaused(opts *bind.FilterOpts) (*SGBFERC721PausedIterator, error) {

	logs, sub, err := _SGBFERC721.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &SGBFERC721PausedIterator{contract: _SGBFERC721.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_SGBFERC721 *SGBFERC721Filterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *SGBFERC721Paused) (event.Subscription, error) {

	logs, sub, err := _SGBFERC721.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SGBFERC721Paused)
				if err := _SGBFERC721.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_SGBFERC721 *SGBFERC721Filterer) ParsePaused(log types.Log) (*SGBFERC721Paused, error) {
	event := new(SGBFERC721Paused)
	if err := _SGBFERC721.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SGBFERC721PermanentURIIterator is returned from FilterPermanentURI and is used to iterate over the raw logs and unpacked data for PermanentURI events raised by the SGBFERC721 contract.
type SGBFERC721PermanentURIIterator struct {
	Event *SGBFERC721PermanentURI // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SGBFERC721PermanentURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SGBFERC721PermanentURI)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SGBFERC721PermanentURI)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SGBFERC721PermanentURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SGBFERC721PermanentURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SGBFERC721PermanentURI represents a PermanentURI event raised by the SGBFERC721 contract.
type SGBFERC721PermanentURI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPermanentURI is a free log retrieval operation binding the contract event 0xa109ba539900bf1b633f956d63c96fc89b814c7287f7aa50a9216d0b55657207.
//
// Solidity: event PermanentURI(string _value, uint256 indexed _id)
func (_SGBFERC721 *SGBFERC721Filterer) FilterPermanentURI(opts *bind.FilterOpts, _id []*big.Int) (*SGBFERC721PermanentURIIterator, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _SGBFERC721.contract.FilterLogs(opts, "PermanentURI", _idRule)
	if err != nil {
		return nil, err
	}
	return &SGBFERC721PermanentURIIterator{contract: _SGBFERC721.contract, event: "PermanentURI", logs: logs, sub: sub}, nil
}

// WatchPermanentURI is a free log subscription operation binding the contract event 0xa109ba539900bf1b633f956d63c96fc89b814c7287f7aa50a9216d0b55657207.
//
// Solidity: event PermanentURI(string _value, uint256 indexed _id)
func (_SGBFERC721 *SGBFERC721Filterer) WatchPermanentURI(opts *bind.WatchOpts, sink chan<- *SGBFERC721PermanentURI, _id []*big.Int) (event.Subscription, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _SGBFERC721.contract.WatchLogs(opts, "PermanentURI", _idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SGBFERC721PermanentURI)
				if err := _SGBFERC721.contract.UnpackLog(event, "PermanentURI", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePermanentURI is a log parse operation binding the contract event 0xa109ba539900bf1b633f956d63c96fc89b814c7287f7aa50a9216d0b55657207.
//
// Solidity: event PermanentURI(string _value, uint256 indexed _id)
func (_SGBFERC721 *SGBFERC721Filterer) ParsePermanentURI(log types.Log) (*SGBFERC721PermanentURI, error) {
	event := new(SGBFERC721PermanentURI)
	if err := _SGBFERC721.contract.UnpackLog(event, "PermanentURI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SGBFERC721TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the SGBFERC721 contract.
type SGBFERC721TransferIterator struct {
	Event *SGBFERC721Transfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SGBFERC721TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SGBFERC721Transfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SGBFERC721Transfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SGBFERC721TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SGBFERC721TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SGBFERC721Transfer represents a Transfer event raised by the SGBFERC721 contract.
type SGBFERC721Transfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_SGBFERC721 *SGBFERC721Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*SGBFERC721TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _SGBFERC721.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &SGBFERC721TransferIterator{contract: _SGBFERC721.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_SGBFERC721 *SGBFERC721Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *SGBFERC721Transfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _SGBFERC721.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SGBFERC721Transfer)
				if err := _SGBFERC721.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_SGBFERC721 *SGBFERC721Filterer) ParseTransfer(log types.Log) (*SGBFERC721Transfer, error) {
	event := new(SGBFERC721Transfer)
	if err := _SGBFERC721.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SGBFERC721UnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the SGBFERC721 contract.
type SGBFERC721UnpausedIterator struct {
	Event *SGBFERC721Unpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SGBFERC721UnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SGBFERC721Unpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SGBFERC721Unpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SGBFERC721UnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SGBFERC721UnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SGBFERC721Unpaused represents a Unpaused event raised by the SGBFERC721 contract.
type SGBFERC721Unpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_SGBFERC721 *SGBFERC721Filterer) FilterUnpaused(opts *bind.FilterOpts) (*SGBFERC721UnpausedIterator, error) {

	logs, sub, err := _SGBFERC721.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &SGBFERC721UnpausedIterator{contract: _SGBFERC721.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_SGBFERC721 *SGBFERC721Filterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *SGBFERC721Unpaused) (event.Subscription, error) {

	logs, sub, err := _SGBFERC721.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SGBFERC721Unpaused)
				if err := _SGBFERC721.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_SGBFERC721 *SGBFERC721Filterer) ParseUnpaused(log types.Log) (*SGBFERC721Unpaused, error) {
	event := new(SGBFERC721Unpaused)
	if err := _SGBFERC721.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SGBFERC721AuthorDividendIterator is returned from FilterAuthorDividend and is used to iterate over the raw logs and unpacked data for AuthorDividend events raised by the SGBFERC721 contract.
type SGBFERC721AuthorDividendIterator struct {
	Event *SGBFERC721AuthorDividend // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SGBFERC721AuthorDividendIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SGBFERC721AuthorDividend)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SGBFERC721AuthorDividend)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SGBFERC721AuthorDividendIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SGBFERC721AuthorDividendIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SGBFERC721AuthorDividend represents a AuthorDividend event raised by the SGBFERC721 contract.
type SGBFERC721AuthorDividend struct {
	Od     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAuthorDividend is a free log retrieval operation binding the contract event 0xdcf32f83d0b615db5d1b4f9e93a35ae7eecd413a19b3c93a0c3cde0fde272458.
//
// Solidity: event authorDividend(address _od, uint256 _amount)
func (_SGBFERC721 *SGBFERC721Filterer) FilterAuthorDividend(opts *bind.FilterOpts) (*SGBFERC721AuthorDividendIterator, error) {

	logs, sub, err := _SGBFERC721.contract.FilterLogs(opts, "authorDividend")
	if err != nil {
		return nil, err
	}
	return &SGBFERC721AuthorDividendIterator{contract: _SGBFERC721.contract, event: "authorDividend", logs: logs, sub: sub}, nil
}

// WatchAuthorDividend is a free log subscription operation binding the contract event 0xdcf32f83d0b615db5d1b4f9e93a35ae7eecd413a19b3c93a0c3cde0fde272458.
//
// Solidity: event authorDividend(address _od, uint256 _amount)
func (_SGBFERC721 *SGBFERC721Filterer) WatchAuthorDividend(opts *bind.WatchOpts, sink chan<- *SGBFERC721AuthorDividend) (event.Subscription, error) {

	logs, sub, err := _SGBFERC721.contract.WatchLogs(opts, "authorDividend")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SGBFERC721AuthorDividend)
				if err := _SGBFERC721.contract.UnpackLog(event, "authorDividend", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAuthorDividend is a log parse operation binding the contract event 0xdcf32f83d0b615db5d1b4f9e93a35ae7eecd413a19b3c93a0c3cde0fde272458.
//
// Solidity: event authorDividend(address _od, uint256 _amount)
func (_SGBFERC721 *SGBFERC721Filterer) ParseAuthorDividend(log types.Log) (*SGBFERC721AuthorDividend, error) {
	event := new(SGBFERC721AuthorDividend)
	if err := _SGBFERC721.contract.UnpackLog(event, "authorDividend", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SGBFERC721OwnerDividendIterator is returned from FilterOwnerDividend and is used to iterate over the raw logs and unpacked data for OwnerDividend events raised by the SGBFERC721 contract.
type SGBFERC721OwnerDividendIterator struct {
	Event *SGBFERC721OwnerDividend // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SGBFERC721OwnerDividendIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SGBFERC721OwnerDividend)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SGBFERC721OwnerDividend)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SGBFERC721OwnerDividendIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SGBFERC721OwnerDividendIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SGBFERC721OwnerDividend represents a OwnerDividend event raised by the SGBFERC721 contract.
type SGBFERC721OwnerDividend struct {
	Od     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterOwnerDividend is a free log retrieval operation binding the contract event 0x0a11769c808fe5441fef8ce036b84f5508912799293d20e39b808d22a2c0309f.
//
// Solidity: event ownerDividend(address _od, uint256 _amount)
func (_SGBFERC721 *SGBFERC721Filterer) FilterOwnerDividend(opts *bind.FilterOpts) (*SGBFERC721OwnerDividendIterator, error) {

	logs, sub, err := _SGBFERC721.contract.FilterLogs(opts, "ownerDividend")
	if err != nil {
		return nil, err
	}
	return &SGBFERC721OwnerDividendIterator{contract: _SGBFERC721.contract, event: "ownerDividend", logs: logs, sub: sub}, nil
}

// WatchOwnerDividend is a free log subscription operation binding the contract event 0x0a11769c808fe5441fef8ce036b84f5508912799293d20e39b808d22a2c0309f.
//
// Solidity: event ownerDividend(address _od, uint256 _amount)
func (_SGBFERC721 *SGBFERC721Filterer) WatchOwnerDividend(opts *bind.WatchOpts, sink chan<- *SGBFERC721OwnerDividend) (event.Subscription, error) {

	logs, sub, err := _SGBFERC721.contract.WatchLogs(opts, "ownerDividend")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SGBFERC721OwnerDividend)
				if err := _SGBFERC721.contract.UnpackLog(event, "ownerDividend", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnerDividend is a log parse operation binding the contract event 0x0a11769c808fe5441fef8ce036b84f5508912799293d20e39b808d22a2c0309f.
//
// Solidity: event ownerDividend(address _od, uint256 _amount)
func (_SGBFERC721 *SGBFERC721Filterer) ParseOwnerDividend(log types.Log) (*SGBFERC721OwnerDividend, error) {
	event := new(SGBFERC721OwnerDividend)
	if err := _SGBFERC721.contract.UnpackLog(event, "ownerDividend", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SGBFERC721SetupStatusIterator is returned from FilterSetupStatus and is used to iterate over the raw logs and unpacked data for SetupStatus events raised by the SGBFERC721 contract.
type SGBFERC721SetupStatusIterator struct {
	Event *SGBFERC721SetupStatus // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SGBFERC721SetupStatusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SGBFERC721SetupStatus)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SGBFERC721SetupStatus)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SGBFERC721SetupStatusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SGBFERC721SetupStatusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SGBFERC721SetupStatus represents a SetupStatus event raised by the SGBFERC721 contract.
type SGBFERC721SetupStatus struct {
	Report string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetupStatus is a free log retrieval operation binding the contract event 0xee2360dec80f538ea533854ba8d98b4f1cdfa230467c0f714689dc2446263773.
//
// Solidity: event setupStatus(string _report)
func (_SGBFERC721 *SGBFERC721Filterer) FilterSetupStatus(opts *bind.FilterOpts) (*SGBFERC721SetupStatusIterator, error) {

	logs, sub, err := _SGBFERC721.contract.FilterLogs(opts, "setupStatus")
	if err != nil {
		return nil, err
	}
	return &SGBFERC721SetupStatusIterator{contract: _SGBFERC721.contract, event: "setupStatus", logs: logs, sub: sub}, nil
}

// WatchSetupStatus is a free log subscription operation binding the contract event 0xee2360dec80f538ea533854ba8d98b4f1cdfa230467c0f714689dc2446263773.
//
// Solidity: event setupStatus(string _report)
func (_SGBFERC721 *SGBFERC721Filterer) WatchSetupStatus(opts *bind.WatchOpts, sink chan<- *SGBFERC721SetupStatus) (event.Subscription, error) {

	logs, sub, err := _SGBFERC721.contract.WatchLogs(opts, "setupStatus")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SGBFERC721SetupStatus)
				if err := _SGBFERC721.contract.UnpackLog(event, "setupStatus", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetupStatus is a log parse operation binding the contract event 0xee2360dec80f538ea533854ba8d98b4f1cdfa230467c0f714689dc2446263773.
//
// Solidity: event setupStatus(string _report)
func (_SGBFERC721 *SGBFERC721Filterer) ParseSetupStatus(log types.Log) (*SGBFERC721SetupStatus, error) {
	event := new(SGBFERC721SetupStatus)
	if err := _SGBFERC721.contract.UnpackLog(event, "setupStatus", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
