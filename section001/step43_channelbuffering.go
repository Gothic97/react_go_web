package main

import (
	"fmt"
	"runtime"
	_ "time"
)

func main() {

	runtime.GOMAXPROCS(1)

	done := make(chan bool, 1)
	count := 4

	go func() {

		for i := 0; i < count; i++ {

			done <- true
			fmt.Println("고루틴 :", i)
			//time.Sleep(100 * time.Millisecond)

		}

	}()

	for i := 0; i < count; i++ {

		<-done
		fmt.Println("메인 함수 :", i)
		//time.Sleep(100 * time.Millisecond)

	}

	fmt.Scanln()

}

// 실행 결과가 책과 다름
