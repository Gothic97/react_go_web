package main

import "fmt"

func sum(a int, b int, c chan int) {

	c <- a + b

}

func main() {

	//var c chan int
	//c = make(chan int)

	c := make(chan int)

	go sum(1, 2, c)

	n := <-c

	fmt.Println(n)

}
