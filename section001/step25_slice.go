package main

import "fmt"

func main() {

	// slice is like arrangement. but length is not fixed. and volume is expended dynamically. Also slice is reference type unlike arrangement.

	// var a []int = make([]int, 5) // alloting space, through 'make' function, into slice that is lengths of five.
	// var b = make([]int, 5) // omit [] and type of data when we declare slice.
	// c := make([]int, 5) // omit [] and type of data, 'var' keyword, when we declare slice.

	// a := []int{32, 29, 78, 16, 81} // initialize the value during creating slice
	//
	// b := []int{
	//   32,
	//   29,
	//   78,
	//   16,
	//   81,  // when elements of slice are enumerated, comma is attached at last element.
	// }
	//
	// var s = make([]int, 5, 10) // length is five. voulme is 10. ( actually slice include arrangement, if this arrangement is expended, it can already allot space to use )
	//                            // if you omit volume, volume is same to length.
	//                            // if volume overflow, and automatically is expended.

	// a := make([]int, 5, 10)
	//
	// fmt.Println(len(a))
	// fmt.Println(cap(a)) // cap function find out volume.

	// a := make([]int, 5, 10)
	//
	// fmt.Println(a[4])
	// fmt.Println(a[5]) // error occured because of index that is out of length.

	// a := []int{1, 2, 3}
	//
	// a = append(a, 4, 5, 6)
	//
	// fmt.Println(a)

	// a := []int{1, 2, 3}
	// b := []int{4, 5, 6}
	//
	// a = append(a, b...) // when slice b is appended to slice a, use 'b...'.
	//
	// fmt.Pirntln(a)

	// a1 := [3]int{1, 2, 3}
	// var b1 [3]int
	//
	// b1 = a1
	// b1[0] = 9
	//
	// fmt.Println(a1)
	// fmt.Println(b1)
	// fmt.Println()
	//
	// a := []int{1, 2, 3}
	// var b []int // declare as a slice
	//
	// b = a
	// b[0] = 9 // because slice is type of reference, a[0] and b[0] are changed all.
	//
	// fmt.Println(a)
	// fmt.Println(b)

	// a := []int{1, 2, 3, 4, 5}
	// b := make([]int, 3)
	//
	// copy(b, a)
	//
	// fmt.Println(a)
	// fmt.Println(b) // [1 2 3]: because length of slice b is three, three elements is copied.

	// because slice is copied, original slice is not changed.
	// a := []int{1, 2, 3}
	// b := make([]int, 3)
	//
	// copy(b, a)
	// b[0] = 9
	//
	// copy(b, a)
	// b[0] = 9
	//
	// fmt.Println(a)
	// fmt.Println(b)

	// Go language divide length and volume to embody dynamical arrangement. thus if element of slice is expended, Go runtime expend volume of slice by selected algorithm.
	// a := []int{1, 2, 3, 4, 5}
	//
	// fmt.Println(len(a), cap(a)) // 5 5: slice whose length is five and volume is five
	//
	// a = append(a, 6, 7)
	//
	// fmt.Println(len(a), cap(a)) // Go runtime icreases volume according to algorithm set.

	// a := []int{1, 2, 3, 4, 5}
	//
	// b := a[0:5] // refer to from 0 to 5 of index of 'a'. last index is too many by one than original slice.
	//
	// fmt.Println(a)
	// fmt.Println(b)

	// a := []int{1, 2, 3, 4, 5}
	//
	// fmt.Println(a[0:3])
	// fmt.Println(a[1:3])
	// fmt.Println(a[2:5])

	// a := []int{1, 2, 3, 4, 5}
	//
	// fmt.Println(a[:])
	// fmt.Println(a[0:])
	// fmt.Println(a[:5])
	// fmt.Println(a[0:len(a)])
	//
	// fmt.Println(a[3:])
	// fmt.Println(a[:3])
	// fmt.Println(a[1:3])

	// a := []int{1, 2, 3, 4, 5}
	//
	// b := a[:2]
	// b[0] = 9
	//
	// fmt.Println(a)
	// fmt.Println(b)

	a := []int{1, 2, 3, 4, 5, 6, 7, 8}

	b := a[0:6:8] // slice_name:[start_index:last_index:capacity] // but when we set the capacity of dulication slice, duplication slice cannot be over than original slice.

	fmt.Println(len(b), cap(b))
	fmt.Println(b)

}
