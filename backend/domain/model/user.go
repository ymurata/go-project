package model

// User ...
type User struct {
	BaseWithHashID
	Name  string `json:"name"`
	Email string `json:"email"`
}

// TableName ...
func (User) TableName() string {
	return "users"
}
