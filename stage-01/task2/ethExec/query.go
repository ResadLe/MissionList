package ethExec

import (
	"BlockchainTask/stage-01/task2/excabi"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"os"
)

// GetX 查询合约计数器值。
func GetX(client *ethclient.Client) {
	contractAddress := os.Getenv("SEPOLIA_CONTRACT_ADDRESS")
	if contractAddress == "" {
		log.Fatal("SEPOLIA_CONTRACT_ADDRESS not set in .env")
	}
	var contractAddr = common.HexToAddress(contractAddress)
	c, err := excabi.NewCounterAbi(contractAddr, client)
	if err != nil {
		log.Fatalf("failed to get NewCounterAbi: %v", err)
	}
	x, err := c.X(nil)
	if err != nil {
		log.Fatalf("failed to get X: %v", err)
	}
	fmt.Println("X = ", x)
}
