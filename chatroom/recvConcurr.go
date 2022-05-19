package main

import "net"

func connTest(conn net.Conn) {
	conn.Write([]byte("hahaha"))
}

func listenTest(listen net.Listener) {
	conn, _ := listen.Accept()
	conn.Write([]byte("hahaha"))
}

func main() {
	listen, _ := net.Listen("tcp", "127.0.0.1:3399")
	for {
		conn, _ := listen.Accept()
		// go listenTest(listen)
		go connTest(conn)
	}
}
