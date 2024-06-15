package main

import (
	"fmt"
	"os"
	"sender/pkg/sender"
	"sync"
)

func main() {
	// 環境変数からアドレスとポートを取得
	multicastAddr := os.Getenv("MULTICAST_ADDR")
	multicastPort := os.Getenv("MULTICAST_PORT")

	if multicastAddr == "" || multicastPort == "" {
		fmt.Println("Not found MULTICAST_ADDR OR MULTICAST_PORT in environment variables")
		return
	}

	// WaitGroupを作成
	var wg sync.WaitGroup

	// マルチキャスト送信をゴルーチンで実行
	wg.Add(1)
	go func() {
		defer wg.Done()
		sender.SendMulticast(multicastAddr, multicastPort)
	}()

	// 全てのゴルーチンの終了を待つ
	wg.Wait()
}
