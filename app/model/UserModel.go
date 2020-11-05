package model

type User struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UserSignUp struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UserLogin struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (User) TableName() string { return "users" }
