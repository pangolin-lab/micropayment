// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package PayToCheck

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

// PayToCheckABI is the input ABI used to generate the binding from.
const PayToCheckABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"claimPayment\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"h\",\"type\":\"bytes32\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"testRecovery\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"add\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\"},{\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"recreateMsg\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"h\",\"type\":\"bytes32\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"testRecoveryNoPrefix\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"t\",\"type\":\"address\"},{\"name\":\"v\",\"type\":\"uint256\"}],\"name\":\"testHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"h\",\"type\":\"bytes32\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"testOwnerSigned\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"returnCheckState\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_sofaToken\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// PayToCheckFuncSigs maps the 4-byte function signature to its string representation.
var PayToCheckFuncSigs = map[string]string{
	"01857261": "claimPayment(uint256,uint256,uint8,bytes32,bytes32)",
	"41c0e1b5": "kill()",
	"088ba4a5": "recreateMsg(address,uint256,uint256,address)",
	"c28cf11d": "returnCheckState(uint256)",
	"b2e954ef": "testHash(address,uint256)",
	"bcabc7f5": "testOwnerSigned(bytes32,uint8,bytes32,bytes32)",
	"0242a5fc": "testRecovery(bytes32,uint8,bytes32,bytes32)",
	"9ac8b02b": "testRecoveryNoPrefix(bytes32,uint8,bytes32,bytes32)",
	"55a373d6": "tokenContract()",
}

// PayToCheckBin is the compiled bytecode used for deploying new contracts.
var PayToCheckBin = "0x6080604052600080546001600160a01b0319163317905534801561002257600080fd5b506040516107c03803806107c08339818101604052602081101561004557600080fd5b5051600280546001600160a01b0319166001600160a01b0390921691909117905561074b806100756000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c806355a373d61161006657806355a373d6146101745780639ac8b02b1461017c578063b2e954ef146101ae578063bcabc7f5146101da578063c28cf11d1461022057610093565b806301857261146100985780630242a5fc146100d2578063088ba4a51461012057806341c0e1b51461016c575b600080fd5b6100d0600480360360a08110156100ae57600080fd5b5080359060208101359060ff6040820135169060608101359060800135610261565b005b610104600480360360808110156100e857600080fd5b5080359060ff60208201351690604081013590606001356103e8565b604080516001600160a01b039092168252519081900360200190f35b61015a6004803603608081101561013657600080fd5b506001600160a01b0381358116916020810135916040820135916060013516610500565b60408051918252519081900360200190f35b6100d0610558565b6101046105b3565b6101046004803603608081101561019257600080fd5b5080359060ff60208201351690604081013590606001356105c2565b61015a600480360360408110156101c457600080fd5b506001600160a01b038135169060200135610637565b61020c600480360360808110156101f057600080fd5b5080359060ff602082013516906040810135906060013561067b565b604080519115158252519081900360200190f35b61023d6004803603602081101561023657600080fd5b5035610701565b6040518082600181111561024d57fe5b60ff16815260200191505060405180910390f35b836000808281526001602081905260409091205460ff169081111561028257fe5b146102c9576040805162461bcd60e51b81526020600482015260126024820152716e6f7420612076616c696465206e6f6e636560701b604482015290519081900360640190fd5b60008581526001602081905260408220805460ff191690911790556102f033888830610500565b6000549091506001600160a01b031661030b828787876105c2565b6001600160a01b03161461035f576040805162461bcd60e51b81526020600482015260166024820152751cda59db985d1d5c994818da1958dac819985a5b195960521b604482015290519081900360640190fd5b6002546040805163a9059cbb60e01b8152336004820152602481018a905290516001600160a01b039092169163a9059cbb916044808201926020929091908290030181600087803b1580156103b357600080fd5b505af11580156103c7573d6000803e3d6000fd5b505050506040513d60208110156103dd57600080fd5b505050505050505050565b600060606040518060400160405280601c81526020017f19457468657265756d205369676e6564204d6573736167653a0a3332000000008152509050600081876040516020018083805190602001908083835b6020831061045a5780518252601f19909201916020918201910161043b565b51815160209384036101000a600019018019909216911617905292019384525060408051808503815284830180835281519184019190912060009182905282860180845281905260ff8d166060870152608086018c905260a086018b90529151919650945060019360c08082019450601f19830192918290030190855afa1580156104e9573d6000803e3d6000fd5b5050604051601f1901519998505050505050505050565b604080516bffffffffffffffffffffffff19606096871b8116602080840191909152603483019690965260548201949094529190941b9091166074820152825160688183030181526088909101909252815191012090565b6000546001600160a01b031633146105b0576040805162461bcd60e51b815260206004820152601660248201527537b7363c9037bbb732b21034b99030b63637bbb2b21760511b604482015290519081900360640190fd5b33ff5b6002546001600160a01b031681565b60008060018686868660405160008152602001604052604051808581526020018460ff1660ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa158015610622573d6000803e3d6000fd5b5050604051601f190151979650505050505050565b6040805160609390931b6bffffffffffffffffffffffff19166020808501919091526034808501939093528151808503909301835260549093019052805191012090565b60008060018686868660405160008152602001604052604051808581526020018460ff1660ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa1580156106db573d6000803e3d6000fd5b5050604051601f1901516000546001600160a01b03908116911614979650505050505050565b60009081526001602052604090205460ff169056fea265627a7a723058209ce067c8b369e78581d664baed77fa0f38e9ef947617ea993d7d27a72308ab7964736f6c634300050a0032"

