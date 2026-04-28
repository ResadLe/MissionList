package ethExec

import (
	"BlockchainTask/stage-01/task2/excabi"
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"os"
	"strconv"
)

// Inc 调用合约 inc()，计数器 +1。
func Inc(client *ethclient.Client, ctx context.Context) {

	c, err := excabi.NewCounterAbi(getContractAddr(), client)
	if err != nil {
		log.Fatalf("NewCounterAbi: %v", err)
	}

	auth, err := buildAuth(getPriKey(), client, getChainId(), ctx)
	if err != nil {
		log.Fatalf("buildAuth err：%v", err)
	}

	inc, err := c.Inc(auth)
	if err != nil {
		log.Fatalf("inc: %v", err)
	}
	fmt.Printf("inc tx: %s\n => Success", inc.Hash().Hex())
}

// IncBy 调用合约 incBy(by)，计数器 +by。
func IncBy(client *ethclient.Client, by *big.Int, ctx context.Context) {
	c, err := excabi.NewCounterAbi(getContractAddr(), client)
	if err != nil {
		log.Fatalf("new contract: %v", err)
	}

	auth, err := buildAuth(getPriKey(), client, getChainId(), ctx)
	if err != nil {
		log.Fatalf("buildAuth err: %v", err)
	}

	incBy, err := c.IncBy(auth, by)
	if err != nil {
		log.Fatalf("incBy: %v", err)
	}
	fmt.Printf("incBy(%d) tx: %s\n => Success", by, incBy.Hash().Hex())
}

func getPriKey() *ecdsa.PrivateKey {
	//初始化连接
	privKeyHex := os.Getenv("ETH_SEPOLIA_PRIVATE_KEY")
	if privKeyHex == "" {
		log.Fatal("PRIVATE_KEY env var required for write operations")
	}

	privKey, err := crypto.HexToECDSA(privKeyHex)
	if err != nil {
		log.Fatalf("parse private key: %v", err)
	}

	return privKey
}

func getContractAddr() common.Address {
	contractAddress := os.Getenv("SEPOLIA_CONTRACT_ADDRESS")
	if contractAddress == "" {
		log.Fatal("SEPOLIA_CONTRACT_ADDRESS not set in .env")
	}
	var contractAddr = common.HexToAddress(contractAddress)
	return contractAddr
}

func getChainId() int64 {
	chainIdStr := os.Getenv("SEPOLIA_CHAIN_ID")
	if chainIdStr == "" {
		log.Fatal("SEPOLIA_CHAIN_ID not set in .env")
	}
	chainId, err := strconv.ParseInt(chainIdStr, 10, 64)
	if err != nil {
		log.Fatal("转换失败:", err)
	}
	return chainId
}
