package main

import (
	"fmt"
)

func vanillaFib(i int, num *[]int) int {
	if len(*num) <= i && i > 2 {
		n := vanillaFib(i-2, num) + vanillaFib(i-1, num)
		*num = append(*num, n)
		fmt.Printf("%v\n", num)
		// return (*num)[i]
	}

	return (*num)[i]
}

func dpFib(i int) int {
	k := 2
	num := []int{0, 1}
	for k <= i {
		tmp := num[k-2] + num[k-1]
		num = append(num, tmp)
		fmt.Printf("%v\n", num)
		k++
	}
	return num[k-1]
}

// here int is int64 (64bite), so only 2^64=10^19
func concurrFib(n int, ch chan<- int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}

func main() {
	n := 100
	slc := []int{0, 1, 1}
	// res := vanillaFib(n, &slc)
	// fmt.Printf(" res is %d\n %v\n", res, slc)
	res1 := dpFib(n)
	fmt.Printf(" res is %d\n %v\n", res1, slc)
	ch := make(chan int)
	go concurrFib(n, ch)
	for c := range ch {
		fmt.Printf("%d\n", c)
	}

}
