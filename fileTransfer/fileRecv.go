package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// listen
	listen, err := net.Listen("tcp", "127.0.0.1:3389")
	if err != nil {
		fmt.Println("listen err")
		return
	}
	defer listen.Close()
	// conn
	conn, err := listen.Accept()
	if err != nil {
		fmt.Println("accept err")
		return
	}
	defer conn.Close()
	// send ok -> client
	conn.Write([]byte("ok"))
	// recv file name
	cmd := os.Args
	if len(cmd) != 2 {
		fmt.Println("command like: go run ***.go filename, now only have", len(cmd), " args")
		fmt.Printf("%v", cmd)
		return
	}
	// create file
	f, err := os.Create(cmd[1])
	if err != nil {
		fmt.Println("open file err")
		return
	}
	defer f.Close()
	// recv file content
	buff_file := make([]byte, 4096)
	for {
		// socket -r-> buff
		n, _ := conn.Read(buff_file)
		// file use io.EOF, conn read use n==0
		if n == 0 {
			fmt.Println("socket read finished")
			return
		}
		// buff -w-> file
		_, err = f.Write(buff_file[:n])
		if err != nil {
			fmt.Println("file write err")
			return
		}
	}
}
