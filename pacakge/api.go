package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
)

// Todo 구조체
type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// JSON 응답을 처리하는 유틸리티 함수
func writeJSONResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		log.Println("Error encoding JSON response:", err)
		return
	}
}

// todosHandler 함수 (go-resty 사용)
func todosHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	client := resty.New()
	var todos []Todo
	_, err := client.R().
		SetResult(&todos).
		Get("https://jsonplaceholder.typicode.com/todos")

	if err != nil {
		http.Error(w, "Error fetching todos from API", http.StatusInternalServerError)
		log.Println("Error fetching todos from API:", err)
		return
	}
	writeJSONResponse(w, todos, http.StatusOK)
	fmt.Println("Sent todos data from JSONPlaceholder API.")
}

func main() {
	http.HandleFunc("/todos", todosHandler)

	fmt.Println("Server starting at :8080...")
	server := &http.Server{
		Addr:           ":8080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(server.ListenAndServe())

}
