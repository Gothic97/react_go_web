package main

import (
	"fmt"
)

func main() {

	/*
	    < switch 'variable' { case 'value': 'code' }

	    switch variable {

	     case value1:
	     // value1일 경우 여기에 있는 것 실행

	     case value2:
	     // value2일 경우 여기에 있는 것 실행

	   default:
	     // 그 외 여기 있는 것 실행

	   }
	*/

	/*
		i := 3

		switch i {

		case 0:
			fmt.Println(0)
		case 1:
			fmt.Println(1)
		case 2:
			fmt.Println(2)
		case 3:
			fmt.Println(3)
		case 4:
			fmt.Println(4)
		default:
			fmt.Println(-1)

		}
	*/

	/* < Using string >
	s := "world"

	switch s {

	case "hello":
		fmt.Println("hello")
	case "world":
		fmt.Println("world")
	default:
		fmt.Println("일치하는 문자열이 없습니다.")

	}
	*/

	/* < Using break > // 브레이크는 작은 리턴 같은 것이다. 손위의 명령을 넘어 명령을 실행시킨다.
	s := "Hello"
	i := 2

	switch i {

	case 1:
		fmt.Println(1)
	case 2:
		if s == "Hello" {
			fmt.Println("Hello 2")
			break
		}

		fmt.Println(2)

	}

	*/

	// <Using fallthrough> // fallthrough는 switch문에서 지속적으로 다음 케이스로 넘어가게 한다.

	i := 3

	switch i {

	case 4:
		fmt.Println("4 이상")
		fallthrough
	case 3:
		fmt.Println("3 이상")
		fallthrough
	case 2:
		fmt.Println("2 이상")
		fallthrough
	case 1:
		fmt.Println("1 이상")
		fallthrough
	case 0:
		fmt.Println("0 이상") // Go 언어 구조상 마지막엔 fallthrough를 넣으면 안 됨.

	}

	// <processing a lot of terms> // 다항처리
	// i := 3
	//
	// switch i {
	//
	// case 2, 4, 6: // when i is 2, 4, 6.
	// 	fmt.Println("짝수")
	// case 1, 3, 5: // when i is 1, 3, 5.
	// 	fmt.Println("홀수")
	//
	// }

	/* <furcate as a condition-formula> // 상태 형식별로 분기

	i := 7

	switch {
	case i >= 5 && i < 10:
		fmt.Println("5 이상, 10 미만")
	case i >= 0 && i < 5:
		fmt.Println("0 이상, 5 미만")
	}


	*/

	/*

	   switch 분기문 안에서 함수를 실행하고 결과값으로 분기할 수 있음. 이때 함수를 호출하고 뒤에 ;(세미콜론)을 붙여줌. 또한 case에서는 값으로 분기할 수 없고 조건식만 사용할 수 있음.

	   * math/rand: random package
	     - Seed: function set up value of seed // 시드값을 설정하는 함수
	     - Intn: create random value. scope of random value is from 0 to the value set up as a parameter. // 랜덤 값을 만든다. 랜덤값의 범위는 0부터 파라미터(매개변수)로 설정된 값까지다.

	   * time: time package
	     - Now: function. that find out present time. // 현재 시간 나타내는 함수
	     - UnixNano: return Unix time as a unit of nano-second. // 나노초 단위로 유닉스 타임을 리턴한다.

	*/

	// rand.Seed(time.Now().UnixNano())
	// switch i := rand.Intn(10); {
	// case i >= 3 && i < 6:
	// 	fmt.Println("3 이상, 6 미만")
	// case i == 9:
	// 	fmt.Println("9")
	// default:
	// 	fmt.Println(i)
	// }

}
