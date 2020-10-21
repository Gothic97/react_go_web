package main

import "fmt"

// 변수 선언 형식 익히기

func main() {
	var i int    //int형의 기본값은 0임.
	var s string //stirng형의 기본값은 ""이나 " "일 것임.

	//i = "I'm int" => int형에 string형을 넣기 불가
	//s = 1 => stirng형에 숫자형 넣기 불가

	i = 1000000
	s = "I'm stirng"

	var age int = 10
	var name string = "Maria"

	fmt.Println(i)
	fmt.Println(s)
	//fmt.Println(s == "")
	fmt.Println(age)
	fmt.Println(name)

	//결론적으로 'var 변수명 자료형'으로 변수를 사용할 수 있음.
	// <짧은 선언문과 비교>
	// 코드에 자료형이 명시돼 있다는 장점이 있음. 곧 추론하면서 자료형을 파악하지 않아도 되는 것.
	// 코드가 길어지는 단점이 있음.
	// s의 기본값은 ""임.
}
