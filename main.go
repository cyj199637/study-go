package main

import (
	"fmt"
	"study-go/chapter01"
	"study-go/chapter02"
)

func main() {
	const greeting string = "Hello World!!!"
	fmt.Println(greeting)

	var greeting2 string = "Hello World!!!"
	fmt.Println(greeting2)

	greeting3 := "Hello World!!!"
	fmt.Println(greeting3)

	word, _ := chapter01.LongestStringAndLen("candy", "calendar", "cheer", "car", "custom")
	fmt.Println(word)

	chapter01.WithPointer()

	chapter01.WithArray()
	chapter01.WithSlice()

	chapter01.WithStruct()

	account := chapter02.NewAccount("Hui")
	account.Deposit(10000)
	err := account.Withdraw(20000)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account)
}
