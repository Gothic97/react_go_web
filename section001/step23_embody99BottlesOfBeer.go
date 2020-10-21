//when the number is more than 2, mark 'bottles'. but when the number is one, mark 'bottle'
//in the last, you shold buy new beer.
//ex) 99 bottles of beer on the wall, 99 bottles of beer.
//    Take one down, pass it around, 98 bottles of beer on the wall.
//    ...
//    1 bottle of beer on the wall, 1 bottle of beer.
//    Take one down, pass it around, No more bottles of beer on the wall.
//    No more bottles of beer on the wall, No more bottles of beer.
//    Go to the store and buy some more, 99 bottles of beer  on the wall.

package main

func main() {

	/*
	   	for i := 99; i >= 0; i-- {

	   		switch {
	   		case i >= 3:
	   			fmt.Println(i, "bottles of beer on the wall, ", i, "bottles of beer.")
	   			fmt.Println("Take one down, pass it around,", i-1, "bottles of beer on the wall.")

	   		case i == 2:
	   			fmt.Println(i, "bottles of beer on the wall, ", i, "bottles of beer.")
	   			fmt.Println("Take one down, pass it around,", i-1, "bottle of beer on the wall.")

	   		case i == 1:
	   			fmt.Println(i, "bottle of beer on the wall, ", i, "bottle of beer.")
	   			fmt.Println("Take one down, pass it around, No more bottles of beer on the wall.")

	   		case i == 0:
	   			fmt.Println("No more bottles of beer on the wall, No more bottles of beer.")
	   			fmt.Println("Go to the store and buy some more, 99 bottles of beer  on the wall.")

	   		}
	     }
	*/

	/*

		for bottles := 99; bottles >= 0; bottles-- {
			switch {
			case bottles > 1:
				fmt.Printf("%d bottles of beer on the wall, %d bottles of beer.\n",
					bottles, bottles)
				s := "bottles"
				if bottles-1 == 1 {
					s = "bottle"
				}
				fmt.Printf("Take one down, pass it around, %d %s of beer on the wall.\n",
					bottles-1, s)
			case bottles == 1:
				fmt.Printf("1 bottle of beer on the wall, 1 bottle of beer.\n")
				fmt.Printf("Take one down, pass it around, No more bottles of beer on the wall.\n")
			default:
				fmt.Printf("No more bottles of beer on the wall, no more bottles of beer.\n")
				fmt.Printf("Go to the store and buy some more, 99 bottles of beer on the wall.\n")
			}
		}

	*/

	// for i := 99; i >= 0; i-- {
	// 	if i >= 3 {
	// 		fmt.Printf("%d bottles of beer on the wall, %d bottles of beer.\n", i, i)
	// 		fmt.Printf("Take one down, pass it around, %d bottles of beer on the wall.\n", i-1)
	// 	} else if i == 2 {
	// 		fmt.Println("2 bottles of beer on the wall, 2 bottles of beer.")
	// 		fmt.Println("Take one down, pass it around, 1 bottle of beer on the wall.")
	// 	} else if i == 1 {
	// 		fmt.Println("1 bottle of beer on the wall, 1 bottle of beer.")
	// 		fmt.Println("Take one down, pass it around, No more bottles of beer on the wall.")
	// 	} else if i == 0 {
	// 		fmt.Println("No more bottles of beer on the wall, no more bottles of beer.")
	// 		fmt.Println("Go to the store and buy some more, 99 bottles of beer on the wall.")
	// 	}
	// }

}
