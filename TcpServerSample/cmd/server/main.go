package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// 监听本地端口
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("error listening:", err.Error())
		os.Exit(1)
	}

	defer listener.Close()

	fmt.Println("TCP server listening on :8080")

	for {
		// 等待客户端连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("error accepting:", err.Error())
			os.Exit(1)
		}

		// 处理客户端连接
		go handleMessage(conn)
	}

}

func handleMessage(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("reading message error:", err.Error())
			break
		}

		fmt.Print("message received:", string(buffer[:n]))

		conn.Write([]byte("message received \n"))
	}
}
