package main

import "fmt"

func main() {

	var s1 string = "Hello, world!\n"
	var s2 string = "안녕하세요\n"
	var s3 string = "\ud55c\uae00"               // 한글 : 유니코드 문자 코드로 저장
	var s4 string = "\U0000d55c\U0000ae00"       // 한글 : 유니코드 문자 코드로 저장
	var s5 string = "\xed\x95\x9c\xea\xb8\x80\n" // 한글 : UTF-8 인코딩의 바이트 값으로 저장

	var s7 string = `안녕하세요
Hello, world!`

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println(s5)
	fmt.Println(s7)

	fmt.Println(len(s3)) // len 함수는 글자의 수를 제한할 때 사용할 수 있을 것으로 보임. 형식상 엄격함이 있어야 할 때, 다양하게 사용할 수 있을 것.

}
