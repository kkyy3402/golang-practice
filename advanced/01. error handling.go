package main

import (
	"errors"
	"fmt"
	"io"
)

func test1() {
	// 표준 에러 예시
	if _, err := io.ReadAll(nil); err != nil {
		fmt.Println("Error:", err) // 에러 출력
	}

	// errors 패키지를 이용한 에러 생성
	err := errors.New("custom error") //에러 어거지로 만들기
	fmt.Println(err)                  // custom error 출력
}

func main() {
	test1()
}
