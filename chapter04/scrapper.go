package chapter04

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strconv"
	"study-go/common"
)

var baseUrl string = "https://kr.indeed.com/jobs?q=java&limit=50"

type Job struct {
	id       string
	title    string
	location string
	salary   string
	summary  string
}

/**
...(ellipsis): 해당 슬라이스의 컬렉션을 표현하는 것
- s... -> s 슬라이스의 모든 요소들의 집합을 의미
*/
func Scrapper() {
	var jobs []Job
	totalCount := GetTotalPagesCount()

	c := make(chan []Job)
	for i := 0; i < totalCount; i++ {
		go GetPage(i, c)
	}

	for i := 0; i < totalCount; i++ {
		extractedJobs := <-c
		jobs = append(jobs, extractedJobs...)
	}

	writeCsv(jobs)
	fmt.Println("Done, extracted ", len(jobs))
}

// GetTotalPagesCount : 전체 페이지 수를 가져오는 함수
func GetTotalPagesCount() int {
	totalCount := 0
	res, err := http.Get(baseUrl)
	common.CheckErr(err)
	common.CheckStatusCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	common.CheckErr(err)

	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		totalCount = s.Find("a").Length()
	})
	return totalCount
}

// GetPage : 지정한 페이지 내에 있는 JD를 다 긁어오는 함수
func GetPage(page int, outerC chan<- []Job) {
	var jobs []Job
	pageUrl := baseUrl + "&start=" + strconv.Itoa(page*50)
	res, err := http.Get(pageUrl)
	common.CheckErr(err)
	common.CheckStatusCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	common.CheckErr(err)

	c := make(chan Job)
	items := doc.Find(".tapItem")
	items.Each(func(i int, item *goquery.Selection) {
		go extractJob(item, c)
	})

	for i := 0; i < items.Length(); i++ {
		job := <-c
		jobs = append(jobs, job)
	}

	outerC <- jobs
}

func extractJob(item *goquery.Selection, c chan<- Job) {
	id, _ := item.Attr("data-jk")
	title := item.Find(".jobTitle>span").Text()
	location := item.Find(".companyInfo>.companyLocation").Text()
	salary := item.Find(".salary-snippet>span").Text()
	summary := item.Find(".job-snippet").Text()
	c <- Job{
		id:       id,
		title:    title,
		location: location,
		salary:   salary,
		summary:  summary,
	}
}
