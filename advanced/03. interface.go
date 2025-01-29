package main

import "fmt"

// 인터페이스 정의
type Shape interface {
	Area() float64
	Perimeter() float64
}

// 구조체 정의
type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

// Rectangle 구조체가 Shape 인터페이스 구현
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Circle 구조체가 Shape 인터페이스 구현
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

// struct가 인터페이스를 가져가는? 형태로 구현됨. 특이허다.

func test1() {
	rect := Rectangle{Width: 5, Height: 10}
	circle := Circle{Radius: 3}

	printShapeInfo(rect)
	printShapeInfo(circle)
}

func printShapeInfo(s Shape) {
	fmt.Printf("Area: %f, Perimeter: %f\n", s.Area(), s.Perimeter())
}

// 빈 인터페이스 형태. JSON을 param으로 받을때 사용하면 좋다고 함.
func printValue(value interface{}) {
	fmt.Printf("Type: %T, Value: %v\n", value, value)
}

func test2() {
	printValue(10)
	printValue("hello")
	printValue(3.14)
	printValue(true)
	printValue(struct{ name string }{name: "Alice"})
}

func main() {
	// test1()
	test2()
}
