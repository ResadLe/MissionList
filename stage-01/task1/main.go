package main

import (
	"BlockchainTask/stage-01/connEth"
	"flag"
	"fmt"
	"log"
	"time"

	"BlockchainTask/stage-01/task1/ethExec"
)

func main() {
	var queryMode bool
	var transferMode bool
	var number string
	var receiver string

	flag.BoolVar(&queryMode, "q", false, "查询模式：查询区块信息")
	flag.BoolVar(&transferMode, "t", false, "交易模式：发送ETH转账")
	flag.StringVar(&number, "n", "", "查询模式: 区块号 / 交易模式: 金额(ETH)")
	flag.StringVar(&receiver, "r", "", "接收者地址（仅交易模式）")
	flag.Parse()

	// 模式校验
	if !queryMode && !transferMode {
		log.Fatal("必须指定 -q（查询模式）或 -t（交易模式）")
	}
	if queryMode && transferMode {
		log.Fatal("-q 和 -t 不能同时指定")
	}
	if transferMode && (number == "" || receiver == "") {
		log.Fatal("交易模式需要 -n（金额）和 -r（接收者地址）")
	}

	// 初始化连接
	ctx, cancel := connEth.NewTimeoutContext(60 * time.Second)
	defer cancel()

	client, _ := connEth.ConnectClient(ctx)
	defer client.Close()

	// 执行任务
	if queryMode {
		fmt.Println("=== Query Mode ===")
		ethExec.RunQuery(ctx, client, number)
	} else if transferMode {
		fmt.Println("=== Transfer Mode ===")
		ethExec.RunTransfer(ctx, client, number, receiver)
	}
}
