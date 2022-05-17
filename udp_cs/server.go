package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	srvAddr, _ := net.ResolveUDPAddr("udp", "127.0.0.1:3379")
	// conn
	conn, err := net.ListenUDP("udp", srvAddr)
	if err != nil {
		fmt.Println("listen err")
		return
	}
	defer conn.Close()
	fmt.Println("waiting for connection...")

	// buff
	buff := make([]byte, 4096)
	for {
		// read
		n, clientAddr, err := conn.ReadFromUDP(buff)
		if err != nil {
			fmt.Println("conn.Read err")
			break
		}
		fmt.Println(clientAddr, "connected")
		// write
		str := time.Now().String() + "server recieve: " + string(buff[:n])
		conn.WriteToUDP([]byte(str), clientAddr)

	}

}
