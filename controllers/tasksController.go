package controllers

import (
	"encoding/json"
	"fmt"
	"go-todo/db"
	"go-todo/middlewares"
	"go-todo/models"
	response "go-todo/utility"
	"net/http"
	"strings"
)

var tasks []models.Task

func CreateTask(w http.ResponseWriter, r *http.Request) {
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	reqToken = splitToken[1]

	claims, _ := middlewares.ExtractTokenClaims(reqToken)

	client_id := claims["client_id"].(string)
	fmt.Println("Client ID", client_id)

	fmt.Println("Create Task", reqToken)
	var task models.Task
	_ = json.NewDecoder(r.Body).Decode(&task)

	taskId := db.InsertTask(task)
	//strTaskId := fmt.Sprintf("%v", taskId)
	fmt.Println("Task ID", taskId)

	userTask := models.UsersTask{
		TaskId: taskId,
		UserId: client_id,
	}

	db.AssignTask(userTask)
	response.JSON(w, task)
}

func ReadTasks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Read Tasks")
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	reqToken = splitToken[1]

	claims, _ := middlewares.ExtractTokenClaims(reqToken)

	client_id := claims["client_id"].(string)

	tasks = db.GetTasksByUser(client_id)

	response.JSONS(w, tasks)
}

func ReadTask(w http.ResponseWriter, r *http.Request, paramId string) {
	fmt.Println("Read Single Task", paramId)
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	reqToken = splitToken[1]

	claims, _ := middlewares.ExtractTokenClaims(reqToken)

	client_id := claims["client_id"].(string)
	/*payload := db.GetTasks()
	for _, p := range payload {
		if p.ID == paramId {
			response.JSON(w, p)
			return
		}
	}*/
	exist := db.CheckUserTask(client_id, paramId)
	if !exist {
		response.ERROR(w, "Permission Denied")
		return
	}
	result := db.GetUserTask(paramId)
	response.JSON(w, result)
}

func UpdateTask(w http.ResponseWriter, r *http.Request, paramId string) {
	fmt.Println("Update Task", paramId)

	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	reqToken = splitToken[1]

	claims, _ := middlewares.ExtractTokenClaims(reqToken)

	client_id := claims["client_id"].(string)

	exist := db.CheckUserTask(client_id, paramId)
	if !exist {
		response.ERROR(w, "Permission Denied")
		return
	}

	var task models.Task
	_ = json.NewDecoder(r.Body).Decode(&task)
	db.UpdateTask(task, paramId)
	response.JSON(w, "Task Updated")
}

func DeleteTask(w http.ResponseWriter, r *http.Request, paramId string) {
	fmt.Println("Delete Task", paramId)

	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	reqToken = splitToken[1]

	claims, _ := middlewares.ExtractTokenClaims(reqToken)

	client_id := claims["client_id"].(string)

	exist := db.CheckUserTask(client_id, paramId)
	if !exist {
		response.ERROR(w, "Permission Denied")
		return
	}

	db.DeleteTask(paramId)
	response.JSON(w, "Task deleted")
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
