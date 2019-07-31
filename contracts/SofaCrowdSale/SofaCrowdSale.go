// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package SofaCrowdSale

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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// IERC20ABI is the input ABI used to generate the binding from.
const IERC20ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// IERC20FuncSigs maps the 4-byte function signature to its string representation.
var IERC20FuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// IERC20 is an auto generated Go binding around an Ethereum contract.
type IERC20 struct {
	IERC20Caller     // Read-only binding to the contract
	IERC20Transactor // Write-only binding to the contract
	IERC20Filterer   // Log filterer for contract events
}

// IERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20Session struct {
	Contract     *IERC20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20CallerSession struct {
	Contract *IERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20TransactorSession struct {
	Contract     *IERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20Raw struct {
	Contract *IERC20 // Generic contract binding to access the raw methods on
}

// IERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20CallerRaw struct {
	Contract *IERC20Caller // Generic read-only contract binding to access the raw methods on
}

// IERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20TransactorRaw struct {
	Contract *IERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20 creates a new instance of IERC20, bound to a specific deployed contract.
func NewIERC20(address common.Address, backend bind.ContractBackend) (*IERC20, error) {
	contract, err := bindIERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
}

// NewIERC20Caller creates a new read-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Caller(address common.Address, caller bind.ContractCaller) (*IERC20Caller, error) {
	contract, err := bindIERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Caller{contract: contract}, nil
}

// NewIERC20Transactor creates a new write-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC20Transactor, error) {
	contract, err := bindIERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Transactor{contract: contract}, nil
}

// NewIERC20Filterer creates a new log filterer instance of IERC20, bound to a specific deployed contract.
func NewIERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC20Filterer, error) {
	contract, err := bindIERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20Filterer{contract: contract}, nil
}

