// continue 문의 용도: 잦은 오류이나, 고치기 어렵고, 프로그램 이용에 크게 문제를 일으키지 않는 오류가 있을 때, 해당 오류를 규정해서, continue문으로 넘기면 좋을 듯함.

package main

import . "fmt"

func main() {

	/*
		for i := 0; i < 5; i++ {
			if i == 2 { // i가 2일 때
				continue // 아래 부분 코드를 실행하지 않고 넘어감
			}

			Println(i)

		}
	*/

Loop:

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if j == 2 {
				continue Loop
			}

			Println(i, j)

		}
	}

	Println("Hello, world!")

}
