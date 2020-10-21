package main

import "fmt"

func main() {

	//var b1 bool = true
	//var b2 bool

	// fmt.Println(b1)
	// fmt.Println(b2) // bool의 기본값은 false임.

	//<논리 연산자> And(&&), OR(||), NOT(!)

	// fmt.Println(true && true)
	// fmt.Println(true && false)
	// fmt.Println(false && false)
	// fmt.Println(true || true)
	// fmt.Println(true || false)
	// fmt.Println(false || false)
	// fmt.Println(!true)
	// fmt.Println(!false)

	//<비교 연산자>

	var num1 int = 3
	var num2 int = 10

	fmt.Println(num1 > num2)
	fmt.Println(num1 < num2)
	fmt.Println(num1 != num2)
	fmt.Println(num1 >= num2)
	fmt.Println(num1 <= num2)
}
