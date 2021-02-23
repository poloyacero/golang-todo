package main

import (
	"fmt"
	"go-todo/controllers"
	"go-todo/middlewares"
	"log"
	"net/http"
)

var mySigningKey = []byte("opensesame")

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Println("test")
}

func main() {
	fmt.Println("Hello World")
	http.Handle("/tasks", middlewares.IsAuthorized(controllers.TaskHandler))
	http.Handle("/tasks/", middlewares.IsAuthorized(controllers.TaskSpecificHandler))
	http.HandleFunc("/auth/", controllers.UserHandler)
	http.Handle("/test", middlewares.IsAuthorized(test))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
