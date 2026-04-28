package main

import (
	"BlockchainTask/stage-01/connEth"
	"BlockchainTask/stage-01/task2/ethExec"
	"flag"
	"fmt"
	"log"
	"math/big"
	"time"
)

func main() {
	var queryMode bool
	var incMode bool
	var incByValue uint64

	flag.BoolVar(&queryMode, "q", false, "查询模式：查询计数器值")
	flag.BoolVar(&incMode, "i", false, "增加模式：计数器 +1")
	flag.Uint64Var(&incByValue, "p", 0, "增加模式：计数器 +指定值")
	flag.Parse()

	//模式校验
	if !queryMode && !incMode && incByValue == 0 {
		log.Fatal("必须指定 -q（查询）、-i（+1）或 -p <值>（+N）")
	}

	// 初始化连接
	ctx, cancel := connEth.NewTimeoutContext(60 * time.Second)
	defer cancel()

	client, _ := connEth.ConnectClient(ctx)
	defer client.Close()

	//查询模式
	if queryMode {
		fmt.Println("=== Query Mode ===")
		ethExec.GetX(client)
	}

	if incMode {
		ethExec.Inc(client, ctx)
	} else if incByValue > 0 {
		ethExec.IncBy(client, big.NewInt(int64(incByValue)), ctx)
	}
}
