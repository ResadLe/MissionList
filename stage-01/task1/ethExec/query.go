package ethExec

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
)

// RunQuery 查询区块信息
func RunQuery(ctx context.Context, client *ethclient.Client, blockNumber string) {
	// 输出ChainID
	chainID, err := client.ChainID(ctx)
	if err != nil {
		log.Fatalf("failed to get chain id: %v", err)
	}
	fmt.Printf("Current ChainID: %s\n", chainID.String())

	// 转换参数并获取区块头
	var blockNumberInput *big.Int
	if blockNumber == "nil" || blockNumber == "" {
		blockNumberInput = nil
	} else {
		n, ok := new(big.Int).SetString(blockNumber, 10)
		if !ok {
			log.Fatalf("invalid block number: %s", blockNumber)
		}
		blockNumberInput = n
	}

	header, err := client.HeaderByNumber(ctx, blockNumberInput)
	if err != nil {
		log.Fatalf("failed to get block header: %v", err)
	}

	fmt.Printf("User input block number: %s\n", blockNumber)
	fmt.Printf("Current block number: %s\n", header.Number)
	fmt.Printf("Block Hash    : %s\n", header.Hash().Hex())
	fmt.Printf("Block Time    : %s\n", time.Unix(int64(header.Time), 0).Format(time.RFC3339))

	// 查询区块交易数
	count, err := client.TransactionCount(ctx, header.Hash())
	if err != nil {
		log.Fatalf("failed to get transaction count: %v", err)
	}
	fmt.Printf("Transaction Count: %d\n", count)
}
