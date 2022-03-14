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

	dict := chapter02.SampleDict{"first": "Buying KeyCap"}
	err2 := dict.Add("first", "Wake up Early")
	if err2 != nil {
		fmt.Println("Add error -> ", err2)
	}
	dict.Add("second", "Wake up Early")
	todo, _ := dict.Search("second")
	fmt.Println(todo)

	err3 := dict.Update("third", "Drink BubbleTea")
	if err3 != nil {
		fmt.Println("Update error -> ", err3)
	}
	dict.Update("second", "Drink BubbleTea")
	todo2, _ := dict.Search("second")
	fmt.Println(todo2)

	err4 := dict.Delete("third")
	if err4 != nil {
		fmt.Println("Delete error -> ", err4)
	}
	dict.Delete("second")
	_, err5 := dict.Search("second")
	fmt.Println("Delete Success -> ", err5)
}
