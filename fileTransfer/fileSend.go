package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

// file send as client in tcp model
func main() {
	// get input
	cmd := os.Args
	if len(cmd) != 2 {
		fmt.Println("command like: go run ***.go args, now only have", len(cmd), " args")
		fmt.Printf("%v", cmd)
		return
	}
	file_path := cmd[1]
	// file name
	fileInfo, err := os.Stat(file_path)
	if err != nil {
		fmt.Println("os.State err")
		return
	}
	fileName := fileInfo.Name()
	// conn socket
	conn, err := net.Dial("tcp", "127.0.0.1:3389")
	if err != nil {
		fmt.Println("net.Dial err")
		return
	}
	defer conn.Close()
	// send file name -> server
	conn.Write([]byte(fileName))
	// read ok <- server send
	// buff
	buff := make([]byte, 1024)
	n, err := conn.Read(buff)

	if err != nil {
		fmt.Println("read error", string(buff))
		return
	}
	if string(buff[:n]) != "ok" {
		fmt.Println("server not ok")
		return
	} else {
		fileSend(file_path, conn)
	}
}

func fileSend(filePath string, conn net.Conn) {
	// read from file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("open error")
	}
	defer file.Close()
	buff := make([]byte, 4096)
	for {
		// file -read-> buff
		n, err := file.Read(buff)
		if err != nil {
			fmt.Println("file read error")
			// file use io.EOF, conn use n==0
			if err == io.EOF {
				fmt.Println("read finished")
				return
			}
		}
		// wirte to socket: buff -write-> socket
		_, err = conn.Write(buff[:n])
		if err != nil {
			fmt.Println("conn write err")
			return
		}
	}
}
