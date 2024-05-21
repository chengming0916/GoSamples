package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	svrAddr := "127.0.0.1:8080"

	// 解析UDP地址
	addr, err := net.ResolveUDPAddr("udp", svrAddr)
	if err != nil {
		fmt.Println("error resolving UDP address:", err)
		os.Exit(1)
	}

	// 创建连接
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Println("error connecting to UDP server:", err)
		os.Exit(1)
	}

	defer conn.Close()

	go func(conn *net.UDPConn) {
		for {
			_, err = conn.Write([]byte("Hello, UDP server"))
			time.Sleep(1 * time.Second)
			if err != nil {
				fmt.Println("error sending to UDP:", err)
				os.Exit(1)
			}
		}
	}(conn)

	// 读取响应
	buffer := make([]byte, 1024)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("error reading from UDP:", err)
			os.Exit(1)
		}

		fmt.Printf("Received %d bytes: %s \n", n, string(buffer[:n]))
	}
}
