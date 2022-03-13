package main

import (
	"fmt"
	"study-go/theory"
)

func main() {
	const greeting string = "Hello World!!!"
	fmt.Println(greeting)

	var greeting2 string = "Hello World!!!"
	fmt.Println(greeting2)

	greeting3 := "Hello World!!!"
	fmt.Println(greeting3)

	word, _ := theory.LongestStringAndLen("candy", "calendar", "cheer", "car", "custom")
	fmt.Println(word)

	theory.WithPointer()

	theory.WithArray()
	theory.WithSlice()

	theory.WithStruct()
}
