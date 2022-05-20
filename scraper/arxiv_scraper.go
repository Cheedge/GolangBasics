package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func errHandle(err error, errMsg string) {
	if err != nil {
		fmt.Println(errMsg, err)
		os.Exit(1)
	}
}

func pageInput() (int, int) {
	var beg_page, end_page int
	fmt.Println("input \"begin page end page\": ")
	if _, err := fmt.Scan(&beg_page, &end_page); err != nil {
		fmt.Println("scan page err: ", err)
		os.Exit(1)
	}

	return beg_page, end_page
}

func scraperHandler(url string, page int) {
	// 2. get
	resp, err := http.Get(url)
	errHandle(err, "http get err")
	// 3. filter
	buff := make([]byte, 4096)
	var content string
	for {
		n, _ := resp.Body.Read(buff)
		if n == 0 {
			break
		}
		content += string(buff[:n])
	}

	defer resp.Body.Close()
	// fmt.Println(content)
	store2File(content, "arxiv"+strconv.Itoa(page)+".txt")
	// return content
}

func store2File(content string, filename string) {
	// path:="~/Desktop/DeepLearning/Golang/"
	os.WriteFile(filename, []byte(content), 0666)
}

func filter() {

}

func main() {
	// 1. url
	beg_page, end_page := pageInput()
	for i := beg_page; i <= end_page; i++ {
		url := "https://arxiv.org/list/physics/pastweek?skip=" + strconv.Itoa(i) + "+&show=100"
		scraperHandler(url, i)

	}

	/* 2. get
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
	*/
}
