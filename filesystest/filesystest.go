package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// open file
	f_r, err := os.Open("/home/sharma/Desktop/DeepLearning/Golang/closure/clousertest1.go")
	if err != nil {
		fmt.Printf("open error: %s\n", err)
		return
	}
	defer f_r.Close()
	// create write file
	f_w, err := os.Create("/home/sharma/Desktop/DeepLearning/Golang/filesystest/writefile.txt")
	if err != nil {
		fmt.Printf("create error: %s\n", err)
		return
	}
	defer f_w.Close()
	// create buffer
	buf := make([]byte, 4096)
	// read to buffer
	for {
		// pos, err := f_r.Read(buf)
		// if err != nil && err == io.EOF {
		// 	fmt.Printf("finished reading and pos=%d\n", pos)
		// 	return
		// }
		if pos, err := f_r.Read(buf); err != nil && err == io.EOF {
			fmt.Printf("finished reading and pos=%d\n", pos)
			return
		} else {
			f_w.Write(buf[:pos])
		}
		// write from buffer
		// f_w.Write(buf[:pos])
	}
}
