package model

// User ...
type User struct {
	Base
	Name  string `json:"name"`
	Email string `json:"email"`
}
