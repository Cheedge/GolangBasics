package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	for i := 0; i < 6; i++ {
		ch <- i
		fmt.Println(<-ch)
	}
}
