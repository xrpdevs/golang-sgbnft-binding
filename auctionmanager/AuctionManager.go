package auctionmanager

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

// AMMetaData contains all meta data concerning the AM contract.
var AMMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"auctionNoWinner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"winningAmount\",\"type\":\"uint256\"}],\"name\":\"auctionWinner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"royRecipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"royPayout\",\"type\":\"uint256\"}],\"name\":\"erc2981Fee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalamount\",\"type\":\"uint256\"}],\"name\":\"increaseBid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"newBid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"listedBy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reserveAmount\",\"type\":\"uint256\"}],\"name\":\"newListing\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"losingBidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newHighBidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newHighBid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"refundedAmount\",\"type\":\"uint256\"}],\"name\":\"outBid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"paymentSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"removalFee\",\"type\":\"uint256\"}],\"name\":\"removedListing\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"auctionBids\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isRunning\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"auctionID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"bidderAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"ownerAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"finishTime\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"finalizeOnReserveMet\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"bidHistory\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"bidderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"earlyFinalize\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"removeListing\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"salesFee\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"bidIncrement\",\"type\":\"uint8\"}],\"name\":\"configureFeePercentages\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserve\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"finalizeOnReserveMet\",\"type\":\"bool\"}],\"name\":\"createAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"earlyFinalizeFeePercentage\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"finalizeAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"finalizeEarly\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getHighBid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"bidderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getOwnerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"tokenOwner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"heldBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_heldBalance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"minNextBidAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"minBid\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextBidMinPercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"removeListingBeforeEndTime\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removeListingFeePercentage\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"salesFeePercentage\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"showFeesAccrued\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"acFees\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"submitNewBid\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AMABI is the input ABI used to generate the binding from.
// Deprecated: Use AMMetaData.ABI instead.
var AMABI = AMMetaData.ABI

// AM is an auto generated Go binding around an Ethereum contract.
type AM struct {
	AMCaller     // Read-only binding to the contract
	AMTransactor // Write-only binding to the contract
	AMFilterer   // Log filterer for contract events
}

// AMCaller is an auto generated read-only Go binding around an Ethereum contract.
type AMCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AMTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AMTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AMFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AMFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AMSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AMSession struct {
	Contract     *AM               // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AMCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AMCallerSession struct {
	Contract *AMCaller     // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AMTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AMTransactorSession struct {
	Contract     *AMTransactor     // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AMRaw is an auto generated low-level Go binding around an Ethereum contract.
type AMRaw struct {
	Contract *AM // Generic contract binding to access the raw methods on
}

// AMCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AMCallerRaw struct {
	Contract *AMCaller // Generic read-only contract binding to access the raw methods on
}

// AMTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AMTransactorRaw struct {
	Contract *AMTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAM creates a new instance of AM, bound to a specific deployed contract.
