package main

import (
	"fmt"
	"net/http"
	"os"
)

func errHandle(err error, errMsg string) {
	if err != nil {
		fmt.Println(errMsg, err)
		os.Exit(1)
	}
}

func main() {
	// 1. url //https://arxiv.org/list/physics/new?skip=100&show=100
	url := "https://arxiv.org/list/physics/new"
	// 2. get
	resp, err := http.Get(url)
	errHandle(err, "http get err")
	// 3. filter
	buff := make([]byte, 4096)
	var body string
	for {
		n, _ := resp.Body.Read(buff)
		if n == 0 {
			break
		}
		body += string(buff[:n])
	}

	defer resp.Body.Close()
	fmt.Println(body)
	// 4. anaylse
}
