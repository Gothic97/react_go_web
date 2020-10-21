package main

import "fmt"

func main() {

	//부동소수점 반올림 오차(Rounding error) 확인

	var a float64 = 10.0

	for i := 0; i < 10; i++ {
		a = 1 - 0.1
	}

	fmt.Println(a)
	fmt.Println(a == 0.9)

	// <가장 빨리 만나는 GO 언어> 41page에 따르면 반올림 오차가 발생해야 하나, 발생하지 않는 것으로 보아,GO언어의 float 데이터 저장 형식에 변화가 있음을 알 수 있음.
	// Q) 부동소수점 반올림 오차(Rounding error)란?
	// A)  수학적으로 실수는 무수히 많음. 그러나 컴퓨터는 유한한 수만 받아들이고 표현할 수 있음. 그러므로 컴퓨터가 표현하는 실수들은 실제 실수와 다를 수 있음.
	//    실제 실수와 컴퓨터의 실수가 다를 때, 이 상황을 두고 '부동소수점 반올림 오차(Rounding error)'라고 하는 것 같음.

}
