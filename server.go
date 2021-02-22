package main

import (
	"fmt"
	"go-server/controllers"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello World")
	http.HandleFunc("/tasks", controllers.TaskHandler)
	http.HandleFunc("/tasks/", controllers.TaskSpecificHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
