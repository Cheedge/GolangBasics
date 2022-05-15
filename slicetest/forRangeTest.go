package main

import "fmt"

func forrange1(slc []int) {
	for i := range slc {
		fmt.Printf("%d\n", i)
	}
}

func main() {
	slc := []int{2, 3, 4, 5, 6}
	forrange1(slc)
}
