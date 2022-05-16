package main

import (
	"fmt"
	"math/rand"
	"time"
)

func producer2(in chan<- int, idx int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 3; i++ {
		r := rand.Intn(999)
		fmt.Printf("===producer%d-%d put %d\n", idx, i, r)
		in <- r
	}
	// defer close(in)
}
func consumer2(out <-chan int, idx int) {
	for i := range out {
		fmt.Printf("consumer%d get %d\n", idx, i)
		// time.Sleep(time.Second * 3)
	}
}
func main() {
	ch := make(chan int)
	// rand.Seed(time.Now().UnixNano())
	for i := 0; i < 6; i++ {
		go producer2(ch, i)
	}

	for i := 0; i < 6; i++ {
		go consumer2(ch, i)
	}
	// time.Sleep(time.Second * 10)
	for {
	}
}
