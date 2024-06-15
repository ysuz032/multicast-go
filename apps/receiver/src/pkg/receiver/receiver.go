package receiver

import (
	"fmt"
	"net"

	"golang.org/x/net/ipv4"
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

	// マルチキャストグループの設定
	group := net.ParseIP(multicastAddr)

	// UDPソケットの作成
	addr := fmt.Sprintf("%s:%s", multicastAddr, multicastPort)
	conn, err := net.ListenPacket("udp4", addr)
	if err != nil {
		fmt.Println("failed to create socket:", err)
		return
	}
	defer conn.Close()

	// マルチキャストグループに参加
	p := ipv4.NewPacketConn(conn)
	if err := p.JoinGroup(iface, &net.UDPAddr{IP: group}); err != nil {
		fmt.Println("failed to join multicast group:", err)
		return
	}
	defer p.LeaveGroup(iface, &net.UDPAddr{IP: group})

	// アプリケーションは、カーネル内のプロトコルスタック間でパケットごとの制御メッセージの送信を設定することがあるため
	// アプリケーションが受信パケットの宛先アドレスを必要とする場合、SetControlMessage of PacketConn を使用して制御メッセージの送信を有効にする
	if err := p.SetControlMessage(ipv4.FlagDst, true); err != nil {
		fmt.Println("failed to enable control message transmissions:", err)
		return
	}

	// 受信ループ
	buffer := make([]byte, 1500)
	for {
		n, cm, src, err := p.ReadFrom(buffer)
		if err != nil {
			fmt.Println("received error:", err)
			continue
		}
		if cm.Dst.IsMulticast() {
			if cm.Dst.Equal(group) {
				fmt.Printf("received: %s from %s\n", string(buffer[:n]), src)
			} else {
				fmt.Printf("skipped: %s from %s\n", string(buffer[:n]), src)
				continue
			}
		} else {
			fmt.Printf("skipped: not multicast")
			continue
		}
	}
}
