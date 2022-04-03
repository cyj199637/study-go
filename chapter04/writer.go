package chapter04

import (
	csv2 "encoding/csv"
	"os"
	"study-go/common"
)

var detailPageUrl = "https://kr.indeed.com/viewjob?jk="

// writeCsv : 긁어온 JD 정보를 CSV 파일로 생성하는 함수
func writeCsv(jobs []Job) {
	file, err := os.Create("jobs.csv")
	common.CheckErr(err)

	w := csv2.NewWriter(file)
	defer w.Flush()

	headers := []string{"link", "title", "location", "salary", "summary"}
	wErr := w.Write(headers)
	common.CheckErr(wErr)

	for _, job := range jobs {
		contents := []string{detailPageUrl + job.id, job.title, job.location, job.salary, job.summary}
		err := w.Write(contents)
		common.CheckErr(err)
	}
}
