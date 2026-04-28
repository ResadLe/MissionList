// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package excabi

/*
============================================================
 Counter 合约 Go 绑定 — 详细注释版
============================================================

合约部署信息：
  - 网络：Sepolia 测试网
  - 地址：0xa33987386298b8Fd4989Fc5D53c179446Cb54FCf

合约功能：一个简单的计数器，包含 3 个函数和 1 个事件

  函数            | 类型        | 说明
  ----------------|-------------|------------------------
  x()             | view        | 读取当前计数值
  inc()           | nonpayable  | 计数值 +1
  incBy(uint256)  | nonpayable  | 计数值 +N

  事件            | 说明
  ----------------|------------------------
  Increment(uint) | 每次增加时触发，记录增加量

类型层级关系（由 abigen 自动生成）：

  CounterAbi ───────────────────────── 完整实例
  ├── CounterAbiCaller ─────────────── 只读（调用 view 函数）
  ├── CounterAbiTransactor ─────────── 只写（发送交易）
  └── CounterAbiFilterer ───────────── 事件日志过滤

  CounterAbiSession ────────────────── 完整 Session（预置 opts）
  ├── CounterAbiCallerSession ──────── 只读 Session
  └── CounterAbiTransactorSession ──── 只写 Session

  CounterAbiRaw ────────────────────── 底层 Raw 接口
  ├── CounterAbiCallerRaw ──────────── 底层只读
  └── CounterAbiTransactorRaw ──────── 底层只写

使用示例：
  // 方式一：每次传 opts（灵活）
  contract, _ := NewCounterAbi(addr, client)
  val, _ := contract.X(nil)
  tx, _ := contract.Inc(auth)

  // 方式二：Session 预置 opts（批量调用更简洁）
  session := &CounterAbiSession{Contract: contract, CallOpts: bind.CallOpts{}, TransactOpts: auth}
  val, _ := session.X()
  tx, _ := session.Inc()

  // 方式三：监听事件
  iter, _ := contract.FilterIncrement(nil)
  for iter.Next() {
      fmt.Println("增加了", iter.Event.By)
  }
*/

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
// 下面这些 import 在本文件中没有直接调用，但生成的辅助代码会用到它们。
// Go 编译器要求所有 import 必须有实际使用，否则报错，所以这里故意引用它们来消除编译错误。
var (
	_ = errors.New            // errors 包用于错误处理
	_ = big.NewInt            // math/big 用于大整数（以太坊的 uint256）
	_ = strings.NewReader     // strings 用于 ABI JSON 解析
	_ = ethereum.NotFound     // ethereum 包常量
	_ = bind.Bind             // bind 核心绑定逻辑
	_ = common.Big1           // common 包常量
	_ = types.BloomLookup     // types 日志相关
	_ = event.NewSubscription // event 事件订阅
	_ = abi.ConvertType       // abi 类型转换
)

// ============================================================
// 元数据
// ============================================================

// CounterAbiMetaData 包含合约的 ABI JSON 元数据。
// 由 abigen 生成，内部存储了合约所有函数和事件的 ABI 描述，
// 后续所有函数调用、事件解析都依赖这份元数据。
//
// 推荐优先使用此字段，而不是直接使用 CounterAbiABI。
var CounterAbiMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"by\",\"type\":\"uint256\"}],\"name\":\"Increment\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"inc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"by\",\"type\":\"uint256\"}],\"name\":\"incBy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"x\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]//sepolia//CounterModule#Counter-0xa33987386298b8Fd4989Fc5D53c179446Cb54FCf",
}

// CounterAbiABI 是旧版兼容字段，已废弃。
// Deprecated: 请使用 CounterAbiMetaData.ABI 替代。
var CounterAbiABI = CounterAbiMetaData.ABI

// ============================================================
// 核心类型（共 12 个类型，分为 3 组）
// ============================================================

/* ---------- 第一组：完整实例 + 三个角色 ---------- */

