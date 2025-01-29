package todo_service

import (
	"fmt"
	"golang-practice/real-usecase/api-server/models"
	"log"

	"github.com/go-resty/resty/v2"
)

func GetTodos() ([]models.Todo, error) {
	client := resty.New()
	var todos []models.Todo
	_, err := client.R().
		SetResult(&todos).
		Get("https://jsonplaceholder.typicode.com/todos")

	if err != nil {
		log.Println("Error fetching todos from API:", err)
		return nil, fmt.Errorf("failed to fetch todos: %w", err)
	}
	return todos, nil
}

func GetTodo(id int) (*models.Todo, error) {
	client := resty.New()
	var todo models.Todo
	_, err := client.R().
		SetResult(&todo).
		Get(fmt.Sprintf("https://jsonplaceholder.typicode.com/todos/%d", id))

	if err != nil {
		log.Println("Error fetching todo from API:", err)
		return nil, fmt.Errorf("failed to fetch todo: %w", err)
	}
	return &todo, nil
}
