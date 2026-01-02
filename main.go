package main

import (
	"fmt"
	"net"
)

func main() {

	listener, err := net.Listen("tcp", "6380")
	if err != nil {
		fmt.Println("error starting server")
		return
	}

	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("Error accepting Connection")
		return
	}

	buf := make([]byte, 1024)

	conn.Read(buf)

}
