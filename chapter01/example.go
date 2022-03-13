package chapter01

import (
	"fmt"
	"strconv"
)

func LongestStringAndLen(words ...string) (longest string, length int) {
	defer fmt.Println("Done")
	for _, word := range words {
		if len(longest) < len(word) {
			longest = word
		}
	}
	length = len(longest)
	return
}

func WithPointer() {
	a := "It's a"
	b := a
	a = "a changed"
	fmt.Println(a, b)

	// stored address of a
	c := &a
	fmt.Println(&a, c)

	// dereference
	fmt.Println(*c)

	*c = "It changed value"
	fmt.Println(a)
}

func WithArray() {
	array := [5]string{"Alice", "Bob", "Charles"}
	fmt.Println(array)

	array[3] = "Den"
	fmt.Println(array, len(array), cap(array))
}

func WithSlice() {
	slice := []int{1, 2, 3}
	slice[0] = 4
	//slice[3] = 5 // index out of range error
	fmt.Println(slice, len(slice), cap(slice))

	withMake := make([]int, 3, 5) // initialize with 0
	withMake[0] = 1
	/**
	처음 정의한 용량을 넘어서 버리면
	기존의 두 배의 용량을 가진 새로운 배열을 만들고 기존 배열을 복사
	*/
	withMake = append(withMake, 1, 2, 3, 4)
	fmt.Println(withMake, len(withMake), cap(withMake))
}

type Person struct {
	name    string
	age     int
	address string
}

func (p Person) String() string {
	return "name: " + p.name + ", age: " + strconv.Itoa(p.age) + ", address: " + p.address
}

func WithStruct() {
	person := Person{name: "Alice", age: 23, address: "New York"}
	fmt.Println(person)

	person2 := Person{name: "Alice", age: 23, address: "New York"}
	fmt.Println(person == person2)
}
