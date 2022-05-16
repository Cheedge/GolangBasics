package main

import (
	"fmt"
	"sync"
	"time"
)

/*lock and channel
 */

var mymutex sync.Mutex // only one mutex lock

func call_func(str string) {
	mymutex.Lock()
	for _, c := range str {
		fmt.Printf("%c", c)
		time.Sleep(time.Millisecond * 100)
	}
	mymutex.Unlock()
}

// var ch = make(chan int)

func call1() {
	call_func("hallo")
	// ch <- 1
}
func call2() {
	// <-ch
	call_func("world")
}

func main() {
	go call1()
	go call2()
	for {
	}
}
