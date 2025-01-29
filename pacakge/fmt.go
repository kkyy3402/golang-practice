package main

import "fmt"

func main() {
	// Print: 기본 출력
	fmt.Print("Hello, ")
	fmt.Print("World!\n")

	// Println: 줄바꿈 포함 출력
	fmt.Println("This is a new line.")

	// Printf: 형식 지정 출력
	name := "Go"
	age := 13
	fmt.Printf("Language: %s, Age: %d\n", name, age)

	// Scan: 사용자 입력 받기. C-style이다. 이거도
	var input string
	fmt.Print("Enter some text: ")
	fmt.Scan(&input)
	fmt.Println("You entered:", input)

	// Scanf: 형식 지정 입력 받기
	var intInput int
	fmt.Print("Enter an integer: ")
	fmt.Scanf("%d", &intInput)
	fmt.Println("You entered an integer: ", intInput)
}
