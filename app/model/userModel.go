package model

// User represents user resources.
type User struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	Country    string `json:"country"`
	Phone      string `json:"phone"`
	PostalCode string `json:"postal_code"`
}

// UpdateUser represents user resources.
type UpdateUser struct {
	Name       string `json:"name" validate:"required"`
	LastName   string `json:"last_name" validate:"required"`
	Password   string `json:"password" validate:"required"`
	Email      string `json:"email" validate:"required,email"`
	Country    string `json:"country" validate:"required"`
	Phone      string `json:"phone" validate:"required"`
	PostalCode string `json:"postal_code" validate:"required"`
}

// CreateUser represents user resources.
type CreateUser struct {
	Name       string `json:"name" validate:"required"`
	LastName   string `json:"last_name" validate:"required"`
	Email      string `json:"email" validate:"required,email"`
	Country    string `json:"country" validate:"required"`
	Phone      string `json:"phone" validate:"required"`
	PostalCode string `json:"postal_code" validate:"required"`
}
