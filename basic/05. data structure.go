package main

import "fmt"

// 배열 할당 후
func test1() {
	var numbers [5]int
	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30
	numbers[3] = 40
	numbers[4] = 50

	for number := range numbers {
		println(number)
	}

	// C-Style의 배열 순회
	for i := 0; i < len(numbers); i++ {
		println("i : ", i, "number : ", numbers[i])
	}
}

func arraySlicing() {
	items := []string{"foo", "bar", "hello"}

	// 1:3 이렇게 슬라이싱하면 인덱스 기준으로 1,2가 된다. 특이허네
	for index, item := range items[1:3] {
		println("index : ", index, ", item : ", item)
	}
}

// 선언 후 할당
// map이 상당히 별로네...
func mapExample1() {
	ages := make(map[string]int)
	ages["age"] = 5
	ages["name"] = 5

	println(ages) //fmt.Println으로 하면 전체가 찍힘
	println(ages["age"])
	println(ages["name"])
}

// 선언과 동시 할당. 우웩
func mapExample2() {
	scores := map[string]int{
		"Math":    10,
		"Science": 50,
	}

	fmt.Println("scores:", scores)

	scores["English"] = 70

	fmt.Println("맵:", scores)
}

type Person struct {
	name string
	age  int
}

func structureExample() {
	person1 := Person{
		age:  14,
		name: "david",
	}

	fmt.Println(person1.age)
	fmt.Println(person1.name)
	fmt.Println(person1)
}

// 흠.......
func mapWithStructure() {
	items := []Person{
		{
			name: "A",
			age:  5,
		},
		{
			name: "B",
			age:  10,
		},
	}
	fmt.Println(items)
}

func main() {
	// test1()
	// arraySlicing()
	// mapExample1()
	// mapExample2()
	// structureExample()
	mapWithStructure()
}
