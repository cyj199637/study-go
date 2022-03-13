package chapter02

import "errors"

var ErrNoMoney = errors.New("The amount is bigger than your balance.")
