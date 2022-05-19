package main

import (
	"fmt"
	"net"
	"os"
)

func errHandle(err error, errMsg string) {
	if err != nil {
		fmt.Println(errMsg, err)
		os.Exit(-1)
	}
}

type User struct {
	addr string
	name string
	ch   chan string
}

func newUser(ad string, na string) User {
	var newuser User
	newuser.addr = ad
	newuser.name = na
	newuser.ch = make(chan string)
	return newuser
}

// var mapList map[string]User
var mapList = make(map[string]User)

// func write2Clients(conn net.Conn, mapList map[string]User) {
// 	for _, u := range mapList {
// 		conn.Write([]byte(<-u.ch))
// 	}
// }

func write2Clients(conn net.Conn, client User) {
	// c <-client.ch is not block
	for c := range client.ch {
		conn.Write([]byte(c))
	}
}

func Manager(MSGch <-chan string, mapList map[string]User) {
	for {
		msg := <-MSGch
		for _, u := range mapList {
			u.ch <- msg
			fmt.Println("Manager msg is ", msg)
		}
	}
}

func connHandler(conn net.Conn, MSGch chan string, mapList map[string]User, client User) {
	// address := conn.RemoteAddr()
	// client := newUser(address.String(), address.String())
	// mapList[client.addr] = client
	// buff
	buff := make([]byte, 4096)
	for {
		// fmt.Println("inside for loop")
		n, err := conn.Read(buff)
		errHandle(err, "conn read err")
		// fmt.Printf("MSGch: %+v\n", MSGch)
		// fmt.Printf("map: %+v\n", mapList)
		// avoid dead lock 1st write into channel
		go Manager(MSGch, mapList)
		MSGch <- client.name + ":  " + string(buff[:n])
		// fmt.Println("after channel <-")
		// user.ch -> client
		// go write2Clients(conn, mapList)
		go write2Clients(conn, client)
		// fmt.Println("after write2clients")
	}
}

func main() {
	// listen
	listen, err := net.Listen("tcp", "127.0.0.1:3399")
	errHandle(err, "listen err")
	defer listen.Close()
	ch := make(chan string)
	// fmt.Printf("ch: %+v\n", ch)
	// mapList = make(map[string]User)
	for {
		conn, err := listen.Accept()
		errHandle(err, "accept err")
		defer conn.Close()
		address := conn.RemoteAddr()
		if _, ok := mapList[address.String()]; ok == false {
			client := newUser(address.String(), address.String())

			// broadcast online member to all others
			// onlinemsg := address.String() + "is online now"
			// for _, u := range mapList {
			// 	u.ch <- onlinemsg
			// 	conn.Write([]byte(<-u.ch))
			// }
			mapList[client.addr] = client
		}

		go connHandler(conn, ch, mapList, mapList[address.String()])
	}
}
