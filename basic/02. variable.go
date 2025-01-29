package main

import "fmt"

// 선언 이후 할당
func test1() {
	var age int
	var name string
	var height float64

	// 기본값 확인
	fmt.Println(age)
	fmt.Println(name)
	fmt.Println(height)

	age = 30
	name = "foo"
	height = 180

	fmt.Println(age)
	fmt.Println(name)
	fmt.Println(height)
}

// 선언과 동시에 할당
func test2() {
	age := 30         // Int
	name := "bar"     // String
	height := 180.0   // Float
	isStudent := true // Boolean

	fmt.Println(age)
	fmt.Println(name)
	fmt.Println(height)
	fmt.Println(isStudent)
}

// 선언과 동시에 할당2
func test3() {
	var age = 30
	var name = "bar"
	var height = 180.0
	var isStudent = true

	fmt.Println(age)
	fmt.Println(name)
	fmt.Println(height)
	fmt.Println(isStudent)
}

func test4() {
	const name = "foo"

	fmt.Println("name : ", name)

	// 타입스크립트의 이런거는 없다고 함
	// fmt.Println("name : ${name}")
}

func main() {
	// test1()
	// test2()
	// test3()
	test4()
}