// CounterAbi 是合约的完整 Go 绑定，包含读、写、事件过滤三种能力。
// 最常用的类型，通过 NewCounterAbi() 创建。
//
// 使用方式：
//
//	contract, _ := NewCounterAbi(addr, client)
//	val, _ := contract.X(nil)       // 读
//	tx, _ := contract.Inc(auth)     // 写
//	iter, _ := contract.FilterIncrement(nil) // 事件
type CounterAbi struct {
	CounterAbiCaller     // 只读角色：调用 view/pure 函数
	CounterAbiTransactor // 只写角色：发送交易修改状态
	CounterAbiFilterer   // 事件角色：过滤和解析合约日志
}

// CounterAbiCaller 是合约的只读绑定。
// 只能调用 view/pure 函数（如 x()），不需要发送交易，不消耗 Gas。
// 通常作为 CounterAbi 的内嵌字段使用，也可以单独创建。
type CounterAbiCaller struct {
	contract *bind.BoundContract // 底层合约包装器，负责实际的 JSON-RPC 调用
}

// CounterAbiTransactor 是合约的只写绑定。
// 只能调用会修改状态的函数（如 inc()、incBy()），需要发送交易并消耗 Gas。
type CounterAbiTransactor struct {
	contract *bind.BoundContract
}

// CounterAbiFilterer 是合约的事件日志过滤绑定。
// 用于查询和订阅合约事件（如 Increment），可以：
//   - Filter*   : 查询历史事件日志
//   - Watch*    : 实时订阅新事件
//   - Parse*    : 解析单条日志数据
type CounterAbiFilterer struct {
	contract *bind.BoundContract
}

/* ---------- 第二组：Session（预置 opts 的便捷类型） ---------- */

// CounterAbiSession 是带预设置调用选项的完整合约绑定。
// 创建时将 CallOpts 和 TransactOpts 固定下来，后续每次调用方法
// 都不需要再传 opts，适合批量操作同一账户/同一区块的场景。
//
// 对比：
//
//	不使用 Session: contract.X(&bind.CallOpts{...})  // 每次都要传
//	使用 Session:   session.X()                       // opts 已预置
type CounterAbiSession struct {
	Contract     *CounterAbi       // 关联的完整合约实例
	CallOpts     bind.CallOpts     // 预置的读选项（如区块号、上下文）
	TransactOpts bind.TransactOpts // 预置的交易选项（如发送者、Gas、Nonce）
}

// CounterAbiCallerSession 是只读的 Session 绑定。
// 预置了 CallOpts，调用 view 函数时无需每次传 opts。
type CounterAbiCallerSession struct {
	Contract *CounterAbiCaller // 关联的只读合约实例
	CallOpts bind.CallOpts     // 预置的读选项
}

// CounterAbiTransactorSession 是只写的 Session 绑定。
// 预置了 TransactOpts，发送交易时无需每次传 opts。
type CounterAbiTransactorSession struct {
	Contract     *CounterAbiTransactor // 关联的只写合约实例
	TransactOpts bind.TransactOpts     // 预置的交易选项
}

/* ---------- 第三组：Raw（底层裸接口） ---------- */

// CounterAbiRaw 是底层裸绑定，提供最基础的 Call/Transfer/Transact 方法。
// 一般不直接使用，而是通过上层的 Caller/Transactor 间接调用。
type CounterAbiRaw struct {
	Contract *CounterAbi
}

// CounterAbiCallerRaw 是底层只读裸绑定。
type CounterAbiCallerRaw struct {
	Contract *CounterAbiCaller
}

// CounterAbiTransactorRaw 是底层只写裸绑定。
type CounterAbiTransactorRaw struct {
	Contract *CounterAbiTransactor
}

// ============================================================
// 构造函数
// ============================================================

