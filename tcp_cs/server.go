package main

import (
	"fmt"
	"net"
)

func main() {
	// listen
	socket, _ := net.Listen("tcp", "127.0.0.1:3379")
	// conn
	conn, _ := socket.Accept()
	// buff
	buff := make([]byte, 4096)
	for {
		// read
		n, err := conn.Read(buff)
		if err != nil {
			fmt.Println("read err")
			return
		}
		// write
		conn.Write(buff[:n])
		fmt.Println(string(buff[:n]))
	}

	// close
	// defer conn.Close()
}
