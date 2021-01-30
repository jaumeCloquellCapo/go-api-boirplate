package model

type User struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	LastName   string `json:"last_name"`
	Password   string `json:"password"`
	Email      string `form:"email" json:"email" validate:"required"`
	Country    string `json:"country"`
	Phone      string `json:"phone"`
	PostalCode string `json:"postal_code"`
}

type UpdateUser struct {
	Name       string `form:"name" json:"name"`
	LastName   string `form:"last_name" json:"last_name"`
	Password   string `form:"password" json:"password"`
	Email      string `form:"email" json:"email"`
	Country    string `form:"country" json:"country"`
	Phone      string `form:"phone" json:"phone"`
	PostalCode string `form:"postal_code" json:"postal_code"`
}

type CreateUser struct {
	Name     string `form:"name" json:"name" binding:"required,max=100"`
	Email    string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required"`
}

type Credentials struct {
	Email    string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required"`
}

func (User) TableName() string { return "users" }