// NewCounterAbi 创建一个绑定到已部署合约的完整 CounterAbi 实例。
//
// 参数：
//
//	address  — 合约地址
//	backend  — 实现了 ContractBackend 接口的客户端（通常 *ethclient.Client）
//	          ContractBackend 同时实现了 ContractCaller、ContractTransactor、ContractFilterer，
//	          因此可以作为三个角色共用同一个客户端。
//
// 返回的 CounterAbi 包含 Caller、Transactor、Filterer 三个内嵌角色，
// 可以同时进行读取、写入、事件监听。
//
// 示例：
//
//	client, _ := ethclient.Dial("https://sepolia.infura.io/v3/YOUR_KEY")
//	addr := common.HexToAddress("0xa33987386298b8Fd4989Fc5D53c179446Cb54FCf")
//	contract, _ := NewCounterAbi(addr, client)
func NewCounterAbi(address common.Address, backend bind.ContractBackend) (*CounterAbi, error) {
	contract, err := bindCounterAbi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CounterAbi{
		CounterAbiCaller:     CounterAbiCaller{contract: contract},
		CounterAbiTransactor: CounterAbiTransactor{contract: contract},
		CounterAbiFilterer:   CounterAbiFilterer{contract: contract},
	}, nil
}

// NewCounterAbiCaller 创建只读合约实例。只需要 ContractCaller（能发起 eth_call 即可）。
// 适用于只需要查询合约状态、不需要发送交易的场景。
func NewCounterAbiCaller(address common.Address, caller bind.ContractCaller) (*CounterAbiCaller, error) {
	contract, err := bindCounterAbi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CounterAbiCaller{contract: contract}, nil
}

// NewCounterAbiTransactor 创建只写合约实例。只需要 ContractTransactor（能发送交易即可）。
// 适用于只负责写入、不需要读取的场景（如仅用于提交交易的服务）。
func NewCounterAbiTransactor(address common.Address, transactor bind.ContractTransactor) (*CounterAbiTransactor, error) {
	contract, err := bindCounterAbi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CounterAbiTransactor{contract: contract}, nil
}

// NewCounterAbiFilterer 创建事件过滤实例。只需要 ContractFilterer（能查询日志即可）。
// 适用于纯事件监听场景，如链上事件通知服务。
func NewCounterAbiFilterer(address common.Address, filterer bind.ContractFilterer) (*CounterAbiFilterer, error) {
	contract, err := bindCounterAbi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CounterAbiFilterer{contract: contract}, nil
}

// bindCounterAbi 是内部辅助函数，将 ABI 元数据与地址、后端绑定为 BoundContract。
// BoundContract 是底层的核心结构，负责：
//   - 将函数名编码为选择器（selector）+ 参数 ABI 编码
//   - 发起 eth_call 或 eth_sendTransaction
//   - 解析返回值
//   - 过滤和解析事件日志
func bindCounterAbi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CounterAbiMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// ============================================================
// 底层 Raw 接口方法（一般不直接调用）
// ============================================================

// CounterAbiRaw.Call 通过底层 BoundContract 发起只读调用。
// 这是最底层的调用方式，传入方法名和参数，结果写入 result。
// 上层类型（Caller/Transactor）的方法内部就是调用这个。
func (_CounterAbi *CounterAbiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CounterAbi.Contract.CounterAbiCaller.contract.Call(opts, result, method, params...)
}

// CounterAbiRaw.Transfer 发起一笔纯转账交易（不指定函数调用）。
// 仅向合约发送 ETH，不触发任何函数。如果合约有 fallback/receive 函数则会被调用。
func (_CounterAbi *CounterAbiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CounterAbi.Contract.CounterAbiTransactor.contract.Transfer(opts)
}

// CounterAbiRaw.Transact 通过底层 BoundContract 发起写入交易。
// 传入方法名和参数，返回交易对象。
func (_CounterAbi *CounterAbiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CounterAbi.Contract.CounterAbiTransactor.contract.Transact(opts, method, params...)
}

