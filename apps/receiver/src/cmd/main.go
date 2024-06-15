package main

import (
	"flag"
	"fmt"
	"os"
	"sync"

	"receiver/pkg/receiver"
)

func main() {
	// 環境変数からアドレスとポートを取得
	multicastAddr := os.Getenv("MULTICAST_ADDR")
	multicastPort := os.Getenv("MULTICAST_PORT")

	// 引数からインターフェース名を取得
	multicastIface := flag.String("iface", "", "network interface to use for multicast")
	flag.Parse()

	if multicastAddr == "" || multicastPort == "" {
		fmt.Println("Not found MULTICAST_ADDR OR MULTICAST_PORT in environment variables")
		return
	}

	// WaitGroupを作成
	var wg sync.WaitGroup

	// マルチキャスト受信をゴルーチンで実行
	wg.Add(1)
	go func() {
		defer wg.Done()
		receiver.ReceiveMulticast(*multicastIface, multicastAddr, multicastPort)
	}()

	// 全てのゴルーチンの終了を待つ
	wg.Wait()
}
