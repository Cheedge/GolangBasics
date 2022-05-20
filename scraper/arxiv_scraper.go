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
	newcontent := filter(content)
	// fmt.Println(content)
	store2File(newcontent, "arxiv"+strconv.Itoa(page)+".txt")
	// return content
}

func store2File(content string, filename string) {
	// path:="~/Desktop/DeepLearning/Golang/"
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
	// `<span class="descriptor">Authors:</span>\s\n(?s:(.*\n))`gm
	// `<a href="/search/physics\?searchtype=author&query=(.*)">(.*)</a>`gm

	RE_title := regexp.MustCompile(`<span class=\"descriptor\">Title:</span>\s(.*)`)
	// RE_authors1 := regexp.MustCompile(`<span class="descriptor">Authors:</span>\s\n(?s:(.*\n))`)
	// RE_authors1 := regexp.MustCompile(`<span class="descriptor">Authors:</span>\s\n(?s:(.*\n))</div>\n<div class="list-comments mathjax">`)
	// RE_authors2 := regexp.MustCompile(`<a href="/search/physics\?searchtype=author&query=(.*)">(.*)</a>`)
	RE_abs := regexp.MustCompile(`<a href="(/abs/.*)" title="Abstract">`)
	abs := RE_abs.FindAllStringSubmatch(content, -1)
	title := RE_title.FindAllStringSubmatch(content, -1)
	// fmt.Print(abs)
	// authors_content := RE_authors1.FindStringSubmatch(content)
	// fmt.Println(authors_content)
	// authors := RE_authors2.FindAllStringSubmatch(authors_content[0][1], -1)
	// fmt.Println(len(authors))
	// fmt.Println(len(authors_content))
	// fmt.Println(title)
	newContent := "Info:\n"
	for i, cont := range title {
		absContent := abstractExtract(i, abs)
		newContent += "Title:\n\t" + cont[1] + "\nAbstract:\n\t" + absContent + "\n"
		// fmt.Println(abs[i][1])

		// for _, auth := range authors {
		// 	newContent += "\t" + auth[2] + "\n"
		// 	// fmt.Println(i)
		// }
	}
	return newContent
}

var ARXIV = "https://arxiv.org"

func main() {
	// 1. url
	beg_page, end_page := pageInput()
	for i := beg_page; i <= end_page; i++ {
		url := ARXIV + "/list/physics/pastweek?skip=" + strconv.Itoa(i) + "+&show=100"
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
