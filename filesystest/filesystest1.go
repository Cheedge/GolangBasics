package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f1, err := os.Open("/home/sharma/Desktop/DeepLearning/Golang/filesystest/awk_a12.dat")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f1.Close()
	f2, err := os.Open("/home/sharma/Desktop/DeepLearning/Golang/filesystest/awk_a13.dat")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f2.Close()
	f3, err := os.Open("/home/sharma/Desktop/DeepLearning/Golang/filesystest/awk_a14.dat")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f3.Close()
	fnew, err := os.Create("/home/sharma/Desktop/DeepLearning/Golang/filesystest/afield12_13_14.dat")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fnew.Close()
	rder1 := bufio.NewReader(f1)
	rder2 := bufio.NewReader(f2)
	rder3 := bufio.NewReader(f3)
	// buf := make([]byte, 4096)
	for i := 0; i < 10001; i++ {
		token1, _, _ := rder1.ReadLine()
		token2, _, _ := rder2.ReadLine()
		token3, _, _ := rder3.ReadLine()
		token := string(token1) + "\t" + string(token2) + "\t" + string(token3) + "\n"
		fnew.WriteString(token)
	}
}
