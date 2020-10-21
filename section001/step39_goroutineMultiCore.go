package main

import (
	"fmt"
	"runtime"
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU()) // runtime.NumCPU is a function find out the number of maximum CPU. runtime.GOMAXPROCS sets up the number of CPU to use.

	fmt.Println(runtime.GOMAXPROCS(0)) // runtime.GOMAXPROCS 함수에 0을 넣으면 설정 값은 바꾸지 않고, 현재 설정 값만 리턴함.

	s := "Hello, world!"

	for i := 0; i < 100; i++ {

		go func(n int) {
			fmt.Println(s, n)
		}(i)

	}

	fmt.Scanln()
}
