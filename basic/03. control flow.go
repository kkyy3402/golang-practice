package main

import "fmt"

// 선언이후 분기처리
func test1() {
	age := 50

	if age > 50 {
		fmt.Println("나이 50초과")
	} else {
		fmt.Println("나이 50이하")
	}
}

// 선언과 동시에 분기처리
func test2() {
	if age := 50; age > 50 {
		fmt.Println("나이 50 초과")
	} else {
		fmt.Println("나이 50 이하")
	}
}

// for문 사용법. C-style이군
func test3() {
	for i := 0; i < 5; i++ {
		println(i)
	}
}

// 다른 for문 스타일. while 없고, 이 스타일로 while 커버하면됨
func test4() {
	i := 0
	for i < 5 {
		println(i)
		i++
	}
}

func test5() {
	array := []int{1, 2, 3, 4, 5}

	println("Printing array..")
	for index, value := range array {
		println("index : ", index, ", value : ", value)
	}

	println("printing element in array : ", array[4])
	// println("printing element in array : ", array[5]) // index error

	// Array 한김에 append 예제
	newArray := append(array, 6)

	println("Printing array..")
	for index, value := range newArray {
		println("index : ", index, ", value : ", value)
	}

}

func test6() {
	grade := "B"

	switch grade {
	case "A":
		println("case a!")
	case "B":
		println("case b!")
	}

}

func main() {
	// test1()
	// test2()
	// test3()
	// test4()
	// test5()
	test6()
}
