package model

// User ...
type User struct {
	BaseWithHashID
	Name  string `json:"name"`
	Email string `json:"email"`
}
