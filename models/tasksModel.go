package models

import "time"

//Title, Description, Due Date, Priority

type Task struct {
	ID          string    `json:"id" bson:"_id,omitempty"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Priority    int       `json:"priority"`
	DueDate     time.Time `json:"duedate"`
	IsComplete  bool      `json:"isComplete"`
}