// bindIERC20 binds a generic wrapper to an already deployed contract.
func bindIERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.IERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_IERC20 *IERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IERC20.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_IERC20 *IERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_IERC20 *IERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_IERC20 *IERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IERC20.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_IERC20 *IERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_IERC20 *IERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_IERC20 *IERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IERC20.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_IERC20 *IERC20Session) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_IERC20 *IERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// IERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20 contract.
type IERC20ApprovalIterator struct {
	Event *IERC20Approval // Event containing the contract specifics and raw log

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
func (it *IERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Approval)
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
		it.Event = new(IERC20Approval)
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
func (it *IERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Approval represents a Approval event raised by the IERC20 contract.
type IERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20ApprovalIterator{contract: _IERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Approval)
				if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) ParseApproval(log types.Log) (*IERC20Approval, error) {
	event := new(IERC20Approval)
	if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20 contract.
type IERC20TransferIterator struct {
	Event *IERC20Transfer // Event containing the contract specifics and raw log

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
func (it *IERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Transfer)
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
		it.Event = new(IERC20Transfer)
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
func (it *IERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Transfer represents a Transfer event raised by the IERC20 contract.
type IERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20TransferIterator{contract: _IERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Transfer)
				if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) ParseTransfer(log types.Log) (*IERC20Transfer, error) {
	event := new(IERC20Transfer)
	if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
var SafeMathBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723058207bb8f8ba86fc021b9037c6eb69ab2aded7fbd3383c2fc4ef45bc264e738390ae64736f6c634300050a0032"

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// SofaCrowdSaleABI is the input ABI used to generate the binding from.
const SofaCrowdSaleABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"openForSale\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_numberOfTokens\",\"type\":\"uint256\"}],\"name\":\"buyToken\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"endSale\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenSold\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"saleOpen\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenPrice\",\"type\":\"uint256\"}],\"name\":\"resetTokenPrice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_sofaToken\",\"type\":\"address\"},{\"name\":\"_initalPrice\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_buyer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Sell\",\"type\":\"event\"}]"

// SofaCrowdSaleFuncSigs maps the 4-byte function signature to its string representation.
var SofaCrowdSaleFuncSigs = map[string]string{
	"2d296bf1": "buyToken(uint256)",
	"380d831b": "endSale()",
	"26138c36": "openForSale()",
	"fc5abac1": "resetTokenPrice(uint256)",
	"99288dbb": "saleOpen()",
	"55a373d6": "tokenContract()",
	"7ff9b596": "tokenPrice()",
	"519ee19e": "tokenSold()",
}

// SofaCrowdSaleBin is the compiled bytecode used for deploying new contracts.
var SofaCrowdSaleBin = "0x608060405234801561001057600080fd5b506040516109013803806109018339818101604052604081101561003357600080fd5b508051602090910151600080546001600160a01b03199081163317909155600180546001600160a01b0390941693909116929092179091556003556108848061007d6000396000f3fe60806040526004361061007b5760003560e01c806355a373d61161004e57806355a373d6146100f05780637ff9b5961461012157806399288dbb14610136578063fc5abac11461015f5761007b565b806326138c36146100805780632d296bf114610097578063380d831b146100b4578063519ee19e146100c9575b600080fd5b34801561008c57600080fd5b50610095610189565b005b610095600480360360208110156100ad57600080fd5b5035610242565b3480156100c057600080fd5b506100956104af565b3480156100d557600080fd5b506100de610697565b60408051918252519081900360200190f35b3480156100fc57600080fd5b5061010561069d565b604080516001600160a01b039092168252519081900360200190f35b34801561012d57600080fd5b506100de6106ac565b34801561014257600080fd5b5061014b6106b2565b604080519115158252519081900360200190f35b34801561016b57600080fd5b506100956004803603602081101561018257600080fd5b50356106bb565b6000546001600160a01b031633146101e4576040805162461bcd60e51b81526020600482015260196024820152781bdc195c985d1a5bdb88191bc81b9bdd08185c1c1c9bdd9959603a1b604482015290519081900360640190fd5b60045460ff1615610233576040805162461bcd60e51b815260206004820152601460248201527339b0b6329034b99030b63932b0b23c9037b832b760611b604482015290519081900360640190fd5b6004805460ff19166001179055565b60045460ff16610289576040805162461bcd60e51b815260206004820152600d60248201526c1b585c9ad95d0818db1bdcd959609a1b604482015290519081900360640190fd5b60035461029d90829063ffffffff61075a16565b34146102f0576040805162461bcd60e51b815260206004820152601860248201527f706c6561736520726563616c63756c6174652076616c75650000000000000000604482015290519081900360640190fd5b600154604080516370a0823160e01b8152306004820152905183926001600160a01b0316916370a08231916024808301926020929190829003018186803b15801561033a57600080fd5b505afa15801561034e573d6000803e3d6000fd5b505050506040513d602081101561036457600080fd5b505110156103a35760405162461bcd60e51b81526004018080602001828103825260288152602001806107bd6028913960400191505060405180910390fd5b6001546040805163a9059cbb60e01b81523360048201526024810184905290516001600160a01b039092169163a9059cbb916044808201926020929091908290030181600087803b1580156103f757600080fd5b505af115801561040b573d6000803e3d6000fd5b505050506040513d602081101561042157600080fd5b5051610469576040805162461bcd60e51b81526020600482015260126024820152711d1c985b9cd858dd1a5bdb8819985a5b195960721b604482015290519081900360640190fd5b6002805482019055604080513381526020810183905281517f5e5e995ce3133561afceaa51a9a154d5db228cd7525d34df5185582c18d3df09929181900390910190a150565b60045460ff166104f6576040805162461bcd60e51b815260206004820152600d60248201526c1b585c9ad95d0818db1bdcd959609a1b604482015290519081900360640190fd5b6000546001600160a01b03163314610551576040805162461bcd60e51b81526020600482015260196024820152781bdc195c985d1a5bdb88191bc81b9bdd08185c1c1c9bdd9959603a1b604482015290519081900360640190fd5b600154600054604080516370a0823160e01b815230600482015290516001600160a01b039384169363a9059cbb93169184916370a0823191602480820192602092909190829003018186803b1580156105a957600080fd5b505afa1580156105bd573d6000803e3d6000fd5b505050506040513d60208110156105d357600080fd5b5051604080516001600160e01b031960e086901b1681526001600160a01b03909316600484015260248301919091525160448083019260209291908290030181600087803b15801561062457600080fd5b505af1158015610638573d6000803e3d6000fd5b505050506040513d602081101561064e57600080fd5b505161068b5760405162461bcd60e51b81526004018080602001828103825260268152602001806108066026913960400191505060405180910390fd5b6004805460ff19169055565b60025481565b6001546001600160a01b031681565b60035481565b60045460ff1681565b6000546001600160a01b03163314610716576040805162461bcd60e51b81526020600482015260196024820152781bdc195c985d1a5bdb88191bc81b9bdd08185c1c1c9bdd9959603a1b604482015290519081900360640190fd5b600081116107555760405162461bcd60e51b815260040180806020018281038252602481526020018061082c6024913960400191505060405180910390fd5b600355565b600082610769575060006107b6565b8282028284828161077657fe5b04146107b35760405162461bcd60e51b81526004018080602001828103825260218152602001806107e56021913960400191505060405180910390fd5b90505b9291505056fe72657175697265206d6f726520746f6b656e207468616e20746869732063616e2070726f76696465536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f776661696c656420746f2072657475726e20756e736f6c6420746f6b656e20746f2061646d696e746f6b656e50726963652073686f756c6420626520706f73697469766520696e20576569a265627a7a723058202c999427cdf556460b9aec54ca43064bf20279dbf7882f37177c917c0ad4980464736f6c634300050a0032"

// DeploySofaCrowdSale deploys a new Ethereum contract, binding an instance of SofaCrowdSale to it.
func DeploySofaCrowdSale(auth *bind.TransactOpts, backend bind.ContractBackend, _sofaToken common.Address, _initalPrice *big.Int) (common.Address, *types.Transaction, *SofaCrowdSale, error) {
	parsed, err := abi.JSON(strings.NewReader(SofaCrowdSaleABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SofaCrowdSaleBin), backend, _sofaToken, _initalPrice)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SofaCrowdSale{SofaCrowdSaleCaller: SofaCrowdSaleCaller{contract: contract}, SofaCrowdSaleTransactor: SofaCrowdSaleTransactor{contract: contract}, SofaCrowdSaleFilterer: SofaCrowdSaleFilterer{contract: contract}}, nil
}

// SofaCrowdSale is an auto generated Go binding around an Ethereum contract.
type SofaCrowdSale struct {
	SofaCrowdSaleCaller     // Read-only binding to the contract
	SofaCrowdSaleTransactor // Write-only binding to the contract
	SofaCrowdSaleFilterer   // Log filterer for contract events
}

// SofaCrowdSaleCaller is an auto generated read-only Go binding around an Ethereum contract.
type SofaCrowdSaleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SofaCrowdSaleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SofaCrowdSaleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SofaCrowdSaleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SofaCrowdSaleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SofaCrowdSaleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SofaCrowdSaleSession struct {
	Contract     *SofaCrowdSale    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SofaCrowdSaleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SofaCrowdSaleCallerSession struct {
	Contract *SofaCrowdSaleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// SofaCrowdSaleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SofaCrowdSaleTransactorSession struct {
	Contract     *SofaCrowdSaleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// SofaCrowdSaleRaw is an auto generated low-level Go binding around an Ethereum contract.
type SofaCrowdSaleRaw struct {
	Contract *SofaCrowdSale // Generic contract binding to access the raw methods on
}

// SofaCrowdSaleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SofaCrowdSaleCallerRaw struct {
	Contract *SofaCrowdSaleCaller // Generic read-only contract binding to access the raw methods on
}

// SofaCrowdSaleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SofaCrowdSaleTransactorRaw struct {
	Contract *SofaCrowdSaleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSofaCrowdSale creates a new instance of SofaCrowdSale, bound to a specific deployed contract.
func NewSofaCrowdSale(address common.Address, backend bind.ContractBackend) (*SofaCrowdSale, error) {
	contract, err := bindSofaCrowdSale(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SofaCrowdSale{SofaCrowdSaleCaller: SofaCrowdSaleCaller{contract: contract}, SofaCrowdSaleTransactor: SofaCrowdSaleTransactor{contract: contract}, SofaCrowdSaleFilterer: SofaCrowdSaleFilterer{contract: contract}}, nil
}

// NewSofaCrowdSaleCaller creates a new read-only instance of SofaCrowdSale, bound to a specific deployed contract.
func NewSofaCrowdSaleCaller(address common.Address, caller bind.ContractCaller) (*SofaCrowdSaleCaller, error) {
	contract, err := bindSofaCrowdSale(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SofaCrowdSaleCaller{contract: contract}, nil
}

// NewSofaCrowdSaleTransactor creates a new write-only instance of SofaCrowdSale, bound to a specific deployed contract.
func NewSofaCrowdSaleTransactor(address common.Address, transactor bind.ContractTransactor) (*SofaCrowdSaleTransactor, error) {
	contract, err := bindSofaCrowdSale(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SofaCrowdSaleTransactor{contract: contract}, nil
}

// NewSofaCrowdSaleFilterer creates a new log filterer instance of SofaCrowdSale, bound to a specific deployed contract.
func NewSofaCrowdSaleFilterer(address common.Address, filterer bind.ContractFilterer) (*SofaCrowdSaleFilterer, error) {
	contract, err := bindSofaCrowdSale(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SofaCrowdSaleFilterer{contract: contract}, nil
}

// bindSofaCrowdSale binds a generic wrapper to an already deployed contract.
func bindSofaCrowdSale(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SofaCrowdSaleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SofaCrowdSale *SofaCrowdSaleRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SofaCrowdSale.Contract.SofaCrowdSaleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SofaCrowdSale *SofaCrowdSaleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SofaCrowdSale.Contract.SofaCrowdSaleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SofaCrowdSale *SofaCrowdSaleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SofaCrowdSale.Contract.SofaCrowdSaleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SofaCrowdSale *SofaCrowdSaleCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SofaCrowdSale.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SofaCrowdSale *SofaCrowdSaleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SofaCrowdSale.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SofaCrowdSale *SofaCrowdSaleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SofaCrowdSale.Contract.contract.Transact(opts, method, params...)
}

// SaleOpen is a free data retrieval call binding the contract method 0x99288dbb.
//
// Solidity: function saleOpen() constant returns(bool)
func (_SofaCrowdSale *SofaCrowdSaleCaller) SaleOpen(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SofaCrowdSale.contract.Call(opts, out, "saleOpen")
	return *ret0, err
}

// SaleOpen is a free data retrieval call binding the contract method 0x99288dbb.
//
// Solidity: function saleOpen() constant returns(bool)
func (_SofaCrowdSale *SofaCrowdSaleSession) SaleOpen() (bool, error) {
	return _SofaCrowdSale.Contract.SaleOpen(&_SofaCrowdSale.CallOpts)
}

// SaleOpen is a free data retrieval call binding the contract method 0x99288dbb.
//
// Solidity: function saleOpen() constant returns(bool)
func (_SofaCrowdSale *SofaCrowdSaleCallerSession) SaleOpen() (bool, error) {
	return _SofaCrowdSale.Contract.SaleOpen(&_SofaCrowdSale.CallOpts)
}

// TokenContract is a free data retrieval call binding the contract method 0x55a373d6.
//
// Solidity: function tokenContract() constant returns(address)
func (_SofaCrowdSale *SofaCrowdSaleCaller) TokenContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SofaCrowdSale.contract.Call(opts, out, "tokenContract")
	return *ret0, err
}

// TokenContract is a free data retrieval call binding the contract method 0x55a373d6.
//
// Solidity: function tokenContract() constant returns(address)
func (_SofaCrowdSale *SofaCrowdSaleSession) TokenContract() (common.Address, error) {
	return _SofaCrowdSale.Contract.TokenContract(&_SofaCrowdSale.CallOpts)
}

// TokenContract is a free data retrieval call binding the contract method 0x55a373d6.
//
// Solidity: function tokenContract() constant returns(address)
func (_SofaCrowdSale *SofaCrowdSaleCallerSession) TokenContract() (common.Address, error) {
	return _SofaCrowdSale.Contract.TokenContract(&_SofaCrowdSale.CallOpts)
}

// TokenPrice is a free data retrieval call binding the contract method 0x7ff9b596.
//
// Solidity: function tokenPrice() constant returns(uint256)
func (_SofaCrowdSale *SofaCrowdSaleCaller) TokenPrice(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SofaCrowdSale.contract.Call(opts, out, "tokenPrice")
	return *ret0, err
}

// TokenPrice is a free data retrieval call binding the contract method 0x7ff9b596.
//
// Solidity: function tokenPrice() constant returns(uint256)
func (_SofaCrowdSale *SofaCrowdSaleSession) TokenPrice() (*big.Int, error) {
	return _SofaCrowdSale.Contract.TokenPrice(&_SofaCrowdSale.CallOpts)
}

// TokenPrice is a free data retrieval call binding the contract method 0x7ff9b596.
//
// Solidity: function tokenPrice() constant returns(uint256)
func (_SofaCrowdSale *SofaCrowdSaleCallerSession) TokenPrice() (*big.Int, error) {
	return _SofaCrowdSale.Contract.TokenPrice(&_SofaCrowdSale.CallOpts)
}

// TokenSold is a free data retrieval call binding the contract method 0x519ee19e.
//
// Solidity: function tokenSold() constant returns(uint256)
func (_SofaCrowdSale *SofaCrowdSaleCaller) TokenSold(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SofaCrowdSale.contract.Call(opts, out, "tokenSold")
	return *ret0, err
}

// TokenSold is a free data retrieval call binding the contract method 0x519ee19e.
//
// Solidity: function tokenSold() constant returns(uint256)
func (_SofaCrowdSale *SofaCrowdSaleSession) TokenSold() (*big.Int, error) {
	return _SofaCrowdSale.Contract.TokenSold(&_SofaCrowdSale.CallOpts)
}

// TokenSold is a free data retrieval call binding the contract method 0x519ee19e.
//
// Solidity: function tokenSold() constant returns(uint256)
func (_SofaCrowdSale *SofaCrowdSaleCallerSession) TokenSold() (*big.Int, error) {
	return _SofaCrowdSale.Contract.TokenSold(&_SofaCrowdSale.CallOpts)
}

// BuyToken is a paid mutator transaction binding the contract method 0x2d296bf1.
//
// Solidity: function buyToken(uint256 _numberOfTokens) returns()
func (_SofaCrowdSale *SofaCrowdSaleTransactor) BuyToken(opts *bind.TransactOpts, _numberOfTokens *big.Int) (*types.Transaction, error) {
	return _SofaCrowdSale.contract.Transact(opts, "buyToken", _numberOfTokens)
}

// BuyToken is a paid mutator transaction binding the contract method 0x2d296bf1.
//
// Solidity: function buyToken(uint256 _numberOfTokens) returns()
func (_SofaCrowdSale *SofaCrowdSaleSession) BuyToken(_numberOfTokens *big.Int) (*types.Transaction, error) {
	return _SofaCrowdSale.Contract.BuyToken(&_SofaCrowdSale.TransactOpts, _numberOfTokens)
}

// BuyToken is a paid mutator transaction binding the contract method 0x2d296bf1.
//
// Solidity: function buyToken(uint256 _numberOfTokens) returns()
func (_SofaCrowdSale *SofaCrowdSaleTransactorSession) BuyToken(_numberOfTokens *big.Int) (*types.Transaction, error) {
	return _SofaCrowdSale.Contract.BuyToken(&_SofaCrowdSale.TransactOpts, _numberOfTokens)
}

// EndSale is a paid mutator transaction binding the contract method 0x380d831b.
//
// Solidity: function endSale() returns()
func (_SofaCrowdSale *SofaCrowdSaleTransactor) EndSale(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SofaCrowdSale.contract.Transact(opts, "endSale")
}

// EndSale is a paid mutator transaction binding the contract method 0x380d831b.
//
// Solidity: function endSale() returns()
func (_SofaCrowdSale *SofaCrowdSaleSession) EndSale() (*types.Transaction, error) {
	return _SofaCrowdSale.Contract.EndSale(&_SofaCrowdSale.TransactOpts)
}

// EndSale is a paid mutator transaction binding the contract method 0x380d831b.
//
// Solidity: function endSale() returns()
func (_SofaCrowdSale *SofaCrowdSaleTransactorSession) EndSale() (*types.Transaction, error) {
	return _SofaCrowdSale.Contract.EndSale(&_SofaCrowdSale.TransactOpts)
}

// OpenForSale is a paid mutator transaction binding the contract method 0x26138c36.
//
// Solidity: function openForSale() returns()
func (_SofaCrowdSale *SofaCrowdSaleTransactor) OpenForSale(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SofaCrowdSale.contract.Transact(opts, "openForSale")
}

// OpenForSale is a paid mutator transaction binding the contract method 0x26138c36.
//
// Solidity: function openForSale() returns()
func (_SofaCrowdSale *SofaCrowdSaleSession) OpenForSale() (*types.Transaction, error) {
	return _SofaCrowdSale.Contract.OpenForSale(&_SofaCrowdSale.TransactOpts)
}

// OpenForSale is a paid mutator transaction binding the contract method 0x26138c36.
//
// Solidity: function openForSale() returns()
func (_SofaCrowdSale *SofaCrowdSaleTransactorSession) OpenForSale() (*types.Transaction, error) {
	return _SofaCrowdSale.Contract.OpenForSale(&_SofaCrowdSale.TransactOpts)
}

// ResetTokenPrice is a paid mutator transaction binding the contract method 0xfc5abac1.
//
// Solidity: function resetTokenPrice(uint256 _tokenPrice) returns()
func (_SofaCrowdSale *SofaCrowdSaleTransactor) ResetTokenPrice(opts *bind.TransactOpts, _tokenPrice *big.Int) (*types.Transaction, error) {
	return _SofaCrowdSale.contract.Transact(opts, "resetTokenPrice", _tokenPrice)
}

// ResetTokenPrice is a paid mutator transaction binding the contract method 0xfc5abac1.
//
// Solidity: function resetTokenPrice(uint256 _tokenPrice) returns()
func (_SofaCrowdSale *SofaCrowdSaleSession) ResetTokenPrice(_tokenPrice *big.Int) (*types.Transaction, error) {
	return _SofaCrowdSale.Contract.ResetTokenPrice(&_SofaCrowdSale.TransactOpts, _tokenPrice)
}

// ResetTokenPrice is a paid mutator transaction binding the contract method 0xfc5abac1.
//
// Solidity: function resetTokenPrice(uint256 _tokenPrice) returns()
func (_SofaCrowdSale *SofaCrowdSaleTransactorSession) ResetTokenPrice(_tokenPrice *big.Int) (*types.Transaction, error) {
	return _SofaCrowdSale.Contract.ResetTokenPrice(&_SofaCrowdSale.TransactOpts, _tokenPrice)
}

// SofaCrowdSaleSellIterator is returned from FilterSell and is used to iterate over the raw logs and unpacked data for Sell events raised by the SofaCrowdSale contract.
type SofaCrowdSaleSellIterator struct {
	Event *SofaCrowdSaleSell // Event containing the contract specifics and raw log

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
func (it *SofaCrowdSaleSellIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SofaCrowdSaleSell)
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
		it.Event = new(SofaCrowdSaleSell)
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
func (it *SofaCrowdSaleSellIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SofaCrowdSaleSellIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SofaCrowdSaleSell represents a Sell event raised by the SofaCrowdSale contract.
type SofaCrowdSaleSell struct {
	Buyer  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSell is a free log retrieval operation binding the contract event 0x5e5e995ce3133561afceaa51a9a154d5db228cd7525d34df5185582c18d3df09.
//
// Solidity: event Sell(address _buyer, uint256 amount)
func (_SofaCrowdSale *SofaCrowdSaleFilterer) FilterSell(opts *bind.FilterOpts) (*SofaCrowdSaleSellIterator, error) {

	logs, sub, err := _SofaCrowdSale.contract.FilterLogs(opts, "Sell")
	if err != nil {
		return nil, err
	}
	return &SofaCrowdSaleSellIterator{contract: _SofaCrowdSale.contract, event: "Sell", logs: logs, sub: sub}, nil
}

// WatchSell is a free log subscription operation binding the contract event 0x5e5e995ce3133561afceaa51a9a154d5db228cd7525d34df5185582c18d3df09.
//
// Solidity: event Sell(address _buyer, uint256 amount)
func (_SofaCrowdSale *SofaCrowdSaleFilterer) WatchSell(opts *bind.WatchOpts, sink chan<- *SofaCrowdSaleSell) (event.Subscription, error) {

	logs, sub, err := _SofaCrowdSale.contract.WatchLogs(opts, "Sell")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SofaCrowdSaleSell)
				if err := _SofaCrowdSale.contract.UnpackLog(event, "Sell", log); err != nil {
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

// ParseSell is a log parse operation binding the contract event 0x5e5e995ce3133561afceaa51a9a154d5db228cd7525d34df5185582c18d3df09.
//
// Solidity: event Sell(address _buyer, uint256 amount)
func (_SofaCrowdSale *SofaCrowdSaleFilterer) ParseSell(log types.Log) (*SofaCrowdSaleSell, error) {
	event := new(SofaCrowdSaleSell)
	if err := _SofaCrowdSale.contract.UnpackLog(event, "Sell", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SofaTokenABI is the input ABI used to generate the binding from.
const SofaTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_initialSupply\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// SofaTokenFuncSigs maps the 4-byte function signature to its string representation.
var SofaTokenFuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"a457c2d7": "decreaseAllowance(address,uint256)",
	"39509351": "increaseAllowance(address,uint256)",
	"06fdde03": "name()",
	"95d89b41": "symbol()",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// SofaTokenBin is the compiled bytecode used for deploying new contracts.
var SofaTokenBin = "0x60c0604052600a60808190527f536f666120546f6b656e0000000000000000000000000000000000000000000060a090815261003e91600291906100cf565b506040805180820190915260048082527f534f4641000000000000000000000000000000000000000000000000000000006020909201918252610083916003916100cf565b5034801561009057600080fd5b506040516109dc3803806109dc833981810160405260208110156100b357600080fd5b505133600090815260208190526040902081905560045561016a565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061011057805160ff191683800117855561013d565b8280016001018555821561013d579182015b8281111561013d578251825591602001919060010190610122565b5061014992915061014d565b5090565b61016791905b808211156101495760008155600101610153565b90565b610863806101796000396000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c806370a082311161006657806370a08231146101dc57806395d89b4114610202578063a457c2d71461020a578063a9059cbb14610236578063dd62ed3e146102625761009e565b806306fdde03146100a3578063095ea7b31461012057806318160ddd1461016057806323b872dd1461017a57806339509351146101b0575b600080fd5b6100ab610290565b6040805160208082528351818301528351919283929083019185019080838360005b838110156100e55781810151838201526020016100cd565b50505050905090810190601f1680156101125780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b61014c6004803603604081101561013657600080fd5b506001600160a01b03813516906020013561031b565b604080519115158252519081900360200190f35b610168610331565b60408051918252519081900360200190f35b61014c6004803603606081101561019057600080fd5b506001600160a01b03813581169160208101359091169060400135610337565b61014c600480360360408110156101c657600080fd5b506001600160a01b03813516906020013561038e565b610168600480360360208110156101f257600080fd5b50356001600160a01b03166103ca565b6100ab6103e5565b61014c6004803603604081101561022057600080fd5b506001600160a01b038135169060200135610440565b61014c6004803603604081101561024c57600080fd5b506001600160a01b03813516906020013561047c565b6101686004803603604081101561027857600080fd5b506001600160a01b0381358116916020013516610489565b6002805460408051602060018416156101000260001901909316849004601f810184900484028201840190925281815292918301828280156103135780601f106102e857610100808354040283529160200191610313565b820191906000526020600020905b8154815290600101906020018083116102f657829003601f168201915b505050505081565b60006103283384846104b4565b50600192915050565b60045490565b60006103448484846105a0565b6001600160a01b03841660009081526001602090815260408083203380855292529091205461038491869161037f908663ffffffff6106e216565b6104b4565b5060019392505050565b3360008181526001602090815260408083206001600160a01b0387168452909152812054909161032891859061037f908663ffffffff61073f16565b6001600160a01b031660009081526020819052604090205490565b6003805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156103135780601f106102e857610100808354040283529160200191610313565b3360008181526001602090815260408083206001600160a01b0387168452909152812054909161032891859061037f908663ffffffff6106e216565b60006103283384846105a0565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b6001600160a01b0383166104f95760405162461bcd60e51b815260040180806020018281038252602481526020018061080b6024913960400191505060405180910390fd5b6001600160a01b03821661053e5760405162461bcd60e51b81526004018080602001828103825260228152602001806107c46022913960400191505060405180910390fd5b6001600160a01b03808416600081815260016020908152604080832094871680845294825291829020859055815185815291517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259281900390910190a3505050565b6001600160a01b0383166105e55760405162461bcd60e51b81526004018080602001828103825260258152602001806107e66025913960400191505060405180910390fd5b6001600160a01b03821661062a5760405162461bcd60e51b81526004018080602001828103825260238152602001806107a16023913960400191505060405180910390fd5b6001600160a01b038316600090815260208190526040902054610653908263ffffffff6106e216565b6001600160a01b038085166000908152602081905260408082209390935590841681522054610688908263ffffffff61073f16565b6001600160a01b038084166000818152602081815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b600082821115610739576040805162461bcd60e51b815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b50900390565b600082820183811015610799576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b939250505056fe45524332303a207472616e7366657220746f20746865207a65726f206164647265737345524332303a20617070726f766520746f20746865207a65726f206164647265737345524332303a207472616e736665722066726f6d20746865207a65726f206164647265737345524332303a20617070726f76652066726f6d20746865207a65726f2061646472657373a265627a7a723058205e98286d9fb90e5ee733353ec49dc72c7f9b144882f8ffd5eb5fd4f8dc0843a364736f6c634300050a0032"

// DeploySofaToken deploys a new Ethereum contract, binding an instance of SofaToken to it.
func DeploySofaToken(auth *bind.TransactOpts, backend bind.ContractBackend, _initialSupply *big.Int) (common.Address, *types.Transaction, *SofaToken, error) {
	parsed, err := abi.JSON(strings.NewReader(SofaTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SofaTokenBin), backend, _initialSupply)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SofaToken{SofaTokenCaller: SofaTokenCaller{contract: contract}, SofaTokenTransactor: SofaTokenTransactor{contract: contract}, SofaTokenFilterer: SofaTokenFilterer{contract: contract}}, nil
}

// SofaToken is an auto generated Go binding around an Ethereum contract.
type SofaToken struct {
	SofaTokenCaller     // Read-only binding to the contract
	SofaTokenTransactor // Write-only binding to the contract
	SofaTokenFilterer   // Log filterer for contract events
}

// SofaTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type SofaTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SofaTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SofaTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SofaTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SofaTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SofaTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SofaTokenSession struct {
	Contract     *SofaToken        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SofaTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SofaTokenCallerSession struct {
	Contract *SofaTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// SofaTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SofaTokenTransactorSession struct {
	Contract     *SofaTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SofaTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type SofaTokenRaw struct {
	Contract *SofaToken // Generic contract binding to access the raw methods on
}

// SofaTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SofaTokenCallerRaw struct {
	Contract *SofaTokenCaller // Generic read-only contract binding to access the raw methods on
}

// SofaTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SofaTokenTransactorRaw struct {
	Contract *SofaTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSofaToken creates a new instance of SofaToken, bound to a specific deployed contract.
func NewSofaToken(address common.Address, backend bind.ContractBackend) (*SofaToken, error) {
	contract, err := bindSofaToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SofaToken{SofaTokenCaller: SofaTokenCaller{contract: contract}, SofaTokenTransactor: SofaTokenTransactor{contract: contract}, SofaTokenFilterer: SofaTokenFilterer{contract: contract}}, nil
}

// NewSofaTokenCaller creates a new read-only instance of SofaToken, bound to a specific deployed contract.
func NewSofaTokenCaller(address common.Address, caller bind.ContractCaller) (*SofaTokenCaller, error) {
	contract, err := bindSofaToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SofaTokenCaller{contract: contract}, nil
}

// NewSofaTokenTransactor creates a new write-only instance of SofaToken, bound to a specific deployed contract.
func NewSofaTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*SofaTokenTransactor, error) {
	contract, err := bindSofaToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SofaTokenTransactor{contract: contract}, nil
}

// NewSofaTokenFilterer creates a new log filterer instance of SofaToken, bound to a specific deployed contract.
func NewSofaTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*SofaTokenFilterer, error) {
	contract, err := bindSofaToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SofaTokenFilterer{contract: contract}, nil
}

// bindSofaToken binds a generic wrapper to an already deployed contract.
func bindSofaToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SofaTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SofaToken *SofaTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SofaToken.Contract.SofaTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SofaToken *SofaTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SofaToken.Contract.SofaTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SofaToken *SofaTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SofaToken.Contract.SofaTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SofaToken *SofaTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SofaToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SofaToken *SofaTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SofaToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SofaToken *SofaTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SofaToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_SofaToken *SofaTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SofaToken.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_SofaToken *SofaTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _SofaToken.Contract.Allowance(&_SofaToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_SofaToken *SofaTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _SofaToken.Contract.Allowance(&_SofaToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_SofaToken *SofaTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SofaToken.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_SofaToken *SofaTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _SofaToken.Contract.BalanceOf(&_SofaToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_SofaToken *SofaTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _SofaToken.Contract.BalanceOf(&_SofaToken.CallOpts, account)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_SofaToken *SofaTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _SofaToken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_SofaToken *SofaTokenSession) Name() (string, error) {
	return _SofaToken.Contract.Name(&_SofaToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_SofaToken *SofaTokenCallerSession) Name() (string, error) {
	return _SofaToken.Contract.Name(&_SofaToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_SofaToken *SofaTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _SofaToken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_SofaToken *SofaTokenSession) Symbol() (string, error) {
	return _SofaToken.Contract.Symbol(&_SofaToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_SofaToken *SofaTokenCallerSession) Symbol() (string, error) {
	return _SofaToken.Contract.Symbol(&_SofaToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_SofaToken *SofaTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SofaToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_SofaToken *SofaTokenSession) TotalSupply() (*big.Int, error) {
	return _SofaToken.Contract.TotalSupply(&_SofaToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_SofaToken *SofaTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _SofaToken.Contract.TotalSupply(&_SofaToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_SofaToken *SofaTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _SofaToken.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_SofaToken *SofaTokenSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _SofaToken.Contract.Approve(&_SofaToken.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_SofaToken *SofaTokenTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _SofaToken.Contract.Approve(&_SofaToken.TransactOpts, spender, value)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_SofaToken *SofaTokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _SofaToken.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_SofaToken *SofaTokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _SofaToken.Contract.DecreaseAllowance(&_SofaToken.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_SofaToken *SofaTokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _SofaToken.Contract.DecreaseAllowance(&_SofaToken.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_SofaToken *SofaTokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _SofaToken.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_SofaToken *SofaTokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _SofaToken.Contract.IncreaseAllowance(&_SofaToken.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_SofaToken *SofaTokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _SofaToken.Contract.IncreaseAllowance(&_SofaToken.TransactOpts, spender, addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_SofaToken *SofaTokenTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SofaToken.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_SofaToken *SofaTokenSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SofaToken.Contract.Transfer(&_SofaToken.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_SofaToken *SofaTokenTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SofaToken.Contract.Transfer(&_SofaToken.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_SofaToken *SofaTokenTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SofaToken.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_SofaToken *SofaTokenSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SofaToken.Contract.TransferFrom(&_SofaToken.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_SofaToken *SofaTokenTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SofaToken.Contract.TransferFrom(&_SofaToken.TransactOpts, sender, recipient, amount)
}

// SofaTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the SofaToken contract.
type SofaTokenApprovalIterator struct {
	Event *SofaTokenApproval // Event containing the contract specifics and raw log

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
func (it *SofaTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SofaTokenApproval)
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
		it.Event = new(SofaTokenApproval)
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
func (it *SofaTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SofaTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SofaTokenApproval represents a Approval event raised by the SofaToken contract.
type SofaTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_SofaToken *SofaTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*SofaTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _SofaToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &SofaTokenApprovalIterator{contract: _SofaToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_SofaToken *SofaTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *SofaTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _SofaToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SofaTokenApproval)
				if err := _SofaToken.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_SofaToken *SofaTokenFilterer) ParseApproval(log types.Log) (*SofaTokenApproval, error) {
	event := new(SofaTokenApproval)
	if err := _SofaToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SofaTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the SofaToken contract.
type SofaTokenTransferIterator struct {
	Event *SofaTokenTransfer // Event containing the contract specifics and raw log

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
func (it *SofaTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SofaTokenTransfer)
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
		it.Event = new(SofaTokenTransfer)
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
func (it *SofaTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SofaTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SofaTokenTransfer represents a Transfer event raised by the SofaToken contract.
type SofaTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_SofaToken *SofaTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SofaTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SofaToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SofaTokenTransferIterator{contract: _SofaToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_SofaToken *SofaTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *SofaTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SofaToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SofaTokenTransfer)
				if err := _SofaToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_SofaToken *SofaTokenFilterer) ParseTransfer(log types.Log) (*SofaTokenTransfer, error) {
	event := new(SofaTokenTransfer)
	if err := _SofaToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
