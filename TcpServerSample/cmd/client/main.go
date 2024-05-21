package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("error connecting:", err.Error())
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("connected to TCP server")

	go func(conn net.Conn) {
		for {
			_, err := conn.Write([]byte("hello server \n"))
			if err != nil {
				fmt.Println("error sending data: ", err.Error())
				break
			}

			time.Sleep(2 * time.Second)
		}
	}(conn)

	buffer := make([]byte, 1024)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("error reading from server:", err.Error())
			break
		}

		fmt.Printf("server response: %s", string(buffer[:n]))

	}
}
