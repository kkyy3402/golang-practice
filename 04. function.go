package main

import "fmt"

func divide(x int, y int) (int, int, error) {
	if y == 0 {
		return 0, 0, fmt.Errorf("division by zero")
	}

	quotient := x / y
	remainder := x % y
	return quotient, remainder, nil
}

func add(numbers ...int) int {
	var sumOfValues = 0
	for _, number := range numbers {
		sumOfValues += number
	}
	return sumOfValues
}

func main() {
	// 나누기 함수 호출
	a, b, c := divide(4, 2)
	println(a)
	println(b)
	println(c)

	// 에러 출력
	d, e, f := divide(4, 0)
	println(d)
	println(e)
	println(f.Error())

	// 더하기
	result := add(1, 2, 3)
	print("result : ", result)
}
