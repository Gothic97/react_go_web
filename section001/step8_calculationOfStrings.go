package main

import "fmt"

func main() {

	var s1 string = "한글"
	var s2 string = "한글"
	var s3 string = "Go"
	// var s4 string = "한글 세종"

	fmt.Println(s1 == s2)
	fmt.Println(s1 + s3 + s2)
	fmt.Println("안녕하세요" + s1)
	// fmt.Println(s4 - s1) => ''유효하지 않은 작동' => 문자까리의 덧셈은 기본 연산식으로 가능하나, 뺄셈은 불가능함.

	var s5 string = "Hello"

	fmt.Printf("%c\n", s5[1])
	//fmt.Println(s5[1])
	fmt.Printf("%c", s5[1])
	//fmt.Printf(s5[1]) => 사용 오류. type에 문제가 있다고 함. Printf 문 내에서 사용해야 하는 type이 있는 듯함.
}
