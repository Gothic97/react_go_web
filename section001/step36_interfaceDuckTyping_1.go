package main

import "fmt"

type Duck struct {
	name string
}

func (d Duck) quack() {

	fmt.Println("꽥~!")

}

func (d Duck) feathers() {

	fmt.Println("오리는 흰색과 회색 털을 가지고 있습니다.")

}

type Person struct {
}

func (p Person) quack() {

	fmt.Println("사람은 오리 흉내를 냅니다. 꽥~!")

}

func (p Person) feathers() {

	fmt.Println("사람은 땅에서 깃털을 주워서 보여줍니다.")

}

type Quacker interface {
	quack()
	feathers()
}

func inTheForest(q Quacker) {

	q.quack()
	q.feathers()

}

func main() {

	var donald Duck = Duck{"quack!"}
	var john Person

	inTheForest(donald)
	inTheForest(john)
	fmt.Println()

	if v, ok := interface{}(donald).(Quacker); ok {
		fmt.Println(v, ok)
	}

}

// 결론적으로 덕타이핑이란 무엇인가?

// 덕타이핑이란 해당 자료형이 어떤 자료형이든 상관 없이 구현된 메서드와 매개변수의 자료형, 리턴값의 자료형이 일치했을 때, 해당 타입을 같은 인터페이스 타입으로 보는 방식이다.
