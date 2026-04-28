package connEth

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

// ConnectClient 加载.env并连接以太坊节点
func ConnectClient(ctx context.Context) (*ethclient.Client, string) {
	// 加载.env文件
	if err := godotenv.Load(); err != nil {
		log.Fatalf("failed to load .env file: %v", err)
	}

	// 获取 RPC URL
	rpcURL := os.Getenv("ETH_SEPOLIA_INFURA_HTTPS_URL")
	if rpcURL == "" {
		log.Fatal("ETH_SEPOLIA_INFURA_HTTPS_URL not set in .env")
	}

	// 连接节点
	client, err := ethclient.DialContext(ctx, rpcURL)
	if err != nil {
		log.Fatalf("failed to connect to Ethereum node: %v", err)
	}

	return client, rpcURL
}

// NewTimeoutContext 创建带超时的上下文
func NewTimeoutContext(timeout time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), timeout)
}
