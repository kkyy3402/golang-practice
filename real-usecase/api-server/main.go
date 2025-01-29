package main

import (
	"fmt"
	todo_controller "golang-practice/real-usecase/api-server/controllers"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	// 라우터 설정
	r := mux.NewRouter()

	// 컨트롤러 라우팅 설정
	todo_controller.RegisterTodoRoutes(r)

	// 서버 시작
	fmt.Println("Server starting at :8080...")
	server := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(server.ListenAndServe())
}
