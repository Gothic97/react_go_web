package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var num1 int8 = 1
	var num2 int16 = 1
	var num3 int32 = 1
	var num4 int64 = 1

	fmt.Println(unsafe.Sizeof(num1))
	fmt.Println(unsafe.Sizeof(num2))
	fmt.Println(unsafe.Sizeof(num3))
	fmt.Println(unsafe.Sizeof(num4))

	// Q1) 왜 이런 결과가 나오는 것일까?

	// Q2) Sizeof 함수는 구체적으로 어떤 기능을 구현할 때 사용하는 것일까?

}
