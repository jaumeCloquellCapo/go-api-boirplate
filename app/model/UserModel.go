package model

type User struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	LastName   string `json:"last_name"`
	Password   string `json:"password"`
	Email      string `json:"email"`
	Country    string `json:"country"`
	Phone      string `json:"phone"`
	PostalCode string `json:"postal_code"`
}

type UpdateUser struct {
	Name       string `json:"name" validate:"required"`
	LastName   string `json:"last_name" validate:"required"`
	Password   string `json:"password" validate:"required"`
	Email      string `json:"email" validate:"required"`
	Country    string `json:"country" validate:"required"`
	Phone      string `json:"phone" validate:"required"`
	PostalCode string `json:"postal_code" validate:"required"`
}

type CreateUser struct {
	Name       string `json:"name" validate:"required"`
	LastName   string `json:"last_name" validate:"required"`
	Email      string `json:"email" validate:"required"`
	Country    string `json:"country" validate:"required"`
	Phone      string `json:"phone" validate:"required"`
	PostalCode string `json:"postal_code" validate:"required"`
	Password   string `json:"password" validate:"required"`
}

type Credentials struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (User) TableName() string { return "users" }
