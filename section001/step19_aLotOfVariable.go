package main

import . "fmt"

func main() {

	for i, j := 0, 0; i < 10; i, j = i+1, j+2 {
		Println(i, j)
	}

	// for i, j := 0, 0; i < 10; i ++, j += {
	//   Println(i, j)
	// } // 변화식에 병렬 할당을 사용하지 않을 경우, 컴파일 에러 발생함.

}
