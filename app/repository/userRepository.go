package repository

import (
	"ApiRest/app/model"
	"database/sql"
)

const userTable = "users"

type userRepository struct {
	db *sql.DB
}

type UserRepository interface {
	FindById(id int) (model.User, error)
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		db,
	}
}

func (r *userRepository) FindById(id int) (user model.User, err error) {

	return
}
