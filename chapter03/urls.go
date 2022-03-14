package chapter03

import (
	"fmt"
	"net/http"
	"study-go/common"
)

func HitUrl(url string) error {
	fmt.Println("Checking ", url)
	res, err := http.Get(url)
	if err != nil || res.StatusCode >= 400 {
		return common.ErrRequestFailed
	}
	return nil
}
