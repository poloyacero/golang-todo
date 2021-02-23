package response

import (
	"encoding/json"
	"go-todo/models"
	"net/http"
)

func JSON(w http.ResponseWriter, any interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(any)
}

/*func JSON(w http.ResponseWriter, user models.User) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}*/

func JSONS(w http.ResponseWriter, tasks []models.Task) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func ERROR(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(message)
}

func MESSAGE(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(message)
}
