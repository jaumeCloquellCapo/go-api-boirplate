package repository

import (
	"ApiRest/app/model"
	"database/sql"
	"gopkg.in/dgrijalva/jwt-go.v3"
	"os"
)

type authRepository struct {
	db *sql.DB
}



type AuthRepository interface {
	CreateToken(user model.User) (string, error)
}

func NewAuthRepository (db *sql.DB) AuthRepository {
	return &authRepository{
		db,
	}
}

func (a authRepository) CreateToken(user model.User) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user.ID
	tk := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tk.SignedString([]byte(os.Getenv("API_SECRET")))
}