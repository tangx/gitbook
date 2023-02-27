package selectdemo

import (
	"context"
	"fmt"
	"time"

	"k8s.io/apiextensions-apiserver/pkg/cmd/server"
)

func selectGrammar(ctx context.Context) {
	channel := make(chan int, 10)
	value := 100

	select {
	case <-ctx.Done():
		// 执行的代码
	case <-time.After(10 * time.Second):
		// 执行的代码
	case val := <-channel:
		fmt.Println(val)
		// 执行的代码
	case channel <- value:
		// 执行代码
	default:
		// 所有通道都没有准备好，执行的代码
	}
}

func foreverSelect(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			// 执行的代码
		case <-time.After(10 * time.Second):
			// 执行的代码
		}
	}
}

func foreverWaiting() {

	// 自己实现了一个服务器
	server.Run()

	// 阻塞以保持服务器始终运行
	select {}
}
