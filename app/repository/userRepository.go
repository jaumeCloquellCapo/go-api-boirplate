package repository

import (
	"ApiRest/app/model"
	"database/sql"
	"log"
)

const userTable = "users"

type userRepository struct {
	db *sql.DB
}

//UserRepository
type UserRepository interface {
	FindById(id int) (model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	CreateUser(model.User) error
}

//NewUserRepository
func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		db,
	}
}

//FindById
func (r *userRepository) FindById(id int) (user model.User, err error) {

	return
}

//GetUserByEmail
func (r *userRepository) GetUserByEmail(email string) (user *model.User, err error) {
	var query = "SELECT * FROM users WHERE EMAIL = $1 LIMIT 1"
	row, err := r.db.Query(query, email)
	defer row.Close()

	if err != nil {
		log.Fatal("GetUserByEmail", err)
	}
	if err := row.Scan(&user); err != nil {
		log.Fatal(err)
	}

	return user, nil
}

func (r *userRepository) CreateUser(user model.User) (err error) {
	var query = "INSERT INTO users (name, password , email) values  ($1, $2, $3)"
	result, err := r.db.Exec(query, user.Name, user.Password, user.Email)
	if err != nil {
		log.Fatal("GetUserByEmail", err)
		return
	}

	if rowsAffected, err := result.RowsAffected(); err != nil {
		log.Println("insertado ", rowsAffected)
	}

	return
}
