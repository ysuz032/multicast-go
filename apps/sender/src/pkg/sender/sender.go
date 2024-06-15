package sender

import (
	"fmt"
	"net"
	"time"
)

// SendMulticast はマルチキャストメッセージを送信します
func SendMulticast(multicastAddr, multicastPort string) {
	// マルチキャストアドレスとポートを設定
	addr := fmt.Sprintf("%s:%s", multicastAddr, multicastPort)
	udpaddr, err := net.ResolveUDPAddr("udp4", addr)
	if err != nil {
		fmt.Println("failed to resolve multicast address:", err)
		return
	}

	// UDPコネクションの作成
	conn, err := net.DialUDP("udp4", nil, udpaddr)
	if err != nil {
		fmt.Println("failed to create conenction:", err)
		return
	}
	defer conn.Close()

	// 送信ループ
	for i := 0; i < 10; i++ {
		message := fmt.Sprintf("message %d", i)
		_, err = conn.Write([]byte(message))
		if err != nil {
			fmt.Println("send error:", err)
			return
		}
		fmt.Printf("send successful: %s\n", message)
		time.Sleep(1 * time.Second)
	}
}
