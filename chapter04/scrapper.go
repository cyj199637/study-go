package chapter04

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strconv"
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

	for i := 0; i < totalCount; i++ {
		extractedJobs := GetPage(i)
		jobs = append(jobs, extractedJobs...)
	}
}

func GetTotalPagesCount() int {
	totalCount := 0
	res, err := http.Get(baseUrl)
	checkErr(err)
	checkStatusCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		totalCount = s.Find("a").Length()
	})
	return totalCount
}

func GetPage(page int) []Job {
	var jobs []Job
	pageUrl := baseUrl + "&start=" + strconv.Itoa(page*50)
	res, err := http.Get(pageUrl)
	checkErr(err)
	checkStatusCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".tapItem").Each(func(i int, item *goquery.Selection) {
		job := extractJob(item)
		jobs = append(jobs, job)
	})

	return jobs
}

func extractJob(item *goquery.Selection) Job {
	id, _ := item.Attr("data-jk")
	title := item.Find(".jobTitle>span").Text()
	location := item.Find(".companyInfo>.companyLocation").Text()
	salary := item.Find(".salary-snippet>span").Text()
	summary := item.Find(".job-snippet").Text()
	return Job{
		id:       id,
		title:    title,
		location: location,
		salary:   salary,
		summary:  summary,
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkStatusCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request Failed With Status: ", res.Status)
	}
}
