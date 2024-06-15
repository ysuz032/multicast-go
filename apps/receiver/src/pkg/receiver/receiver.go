package receiver

import (
	"fmt"
	"net"
)

// GetDefaultInterface gets the default network interface
func GetDefaultInterface() (string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range interfaces {
		if (iface.Flags&net.FlagUp) != 0 && (iface.Flags&net.FlagLoopback) == 0 {
			return iface.Name, nil
		}
	}
	return "", nil
}

// ReceiveMulticast はマルチキャストメッセージを受信します
func ReceiveMulticast(multicastIface, multicastAddr, multicastPort string) {

	// インターフェース名を取得
	if multicastIface == "" {
		var err error
		multicastIface, err = GetDefaultInterface()
		if err != nil {
			fmt.Println("failed to get default Interface", err)
			return
		}
	}
	fmt.Println("use interface:", multicastIface)

	// インターフェースを取得
	iface, err := net.InterfaceByName(multicastIface)
	if err != nil {
		fmt.Println("failed to resolve Interface", err)
		return
	}

	// マルチキャストアドレスとポートを設定
	addr := fmt.Sprintf("%s:%s", multicastAddr, multicastPort)
	udpAddr, err := net.ResolveUDPAddr("udp4", addr)
	if err != nil {
		fmt.Println("failed to resolve multicast address:", err)
		return
	}

	// UDPソケットの作成
	conn, err := net.ListenMulticastUDP("udp4", iface, udpAddr)
	if err != nil {
		fmt.Println("failed to create socket:", err)
		return
	}
	defer conn.Close()

	// 受信ループ
	buffer := make([]byte, 1024)
	for {
		n, src, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("received error:", err)
			continue
		}
		fmt.Printf("received: %s from %s", string(buffer[:n]), src)
	}
}
