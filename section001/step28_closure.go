// definition of "closure" is function that can declare and define function in other function.
// and can access variable that is declared outside of function.

// 클로저 == 이미 생명 주기가 끝난 외부 함수의 변수를 참조하는 함수

package main

import "fmt"

func calc() func(x int) int { // => 1, 2

	a, b := 3, 5
	return func(x int) int {
		return a*x + b // => 3
	}

}

func main() {

	f := calc()

	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))
	fmt.Println(f(4))
	fmt.Println(f(5))

}
