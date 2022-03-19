package chapter03

import (
	"fmt"
	"net/http"
)

type ConnectionResult struct {
	Url    string
	Status string
}

func UrlChecker() {
	results := make(map[string]string)
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

	ch := make(chan ConnectionResult)
	for _, url := range urls {
		go HitUrl(url, ch)
	}

	for range urls {
		result := <-ch
		results[result.Url] = result.Status
	}

	for url, status := range results {
		fmt.Println(url, " -> ", status)
	}
}

/**
chan<- : This channel is send-only.
*/
func HitUrl(url string, ch chan<- ConnectionResult) {
	fmt.Println("Checking ", url)
	res, _ := http.Get(url)
	ch <- ConnectionResult{Url: url, Status: res.Status}
}
