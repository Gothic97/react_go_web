// print from 0 to 100.
// three times print Fizz
// five times print Buzz
// common multiple between 3 and 5 print FizzBuzz

package main

import "fmt"

func main() {

	for i := 0; i <= 100; i++ {

		switch {

		case i == 0:
			fmt.Println(i)
		case i%3 == 0 && i%5 == 0:
			fmt.Println("FizzBuzz")
		case i%5 == 0:
			fmt.Println("Buzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		default:
			fmt.Println(i)

		}

	}

}
