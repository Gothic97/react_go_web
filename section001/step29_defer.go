package main

import "fmt"
import "os"

// func hello() {
//
// 	fmt.Println("Hello")
//
// }
//
// func world() {
//
// 	fmt.Println("world")
//
// }

// func HelloWorld() {
//
// 	defer func() {
// 		fmt.Println("world")
// 	}()
//
// 	func() {
// 		fmt.Println("Hello")
// 	}()
//
// }

func ReadHello() {

	file, err := os.Open("step29_hello.txt")
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	buf := make([]byte, 100)
	if _, err = file.Read(buf); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(buf))

}

func main() {

	// defer world() // if you use defer, function of defer is summoned just before finish of current function.
	// hello()
	// hello()
	// hello()

	// HelloWorld()

	// for i := 0; i < 5; i++ {
	//
	// 	defer fmt.Printf("%d ", i)
	//
	// }

	ReadHello()

}
