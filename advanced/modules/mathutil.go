package modules

// Add 함수는 두 정수의 합을 반환합니다.
func Add(x int, y int) int {
	return x + y
}

// PrivateAdd 함수 (패키지 내부에서만 사용 가능)
func privateAdd(x int, y int) int {
	return x + y
}
