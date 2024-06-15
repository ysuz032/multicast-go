package receiver

import (
	"fmt"
	"net"
)

// ReceiveMulticast はマルチキャストメッセージを受信します
func ReceiveMulticast(multicastAddr, multicastPort string) {
	addr := fmt.Sprintf("%s:%s", multicastAddr, multicastPort)
	udpAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		fmt.Println("failed to resolve multicast address:", err)
		return
	}

	conn, err := net.ListenMulticastUDP("udp", nil, udpAddr)
	if err != nil {
		fmt.Println("failed to create socket:", err)
		return
	}
	defer conn.Close()

	buf := make([]byte, 1024)
	for {
		n, src, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("received error:", err)
			continue
		}
		fmt.Printf("received: %s from %s\n", string(buf[:n]), src)
	}
}
