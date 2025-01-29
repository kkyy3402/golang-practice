package main

import (
	"fmt"
	"io"
	"os"
)

// writeFile 함수: 파일에 데이터를 쓰는 함수
func writeFile(filename string, data string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer file.Close()

	_, err = file.WriteString(data)
	if err != nil {
		return fmt.Errorf("error writing to file: %w", err)
	}
	return nil
}

// readFile 함수: 파일에서 데이터를 읽는 함수
func readFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	buf := make([]byte, 128)
	var content string
	for {
		n, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", fmt.Errorf("error reading from file: %w", err)
		}
		content += string(buf[:n])
	}
	return content, nil
}

func main() {

	filename := "../assets/example.txt"
	initialData := "파일 씁니다.."

	// 파일 쓰기
	err := writeFile(filename, initialData)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
	fmt.Println("Successfully wrote initial text to", filename)

	// 파일 읽기
	content, err := readFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("Content of", filename, ":\n", content)
	fmt.Println("Successfully read from", filename)
}