// CounterAbiCallerRaw.Call 只读方的底层调用。
func (_CounterAbi *CounterAbiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CounterAbi.Contract.contract.Call(opts, result, method, params...)
}

// CounterAbiTransactorRaw.Transfer 只写方的底层转账。
func (_CounterAbi *CounterAbiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CounterAbi.Contract.contract.Transfer(opts)
}

// CounterAbiTransactorRaw.Transact 只写方的底层交易。
func (_CounterAbi *CounterAbiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CounterAbi.Contract.contract.Transact(opts, method, params...)
}

// ============================================================
// 合约函数绑定 — x() 读取计数器
// ============================================================
// Solidity: function x() view returns(uint256)
// 函数选择器: 0x0c55699c
// CounterAbiCaller.X 读取合约中的计数器值。
//
// 底层流程：
//  1. 将函数名 "x" 编码为 4 字节选择器 0x0c55699c
//  2. 发起 eth_call JSON-RPC 请求（不消耗 Gas）
//  3. 将返回的 ABI 编码数据解码为 *big.Int
//
// 参数 opts 可指定区块号（如查询历史状态），传 nil 则使用最新区块。
func (_CounterAbi *CounterAbiCaller) X(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CounterAbi.contract.Call(opts, &out, "x")

	if err != nil {
		return *new(*big.Int), err
	}

	// abi.ConvertType 将 interface{} 转换为 *big.Int
	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CounterAbiSession.X 通过 Session 读取计数器值。
// opts 已在创建 Session 时预置，此处无需传参，写法更简洁。
func (_CounterAbi *CounterAbiSession) X() (*big.Int, error) {
	return _CounterAbi.Contract.X(&_CounterAbi.CallOpts)
}

// CounterAbiCallerSession.X 通过只读 Session 读取计数器值。
func (_CounterAbi *CounterAbiCallerSession) X() (*big.Int, error) {
	return _CounterAbi.Contract.X(&_CounterAbi.CallOpts)
}

// ============================================================
// 合约函数绑定 — inc() 计数器 +1
// ============================================================
// Solidity: function inc() returns()
// 函数选择器: 0x371303c0
// 此方法会修改链上状态，需要发送交易并消耗 Gas。

// CounterAbiTransactor.Inc 调用合约的 inc() 函数，使计数器值 +1。
//
// 底层流程：
//  1. 将函数名 "inc" 编码为 4 字节选择器 0x371303c0（无参数，只需选择器）
//  2. 通过 TransactOpts（包含发送者地址、Gas 限制、私钥签名等）签名交易
//  3. 发送 eth_sendRawTransaction 到网络
//  4. 返回交易对象（此时交易尚未被打包，需等待确认）
//
// 参数 opts 必须包含：
//   - From:   发送者地址
//   - Signer: 签名函数（用于签署交易）
func (_CounterAbi *CounterAbiTransactor) Inc(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CounterAbi.contract.Transact(opts, "inc")
}

// CounterAbiSession.Inc 通过 Session 发送 inc 交易，opts 已预置。
func (_CounterAbi *CounterAbiSession) Inc() (*types.Transaction, error) {
	return _CounterAbi.Contract.Inc(&_CounterAbi.TransactOpts)
}

// CounterAbiTransactorSession.Inc 通过只写 Session 发送 inc 交易。
func (_CounterAbi *CounterAbiTransactorSession) Inc() (*types.Transaction, error) {
	return _CounterAbi.Contract.Inc(&_CounterAbi.TransactOpts)
}

// ============================================================
// 合约函数绑定 — incBy(uint256) 计数器 +N
// ============================================================
// Solidity: function incBy(uint256 by) returns()
// 函数选择器: 0x70119d06
// 此方法会修改链上状态，需要发送交易并消耗 Gas。

// CounterAbiTransactor.IncBy 调用合约的 incBy(by) 函数，使计数器值增加指定数量。
//
// 底层流程：
//  1. 将函数名 "incBy" 编码为 4 字节选择器 0x70119d06
//  2. 将参数 by（*big.Int）按 ABI 规范编码为 32 字节大端序
//  3. 选择器 + 参数编码拼接为完整的交易 data 字段
//  4. 签名并发送交易
//
// 示例：
//
//	tx, err := contract.IncBy(auth, big.NewInt(10)) // 计数器 +10
func (_CounterAbi *CounterAbiTransactor) IncBy(opts *bind.TransactOpts, by *big.Int) (*types.Transaction, error) {
	return _CounterAbi.contract.Transact(opts, "incBy", by)
}

// CounterAbiSession.IncBy 通过 Session 发送 incBy 交易。
func (_CounterAbi *CounterAbiSession) IncBy(by *big.Int) (*types.Transaction, error) {
	return _CounterAbi.Contract.IncBy(&_CounterAbi.TransactOpts, by)
}

// CounterAbiTransactorSession.IncBy 通过只写 Session 发送 incBy 交易。
func (_CounterAbi *CounterAbiTransactorSession) IncBy(by *big.Int) (*types.Transaction, error) {
	return _CounterAbi.Contract.IncBy(&_CounterAbi.TransactOpts, by)
}

// ============================================================
// 事件绑定 — Increment(uint256 by)
// ============================================================
// Solidity: event Increment(uint256 by)
// 事件哈希: 0x51af157c2eee40f68107a47a49c32fbbeb0a3c9e5cd37aa56e88e6be92368a81
//
// 此事件在每次调用 inc() 或 incBy() 时触发，记录本次增加了多少。
//
// 事件在以太坊中的工作原理：
//   - 合约执行 LOG 指令将事件数据写入收据的日志中
//   - 日志包含：事件哈希（topic[0]）、参数数据、合约地址、区块号等
//   - 未 indexed 的参数按 ABI 编码后直接存储在日志的 data 字段
//   - indexed 的参数（最多 3 个）作为 topic[1]~topic[3] 存储，可用于过滤
//   - 本事件的 by 参数未 indexed，因此只能读取，不能按值过滤

// CounterAbiIncrement 表示一条 Increment 事件的具体数据。
// By 字段记录了本次增加的数值，Raw 字段包含原始日志信息（区块号、交易哈希等）。
type CounterAbiIncrement struct {
	By  *big.Int
	Raw types.Log // 原始区块链日志，包含区块号、交易哈希、日志索引等上下文信息
}

// CounterAbiIncrementIterator 是事件日志的迭代器。
// 由 FilterIncrement 返回，用于逐条遍历匹配的事件日志。
//
// 使用模式：
//
//	iter, _ := contract.FilterIncrement(nil)
//	for iter.Next() {
//	    log.Println(iter.Event.By) // 处理每条事件
//	}
//	if err := iter.Error(); err != nil { ... } // 检查是否有错误
//	iter.Close() // 释放订阅资源
type CounterAbiIncrementIterator struct {
	Event *CounterAbiIncrement // 当前事件数据（调用 Next() 后填充）

	contract *bind.BoundContract // 用于解析日志数据的合约实例
	event    string              // 事件名称 "Increment"

	logs chan types.Log        // 接收合约事件的日志通道
	sub  ethereum.Subscription // 底层订阅，用于接收错误和终止信号
	done bool                  // 订阅是否已完成
	fail error                 // 迭代过程中发生的错误
}

// CounterAbiIncrementIterator.Next 将迭代器推进到下一条事件。
// 返回 true 表示找到了新事件（数据已填充到 it.Event），
// 返回 false 表示没有更多事件或发生了错误。
//
// 内部逻辑：
//  1. 如果之前迭代失败，直接返回 false
//  2. 如果订阅已完成，从通道中消费剩余日志
//  3. 否则阻塞等待新日志或订阅错误
//  4. 收到日志后调用 UnpackLog 将原始 ABI 数据解码为 Go 结构体
func (it *CounterAbiIncrementIterator) Next() bool {
	// 如果迭代失败，停止迭代
	if it.fail != nil {
		return false
	}
	// 如果订阅已完成，直接消费通道中剩余的日志
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CounterAbiIncrement)
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
	// 订阅仍在进行，等待新日志或错误
	select {
	case log := <-it.logs:
		it.Event = new(CounterAbiIncrement)
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

// CounterAbiIncrementIterator.Error 返回迭代过程中的错误。
// 应在 Next() 返回 false 后调用，确认是正常结束还是发生了错误。
func (it *CounterAbiIncrementIterator) Error() error {
	return it.fail
}

// CounterAbiIncrementIterator.Close 关闭迭代器，取消底层订阅，释放资源。
// 迭代完成后应调用此方法避免资源泄漏。
func (it *CounterAbiIncrementIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CounterAbiFilterer.FilterIncrement 查询历史 Increment 事件。
//
// 参数 opts 可指定：
//   - Start/End: 查询的区块范围
//   - Context: 用于取消查询的上下文
//
// 返回一个迭代器，逐条遍历匹配的事件日志。
//
// 示例 — 查询所有历史 Increment 事件：
//
//	iter, _ := contract.FilterIncrement(nil)
//	for iter.Next() {
//	    fmt.Printf("区块 %d: 增加了 %d\n", iter.Event.Raw.BlockNumber, iter.Event.By)
//	}
func (_CounterAbi *CounterAbiFilterer) FilterIncrement(opts *bind.FilterOpts) (*CounterAbiIncrementIterator, error) {

	logs, sub, err := _CounterAbi.contract.FilterLogs(opts, "Increment")
	if err != nil {
		return nil, err
	}
	return &CounterAbiIncrementIterator{contract: _CounterAbi.contract, event: "Increment", logs: logs, sub: sub}, nil
}

// CounterAbiFilterer.WatchIncrement 实时订阅新的 Increment 事件。
//
// 与 FilterIncrement 的区别：
//   - Filter: 查询已有的历史日志，查完就结束
//   - Watch: 持续监听新产生的事件，直到主动取消订阅
//
// 参数：
//
//	opts  — 可指定起始区块（传 nil 从最新区块开始）
//	sink  — 事件推送通道，新事件会发送到这里
//
// 返回一个 Subscription，调用 Cancel() 可停止订阅。
//
// 示例 — 实时监听：
//
//	sink := make(chan *CounterAbiIncrement)
//	sub, _ := contract.WatchIncrement(nil, sink)
//	defer sub.Unsubscribe()
//	for {
//	    select {
//	    case event := <-sink:
//	        fmt.Println("新事件: 增加了", event.By)
//	    case err := <-sub.Err():
//	        log.Fatal(err)
//	    }
//	}
func (_CounterAbi *CounterAbiFilterer) WatchIncrement(opts *bind.WatchOpts, sink chan<- *CounterAbiIncrement) (event.Subscription, error) {

	logs, sub, err := _CounterAbi.contract.WatchLogs(opts, "Increment")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// 新日志到达，解析事件数据并推送到用户通道
				event := new(CounterAbiIncrement)
				if err := _CounterAbi.contract.UnpackLog(event, "Increment", log); err != nil {
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

// CounterAbiFilterer.ParseIncrement 解析单条原始日志。
//
// 当你已经有 types.Log（比如从其他途径获取的日志数据），
// 可以用此方法将其解析为结构化的 CounterAbiIncrement。
//
// 示例：
//
//	log := ... // 从某处获得的原始日志
//	event, err := contract.ParseIncrement(log)
//	if err == nil {
//	    fmt.Println("增加了", event.By)
//	}
func (_CounterAbi *CounterAbiFilterer) ParseIncrement(log types.Log) (*CounterAbiIncrement, error) {
	event := new(CounterAbiIncrement)
	if err := _CounterAbi.contract.UnpackLog(event, "Increment", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
