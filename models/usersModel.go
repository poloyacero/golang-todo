package models

type User struct {
	ID       string `json:"id" bson:"_id,omitempty"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
