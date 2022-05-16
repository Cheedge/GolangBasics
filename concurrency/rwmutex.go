package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var rwmutex sync.RWMutex

func producer3(ch chan<- int, idx int) {
	for i := 0; i < 3; i++ {
		num := rand.Intn(999)
		rwmutex.Lock()
		fmt.Printf("prod%d-%d: %d\n", idx, i, num)
		ch <- num
		rwmutex.Unlock()
	}
	// defer close(ch)
}

func consumer3(ch <-chan int, idx int) {
	// for i := range ch {
	// 	rwmutex.RLock()
	// 	fmt.Printf("consu: %d\n", i)
	// 	rwmutex.RUnlock()
	// }
	for i := 0; i < 3; i++ {
		// rwmutex.RLock()
		num := <-ch
		fmt.Printf("con%d: %d\n", idx, num)
		// rwmutex.RUnlock()
	}
}

func main() {
	ch := make(chan int)
	for i := 0; i < 2; i++ {
		go producer3(ch, i)
	}
	for i := 0; i < 2; i++ {
		go consumer3(ch, i)
	}
	time.Sleep(time.Second * 5)
	// for {
	// }
}
