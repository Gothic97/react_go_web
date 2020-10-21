package main

import "fmt"

func main() {

	/*
	   const (

	     Sunday       = 0
	     Monday       = 1
	     Tuesday      = 2
	     Wednesday    = 3
	     Thursday     = 4
	     Friday       = 5
	     Saturday     = 6
	     numberOfDays = 7

	   )
	*/

	const (
		Sunday = iota //0
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		numberOfDays
	)

	fmt.Println(Thursday)
	fmt.Println(numberOfDays)

	const (
		bit0, mask0 = 1 << iota, 1<<iota - 1
		bit1, mask1
		_, _
		bit3, mask3
	)

	fmt.Println(bit0, mask0)
	fmt.Println(bit1, mask1)
	fmt.Println(bit3, mask3)

}
