package chapter04

import (
	csv "github.com/tsak/concurrent-csv-writer"
	"study-go/common"
)

var detailPageUrl = "https://kr.indeed.com/viewjob?jk="

// writeCsv : 긁어온 JD 정보를 CSV 파일로 생성하는 함수
func writeCsv(jobs []Job) {
	w, err := csv.NewCsvWriter("jobs.csv")
	common.CheckErr(err)

	defer w.Flush()

	headers := []string{"link", "title", "location", "salary", "summary"}
	wErr := w.Write(headers)
	common.CheckErr(wErr)

	for _, job := range jobs {
		go writeRow(w, job)
	}
}

func writeRow(w *csv.CsvWriter, job Job) {
	contents := []string{detailPageUrl + job.id, job.title, job.location, job.salary, job.summary}
	err := w.Write(contents)
	common.CheckErr(err)
}
