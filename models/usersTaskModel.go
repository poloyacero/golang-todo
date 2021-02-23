package models

type UsersTask struct {
	ID     string      `json:"id" bson:"_id,omitempty"`
	TaskId interface{} `json:"taskId"`
	UserId string      `json:"userId"`
}
