package main

import "fmt"

//용도 : If, for문 테스트 및 변수의 지역성 파악

func main() {
	fmt.Println("Chapter.1")

	i := 10

	if i >= 5 {
		fmt.Println("5 이상")
	}

	fmt.Println(i)

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println(i)
}
