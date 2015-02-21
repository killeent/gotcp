package main

import (
	"fmt"
	"net"
)

func main() {
	server, err := net.Listen("tcp", ":12321")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		conn, err := server.Accept()
		fmt.Println("Connection Received!")
		if err != nil {
			fmt.Println(err)
		}
		conn.Write([]byte("goodbye\n"))
		conn.Close()
	}
}
