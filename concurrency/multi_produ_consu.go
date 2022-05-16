package main

import "fmt"

// n producer and n consumer
func producer1(in chan<- int) {
	for i := 0; i < 6; i++ {
		in <- i
		fmt.Printf("producer put into channel %d\n", i)
	}
	close(in)
}

func consumer1(out <-chan int) {
	for i := range out {
		fmt.Printf("consumer get %d from channel\n", i)
	}
}

func main() {
	ch := make(chan int)
	go producer1(ch)
	consumer1(ch)
}
