package main

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
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

func scraperHandler(url string, page int, concurr_token chan<- int) {
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
	newcontent := filter(content)

	store2File(newcontent, "Arxiv_concurr"+strconv.Itoa(page)+".txt")
	concurr_token <- page
}

func store2File(content string, filename string) {
	os.WriteFile(filename, []byte(content), 0666)
}

func absFilter(abs string) string {
	RE := regexp.MustCompile(`<span class="descriptor">Abstract:</span>(?s:(.*))</blockquote>`)
	absContent := RE.FindAllStringSubmatch(abs, 1)
	return absContent[0][1]
}

func abstractExtract(i int, abs [][]string) string {
	url := ARXIV + abs[i][1]
	resp, err := http.Get(url)
	errHandle(err, "abstract extraciton error")
	buff := make([]byte, 4096)
	var absContent string
	for {
		n, _ := resp.Body.Read(buff)
		if n == 0 {
			break
		}
		absContent += string(buff[:n])
	}
	absAfterFilte := absFilter(absContent)
	return absAfterFilte
}

func filter(content string) string {
	RE_title := regexp.MustCompile(`<span class=\"descriptor\">Title:</span>\s(.*)`)

	RE_abs := regexp.MustCompile(`<a href="(/abs/.*)" title="Abstract">`)
	abs := RE_abs.FindAllStringSubmatch(content, -1)
	title := RE_title.FindAllStringSubmatch(content, -1)

	newContent := "Info:\n"
	for i, cont := range title {
		absContent := abstractExtract(i, abs)
		newContent += "Title:\n\t" + cont[1] + "\nAbstract:\n\t" + absContent + "\n"

	}
	return newContent
}

var ARXIV = "https://arxiv.org"

func main() {
	// 0. concurrency
	concurr_token := make(chan int)
	// 1. url
	beg_page, end_page := pageInput()
	for i := beg_page; i <= end_page; i++ {
		url := ARXIV + "/list/physics/pastweek?skip=" + strconv.Itoa(i) + "+&show=100"
		go scraperHandler(url, i, concurr_token)
	}
	for i := beg_page; i <= end_page; i++ {
		fmt.Printf("%d page has been finished.\n", <-concurr_token)
	}

}
