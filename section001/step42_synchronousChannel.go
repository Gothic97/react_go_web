package main

import (
	"fmt"
	"time"
)

func main() {

	done := make(chan bool)
	count := 3

	go func() {

		for i := 0; i < count; i++ {

			done <- true
			fmt.Println("고루틴 :", i)
			time.Sleep(1 * time.Second)

		}

	}()

	for i := 0; i < count; i++ {

		<-done
		fmt.Println("메인 함수:", i)

	}

}

// Q) 왜 time을 넣지 않으면, 고루틴과 메인 함수가 정상적으로 동기화되지 않을까?
// Q) why goroutine and main function is not synchronized without time pakcage?
