package main

import (
	"log"
	"net"
	"os"

	"golang.org/x/crypto/ssh"
)

func main() {

	config := &ssh.ServerConfig{
		// 密码验证回调函数
		PasswordCallback: func(c ssh.ConnMetadata, pass []byte) (*ssh.Permissions, error) {
			return nil, nil
		},
		// NoClientAuth:  true,                 // 客户端不验证, 任何客户端都可以连接
		// ServerVersion: "SSH-2.0-OWN-SERVER", // "SSH-2.0-"，SSH版本
	}

	// 秘钥用于SSH交互双方进行 Diffie-hellman 秘钥交换验证
	private, err := os.ReadFile("~/.ssh/id_rsa")
	if err != nil {
		log.Fatalln("failed to read private key")
		return
	}

	b, err := ssh.ParsePrivateKey(private)

	if err != nil {
		log.Fatalln("failed to parse private key")
		return
	}

	config.AddHostKey(b)

	listener, err := net.Listen("tcp", "0.0.0.0:22")
	if err != nil {
		log.Fatalf("failed to listen on 22 %s \n", err.Error())
		return
	}

	log.Println("listen on 0.0.0.0:22")

	// 接受所有连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("failed to accept incoming connection (%s) \n", err)
			continue
		}

		go handleConnection(conn)

	}

}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	//TODO:
}
