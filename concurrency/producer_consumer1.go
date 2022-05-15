package main

import (
	"fmt"
)

func producer(chpro chan<- int) {
	for i := 0; i < 6; i++ {
		chpro <- i
		fmt.Printf("chpro = %+v\n", chpro)
		fmt.Printf("producer i=%d\n", i)
	}
	close(chpro)
}

func consumer(chcons <-chan int) {
	for ch := range chcons {
		// time.Sleep(10000)
		fmt.Printf("chcons = %+v\n", ch)
		fmt.Printf("consumer i=%d\n", ch)
	}

}

func main() {
	ch := make(chan int)
	// fmt.Printf("ch = %+v, add = %p\n", ch, &ch)
	// n gorouines(wrong!!)
	// for i := 0; i < 6; i++ {
	// 	go producer(ch, i)
	// 	go consumer(ch, i)
	// }
	go producer(ch)

	consumer(ch)

	// fmt.Printf("%d", <-ch)
	// close(ch)
}
