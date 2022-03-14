package common

import "errors"

var (
	ErrNoMoney       = errors.New("The amount is bigger than your balance.")
	ErrNotFound      = errors.New("Not Found")
	ErrAlreadyExists = errors.New("Already Exists")
	ErrRequestFailed = errors.New("Request Failed")
)