// DeployPayToCheck deploys a new Ethereum contract, binding an instance of PayToCheck to it.
func DeployPayToCheck(auth *bind.TransactOpts, backend bind.ContractBackend, _sofaToken common.Address) (common.Address, *types.Transaction, *PayToCheck, error) {
	parsed, err := abi.JSON(strings.NewReader(PayToCheckABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PayToCheckBin), backend, _sofaToken)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PayToCheck{PayToCheckCaller: PayToCheckCaller{contract: contract}, PayToCheckTransactor: PayToCheckTransactor{contract: contract}, PayToCheckFilterer: PayToCheckFilterer{contract: contract}}, nil
}

// PayToCheck is an auto generated Go binding around an Ethereum contract.
type PayToCheck struct {
	PayToCheckCaller     // Read-only binding to the contract
	PayToCheckTransactor // Write-only binding to the contract
	PayToCheckFilterer   // Log filterer for contract events
}

// PayToCheckCaller is an auto generated read-only Go binding around an Ethereum contract.
type PayToCheckCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PayToCheckTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PayToCheckTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PayToCheckFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PayToCheckFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PayToCheckSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PayToCheckSession struct {
	Contract     *PayToCheck       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PayToCheckCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PayToCheckCallerSession struct {
	Contract *PayToCheckCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// PayToCheckTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PayToCheckTransactorSession struct {
	Contract     *PayToCheckTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// PayToCheckRaw is an auto generated low-level Go binding around an Ethereum contract.
type PayToCheckRaw struct {
	Contract *PayToCheck // Generic contract binding to access the raw methods on
}

// PayToCheckCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PayToCheckCallerRaw struct {
	Contract *PayToCheckCaller // Generic read-only contract binding to access the raw methods on
}

// PayToCheckTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PayToCheckTransactorRaw struct {
	Contract *PayToCheckTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPayToCheck creates a new instance of PayToCheck, bound to a specific deployed contract.
func NewPayToCheck(address common.Address, backend bind.ContractBackend) (*PayToCheck, error) {
	contract, err := bindPayToCheck(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PayToCheck{PayToCheckCaller: PayToCheckCaller{contract: contract}, PayToCheckTransactor: PayToCheckTransactor{contract: contract}, PayToCheckFilterer: PayToCheckFilterer{contract: contract}}, nil
}

// NewPayToCheckCaller creates a new read-only instance of PayToCheck, bound to a specific deployed contract.
func NewPayToCheckCaller(address common.Address, caller bind.ContractCaller) (*PayToCheckCaller, error) {
	contract, err := bindPayToCheck(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PayToCheckCaller{contract: contract}, nil
}

// NewPayToCheckTransactor creates a new write-only instance of PayToCheck, bound to a specific deployed contract.
func NewPayToCheckTransactor(address common.Address, transactor bind.ContractTransactor) (*PayToCheckTransactor, error) {
	contract, err := bindPayToCheck(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PayToCheckTransactor{contract: contract}, nil
}

// NewPayToCheckFilterer creates a new log filterer instance of PayToCheck, bound to a specific deployed contract.
func NewPayToCheckFilterer(address common.Address, filterer bind.ContractFilterer) (*PayToCheckFilterer, error) {
	contract, err := bindPayToCheck(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PayToCheckFilterer{contract: contract}, nil
}

// bindPayToCheck binds a generic wrapper to an already deployed contract.
func bindPayToCheck(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PayToCheckABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PayToCheck *PayToCheckRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PayToCheck.Contract.PayToCheckCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PayToCheck *PayToCheckRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PayToCheck.Contract.PayToCheckTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PayToCheck *PayToCheckRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PayToCheck.Contract.PayToCheckTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PayToCheck *PayToCheckCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PayToCheck.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PayToCheck *PayToCheckTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PayToCheck.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PayToCheck *PayToCheckTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PayToCheck.Contract.contract.Transact(opts, method, params...)
}

// RecreateMsg is a free data retrieval call binding the contract method 0x088ba4a5.
//
// Solidity: function recreateMsg(address add, uint256 amount, uint256 nonce, address contractAddress) constant returns(bytes32)
func (_PayToCheck *PayToCheckCaller) RecreateMsg(opts *bind.CallOpts, add common.Address, amount *big.Int, nonce *big.Int, contractAddress common.Address) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _PayToCheck.contract.Call(opts, out, "recreateMsg", add, amount, nonce, contractAddress)
	return *ret0, err
}

// RecreateMsg is a free data retrieval call binding the contract method 0x088ba4a5.
//
// Solidity: function recreateMsg(address add, uint256 amount, uint256 nonce, address contractAddress) constant returns(bytes32)
func (_PayToCheck *PayToCheckSession) RecreateMsg(add common.Address, amount *big.Int, nonce *big.Int, contractAddress common.Address) ([32]byte, error) {
	return _PayToCheck.Contract.RecreateMsg(&_PayToCheck.CallOpts, add, amount, nonce, contractAddress)
}

// RecreateMsg is a free data retrieval call binding the contract method 0x088ba4a5.
//
// Solidity: function recreateMsg(address add, uint256 amount, uint256 nonce, address contractAddress) constant returns(bytes32)
func (_PayToCheck *PayToCheckCallerSession) RecreateMsg(add common.Address, amount *big.Int, nonce *big.Int, contractAddress common.Address) ([32]byte, error) {
	return _PayToCheck.Contract.RecreateMsg(&_PayToCheck.CallOpts, add, amount, nonce, contractAddress)
}

// ReturnCheckState is a free data retrieval call binding the contract method 0xc28cf11d.
//
// Solidity: function returnCheckState(uint256 nonce) constant returns(uint8)
func (_PayToCheck *PayToCheckCaller) ReturnCheckState(opts *bind.CallOpts, nonce *big.Int) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _PayToCheck.contract.Call(opts, out, "returnCheckState", nonce)
	return *ret0, err
}

// ReturnCheckState is a free data retrieval call binding the contract method 0xc28cf11d.
//
// Solidity: function returnCheckState(uint256 nonce) constant returns(uint8)
func (_PayToCheck *PayToCheckSession) ReturnCheckState(nonce *big.Int) (uint8, error) {
	return _PayToCheck.Contract.ReturnCheckState(&_PayToCheck.CallOpts, nonce)
}

// ReturnCheckState is a free data retrieval call binding the contract method 0xc28cf11d.
//
// Solidity: function returnCheckState(uint256 nonce) constant returns(uint8)
func (_PayToCheck *PayToCheckCallerSession) ReturnCheckState(nonce *big.Int) (uint8, error) {
	return _PayToCheck.Contract.ReturnCheckState(&_PayToCheck.CallOpts, nonce)
}

// TestHash is a free data retrieval call binding the contract method 0xb2e954ef.
//
// Solidity: function testHash(address t, uint256 v) constant returns(bytes32)
func (_PayToCheck *PayToCheckCaller) TestHash(opts *bind.CallOpts, t common.Address, v *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _PayToCheck.contract.Call(opts, out, "testHash", t, v)
	return *ret0, err
}

// TestHash is a free data retrieval call binding the contract method 0xb2e954ef.
//
// Solidity: function testHash(address t, uint256 v) constant returns(bytes32)
func (_PayToCheck *PayToCheckSession) TestHash(t common.Address, v *big.Int) ([32]byte, error) {
	return _PayToCheck.Contract.TestHash(&_PayToCheck.CallOpts, t, v)
}

// TestHash is a free data retrieval call binding the contract method 0xb2e954ef.
//
// Solidity: function testHash(address t, uint256 v) constant returns(bytes32)
func (_PayToCheck *PayToCheckCallerSession) TestHash(t common.Address, v *big.Int) ([32]byte, error) {
	return _PayToCheck.Contract.TestHash(&_PayToCheck.CallOpts, t, v)
}

// TestOwnerSigned is a free data retrieval call binding the contract method 0xbcabc7f5.
//
// Solidity: function testOwnerSigned(bytes32 h, uint8 v, bytes32 r, bytes32 s) constant returns(bool)
func (_PayToCheck *PayToCheckCaller) TestOwnerSigned(opts *bind.CallOpts, h [32]byte, v uint8, r [32]byte, s [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PayToCheck.contract.Call(opts, out, "testOwnerSigned", h, v, r, s)
	return *ret0, err
}

// TestOwnerSigned is a free data retrieval call binding the contract method 0xbcabc7f5.
//
// Solidity: function testOwnerSigned(bytes32 h, uint8 v, bytes32 r, bytes32 s) constant returns(bool)
func (_PayToCheck *PayToCheckSession) TestOwnerSigned(h [32]byte, v uint8, r [32]byte, s [32]byte) (bool, error) {
	return _PayToCheck.Contract.TestOwnerSigned(&_PayToCheck.CallOpts, h, v, r, s)
}

// TestOwnerSigned is a free data retrieval call binding the contract method 0xbcabc7f5.
//
// Solidity: function testOwnerSigned(bytes32 h, uint8 v, bytes32 r, bytes32 s) constant returns(bool)
func (_PayToCheck *PayToCheckCallerSession) TestOwnerSigned(h [32]byte, v uint8, r [32]byte, s [32]byte) (bool, error) {
	return _PayToCheck.Contract.TestOwnerSigned(&_PayToCheck.CallOpts, h, v, r, s)
}

// TestRecovery is a free data retrieval call binding the contract method 0x0242a5fc.
//
// Solidity: function testRecovery(bytes32 h, uint8 v, bytes32 r, bytes32 s) constant returns(address)
func (_PayToCheck *PayToCheckCaller) TestRecovery(opts *bind.CallOpts, h [32]byte, v uint8, r [32]byte, s [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PayToCheck.contract.Call(opts, out, "testRecovery", h, v, r, s)
	return *ret0, err
}

// TestRecovery is a free data retrieval call binding the contract method 0x0242a5fc.
//
// Solidity: function testRecovery(bytes32 h, uint8 v, bytes32 r, bytes32 s) constant returns(address)
func (_PayToCheck *PayToCheckSession) TestRecovery(h [32]byte, v uint8, r [32]byte, s [32]byte) (common.Address, error) {
	return _PayToCheck.Contract.TestRecovery(&_PayToCheck.CallOpts, h, v, r, s)
}

// TestRecovery is a free data retrieval call binding the contract method 0x0242a5fc.
//
// Solidity: function testRecovery(bytes32 h, uint8 v, bytes32 r, bytes32 s) constant returns(address)
func (_PayToCheck *PayToCheckCallerSession) TestRecovery(h [32]byte, v uint8, r [32]byte, s [32]byte) (common.Address, error) {
	return _PayToCheck.Contract.TestRecovery(&_PayToCheck.CallOpts, h, v, r, s)
}

// TestRecoveryNoPrefix is a free data retrieval call binding the contract method 0x9ac8b02b.
//
// Solidity: function testRecoveryNoPrefix(bytes32 h, uint8 v, bytes32 r, bytes32 s) constant returns(address)
func (_PayToCheck *PayToCheckCaller) TestRecoveryNoPrefix(opts *bind.CallOpts, h [32]byte, v uint8, r [32]byte, s [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PayToCheck.contract.Call(opts, out, "testRecoveryNoPrefix", h, v, r, s)
	return *ret0, err
}

// TestRecoveryNoPrefix is a free data retrieval call binding the contract method 0x9ac8b02b.
//
// Solidity: function testRecoveryNoPrefix(bytes32 h, uint8 v, bytes32 r, bytes32 s) constant returns(address)
func (_PayToCheck *PayToCheckSession) TestRecoveryNoPrefix(h [32]byte, v uint8, r [32]byte, s [32]byte) (common.Address, error) {
	return _PayToCheck.Contract.TestRecoveryNoPrefix(&_PayToCheck.CallOpts, h, v, r, s)
}

// TestRecoveryNoPrefix is a free data retrieval call binding the contract method 0x9ac8b02b.
//
// Solidity: function testRecoveryNoPrefix(bytes32 h, uint8 v, bytes32 r, bytes32 s) constant returns(address)
func (_PayToCheck *PayToCheckCallerSession) TestRecoveryNoPrefix(h [32]byte, v uint8, r [32]byte, s [32]byte) (common.Address, error) {
	return _PayToCheck.Contract.TestRecoveryNoPrefix(&_PayToCheck.CallOpts, h, v, r, s)
}

// TokenContract is a free data retrieval call binding the contract method 0x55a373d6.
//
// Solidity: function tokenContract() constant returns(address)
func (_PayToCheck *PayToCheckCaller) TokenContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PayToCheck.contract.Call(opts, out, "tokenContract")
	return *ret0, err
}

// TokenContract is a free data retrieval call binding the contract method 0x55a373d6.
//
// Solidity: function tokenContract() constant returns(address)
func (_PayToCheck *PayToCheckSession) TokenContract() (common.Address, error) {
	return _PayToCheck.Contract.TokenContract(&_PayToCheck.CallOpts)
}

// TokenContract is a free data retrieval call binding the contract method 0x55a373d6.
//
// Solidity: function tokenContract() constant returns(address)
func (_PayToCheck *PayToCheckCallerSession) TokenContract() (common.Address, error) {
	return _PayToCheck.Contract.TokenContract(&_PayToCheck.CallOpts)
}

// ClaimPayment is a paid mutator transaction binding the contract method 0x01857261.
//
// Solidity: function claimPayment(uint256 amount, uint256 nonce, uint8 v, bytes32 r, bytes32 s) returns()
func (_PayToCheck *PayToCheckTransactor) ClaimPayment(opts *bind.TransactOpts, amount *big.Int, nonce *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _PayToCheck.contract.Transact(opts, "claimPayment", amount, nonce, v, r, s)
}

// ClaimPayment is a paid mutator transaction binding the contract method 0x01857261.
//
// Solidity: function claimPayment(uint256 amount, uint256 nonce, uint8 v, bytes32 r, bytes32 s) returns()
func (_PayToCheck *PayToCheckSession) ClaimPayment(amount *big.Int, nonce *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _PayToCheck.Contract.ClaimPayment(&_PayToCheck.TransactOpts, amount, nonce, v, r, s)
}

// ClaimPayment is a paid mutator transaction binding the contract method 0x01857261.
//
// Solidity: function claimPayment(uint256 amount, uint256 nonce, uint8 v, bytes32 r, bytes32 s) returns()
func (_PayToCheck *PayToCheckTransactorSession) ClaimPayment(amount *big.Int, nonce *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _PayToCheck.Contract.ClaimPayment(&_PayToCheck.TransactOpts, amount, nonce, v, r, s)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_PayToCheck *PayToCheckTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PayToCheck.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_PayToCheck *PayToCheckSession) Kill() (*types.Transaction, error) {
	return _PayToCheck.Contract.Kill(&_PayToCheck.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_PayToCheck *PayToCheckTransactorSession) Kill() (*types.Transaction, error) {
	return _PayToCheck.Contract.Kill(&_PayToCheck.TransactOpts)
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
