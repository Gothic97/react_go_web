package main

import "fmt"

type Dog struct {
	name	string
	age		int
}

type Cat struct {
	name	string
	age		int
	color	string
}

// Dog 리시버 메소드

func (a Dog) bites() {
	fmt.Println(a.name, "이 짖었다!")
}

func (a Dog) run() {
	fmt.Println(a.name, "이 달린다!")
}

func (a Dog) lol() {
	fmt.Printf("%s이 웃는다!\n", a.name)
}

// Cat 리시버 메소드

func (a Cat) bites() {
	fmt.Println(a.name, "이 짖었다!")
}

func (a Cat) run() {
	fmt.Println(a.name, "이 달린다!")
}

func (a Cat) lol() {
	fmt.Println(a.name, "이 웃는다!")
}

func (a Cat) cries() {
	fmt.Println(a.name, "이 운다!")
}

type Behavior interface {
	bites()
	run()
	lol()
}

func CreaterD(name string, age int) Dog {
	return Dog{name, age}
}

func CreaterC(name string, age int, color string) Cat {
	return Cat{name, age, color}
}

func DuckType (i Behavior) {
	i.bites()
	i.run()
	i.lol()
}

func main() {

	Dog1 := CreaterD("Daniel", 3)
	Cat1 := CreaterC("Holy", 3, "빨강")

	DuckType(Dog1)
	DuckType(Cat1)

}













