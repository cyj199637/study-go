package main

import (
	"fmt"
	"study-go/chapter03"
)

func main() {
	var results = make(map[string]string)
	urls := []string{
		"https://www.naver.com/",
		"https://www.google.com/",
		"https://www.youtube.com/",
		"https://www.instagram.com/",
		"https://www.amazon.com/",
		"https://www.ebay.com/",
		"https://www.coupang.com/",
		"https://www.netflix.com/kr/",
		"https://www.disneyplus.com/home",
		"https://nintendo.com/",
	}

	for _, url := range urls {
		err := chapter03.HitUrl(url)
		if err != nil {
			results[url] = "FAILED"
		} else {
			results[url] = "OK"
		}
	}

	for url, result := range results {
		fmt.Println(url, " -> ", result)
	}
}
