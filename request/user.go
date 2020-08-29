package request

// UserID ...
type UserID struct {
	ID int64 `param:"id"`
}

// UserCreate ...
type UserCreate struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// UserUpdate ...
type UserUpdate struct {
	ID   int64  `param:"id"`
	Name string `json:"name"`
}
