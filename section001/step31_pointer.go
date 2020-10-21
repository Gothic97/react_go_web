package main

import "fmt"

// func hello(n int) {
//
// 	n = 2
//
// }

func hello(n *int) {
	*n = 2
}

func main() {

	// var numPtr *int
	//
	// fmt.Println(numPtr)

	// var numPtr *int = new(int)
	//
	// fmt.Println(numPtr)

	// var numPtr *int = new(int)
	//
	// *numPtr = 1
	//
	// fmt.Println(*numPtr)

	// var num int = 1
	// var numPtr *int = &num
	//
	// fmt.Println(numPtr)
	// fmt.Println(&num)

	// var numPtr *int = new(int)
	//
	// numPtr++              // In Go, calculation of pointer is not admitted.
	// numPtr = 0xc0820062d0 // In Go, address of memory cannot be substituted into something.
	//
	// fmt.Println(numPtr)

	// var n int = 1
	//
	// hello(n)
	//
	// fmt.Println(n)

	var n int = 1

	hello(&n)

	fmt.Println(n)

}
