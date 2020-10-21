package main

import "fmt"

func main() {

	/*
	   	const age int = 10
	   	const name string = "Maria"
	   	// const score int // 컴파일 에러

	   	age = 20 // 컴파일 에러
	     name = "Grace" // 컴파일 에러
	*/

	/*

	   const age = 10 // int
	   const name = "Maria" // string
	   // const 10minites = 10 // 상수명을 숫자로 시작할 수 없음.
	   // const address //컴파일 에러

	*/

	// const x, y int = 30, 50     // x = 30, y = 50
	// const age, name = 10, "Maria" // age = 10, name = "Maria"

	const (
		x, y      int = 30, 50
		age, name     = 10, "Maria"
	)

	fmt.Println("x :", x)
	fmt.Println("y :", y)
	fmt.Println("age :", age)
	fmt.Println("name :", name)

}