func NewAM(address common.Address, backend bind.ContractBackend) (*AM, error) {
	contract, err := bindAM(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AM{AMCaller: AMCaller{contract: contract}, AMTransactor: AMTransactor{contract: contract}, AMFilterer: AMFilterer{contract: contract}}, nil
}

// NewAMCaller creates a new read-only instance of AM, bound to a specific deployed contract.
func NewAMCaller(address common.Address, caller bind.ContractCaller) (*AMCaller, error) {
	contract, err := bindAM(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AMCaller{contract: contract}, nil
}

// NewAMTransactor creates a new write-only instance of AM, bound to a specific deployed contract.
func NewAMTransactor(address common.Address, transactor bind.ContractTransactor) (*AMTransactor, error) {
	contract, err := bindAM(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AMTransactor{contract: contract}, nil
}

// NewAMFilterer creates a new log filterer instance of AM, bound to a specific deployed contract.
func NewAMFilterer(address common.Address, filterer bind.ContractFilterer) (*AMFilterer, error) {
	contract, err := bindAM(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AMFilterer{contract: contract}, nil
}

// bindAM binds a generic wrapper to an already deployed contract.
func bindAM(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AMABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AM *AMRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AM.Contract.AMCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AM *AMRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AM.Contract.AMTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AM *AMRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AM.Contract.AMTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AM *AMCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AM.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AM *AMTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AM.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AM *AMTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AM.Contract.contract.Transact(opts, method, params...)
}

// AuctionBids is a free data retrieval call binding the contract method 0x299a0e1e.
//
// Solidity: function auctionBids(address , uint256 ) view returns(bool isRunning, uint256 auctionID, address bidderAddress, address ownerAddress, uint256 bidAmount, uint256 timestamp, uint256 reserveAmount, uint256 finishTime, bool finalizeOnReserveMet)
func (_AM *AMCaller) AuctionBids(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	IsRunning            bool
	AuctionID            *big.Int
	BidderAddress        common.Address
	OwnerAddress         common.Address
	BidAmount            *big.Int
	Timestamp            *big.Int
	ReserveAmount        *big.Int
	FinishTime           *big.Int
	FinalizeOnReserveMet bool
}, error) {
	var out []interface{}
	err := _AM.contract.Call(opts, &out, "auctionBids", arg0, arg1)

	outstruct := new(struct {
		IsRunning            bool
		AuctionID            *big.Int
		BidderAddress        common.Address
		OwnerAddress         common.Address
		BidAmount            *big.Int
		Timestamp            *big.Int
		ReserveAmount        *big.Int
		FinishTime           *big.Int
		FinalizeOnReserveMet bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsRunning = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.AuctionID = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BidderAddress = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.OwnerAddress = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.BidAmount = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.ReserveAmount = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.FinishTime = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.FinalizeOnReserveMet = *abi.ConvertType(out[8], new(bool)).(*bool)

	return *outstruct, err

}

// AuctionBids is a free data retrieval call binding the contract method 0x299a0e1e.
//
// Solidity: function auctionBids(address , uint256 ) view returns(bool isRunning, uint256 auctionID, address bidderAddress, address ownerAddress, uint256 bidAmount, uint256 timestamp, uint256 reserveAmount, uint256 finishTime, bool finalizeOnReserveMet)
func (_AM *AMSession) AuctionBids(arg0 common.Address, arg1 *big.Int) (struct {
	IsRunning            bool
	AuctionID            *big.Int
	BidderAddress        common.Address
	OwnerAddress         common.Address
	BidAmount            *big.Int
	Timestamp            *big.Int
	ReserveAmount        *big.Int
	FinishTime           *big.Int
	FinalizeOnReserveMet bool
}, error) {
	return _AM.Contract.AuctionBids(&_AM.CallOpts, arg0, arg1)
}

// AuctionBids is a free data retrieval call binding the contract method 0x299a0e1e.
//
// Solidity: function auctionBids(address , uint256 ) view returns(bool isRunning, uint256 auctionID, address bidderAddress, address ownerAddress, uint256 bidAmount, uint256 timestamp, uint256 reserveAmount, uint256 finishTime, bool finalizeOnReserveMet)
func (_AM *AMCallerSession) AuctionBids(arg0 common.Address, arg1 *big.Int) (struct {
	IsRunning            bool
	AuctionID            *big.Int
	BidderAddress        common.Address
	OwnerAddress         common.Address
	BidAmount            *big.Int
	Timestamp            *big.Int
	ReserveAmount        *big.Int
	FinishTime           *big.Int
	FinalizeOnReserveMet bool
}, error) {
	return _AM.Contract.AuctionBids(&_AM.CallOpts, arg0, arg1)
}

// BidHistory is a free data retrieval call binding the contract method 0x61ab546e.
//
// Solidity: function bidHistory(uint256 , uint256 ) view returns(uint256 timestamp, address bidderAddress, uint256 bidAmount)
func (_AM *AMCaller) BidHistory(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (struct {
	Timestamp     *big.Int
	BidderAddress common.Address
	BidAmount     *big.Int
}, error) {
	var out []interface{}
	err := _AM.contract.Call(opts, &out, "bidHistory", arg0, arg1)

	outstruct := new(struct {
		Timestamp     *big.Int
		BidderAddress common.Address
		BidAmount     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Timestamp = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.BidderAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.BidAmount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// BidHistory is a free data retrieval call binding the contract method 0x61ab546e.
//
// Solidity: function bidHistory(uint256 , uint256 ) view returns(uint256 timestamp, address bidderAddress, uint256 bidAmount)
func (_AM *AMSession) BidHistory(arg0 *big.Int, arg1 *big.Int) (struct {
	Timestamp     *big.Int
	BidderAddress common.Address
	BidAmount     *big.Int
}, error) {
	return _AM.Contract.BidHistory(&_AM.CallOpts, arg0, arg1)
}

// BidHistory is a free data retrieval call binding the contract method 0x61ab546e.
//
// Solidity: function bidHistory(uint256 , uint256 ) view returns(uint256 timestamp, address bidderAddress, uint256 bidAmount)
func (_AM *AMCallerSession) BidHistory(arg0 *big.Int, arg1 *big.Int) (struct {
	Timestamp     *big.Int
	BidderAddress common.Address
	BidAmount     *big.Int
}, error) {
	return _AM.Contract.BidHistory(&_AM.CallOpts, arg0, arg1)
}

// EarlyFinalizeFeePercentage is a free data retrieval call binding the contract method 0xbb73a6d3.
//
// Solidity: function earlyFinalizeFeePercentage() view returns(uint8)
func (_AM *AMCaller) EarlyFinalizeFeePercentage(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _AM.contract.Call(opts, &out, "earlyFinalizeFeePercentage")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// EarlyFinalizeFeePercentage is a free data retrieval call binding the contract method 0xbb73a6d3.
//
// Solidity: function earlyFinalizeFeePercentage() view returns(uint8)
func (_AM *AMSession) EarlyFinalizeFeePercentage() (uint8, error) {
	return _AM.Contract.EarlyFinalizeFeePercentage(&_AM.CallOpts)
}

// EarlyFinalizeFeePercentage is a free data retrieval call binding the contract method 0xbb73a6d3.
//
// Solidity: function earlyFinalizeFeePercentage() view returns(uint8)
func (_AM *AMCallerSession) EarlyFinalizeFeePercentage() (uint8, error) {
	return _AM.Contract.EarlyFinalizeFeePercentage(&_AM.CallOpts)
}

// GetHighBid is a free data retrieval call binding the contract method 0x4c9e732a.
//
// Solidity: function getHighBid(address nftContract, uint256 tokenId) view returns(uint256 bidAmount, address bidderAddress, uint256 timestamp)
func (_AM *AMCaller) GetHighBid(opts *bind.CallOpts, nftContract common.Address, tokenId *big.Int) (struct {
	BidAmount     *big.Int
	BidderAddress common.Address
	Timestamp     *big.Int
}, error) {
	var out []interface{}
	err := _AM.contract.Call(opts, &out, "getHighBid", nftContract, tokenId)

	outstruct := new(struct {
		BidAmount     *big.Int
		BidderAddress common.Address
		Timestamp     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BidAmount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.BidderAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Timestamp = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetHighBid is a free data retrieval call binding the contract method 0x4c9e732a.
//
// Solidity: function getHighBid(address nftContract, uint256 tokenId) view returns(uint256 bidAmount, address bidderAddress, uint256 timestamp)
func (_AM *AMSession) GetHighBid(nftContract common.Address, tokenId *big.Int) (struct {
	BidAmount     *big.Int
	BidderAddress common.Address
	Timestamp     *big.Int
}, error) {
	return _AM.Contract.GetHighBid(&_AM.CallOpts, nftContract, tokenId)
}

// GetHighBid is a free data retrieval call binding the contract method 0x4c9e732a.
//
// Solidity: function getHighBid(address nftContract, uint256 tokenId) view returns(uint256 bidAmount, address bidderAddress, uint256 timestamp)
func (_AM *AMCallerSession) GetHighBid(nftContract common.Address, tokenId *big.Int) (struct {
	BidAmount     *big.Int
	BidderAddress common.Address
	Timestamp     *big.Int
}, error) {
	return _AM.Contract.GetHighBid(&_AM.CallOpts, nftContract, tokenId)
}

// GetOwnerOf is a free data retrieval call binding the contract method 0xd5d03e21.
//
// Solidity: function getOwnerOf(address nftContract, uint256 tokenId) view returns(address tokenOwner)
func (_AM *AMCaller) GetOwnerOf(opts *bind.CallOpts, nftContract common.Address, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AM.contract.Call(opts, &out, "getOwnerOf", nftContract, tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwnerOf is a free data retrieval call binding the contract method 0xd5d03e21.
//
// Solidity: function getOwnerOf(address nftContract, uint256 tokenId) view returns(address tokenOwner)
func (_AM *AMSession) GetOwnerOf(nftContract common.Address, tokenId *big.Int) (common.Address, error) {
	return _AM.Contract.GetOwnerOf(&_AM.CallOpts, nftContract, tokenId)
}

// GetOwnerOf is a free data retrieval call binding the contract method 0xd5d03e21.
//
// Solidity: function getOwnerOf(address nftContract, uint256 tokenId) view returns(address tokenOwner)
func (_AM *AMCallerSession) GetOwnerOf(nftContract common.Address, tokenId *big.Int) (common.Address, error) {
	return _AM.Contract.GetOwnerOf(&_AM.CallOpts, nftContract, tokenId)
}

// HeldBalance is a free data retrieval call binding the contract method 0x6673c255.
//
// Solidity: function heldBalance() view returns(uint256 _heldBalance)
func (_AM *AMCaller) HeldBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AM.contract.Call(opts, &out, "heldBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HeldBalance is a free data retrieval call binding the contract method 0x6673c255.
//
// Solidity: function heldBalance() view returns(uint256 _heldBalance)
func (_AM *AMSession) HeldBalance() (*big.Int, error) {
	return _AM.Contract.HeldBalance(&_AM.CallOpts)
}

// HeldBalance is a free data retrieval call binding the contract method 0x6673c255.
//
// Solidity: function heldBalance() view returns(uint256 _heldBalance)
func (_AM *AMCallerSession) HeldBalance() (*big.Int, error) {
	return _AM.Contract.HeldBalance(&_AM.CallOpts)
}

// MinNextBidAmount is a free data retrieval call binding the contract method 0x2ec0bf96.
//
// Solidity: function minNextBidAmount(address nftContract, uint256 tokenId) view returns(uint256 minBid)
func (_AM *AMCaller) MinNextBidAmount(opts *bind.CallOpts, nftContract common.Address, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AM.contract.Call(opts, &out, "minNextBidAmount", nftContract, tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinNextBidAmount is a free data retrieval call binding the contract method 0x2ec0bf96.
//
// Solidity: function minNextBidAmount(address nftContract, uint256 tokenId) view returns(uint256 minBid)
func (_AM *AMSession) MinNextBidAmount(nftContract common.Address, tokenId *big.Int) (*big.Int, error) {
	return _AM.Contract.MinNextBidAmount(&_AM.CallOpts, nftContract, tokenId)
}

// MinNextBidAmount is a free data retrieval call binding the contract method 0x2ec0bf96.
//
// Solidity: function minNextBidAmount(address nftContract, uint256 tokenId) view returns(uint256 minBid)
func (_AM *AMCallerSession) MinNextBidAmount(nftContract common.Address, tokenId *big.Int) (*big.Int, error) {
	return _AM.Contract.MinNextBidAmount(&_AM.CallOpts, nftContract, tokenId)
}

// NextBidMinPercentage is a free data retrieval call binding the contract method 0x11a0c101.
//
// Solidity: function nextBidMinPercentage() view returns(uint256)
func (_AM *AMCaller) NextBidMinPercentage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AM.contract.Call(opts, &out, "nextBidMinPercentage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextBidMinPercentage is a free data retrieval call binding the contract method 0x11a0c101.
//
// Solidity: function nextBidMinPercentage() view returns(uint256)
func (_AM *AMSession) NextBidMinPercentage() (*big.Int, error) {
	return _AM.Contract.NextBidMinPercentage(&_AM.CallOpts)
}

// NextBidMinPercentage is a free data retrieval call binding the contract method 0x11a0c101.
//
// Solidity: function nextBidMinPercentage() view returns(uint256)
func (_AM *AMCallerSession) NextBidMinPercentage() (*big.Int, error) {
	return _AM.Contract.NextBidMinPercentage(&_AM.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AM *AMCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AM.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AM *AMSession) Owner() (common.Address, error) {
	return _AM.Contract.Owner(&_AM.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AM *AMCallerSession) Owner() (common.Address, error) {
	return _AM.Contract.Owner(&_AM.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_AM *AMCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _AM.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_AM *AMSession) Paused() (bool, error) {
	return _AM.Contract.Paused(&_AM.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_AM *AMCallerSession) Paused() (bool, error) {
	return _AM.Contract.Paused(&_AM.CallOpts)
}

// RemoveListingFeePercentage is a free data retrieval call binding the contract method 0xf9220d79.
//
// Solidity: function removeListingFeePercentage() view returns(uint8)
func (_AM *AMCaller) RemoveListingFeePercentage(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _AM.contract.Call(opts, &out, "removeListingFeePercentage")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// RemoveListingFeePercentage is a free data retrieval call binding the contract method 0xf9220d79.
//
// Solidity: function removeListingFeePercentage() view returns(uint8)
func (_AM *AMSession) RemoveListingFeePercentage() (uint8, error) {
	return _AM.Contract.RemoveListingFeePercentage(&_AM.CallOpts)
}

// RemoveListingFeePercentage is a free data retrieval call binding the contract method 0xf9220d79.
//
// Solidity: function removeListingFeePercentage() view returns(uint8)
func (_AM *AMCallerSession) RemoveListingFeePercentage() (uint8, error) {
	return _AM.Contract.RemoveListingFeePercentage(&_AM.CallOpts)
}

// SalesFeePercentage is a free data retrieval call binding the contract method 0x25b1f021.
//
// Solidity: function salesFeePercentage() view returns(uint8)
func (_AM *AMCaller) SalesFeePercentage(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _AM.contract.Call(opts, &out, "salesFeePercentage")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// SalesFeePercentage is a free data retrieval call binding the contract method 0x25b1f021.
//
// Solidity: function salesFeePercentage() view returns(uint8)
func (_AM *AMSession) SalesFeePercentage() (uint8, error) {
	return _AM.Contract.SalesFeePercentage(&_AM.CallOpts)
}

// SalesFeePercentage is a free data retrieval call binding the contract method 0x25b1f021.
//
// Solidity: function salesFeePercentage() view returns(uint8)
func (_AM *AMCallerSession) SalesFeePercentage() (uint8, error) {
	return _AM.Contract.SalesFeePercentage(&_AM.CallOpts)
}

// ConfigureFeePercentages is a paid mutator transaction binding the contract method 0x7557d7b8.
//
// Solidity: function configureFeePercentages(uint8 earlyFinalize, uint8 removeListing, uint8 salesFee, uint8 bidIncrement) returns()
func (_AM *AMTransactor) ConfigureFeePercentages(opts *bind.TransactOpts, earlyFinalize uint8, removeListing uint8, salesFee uint8, bidIncrement uint8) (*types.Transaction, error) {
	return _AM.contract.Transact(opts, "configureFeePercentages", earlyFinalize, removeListing, salesFee, bidIncrement)
}

// ConfigureFeePercentages is a paid mutator transaction binding the contract method 0x7557d7b8.
//
// Solidity: function configureFeePercentages(uint8 earlyFinalize, uint8 removeListing, uint8 salesFee, uint8 bidIncrement) returns()
func (_AM *AMSession) ConfigureFeePercentages(earlyFinalize uint8, removeListing uint8, salesFee uint8, bidIncrement uint8) (*types.Transaction, error) {
	return _AM.Contract.ConfigureFeePercentages(&_AM.TransactOpts, earlyFinalize, removeListing, salesFee, bidIncrement)
}

// ConfigureFeePercentages is a paid mutator transaction binding the contract method 0x7557d7b8.
//
// Solidity: function configureFeePercentages(uint8 earlyFinalize, uint8 removeListing, uint8 salesFee, uint8 bidIncrement) returns()
func (_AM *AMTransactorSession) ConfigureFeePercentages(earlyFinalize uint8, removeListing uint8, salesFee uint8, bidIncrement uint8) (*types.Transaction, error) {
	return _AM.Contract.ConfigureFeePercentages(&_AM.TransactOpts, earlyFinalize, removeListing, salesFee, bidIncrement)
}

// CreateAuction is a paid mutator transaction binding the contract method 0x41370f0d.
//
// Solidity: function createAuction(address nftContract, uint256 tokenId, uint256 duration, uint256 reserve, bool finalizeOnReserveMet) returns()
func (_AM *AMTransactor) CreateAuction(opts *bind.TransactOpts, nftContract common.Address, tokenId *big.Int, duration *big.Int, reserve *big.Int, finalizeOnReserveMet bool) (*types.Transaction, error) {
	return _AM.contract.Transact(opts, "createAuction", nftContract, tokenId, duration, reserve, finalizeOnReserveMet)
}

// CreateAuction is a paid mutator transaction binding the contract method 0x41370f0d.
//
// Solidity: function createAuction(address nftContract, uint256 tokenId, uint256 duration, uint256 reserve, bool finalizeOnReserveMet) returns()
func (_AM *AMSession) CreateAuction(nftContract common.Address, tokenId *big.Int, duration *big.Int, reserve *big.Int, finalizeOnReserveMet bool) (*types.Transaction, error) {
	return _AM.Contract.CreateAuction(&_AM.TransactOpts, nftContract, tokenId, duration, reserve, finalizeOnReserveMet)
}

// CreateAuction is a paid mutator transaction binding the contract method 0x41370f0d.
//
// Solidity: function createAuction(address nftContract, uint256 tokenId, uint256 duration, uint256 reserve, bool finalizeOnReserveMet) returns()
func (_AM *AMTransactorSession) CreateAuction(nftContract common.Address, tokenId *big.Int, duration *big.Int, reserve *big.Int, finalizeOnReserveMet bool) (*types.Transaction, error) {
	return _AM.Contract.CreateAuction(&_AM.TransactOpts, nftContract, tokenId, duration, reserve, finalizeOnReserveMet)
}

// FinalizeAuction is a paid mutator transaction binding the contract method 0x68905116.
//
// Solidity: function finalizeAuction(address nftContract, uint256 tokenId) returns()
func (_AM *AMTransactor) FinalizeAuction(opts *bind.TransactOpts, nftContract common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _AM.contract.Transact(opts, "finalizeAuction", nftContract, tokenId)
}

// FinalizeAuction is a paid mutator transaction binding the contract method 0x68905116.
//
// Solidity: function finalizeAuction(address nftContract, uint256 tokenId) returns()
func (_AM *AMSession) FinalizeAuction(nftContract common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _AM.Contract.FinalizeAuction(&_AM.TransactOpts, nftContract, tokenId)
}

// FinalizeAuction is a paid mutator transaction binding the contract method 0x68905116.
//
// Solidity: function finalizeAuction(address nftContract, uint256 tokenId) returns()
func (_AM *AMTransactorSession) FinalizeAuction(nftContract common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _AM.Contract.FinalizeAuction(&_AM.TransactOpts, nftContract, tokenId)
}

// FinalizeEarly is a paid mutator transaction binding the contract method 0x8ea5c3d9.
//
// Solidity: function finalizeEarly(address nftContract, uint256 tokenId) returns()
func (_AM *AMTransactor) FinalizeEarly(opts *bind.TransactOpts, nftContract common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _AM.contract.Transact(opts, "finalizeEarly", nftContract, tokenId)
}

// FinalizeEarly is a paid mutator transaction binding the contract method 0x8ea5c3d9.
//
// Solidity: function finalizeEarly(address nftContract, uint256 tokenId) returns()
func (_AM *AMSession) FinalizeEarly(nftContract common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _AM.Contract.FinalizeEarly(&_AM.TransactOpts, nftContract, tokenId)
}

// FinalizeEarly is a paid mutator transaction binding the contract method 0x8ea5c3d9.
//
// Solidity: function finalizeEarly(address nftContract, uint256 tokenId) returns()
func (_AM *AMTransactorSession) FinalizeEarly(nftContract common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _AM.Contract.FinalizeEarly(&_AM.TransactOpts, nftContract, tokenId)
}

// RemoveListingBeforeEndTime is a paid mutator transaction binding the contract method 0xf6e1c2d0.
//
// Solidity: function removeListingBeforeEndTime(address nftContract, uint256 tokenId) payable returns()
func (_AM *AMTransactor) RemoveListingBeforeEndTime(opts *bind.TransactOpts, nftContract common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _AM.contract.Transact(opts, "removeListingBeforeEndTime", nftContract, tokenId)
}

// RemoveListingBeforeEndTime is a paid mutator transaction binding the contract method 0xf6e1c2d0.
//
// Solidity: function removeListingBeforeEndTime(address nftContract, uint256 tokenId) payable returns()
func (_AM *AMSession) RemoveListingBeforeEndTime(nftContract common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _AM.Contract.RemoveListingBeforeEndTime(&_AM.TransactOpts, nftContract, tokenId)
}

// RemoveListingBeforeEndTime is a paid mutator transaction binding the contract method 0xf6e1c2d0.
//
// Solidity: function removeListingBeforeEndTime(address nftContract, uint256 tokenId) payable returns()
func (_AM *AMTransactorSession) RemoveListingBeforeEndTime(nftContract common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _AM.Contract.RemoveListingBeforeEndTime(&_AM.TransactOpts, nftContract, tokenId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AM *AMTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AM.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AM *AMSession) RenounceOwnership() (*types.Transaction, error) {
	return _AM.Contract.RenounceOwnership(&_AM.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AM *AMTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AM.Contract.RenounceOwnership(&_AM.TransactOpts)
}

// Setup is a paid mutator transaction binding the contract method 0xba0bba40.
//
// Solidity: function setup() returns()
func (_AM *AMTransactor) Setup(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AM.contract.Transact(opts, "setup")
}

// Setup is a paid mutator transaction binding the contract method 0xba0bba40.
//
// Solidity: function setup() returns()
func (_AM *AMSession) Setup() (*types.Transaction, error) {
	return _AM.Contract.Setup(&_AM.TransactOpts)
}

// Setup is a paid mutator transaction binding the contract method 0xba0bba40.
//
// Solidity: function setup() returns()
func (_AM *AMTransactorSession) Setup() (*types.Transaction, error) {
	return _AM.Contract.Setup(&_AM.TransactOpts)
}

// ShowFeesAccrued is a paid mutator transaction binding the contract method 0xffac78b4.
//
// Solidity: function showFeesAccrued() returns(uint256 acFees)
func (_AM *AMTransactor) ShowFeesAccrued(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AM.contract.Transact(opts, "showFeesAccrued")
}

// ShowFeesAccrued is a paid mutator transaction binding the contract method 0xffac78b4.
//
// Solidity: function showFeesAccrued() returns(uint256 acFees)
func (_AM *AMSession) ShowFeesAccrued() (*types.Transaction, error) {
	return _AM.Contract.ShowFeesAccrued(&_AM.TransactOpts)
}

// ShowFeesAccrued is a paid mutator transaction binding the contract method 0xffac78b4.
//
// Solidity: function showFeesAccrued() returns(uint256 acFees)
func (_AM *AMTransactorSession) ShowFeesAccrued() (*types.Transaction, error) {
	return _AM.Contract.ShowFeesAccrued(&_AM.TransactOpts)
}

// SubmitNewBid is a paid mutator transaction binding the contract method 0x326557bc.
//
// Solidity: function submitNewBid(address nftContract, uint256 tokenId) payable returns()
func (_AM *AMTransactor) SubmitNewBid(opts *bind.TransactOpts, nftContract common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _AM.contract.Transact(opts, "submitNewBid", nftContract, tokenId)
}

// SubmitNewBid is a paid mutator transaction binding the contract method 0x326557bc.
//
// Solidity: function submitNewBid(address nftContract, uint256 tokenId) payable returns()
func (_AM *AMSession) SubmitNewBid(nftContract common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _AM.Contract.SubmitNewBid(&_AM.TransactOpts, nftContract, tokenId)
}

// SubmitNewBid is a paid mutator transaction binding the contract method 0x326557bc.
//
// Solidity: function submitNewBid(address nftContract, uint256 tokenId) payable returns()
func (_AM *AMTransactorSession) SubmitNewBid(nftContract common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _AM.Contract.SubmitNewBid(&_AM.TransactOpts, nftContract, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AM *AMTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AM.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AM *AMSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AM.Contract.TransferOwnership(&_AM.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AM *AMTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AM.Contract.TransferOwnership(&_AM.TransactOpts, newOwner)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x853828b6.
//
// Solidity: function withdrawAll() returns()
func (_AM *AMTransactor) WithdrawAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AM.contract.Transact(opts, "withdrawAll")
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x853828b6.
//
// Solidity: function withdrawAll() returns()
func (_AM *AMSession) WithdrawAll() (*types.Transaction, error) {
	return _AM.Contract.WithdrawAll(&_AM.TransactOpts)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x853828b6.
//
// Solidity: function withdrawAll() returns()
func (_AM *AMTransactorSession) WithdrawAll() (*types.Transaction, error) {
	return _AM.Contract.WithdrawAll(&_AM.TransactOpts)
}

// AMOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AM contract.
type AMOwnershipTransferredIterator struct {
	Event *AMOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AMOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AMOwnershipTransferred)
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
		it.Event = new(AMOwnershipTransferred)
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
func (it *AMOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AMOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AMOwnershipTransferred represents a OwnershipTransferred event raised by the AM contract.
type AMOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AM *AMFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AMOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AM.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AMOwnershipTransferredIterator{contract: _AM.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AM *AMFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AMOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AM.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AMOwnershipTransferred)
				if err := _AM.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_AM *AMFilterer) ParseOwnershipTransferred(log types.Log) (*AMOwnershipTransferred, error) {
	event := new(AMOwnershipTransferred)
	if err := _AM.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AMPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the AM contract.
type AMPausedIterator struct {
	Event *AMPaused // Event containing the contract specifics and raw log

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
func (it *AMPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AMPaused)
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
		it.Event = new(AMPaused)
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
func (it *AMPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AMPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AMPaused represents a Paused event raised by the AM contract.
type AMPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_AM *AMFilterer) FilterPaused(opts *bind.FilterOpts) (*AMPausedIterator, error) {

	logs, sub, err := _AM.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &AMPausedIterator{contract: _AM.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_AM *AMFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *AMPaused) (event.Subscription, error) {

	logs, sub, err := _AM.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AMPaused)
				if err := _AM.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_AM *AMFilterer) ParsePaused(log types.Log) (*AMPaused, error) {
	event := new(AMPaused)
	if err := _AM.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AMUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the AM contract.
type AMUnpausedIterator struct {
	Event *AMUnpaused // Event containing the contract specifics and raw log

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
func (it *AMUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AMUnpaused)
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
		it.Event = new(AMUnpaused)
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
func (it *AMUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AMUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AMUnpaused represents a Unpaused event raised by the AM contract.
type AMUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_AM *AMFilterer) FilterUnpaused(opts *bind.FilterOpts) (*AMUnpausedIterator, error) {

	logs, sub, err := _AM.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &AMUnpausedIterator{contract: _AM.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_AM *AMFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *AMUnpaused) (event.Subscription, error) {

	logs, sub, err := _AM.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AMUnpaused)
				if err := _AM.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_AM *AMFilterer) ParseUnpaused(log types.Log) (*AMUnpaused, error) {
	event := new(AMUnpaused)
	if err := _AM.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AMAuctionNoWinnerIterator is returned from FilterAuctionNoWinner and is used to iterate over the raw logs and unpacked data for AuctionNoWinner events raised by the AM contract.
type AMAuctionNoWinnerIterator struct {
	Event *AMAuctionNoWinner // Event containing the contract specifics and raw log

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
func (it *AMAuctionNoWinnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AMAuctionNoWinner)
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
		it.Event = new(AMAuctionNoWinner)
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
func (it *AMAuctionNoWinnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AMAuctionNoWinnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AMAuctionNoWinner represents a AuctionNoWinner event raised by the AM contract.
type AMAuctionNoWinner struct {
	AuctionId   *big.Int
	NftContract common.Address
	TokenId     *big.Int
	Reason      string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAuctionNoWinner is a free log retrieval operation binding the contract event 0xbf6ba96977596686e6a419d488d0a93ea4d5fa899ad16112981994d3d2dd5336.
//
// Solidity: event auctionNoWinner(uint256 indexed auctionId, address nftContract, uint256 tokenId, string reason)
func (_AM *AMFilterer) FilterAuctionNoWinner(opts *bind.FilterOpts, auctionId []*big.Int) (*AMAuctionNoWinnerIterator, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}

	logs, sub, err := _AM.contract.FilterLogs(opts, "auctionNoWinner", auctionIdRule)
	if err != nil {
		return nil, err
	}
	return &AMAuctionNoWinnerIterator{contract: _AM.contract, event: "auctionNoWinner", logs: logs, sub: sub}, nil
}

// WatchAuctionNoWinner is a free log subscription operation binding the contract event 0xbf6ba96977596686e6a419d488d0a93ea4d5fa899ad16112981994d3d2dd5336.
//
// Solidity: event auctionNoWinner(uint256 indexed auctionId, address nftContract, uint256 tokenId, string reason)
func (_AM *AMFilterer) WatchAuctionNoWinner(opts *bind.WatchOpts, sink chan<- *AMAuctionNoWinner, auctionId []*big.Int) (event.Subscription, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}

	logs, sub, err := _AM.contract.WatchLogs(opts, "auctionNoWinner", auctionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AMAuctionNoWinner)
				if err := _AM.contract.UnpackLog(event, "auctionNoWinner", log); err != nil {
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

// ParseAuctionNoWinner is a log parse operation binding the contract event 0xbf6ba96977596686e6a419d488d0a93ea4d5fa899ad16112981994d3d2dd5336.
//
// Solidity: event auctionNoWinner(uint256 indexed auctionId, address nftContract, uint256 tokenId, string reason)
func (_AM *AMFilterer) ParseAuctionNoWinner(log types.Log) (*AMAuctionNoWinner, error) {
	event := new(AMAuctionNoWinner)
	if err := _AM.contract.UnpackLog(event, "auctionNoWinner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AMAuctionWinnerIterator is returned from FilterAuctionWinner and is used to iterate over the raw logs and unpacked data for AuctionWinner events raised by the AM contract.
type AMAuctionWinnerIterator struct {
	Event *AMAuctionWinner // Event containing the contract specifics and raw log

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
func (it *AMAuctionWinnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AMAuctionWinner)
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
		it.Event = new(AMAuctionWinner)
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
func (it *AMAuctionWinnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AMAuctionWinnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AMAuctionWinner represents a AuctionWinner event raised by the AM contract.
type AMAuctionWinner struct {
	AuctionId     *big.Int
	NftContract   common.Address
	TokenId       *big.Int
	Bidder        common.Address
	WinningAmount *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAuctionWinner is a free log retrieval operation binding the contract event 0x2d48ab58ab9be76f496342bbe284d9573ab2758fb16b357730bd11f527d099c4.
//
// Solidity: event auctionWinner(uint256 indexed auctionId, address nftContract, uint256 tokenId, address bidder, uint256 winningAmount)
func (_AM *AMFilterer) FilterAuctionWinner(opts *bind.FilterOpts, auctionId []*big.Int) (*AMAuctionWinnerIterator, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}

	logs, sub, err := _AM.contract.FilterLogs(opts, "auctionWinner", auctionIdRule)
	if err != nil {
		return nil, err
	}
	return &AMAuctionWinnerIterator{contract: _AM.contract, event: "auctionWinner", logs: logs, sub: sub}, nil
}

// WatchAuctionWinner is a free log subscription operation binding the contract event 0x2d48ab58ab9be76f496342bbe284d9573ab2758fb16b357730bd11f527d099c4.
//
// Solidity: event auctionWinner(uint256 indexed auctionId, address nftContract, uint256 tokenId, address bidder, uint256 winningAmount)
func (_AM *AMFilterer) WatchAuctionWinner(opts *bind.WatchOpts, sink chan<- *AMAuctionWinner, auctionId []*big.Int) (event.Subscription, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}

	logs, sub, err := _AM.contract.WatchLogs(opts, "auctionWinner", auctionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AMAuctionWinner)
				if err := _AM.contract.UnpackLog(event, "auctionWinner", log); err != nil {
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

// ParseAuctionWinner is a log parse operation binding the contract event 0x2d48ab58ab9be76f496342bbe284d9573ab2758fb16b357730bd11f527d099c4.
//
// Solidity: event auctionWinner(uint256 indexed auctionId, address nftContract, uint256 tokenId, address bidder, uint256 winningAmount)
func (_AM *AMFilterer) ParseAuctionWinner(log types.Log) (*AMAuctionWinner, error) {
	event := new(AMAuctionWinner)
	if err := _AM.contract.UnpackLog(event, "auctionWinner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AMErc2981FeeIterator is returned from FilterErc2981Fee and is used to iterate over the raw logs and unpacked data for Erc2981Fee events raised by the AM contract.
type AMErc2981FeeIterator struct {
	Event *AMErc2981Fee // Event containing the contract specifics and raw log

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
func (it *AMErc2981FeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AMErc2981Fee)
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
		it.Event = new(AMErc2981Fee)
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
func (it *AMErc2981FeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AMErc2981FeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AMErc2981Fee represents a Erc2981Fee event raised by the AM contract.
type AMErc2981Fee struct {
	AuctionId    *big.Int
	NftContract  common.Address
	RoyRecipient common.Address
	RoyPayout    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterErc2981Fee is a free log retrieval operation binding the contract event 0xf8acbc058a06a3b59bc2ce79d00ea0415c2020c32ccab1c04000c0a1206c85d8.
//
// Solidity: event erc2981Fee(uint256 indexed auctionId, address nftContract, address royRecipient, uint256 royPayout)
func (_AM *AMFilterer) FilterErc2981Fee(opts *bind.FilterOpts, auctionId []*big.Int) (*AMErc2981FeeIterator, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}

	logs, sub, err := _AM.contract.FilterLogs(opts, "erc2981Fee", auctionIdRule)
	if err != nil {
		return nil, err
	}
	return &AMErc2981FeeIterator{contract: _AM.contract, event: "erc2981Fee", logs: logs, sub: sub}, nil
}

// WatchErc2981Fee is a free log subscription operation binding the contract event 0xf8acbc058a06a3b59bc2ce79d00ea0415c2020c32ccab1c04000c0a1206c85d8.
//
// Solidity: event erc2981Fee(uint256 indexed auctionId, address nftContract, address royRecipient, uint256 royPayout)
func (_AM *AMFilterer) WatchErc2981Fee(opts *bind.WatchOpts, sink chan<- *AMErc2981Fee, auctionId []*big.Int) (event.Subscription, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}

	logs, sub, err := _AM.contract.WatchLogs(opts, "erc2981Fee", auctionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AMErc2981Fee)
				if err := _AM.contract.UnpackLog(event, "erc2981Fee", log); err != nil {
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

// ParseErc2981Fee is a log parse operation binding the contract event 0xf8acbc058a06a3b59bc2ce79d00ea0415c2020c32ccab1c04000c0a1206c85d8.
//
// Solidity: event erc2981Fee(uint256 indexed auctionId, address nftContract, address royRecipient, uint256 royPayout)
func (_AM *AMFilterer) ParseErc2981Fee(log types.Log) (*AMErc2981Fee, error) {
	event := new(AMErc2981Fee)
	if err := _AM.contract.UnpackLog(event, "erc2981Fee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AMIncreaseBidIterator is returned from FilterIncreaseBid and is used to iterate over the raw logs and unpacked data for IncreaseBid events raised by the AM contract.
type AMIncreaseBidIterator struct {
	Event *AMIncreaseBid // Event containing the contract specifics and raw log

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
func (it *AMIncreaseBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AMIncreaseBid)
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
		it.Event = new(AMIncreaseBid)
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
func (it *AMIncreaseBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AMIncreaseBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AMIncreaseBid represents a IncreaseBid event raised by the AM contract.
type AMIncreaseBid struct {
	AuctionId   *big.Int
	NftContract common.Address
	TokenId     *big.Int
	Bidder      common.Address
	Totalamount *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterIncreaseBid is a free log retrieval operation binding the contract event 0xef69aef98778085413c55d2eb5c3eea407679c3f613ee0ed7a01713d5ee5b805.
//
// Solidity: event increaseBid(uint256 indexed auctionId, address nftContract, uint256 tokenId, address bidder, uint256 totalamount)
func (_AM *AMFilterer) FilterIncreaseBid(opts *bind.FilterOpts, auctionId []*big.Int) (*AMIncreaseBidIterator, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}

	logs, sub, err := _AM.contract.FilterLogs(opts, "increaseBid", auctionIdRule)
	if err != nil {
		return nil, err
	}
	return &AMIncreaseBidIterator{contract: _AM.contract, event: "increaseBid", logs: logs, sub: sub}, nil
}

// WatchIncreaseBid is a free log subscription operation binding the contract event 0xef69aef98778085413c55d2eb5c3eea407679c3f613ee0ed7a01713d5ee5b805.
//
// Solidity: event increaseBid(uint256 indexed auctionId, address nftContract, uint256 tokenId, address bidder, uint256 totalamount)
func (_AM *AMFilterer) WatchIncreaseBid(opts *bind.WatchOpts, sink chan<- *AMIncreaseBid, auctionId []*big.Int) (event.Subscription, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}

	logs, sub, err := _AM.contract.WatchLogs(opts, "increaseBid", auctionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AMIncreaseBid)
				if err := _AM.contract.UnpackLog(event, "increaseBid", log); err != nil {
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

// ParseIncreaseBid is a log parse operation binding the contract event 0xef69aef98778085413c55d2eb5c3eea407679c3f613ee0ed7a01713d5ee5b805.
//
// Solidity: event increaseBid(uint256 indexed auctionId, address nftContract, uint256 tokenId, address bidder, uint256 totalamount)
func (_AM *AMFilterer) ParseIncreaseBid(log types.Log) (*AMIncreaseBid, error) {
	event := new(AMIncreaseBid)
	if err := _AM.contract.UnpackLog(event, "increaseBid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AMNewBidIterator is returned from FilterNewBid and is used to iterate over the raw logs and unpacked data for NewBid events raised by the AM contract.
type AMNewBidIterator struct {
	Event *AMNewBid // Event containing the contract specifics and raw log

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
func (it *AMNewBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AMNewBid)
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
		it.Event = new(AMNewBid)
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
func (it *AMNewBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AMNewBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AMNewBid represents a NewBid event raised by the AM contract.
type AMNewBid struct {
	AuctionId   *big.Int
	NftContract common.Address
	TokenId     *big.Int
	Bidder      common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewBid is a free log retrieval operation binding the contract event 0xcbc020507172138f14c003a554717cb8f4f65eda7de7930906d657d4ba89a8f9.
//
// Solidity: event newBid(uint256 indexed auctionId, address nftContract, uint256 tokenId, address bidder, uint256 amount)
func (_AM *AMFilterer) FilterNewBid(opts *bind.FilterOpts, auctionId []*big.Int) (*AMNewBidIterator, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}

	logs, sub, err := _AM.contract.FilterLogs(opts, "newBid", auctionIdRule)
	if err != nil {
		return nil, err
	}
	return &AMNewBidIterator{contract: _AM.contract, event: "newBid", logs: logs, sub: sub}, nil
}

// WatchNewBid is a free log subscription operation binding the contract event 0xcbc020507172138f14c003a554717cb8f4f65eda7de7930906d657d4ba89a8f9.
//
// Solidity: event newBid(uint256 indexed auctionId, address nftContract, uint256 tokenId, address bidder, uint256 amount)
func (_AM *AMFilterer) WatchNewBid(opts *bind.WatchOpts, sink chan<- *AMNewBid, auctionId []*big.Int) (event.Subscription, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}

	logs, sub, err := _AM.contract.WatchLogs(opts, "newBid", auctionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AMNewBid)
				if err := _AM.contract.UnpackLog(event, "newBid", log); err != nil {
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

// ParseNewBid is a log parse operation binding the contract event 0xcbc020507172138f14c003a554717cb8f4f65eda7de7930906d657d4ba89a8f9.
//
// Solidity: event newBid(uint256 indexed auctionId, address nftContract, uint256 tokenId, address bidder, uint256 amount)
func (_AM *AMFilterer) ParseNewBid(log types.Log) (*AMNewBid, error) {
	event := new(AMNewBid)
	if err := _AM.contract.UnpackLog(event, "newBid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AMNewListingIterator is returned from FilterNewListing and is used to iterate over the raw logs and unpacked data for NewListing events raised by the AM contract.
type AMNewListingIterator struct {
	Event *AMNewListing // Event containing the contract specifics and raw log

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
func (it *AMNewListingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AMNewListing)
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
		it.Event = new(AMNewListing)
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
func (it *AMNewListingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AMNewListingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AMNewListing represents a NewListing event raised by the AM contract.
type AMNewListing struct {
	AuctionId     *big.Int
	NftContract   common.Address
	TokenId       *big.Int
	ListedBy      common.Address
	ReserveAmount *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNewListing is a free log retrieval operation binding the contract event 0x40e15dd727aca87213075dfba1b3356fab0033a3c2463ece1d5f0451dc8b0566.
//
// Solidity: event newListing(uint256 indexed auctionId, address nftContract, uint256 tokenId, address listedBy, uint256 reserveAmount)
func (_AM *AMFilterer) FilterNewListing(opts *bind.FilterOpts, auctionId []*big.Int) (*AMNewListingIterator, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}

	logs, sub, err := _AM.contract.FilterLogs(opts, "newListing", auctionIdRule)
	if err != nil {
		return nil, err
	}
	return &AMNewListingIterator{contract: _AM.contract, event: "newListing", logs: logs, sub: sub}, nil
}

// WatchNewListing is a free log subscription operation binding the contract event 0x40e15dd727aca87213075dfba1b3356fab0033a3c2463ece1d5f0451dc8b0566.
//
// Solidity: event newListing(uint256 indexed auctionId, address nftContract, uint256 tokenId, address listedBy, uint256 reserveAmount)
func (_AM *AMFilterer) WatchNewListing(opts *bind.WatchOpts, sink chan<- *AMNewListing, auctionId []*big.Int) (event.Subscription, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}

	logs, sub, err := _AM.contract.WatchLogs(opts, "newListing", auctionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AMNewListing)
				if err := _AM.contract.UnpackLog(event, "newListing", log); err != nil {
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

// ParseNewListing is a log parse operation binding the contract event 0x40e15dd727aca87213075dfba1b3356fab0033a3c2463ece1d5f0451dc8b0566.
//
// Solidity: event newListing(uint256 indexed auctionId, address nftContract, uint256 tokenId, address listedBy, uint256 reserveAmount)
func (_AM *AMFilterer) ParseNewListing(log types.Log) (*AMNewListing, error) {
	event := new(AMNewListing)
	if err := _AM.contract.UnpackLog(event, "newListing", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AMOutBidIterator is returned from FilterOutBid and is used to iterate over the raw logs and unpacked data for OutBid events raised by the AM contract.
type AMOutBidIterator struct {
	Event *AMOutBid // Event containing the contract specifics and raw log

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
func (it *AMOutBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AMOutBid)
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
		it.Event = new(AMOutBid)
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
func (it *AMOutBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AMOutBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AMOutBid represents a OutBid event raised by the AM contract.
type AMOutBid struct {
	AuctionId      *big.Int
	NftContract    common.Address
	TokenId        *big.Int
	LosingBidder   common.Address
	NewHighBidder  common.Address
	NewHighBid     *big.Int
	RefundedAmount *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOutBid is a free log retrieval operation binding the contract event 0xe47c8ee2eb139c89965e1648783acc224523d728651e66bad1b7b2a4c9d29437.
//
// Solidity: event outBid(uint256 indexed auctionId, address nftContract, uint256 tokenId, address losingBidder, address newHighBidder, uint256 newHighBid, uint256 refundedAmount)
func (_AM *AMFilterer) FilterOutBid(opts *bind.FilterOpts, auctionId []*big.Int) (*AMOutBidIterator, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}

	logs, sub, err := _AM.contract.FilterLogs(opts, "outBid", auctionIdRule)
	if err != nil {
		return nil, err
	}
	return &AMOutBidIterator{contract: _AM.contract, event: "outBid", logs: logs, sub: sub}, nil
}

// WatchOutBid is a free log subscription operation binding the contract event 0xe47c8ee2eb139c89965e1648783acc224523d728651e66bad1b7b2a4c9d29437.
//
// Solidity: event outBid(uint256 indexed auctionId, address nftContract, uint256 tokenId, address losingBidder, address newHighBidder, uint256 newHighBid, uint256 refundedAmount)
func (_AM *AMFilterer) WatchOutBid(opts *bind.WatchOpts, sink chan<- *AMOutBid, auctionId []*big.Int) (event.Subscription, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}

	logs, sub, err := _AM.contract.WatchLogs(opts, "outBid", auctionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AMOutBid)
				if err := _AM.contract.UnpackLog(event, "outBid", log); err != nil {
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

// ParseOutBid is a log parse operation binding the contract event 0xe47c8ee2eb139c89965e1648783acc224523d728651e66bad1b7b2a4c9d29437.
//
// Solidity: event outBid(uint256 indexed auctionId, address nftContract, uint256 tokenId, address losingBidder, address newHighBidder, uint256 newHighBid, uint256 refundedAmount)
func (_AM *AMFilterer) ParseOutBid(log types.Log) (*AMOutBid, error) {
	event := new(AMOutBid)
	if err := _AM.contract.UnpackLog(event, "outBid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AMPaymentSentIterator is returned from FilterPaymentSent and is used to iterate over the raw logs and unpacked data for PaymentSent events raised by the AM contract.
type AMPaymentSentIterator struct {
	Event *AMPaymentSent // Event containing the contract specifics and raw log

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
func (it *AMPaymentSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AMPaymentSent)
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
		it.Event = new(AMPaymentSent)
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
func (it *AMPaymentSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AMPaymentSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AMPaymentSent represents a PaymentSent event raised by the AM contract.
type AMPaymentSent struct {
	AuctionId *big.Int
	Payee     common.Address
	Amount    *big.Int
	Reason    string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPaymentSent is a free log retrieval operation binding the contract event 0x2f6183933735cb1c3926e8aac0944e51f79cf0793368400136722ebd82afe69d.
//
// Solidity: event paymentSent(uint256 indexed auctionId, address payee, uint256 amount, string reason)
func (_AM *AMFilterer) FilterPaymentSent(opts *bind.FilterOpts, auctionId []*big.Int) (*AMPaymentSentIterator, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}

	logs, sub, err := _AM.contract.FilterLogs(opts, "paymentSent", auctionIdRule)
	if err != nil {
		return nil, err
	}
	return &AMPaymentSentIterator{contract: _AM.contract, event: "paymentSent", logs: logs, sub: sub}, nil
}

// WatchPaymentSent is a free log subscription operation binding the contract event 0x2f6183933735cb1c3926e8aac0944e51f79cf0793368400136722ebd82afe69d.
//
// Solidity: event paymentSent(uint256 indexed auctionId, address payee, uint256 amount, string reason)
func (_AM *AMFilterer) WatchPaymentSent(opts *bind.WatchOpts, sink chan<- *AMPaymentSent, auctionId []*big.Int) (event.Subscription, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}

	logs, sub, err := _AM.contract.WatchLogs(opts, "paymentSent", auctionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AMPaymentSent)
				if err := _AM.contract.UnpackLog(event, "paymentSent", log); err != nil {
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

// ParsePaymentSent is a log parse operation binding the contract event 0x2f6183933735cb1c3926e8aac0944e51f79cf0793368400136722ebd82afe69d.
//
// Solidity: event paymentSent(uint256 indexed auctionId, address payee, uint256 amount, string reason)
func (_AM *AMFilterer) ParsePaymentSent(log types.Log) (*AMPaymentSent, error) {
	event := new(AMPaymentSent)
	if err := _AM.contract.UnpackLog(event, "paymentSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AMRemovedListingIterator is returned from FilterRemovedListing and is used to iterate over the raw logs and unpacked data for RemovedListing events raised by the AM contract.
type AMRemovedListingIterator struct {
	Event *AMRemovedListing // Event containing the contract specifics and raw log

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
func (it *AMRemovedListingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AMRemovedListing)
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
		it.Event = new(AMRemovedListing)
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
func (it *AMRemovedListingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AMRemovedListingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AMRemovedListing represents a RemovedListing event raised by the AM contract.
type AMRemovedListing struct {
	AuctionId   *big.Int
	NftContract common.Address
	TokenId     *big.Int
	TokenOwner  common.Address
	RemovalFee  *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRemovedListing is a free log retrieval operation binding the contract event 0x829c04137f89db6c78df6af22fea4e44d2dbbb97d8b9c0819559c07d90e0d8f5.
//
// Solidity: event removedListing(uint256 indexed auctionId, address nftContract, uint256 tokenId, address tokenOwner, uint256 removalFee)
func (_AM *AMFilterer) FilterRemovedListing(opts *bind.FilterOpts, auctionId []*big.Int) (*AMRemovedListingIterator, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}

	logs, sub, err := _AM.contract.FilterLogs(opts, "removedListing", auctionIdRule)
	if err != nil {
		return nil, err
	}
	return &AMRemovedListingIterator{contract: _AM.contract, event: "removedListing", logs: logs, sub: sub}, nil
}

// WatchRemovedListing is a free log subscription operation binding the contract event 0x829c04137f89db6c78df6af22fea4e44d2dbbb97d8b9c0819559c07d90e0d8f5.
//
// Solidity: event removedListing(uint256 indexed auctionId, address nftContract, uint256 tokenId, address tokenOwner, uint256 removalFee)
func (_AM *AMFilterer) WatchRemovedListing(opts *bind.WatchOpts, sink chan<- *AMRemovedListing, auctionId []*big.Int) (event.Subscription, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}

	logs, sub, err := _AM.contract.WatchLogs(opts, "removedListing", auctionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AMRemovedListing)
				if err := _AM.contract.UnpackLog(event, "removedListing", log); err != nil {
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

// ParseRemovedListing is a log parse operation binding the contract event 0x829c04137f89db6c78df6af22fea4e44d2dbbb97d8b9c0819559c07d90e0d8f5.
//
// Solidity: event removedListing(uint256 indexed auctionId, address nftContract, uint256 tokenId, address tokenOwner, uint256 removalFee)
func (_AM *AMFilterer) ParseRemovedListing(log types.Log) (*AMRemovedListing, error) {
	event := new(AMRemovedListing)
	if err := _AM.contract.UnpackLog(event, "removedListing", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
