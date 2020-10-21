package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (_ *Person) greeting() {

	fmt.Println("Hello~")

}

type Student struct {
	Person
	school string
	grade  int
}

func (_ *Student) greeting() {

	fmt.Println("Hello~ Student!")

}

func main() {

	var s Student

	s.Person.greeting()
	s.greeting()

}
