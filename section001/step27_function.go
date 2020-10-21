package main

import "fmt"

// func hello() { // declare about hello function
//
// 	fmt.Println("Hello, world!")
//
// }

// func sum(a int, b int) int { // int 형 매개변수 a, b
// 	// 그리고 int 형 리턴값을 가지는 함수 정의
// 	return a + b
//
// }

// func sum(a int, b int) (r int) {
//
// 	r = a + b // substitute value at variable 'r'
// 	// in Go, we can set up return_value in advance.
// 	return
//
// }

// func SumAndDiff(a int, b int) (int, int) { // definition about function of two return value of int type
//
// 	return a + b, a - b
//
// }

// func SumAndDiff(a int, b int) (int, int) {
// 	return a + b, a - b
// }

// func hello() (int, int, int, int, int) {
// 	return 1, 2, 3, 4, 5
// }

// func SumAndDiff(a int, b int) (sum int, diff int) {
//
// 	sum = a + b
// 	diff = a - b
// 	return
//
// }

// func sum(n ...int) int {
//
// 	total := 0
// 	for _, value := range n {
// 		total += value
// 	}
//
// 	return total
//
// }

// func factorial(n uint64) uint64 {
//
// 	if n == 0 {
// 		return 1
// 	}
//
// 	return n * factorial(n-1)
//
// }

// func sum(a int, b int) int {
// 	return a + b
// }
//
// func diff(a int, b int) int {
//
// 	return a - b
//
// }

func main() {

	// hello()

	// func function_name(parameter type) returnValue_type {}

	// r := sum(1, 2)
	// fmt.Println(r)

	// r := sum(1, 2)
	// fmt.Println(r) // 3

	// sum, diff := SumAndDiff(6, 2) // substitute two return value in to two variables
	// fmt.Println(sum, diff) // In this way, return_value is returned according to an order that is designed in function. thus if you design value to be returned, the first retrun_value is saved.
	// if you want to use the second return_value as omitting the first return_value, you can use '_'.

	// _, diff := SumAndDiff(6, 2)
	// fmt.Println(diff)

	// a, _, c, _, e := hello()
	// fmt.Println(a, c, e)

	// sum, diff := SumAndDiff(6, 2)
	// fmt.Println(sum, diff)

	// r := sum(1, 2, 3, 4, 5)
	// fmt.Println(r)

	// n := []int{1, 2, 3, 4, 5}
	// r := sum(n...) // in this way, variable factor function must receive only a value of int type. and cannot receive slice.
	// // thus we must affix '...' in the last. if you affix '...', GO alrorithm hand over factors in the slice.
	//
	// fmt.Println(r)

	// fmt.Println(factorial(5))

	// var hello func(a int, b int) int = sum
	// world := sum // Go language can automatically infer type of variable, when we declare and substitute variables.
	//
	// fmt.Println(hello(1, 2))
	// fmt.Println(world(1, 2))

	// slice = []func(parameter_name type) return_value_type{function_name1, function_name2}

	// f := []func(int, int) int{sum, diff}
	//
	// fmt.Println(f[0](1, 2)) // function is summoned as a first factor of arragement.
	// fmt.Println(f[1](1, 2)) // function is summoned as a second factor of arragement.

	// f := map[string]func(int, int) int{
	//
	// 	"sum":  sum,
	// 	"diff": diff,
	// }
	//
	// fmt.Println(f["sum"](1, 2))
	// fmt.Println(f["diff"](1, 2)) // Q) if I want to summon the whole slice, How can I do this without function of rounds?

	func() {
		fmt.Println("Hello, world!")
	}()

	func(s string) {
		fmt.Println(s)
	}("Hello, world!")

	r := func(a int, b int) int {
		return a + b
	}(1, 2)

	fmt.Println(r)

	// through the anonymous function, we can reduce volume of code.
}
