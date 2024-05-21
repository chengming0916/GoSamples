package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	// 监听地址和端口
	addr, err := net.ResolveUDPAddr("udp", ":8080")
	if err != nil {
		fmt.Println("error resolving UDP address:", err)
		os.Exit(1)
	}

	// 监听UDP端口
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("error listening on UDP port:", err)
		os.Exit(1)
	}

	defer conn.Close()

	fmt.Println("UDP server listening on ", addr)

	buffer := make([]byte, 1024)

	for {
		n, remoteAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("error reading from UDP:", err)
			continue
		}
		go func(n int, remoteAddr *net.UDPAddr, buffer []byte) {
			fmt.Printf("received %d bytes from %s: %s \n", n, remoteAddr.String(), string(buffer[:n]))
			time.Sleep(10 * time.Second) // 模拟耗时操作
			_, err = conn.WriteToUDP([]byte("message received"), remoteAddr)
			if err != nil {
				fmt.Println("error writing to UDP: ", err)
			}
		}(n, remoteAddr, buffer)

	}
}
