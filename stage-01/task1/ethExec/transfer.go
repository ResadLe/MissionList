package ethExec

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// RunTransfer 发送ETH转账并等待确认
func RunTransfer(ctx context.Context, client *ethclient.Client, amountEth string, toAddress string) {
	// 输出ChainID
	chainID, err := client.ChainID(ctx)
	if err != nil {
		log.Fatalf("failed to get chain id: %v", err)
	}
	fmt.Printf("Current ChainID: %s\n", chainID.String())

	// 加载私钥
	myPrivateKeyHex := os.Getenv("ETH_SEPOLIA_PRIVATE_KEY")
	if myPrivateKeyHex == "" {
		log.Fatal("ETH_SEPOLIA_PRIVATE_KEY not set in .env")
	}

	myPrivateKey, err := crypto.HexToECDSA(trim0x(myPrivateKeyHex))
	if err != nil {
		log.Fatalf("failed to parse private key: %v", err)
	}

	//获取发送方地址
	myPublicKey := myPrivateKey.Public()
	myPublicKeyECDSA, ok := myPublicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("failed to get public key")
	}

	toAddr := common.HexToAddress(toAddress)

	senderAddress := crypto.PubkeyToAddress(*myPublicKeyECDSA)
	fmt.Printf("From: %s\n", senderAddress.Hex())
	fmt.Printf("To:   %s\n", toAddress)
	fmt.Printf("Amount: %s ETH\n", amountEth)

	//---------------------------------------------------------------

	// 获取 nonce
	nonce, err := client.PendingNonceAt(ctx, senderAddress)
	if err != nil {
		log.Fatalf("failed to get nonce: %v", err)
	}

	// 获取建议的 Gas 价格（使用 EIP-1559 动态费用）
	gasTipCap, err := client.SuggestGasTipCap(ctx)
	if err != nil {
		log.Fatalf("failed to get gas tip cap: %v", err)
	}

	// 获取 base fee，计算 fee cap
	header, err := client.HeaderByNumber(ctx, nil)
	if err != nil {
		log.Fatalf("failed to get header: %v", err)
	}

	baseFee := header.BaseFee
	if baseFee == nil {
		// 如果不支持 EIP-1559，使用传统 gas price
		gasPrice, err := client.SuggestGasPrice(ctx)
		if err != nil {
			log.Fatalf("failed to get gas price: %v", err)
		}
		baseFee = gasPrice
	}

	// fee cap = base fee * 2 + tip cap（简单策略）
	gasFeeCap := new(big.Int).Add(
		new(big.Int).Mul(baseFee, big.NewInt(2)),
		gasTipCap,
	)

	// 估算 Gas Limit（普通转账固定为 21000）
	gasLimit := uint64(21000)

	// 转换 ETH 金额为 Wei
	// amountEth * 1e18
	// 转换金额：ETH -> Wei
	amount, ok := new(big.Float).SetString(amountEth)
	if !ok {
		log.Fatalf("invalid amount: %s", amountEth)
	}
	amountWei := new(big.Float).Mul(amount, big.NewFloat(1e18))
	valueWei, _ := amountWei.Int(nil)

	// 检查余额是否足够
	balance, err := client.BalanceAt(ctx, senderAddress, nil)
	if err != nil {
		log.Fatalf("failed to get balance: %v", err)
	}

	// 计算总费用：value + gasFeeCap * gasLimit
	totalCost := new(big.Int).Add(
		valueWei,
		new(big.Int).Mul(gasFeeCap, big.NewInt(int64(gasLimit))),
	)

	if balance.Cmp(totalCost) < 0 {
		log.Fatalf("insufficient balance: have %s wei, need %s wei", balance.String(), totalCost.String())
	}

	// 构造交易（EIP-1559 动态费用交易）
	txData := &types.DynamicFeeTx{
		ChainID:   chainID,
		Nonce:     nonce,
		GasTipCap: gasTipCap,
		GasFeeCap: gasFeeCap,
		Gas:       gasLimit,
		To:        &toAddr,
		Value:     valueWei,
		Data:      nil,
	}
	tx := types.NewTx(txData)

	// 签名交易
	signer := types.NewLondonSigner(chainID)
	signedTx, err := types.SignTx(tx, signer, myPrivateKey)
	if err != nil {
		log.Fatalf("failed to sign transaction: %v", err)
	}

	// 发送交易
	if err := client.SendTransaction(ctx, signedTx); err != nil {
		log.Fatalf("failed to send transaction: %v", err)
	}

	// 输出交易信息
	fmt.Println("=== Transaction Sent ===")
	fmt.Printf("From       : %s\n", senderAddress.Hex())
	fmt.Printf("To         : %s\n", toAddr.Hex())
	fmt.Printf("Value      : %s ETH (%s Wei)\n", fmt.Sprintf("%.6f", amountEth), valueWei.String())
	fmt.Printf("Gas Limit  : %d\n", gasLimit)
	fmt.Printf("Gas Tip Cap: %s Wei\n", gasTipCap.String())
	fmt.Printf("Gas Fee Cap: %s Wei\n", gasFeeCap.String())
	fmt.Printf("Nonce      : %d\n", nonce)
	fmt.Printf("Tx Hash    : %s\n", signedTx.Hash().Hex())
	fmt.Println("\nTransaction is pending. Use --tx flag to query status:")
	fmt.Printf("  go run main.go --tx %s\n", signedTx.Hash().Hex())

	fmt.Println("=======================================================")
	fmt.Printf("Tx Hash: %s\n", signedTx.Hash().Hex())

	// 等待确认
	fmt.Println("Waiting for confirmation...")
	receipt, err := bind.WaitMined(ctx, client, signedTx)
	if err != nil {
		log.Fatalf("failed to wait for confirmation: %v", err)
	}

	fmt.Printf("Confirmed! Block: %d\n", receipt.BlockNumber)
	fmt.Printf("Gas Used: %d\n", receipt.GasUsed)
	fmt.Printf("Status: %d (1=success, 0=failed)\n", receipt.Status)
}

// trim0x 移除十六进制字符串前缀 "0x"
func trim0x(s string) string {
	if len(s) >= 2 && s[:2] == "0x" {
		return s[2:]
	}
	return s
}
