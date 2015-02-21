package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	server, err := net.Listen("tcp", ":12321")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := server.Accept()
		fmt.Println("Connection Received!")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintln(conn, "goodbye!")
		conn.Close()
	}
}
