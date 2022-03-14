package chapter03

import (
	"fmt"
	"time"
)

/**
goroutine은 main 함수가 실행하는 동안에만 동작함
*/
func ExampleGoroutines() {
	go countWithWord("Black")
	countWithWord("Pink")
}

/**
main 함수에서 실행할 코드가 없다면 아래 goroutine들은 동작하지 않음
*/
func FailedExampleGoroutines() {
	go countWithWord("Black")
	go countWithWord("Pink")
}

func Example2Goroutines() {
	go countWithWord("Black")
	go countWithWord("Pink")
	time.Sleep(time.Millisecond)
}

func countWithWord(word string) {
	for i := 0; i < 20; i++ {
		fmt.Println(word, " -> ", i)
	}
}
