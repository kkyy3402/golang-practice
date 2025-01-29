package main

import "fmt"

// 포인터는 C랑 똑같음
func pointerExample() {
	var num int = 10
	var ptr *int // int 타입 포인터 변수 선언

	ptr = &num

	fmt.Println(ptr)
	fmt.Println(*ptr)
}

// array에 포인터 적용해보기
func pointerExample2() {
	var numbers [5]int = [5]int{10, 20, 30, 40, 50}

	var ptr *int = &numbers[0]
	fmt.Println(ptr)
	fmt.Println(*ptr)
}

// 값에 의한 전달
func callByValue(num int) {
	num = num * 2
}

// 참조에 의한 전달
func callByPointer(ptr *int) {
	*ptr = *ptr * 2
}

func main() {
	// pointerExample()
	// pointerExample2()

	// CallbyRefenceTest
	a := 5
	callByValue(a)
	fmt.Println("call by value result : ", a)

	// CallbyValueTest
	b := 5
	callByPointer(&b)
	fmt.Println("call by reference result : ", b)
}
