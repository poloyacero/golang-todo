package controllers

import (
	"encoding/json"
	"fmt"
	"go-server/db"
	"go-server/models"
	"net/http"
	"strings"
)

var tasks []models.Task

func CreateTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create Task")
	var task models.Task
	_ = json.NewDecoder(r.Body).Decode(&task)
	db.InsertTask(task)
	json.NewEncoder(w).Encode(task)
}

func ReadTasks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Read Tasks")
	tasks := db.GetTasks()
	json.NewEncoder(w).Encode(tasks)
}

func ReadTask(w http.ResponseWriter, r *http.Request, paramId string) {
	fmt.Println("Read Single Task", paramId)
	payload := db.GetTasks()
	for _, p := range payload {
		if p.ID == paramId {
			json.NewEncoder(w).Encode(p)
			return
		}
	}
	json.NewEncoder(w).Encode("Task not found")
}

func UpdateTask(w http.ResponseWriter, r *http.Request, paramId string) {
	fmt.Println("Update Task", paramId)

}

func DeleteTask(w http.ResponseWriter, r *http.Request, paramId string) {
	fmt.Println("Delete Task", paramId)
	db.DeleteTask(paramId)
	json.NewEncoder(w).Encode("Task deleted")
}

func TaskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		ReadTasks(w, r)
		return

	case "POST":
		CreateTask(w, r)
		return

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode("Method Not Allowed")
		return
	}
}

func TaskSpecificHandler(w http.ResponseWriter, r *http.Request) {
	// array of split the url
	url := strings.Split(r.URL.String(), "/")

	paramId := url[2]

	if len(url) != 3 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("404 Not Found")
		return
	}

	switch r.Method {
	case "GET":
		ReadTask(w, r, paramId)
		return

	case "DELETE":
		DeleteTask(w, r, paramId)
		return

	case "PUT":
		UpdateTask(w, r, paramId)
		return

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode("Method not allowed")
		return
	}
}
