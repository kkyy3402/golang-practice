package todo_controller

import (
	"encoding/json"
	"fmt"
	todo_service "golang-practice/real-usecase/api-server/services"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// JSON 응답을 처리하는 유틸리티 함수
func writeJSONResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		log.Println("Error encoding JSON response:", err)
		return
	}
}

// RegisterTodoRoutes 함수
func RegisterTodoRoutes(r *mux.Router) {
	// API 엔드포인트 정의
	r.HandleFunc("/todos", GetTodos).Methods("GET")
	r.HandleFunc("/todos/{id}", GetTodo).Methods("GET")
}

// GetTodos 핸들러
func GetTodos(w http.ResponseWriter, r *http.Request) {

	todos, err := todo_service.GetTodos()
	if err != nil {
		http.Error(w, "Error fetching todos from API", http.StatusInternalServerError)
		log.Println("Error fetching todos from API:", err)
		return
	}
	writeJSONResponse(w, todos, http.StatusOK)
	fmt.Println("Sent todos data from JSONPlaceholder API.")
}

// GetTodo 핸들러
func GetTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		http.Error(w, "Todo ID is missing in path", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid todo ID", http.StatusBadRequest)
		return
	}
	todo, err := todo_service.GetTodo(id)
	if err != nil {
		http.Error(w, "Error fetching todo from API", http.StatusInternalServerError)
		log.Println("Error fetching todo from API:", err)
		return
	}
	writeJSONResponse(w, todo, http.StatusOK)
	fmt.Printf("Sent todo with id %d data from JSONPlaceholder API\n", id)
}
