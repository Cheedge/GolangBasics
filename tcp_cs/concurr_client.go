package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// conn socket
	conn, err := net.Dial("tcp", "127.0.0.1:3379")

	if err != nil {
		fmt.Println("net.Dial err")
		return
	}
	defer conn.Close()
	// write
	go func() {
		// buff
		buff_w := make([]byte, 4096)
		for {
			n, _ := os.Stdin.Read(buff_w)
			conn.Write(buff_w[:n])
		}
	}()

	// read
	buff_r := make([]byte, 4096)
	for {
		n, err := conn.Read(buff_r)
		// avoid close server, then continue writing
		if n == 0 {
			fmt.Println("server has been shut down")
			break
		}
		if err != nil {
			fmt.Println("read err")
			break
		}
		fmt.Println(string(buff_r[:n]))
	}
}
