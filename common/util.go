package common

import (
	"log"
	"net/http"
)

func CheckErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func CheckStatusCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request Failed With Status: ", res.Status)
	}
}
