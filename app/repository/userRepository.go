package repository

import (
	"ApiRest/app/model"
	"database/sql"
	"log"
)

type userRepository struct {
	db *sql.DB
}

//UserRepository
type UserRepository interface {
	GetUserById(id int) (model.User, error)
	GetUserByEmail(email string) (model.User, error)
	CreateUser(model.User) error
}

//NewUserRepository
func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		db,
	}
}

//FindById
func (r *userRepository) GetUserById(id int) (user model.User, err error) {
	user = model.User{}
	var query = "SELECT id, email, name, password FROM users WHERE id = ?"
	row := r.db.QueryRow(query, id)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Println(err)
		}
		log.Println("Error", err.Error())
		return model.User{}, err
	}

	if err := row.Scan(&user.ID, &user.Email, &user.Name, &user.Password); err != nil {
		log.Println("Error", err.Error())
		return model.User{}, err
	}

	return user, nil
}

//GetUserByEmail
func (r *userRepository) GetUserByEmail(email string) (user model.User, err error) {

	user = model.User{}
	var query = "SELECT id, email, name, password FROM users WHERE email = ?"
	row := r.db.QueryRow(query, email)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Println(err)
		}

		return model.User{}, err
	}

	if err := row.Scan(&user.ID, &user.Email, &user.Name, &user.Password); err != nil {
		log.Println("Error", err.Error())
		return model.User{}, err
	}

	return user, nil
}

func (r *userRepository) CreateUser(user model.User) (err error) {
	query := "INSERT INTO users (name, password , email) values  ($1, $2, $3)"
	result, err := r.db.Exec(query, user.Name, user.Password, user.Email)
	if err != nil {
		log.Println(err)
		return
	}

	if rowsAffected, err := result.RowsAffected(); err != nil {
		log.Println(rowsAffected)
	}

	return
}
