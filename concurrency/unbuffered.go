package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 6; i++ {
			fmt.Printf("goroutine i=%d\n", i)
			ch <- i
		}
		close(ch) // for range, can see close must be set
	}()
	time.Sleep(10)
	// for i := 0; i < 6; i++ {
	// 	// <-ch
	// 	fmt.Printf("main i=%d\n", <-ch)
	// }
	// this situation can see close must be set else:fatal error: all goroutines are asleep - deadlock!
	for i := range ch {
		// <-ch
		fmt.Printf("main i=%d\n", i)
	}
}
