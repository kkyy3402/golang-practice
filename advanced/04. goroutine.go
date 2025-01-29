package main

import (
	"fmt"
	"time"
)

func printNumber(num int) {
	fmt.Println("Go routine:", num)
}

func main() {
	fmt.Println("start main")
	for i := 0; i < 5; i++ {
		// go routine으로 비동기로 동작하게 된다.
		go printNumber(i)
	}
	time.Sleep(2 * time.Second)
	fmt.Println("end main")
}
