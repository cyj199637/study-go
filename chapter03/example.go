package chapter03

import (
	"fmt"
	"time"
)

/**
goroutine: Go 런타임이 관리하는 Lightweight 논리적 쓰레드
- 여러 코드들을 asynchronously + concurrently 하게 실행함
- main 함수가 실행하는 동안에만 동작함

즉, 프로그램에서 1, 2, 3, 4, 5라는 함수를 실행한다고 했을 때
	- goroutine 사용 전: 프로그램 실행 시간 = (1 + 2 + 3 + 4 + 5) 의 실행 시간 총합
	- goroutine 사용 후: 프로그램 실행 시간 = (1, 2, 3, 4, 5) 의 실행 시간 중 가장 긴 실행 시간
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

/**
Channel: goroutine과 main 함수 혹은 서로 다른 goroutine 간에 정보를 주고받는 통로
- Channel에서 메시지를 수신하는 것은 blocking operation 이다.

- 1번 예시: 함수를 goroutine 으로 실행은 하지만, 다음 라인에서 채널로부터 메시지를 받기 위해 대기하고 있기 때문에
		  다음 goroutine 함수를 병렬적으로 실행하지 못하게 됨
- 2번 예시: goroutine 으로 함수를 실행하는 for 문이 끝난 후에 채널로부터 메시지를 받기 위해 대기하고 있기 때문에,
		  모든 goroutine 함수를 병렬적으로 실행할 수 있게 됨
*/
func ExampleChannels1() {
	names := []string{"JjangGu", "CheolSu", "YuRi", "MaengGu", "Hunbalnom"}
	ch := make(chan string)
	for _, name := range names {
		go returnWithMessages(name, ch)
		fmt.Println(<-ch)
	}
}

func ExampleChannels2() {
	names := []string{"JjangGu", "CheolSu", "YuRi", "MaengGu", "Hunbalnom"}
	ch := make(chan string)
	for _, name := range names {
		go returnWithMessages(name, ch)
	}

	for range names {
		fmt.Println(<-ch)
	}
}

func returnWithMessages(name string, ch chan string) {
	fmt.Println("Waiting for message...")
	time.Sleep(time.Second * 5)
	ch <- name + " is the agent of Tteok-Ip Security Guard."
}
