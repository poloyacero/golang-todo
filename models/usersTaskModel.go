package models

type UsersTask struct {
	ID     string `json:"id" bson:"_id,omitempty"`
	TaskId string `json:"taskId"`
	UserId string `json:"userId"`
}
