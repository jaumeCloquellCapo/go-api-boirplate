package model

type User struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	LastName   string `json:"lastName"`
	Password   string `json:"password"`
	Email      string `json:"email"`
	Country    string `json:"country"`
	Phone      string `json:"phone"`
	PostalCode string `json:"postalCode"`
}

type UserSignUp struct {
	Name     string `form:"name" json:"name" binding:"required,max=100"`
	Email    string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required"`
}

type UserLogin struct {
	Email    string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required"`
}

func (User) TableName() string { return "users" }
