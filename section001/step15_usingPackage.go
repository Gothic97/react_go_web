package main

import . "fmt" // fmt를 전역 패키지로 가져옴
import . "runtime"

// import f "fmt" // fmt를 f로 가져옴
// import _ // 패키지를 가져온 뒤 사용하지 않을 때

func main() {

	Println("Hello word!")
	Println("CPU Count : ", NumCPU()) // 전역 패키지는 복수사용이 가능함

}
