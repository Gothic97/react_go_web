package main

import "fmt"

func main() {

	// var a [5]int
	//
	// a[2] = 7
	//
	// fmt.Println(a)

	// var a [5]int = [5]int{32, 29, 78, 16, 81}
	// var b = [5]int{32, 29, 78, 16, 81}
	// c := [5]int{32, 29, 78, 16, 81}

	// a := [5]int{32, 29, 78, 16, 81}
	//
	// b := [5]int{
	// 	32,
	// 	29,
	// 	78,
	// 	16,
	// 	81, // when these constituents are listed, affix comma in last element.
	// }
	//
	// c := [...]int{32, 29, 78, 16, 81} // there are five elements to initialize, due to using '...', volume of arrangement is five.
	//
	// d := [...]string{"Maria", "Andrew", "John"}

	// a := [3][3]int{
	// 	{1, 2, 3},
	// 	{4, 5, 6},
	// 	{7, 8, 9},
	// }
	//
	// fmt.Println(a)

	// a := [5]int{32, 29, 78, 16, 81}
	//
	// for i := 0; i < len(a); i++ { // function "len" find out length of arrangement
	// 	fmt.Println(a[i])
	// }

	// a := [5]int{32, 29, 78, 16, 81}
	//
	// for i, value := range a { // value of arrangement element get in to "value". and index get in to "i".
	// 	fmt.Println(i, value)
	// } // if you use range keyword, value of index and value of "value" get in to variables automatically.

	// a := [5]int{32, 29, 78, 16, 81}
	//
	// for value := range a { // Instead of value, index get in to "value".
	//
	// 	fmt.Println(value)
	//
	// }

	// a := [5]int{32, 29, 78, 16, 81}
	//
	// for _, value := range a { // omit index. value of arrangement elements get in to the value.
	//
	// 	fmt.Println(value)
	//
	// }

	a := [5]int{1, 2, 3, 4, 5}

	b := a // if arrangement substitute to variable, the whole of arrangement is duplicated.

	fmt.Println(a)
	fmt.Println(b)

}
