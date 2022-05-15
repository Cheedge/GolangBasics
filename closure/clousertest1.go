package main

import "fmt"

func myshift(base int) func(int) int {
	shif := 99
	return func(i int) int {
		base += shif + i
		fmt.Printf("inside func ==> %d, %p\n", base, &base)
		return base
	}
}

func main() {
	base := 100
	myfunc := myshift(base)
	myfunc(2)
	fmt.Printf("main ==> %d, %p\n", base, &base)
	myfunc(3)
	fmt.Printf("main ==> %d, %p\n", base, &base)
}
