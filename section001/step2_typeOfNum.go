/*

uint8 : 부호 없는(unsigned) 8비트, 1바이트 정수
init8 : 부호 있는(signed) 8비트, 1바이트 정수
uint : 32비트 시스템에서는 unt32, 64비트 시스템에서는 uint64
uintptr : unit와 크기가 동일. 포인터를 저장할 때 사용
float32 : IEEE-754 32비트 단정밀도 부동소수점. 7자리 정밀도 보장
float64 : IEEE-754 64비트 배정밀도 부동소수점. 15자리 정밀도 보장
complex64 : float32 크기의 실수부와 허수부로 된 복소수
byte : unit8과 크기가 동일. 바이트 단위로 저장할 때 사용
rune : unit32와 크기가 동일. 유니코드 문자 코드(Code point)를 저장할 때 사용

*/

package main

import "fmt"

func main() {

	// 8진수 및 16진수의 저장방식 확인

	var num1 int = 0723
	// var num2 := 0723 => 오류
	var num3 int = 0x32fa2c75
	// var num4 := 0x32fa2c75 => 오류

	fmt.Println("ex1 : ", num1)
	// fmt.Println("ex1 : ", num2)
	fmt.Println("ex1 : ", num3)
	// fmt.Println("ex1 : ", num4)
	fmt.Println("")

	// 진수별 데이터 표현 방식 이해

	// var num5 int = 0000 1010 => 오류
	var num6 int = 012
	var num7 int = 0xa

	// fmt.Println("ex2 : ", num5)
	fmt.Println("ex2 : ", num6)
	fmt.Println("ex2 : ", num7)

	// 진수별 데이터 표현 방식에 대한 추가 공부가 필요함. 허나 당장 중요한 것임은 아니므로 보류

}
