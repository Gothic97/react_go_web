package main

import "fmt"

type Rectangle struct {
	width  int
	height int
}

// func NewRectangle(width, height int) *Rectangle {
//
// 	return &Rectangle{width, height}
//
// }

// func rectangleArea(rect *Rectangle) int {
//
// 	return rect.width * rect.height
//
// }

// func rectangleArea(rect *Rectangle) int {
//
//   return rect.width * rect.height
//
// }

// func rectangleScaleA(rect *Rectangle, factor int) {
//
// 	rect.width = rect.width * factor
// 	rect.height = rect.height * factor
//
// }
//
// func rectangleScaleB(rect Rectangle, factor int) {
//
// 	rect.width = rect.width * factor
// 	rect.height = rect.height * factor
//
// }

// func (rect *Rectangle) area() int {
//
// 	return rect.width * rect.height
//
// }

// func (rect *Rectangle) scaleA(factor int) {
//
// 	rect.width = rect.width * factor
// 	rect.height = rect.height * factor
//
// }
//
// func (rect Rectangle) scaleB(factor int) {
//
// 	rect.width = rect.width * factor
// 	rect.height = rect.height * factor
//
// }

func (_ Rectangle) dummy() {

	fmt.Println("dummy")

}

func main() {

	// var rect1 Rectangle
	// var rect2 *Rectangle = new(Rectangle)
	//
	// rect1.height = 20
	// rect2.height = 62
	//
	// fmt.Println(rect1)
	// fmt.Println(rect2)

	// rect := NewRectangle(20, 10)
	// fmt.Println(rect)

	// rect := &Rectangle{20, 10}
	// fmt.Println(rect)

	// rect := Rectangle{30, 30}
	//
	// area := rectangleArea(&rect)
	//
	// fmt.Println(area)

	// rect1 := Rectangle{30, 30}
	// rectangleScaleA(&rect1, 10)
	// fmt.Println(rect1)
	//
	// rect2 := Rectangle{30, 30}
	// rectangleScaleB(rect2, 10)
	// fmt.Println(rect2)

	// 객체지향 관련 용어 정리
	// 객체: 완성된 자립형 메소드. 변수끼리의 참조 및 정의 혼동과 같은 문제 상황에 대한 해결책으로서, 코드의 가독성과 유지보수의 간편성 등의 장점을 가지고 있음.
	// 클래스: 객체 그 자체임. 연관된 변수들과 메소드들을 총칭하는 단어.
	// 메소드: 각 변수들을 저장하는 형식을 가진 함수. 건설에서, 거푸집 내의 정사각형 모양의 독립 가능한 개별 요소를 메소드라고 생각하면 된다.
	// 인스턴스: 메소드를 통해 구체화된 데이터의 집합. 이때 구분 단위는 메소드이다. 건설에서, 거푸집으로 찍어낸 정사각형 모양의 개별 요소를 인스턴스라고 생각하면 된다.

	// rect := Rectangle{10, 20}
	// fmt.Println(rect.area())

	// rect3 := Rectangle{30, 30}
	// rect3.scaleA(10)
	// fmt.Println(rect3)
	//
	// rect4 := Rectangle{30, 30}
	// rect4.scaleB(10)
	// fmt.Println(rect4)

	var r Rectangle
	r.dummy()

}
