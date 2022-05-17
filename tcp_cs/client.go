package main

import (
	"fmt"
	"net"
)

func main() {
	// conn
	conn, err := net.Dial("tcp", "127.0.0.1:3379")
	if err != nil {
		fmt.Println("dial error: ", err)
		return
	}
	// buff
	buff := make([]byte, 4096)
	// write
	conn.Write([]byte("client ask for connect..."))

	// read
	n, err := conn.Read(buff)
	if err != nil {
		fmt.Println("read err")
		return
	}
	fmt.Println(string(buff[:n]))

	// close
	defer conn.Close()
}
