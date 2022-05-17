package main

import (
	"fmt"
	"net"
	"strings"
)

func concurr_handler(conn net.Conn) {
	defer conn.Close()
	addr := conn.RemoteAddr().String()
	fmt.Println(addr, "service has been connected")
	// buff
	buff := make([]byte, 4096)
	for {
		// read
		n, err := conn.Read(buff)
		if n == 0 {
			fmt.Println("client has been shutdown")
			break
		}
		if err != nil {
			fmt.Println("read err")
			return
		}
		// write

		conn.Write([]byte(strings.ToUpper(string(buff[:n]))))
		fmt.Println(string(buff[:n]))
	}
}

func main() {
	// listen
	listen, err := net.Listen("tcp", "127.0.0.1:3379")
	if err != nil {
		fmt.Println("listen err!")
		return
	}
	defer listen.Close()
	fmt.Println("waiting for clients connection")
	// need multiple conn socket
	for {
		// conn
		conn, _ := listen.Accept()
		// defer conn.Close() // Notice here need to close inside func concurr_handler()
		go concurr_handler(conn)
	}

}
