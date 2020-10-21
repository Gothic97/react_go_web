package main

import f "fmt"

func main() {

	/*
		var a int = 1

		if a == 1 {

			f.Println("Error 1")
			return

		}

		if a == 2 {

			f.Println("Error 2")
			return

		}

		if a == 3 {

			f.Println("Error 1")
			return

		}

		f.Println(a)
		f.Println("Success") // 정상 실행

		// Q) 리턴은 무엇인가?
		// A) 리턴은 카운터 같은 것이다. 구체적으로, 리턴은 '함수 밖으로'를 함의한다.

	*/

	/*
	   	var a int = 1

	   	if a == 1 {

	   		goto ERROR1 // if it's state of affairs of ERROR1, Go to ERROR1

	   	}

	   	if a == 2 {

	   		goto ERROR2 // if it's state of affairs of ERROR1, Go to ERROR1

	   	}

	   	if a == 3 {

	   		goto ERROR1 // if it's state of affairs of ERROR1, Go to ERROR1

	   	}

	   	f.Println(a)
	   	f.Println("Success")
	   	return

	   ERROR1:
	   	f.Println("Error 1")
	   	return

	   ERROR2:
	   	f.Println("Error 2")
	   	return
	*/

	var a int = 1

	if a == 1 {

		goto ERROR // 레벨이라는 라벨으로
		b := 1
		f.Println(b)

	}

ERROR:

	f.Println("Error")

}
