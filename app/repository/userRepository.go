package repository

import (
	error2 "ApiRest/app/error"
	"ApiRest/app/model"
	"ApiRest/provider"
	"database/sql"
	"fmt"
	"log"
)

type userRepository struct {
	db *provider.DbStore
}

//UserRepository
type UserRepositoryInterface interface {
	GetUsers() ([]model.User, error)
	GetUserById(id int) (user *model.User, err error)
	RemoveUserById(id int) error
	UpdateUserById(id int, user model.UpdateUser) error
	GetUserByEmail(email string) (user *model.User, err error)
	CreateUser(model.CreateUser) (user *model.User, err error)
}

//NewUserRepository
func NewUserRepository(db *provider.DbStore) UserRepositoryInterface {
	return &userRepository{
		db,
	}
}

//FindById
func (r *userRepository) GetUserById(id int) (user *model.User, err error) {
	user = &model.User{}

	var query = "SELECT id, email, name, password FROM users WHERE id = ?"
	row := r.db.QueryRow(query, id)

	if err := row.Scan(&user.ID, &user.Email, &user.Name, &user.Password); err != nil {
		if err == sql.ErrNoRows {
			return nil, error2.NewErrorNotFound(fmt.Sprintf("Error: User not found by ID %d", id))
		}

		return nil, err
	}

	return user, nil
}

func (r *userRepository) RemoveUserById(id int) error {

	_, err := r.db.Exec(`DELETE FROM users WHERE id = $1;`, id)
	if err != nil {
		panic(err)
	}

	if err != nil {
		fmt.Print(err.Error())

	}

	return err
}

func (r *userRepository) UpdateUserById(id int, user model.UpdateUser) error {
	result, err := r.db.Exec("UPDATE users SET name = ?, email = ?,  = ?, last_name = ?, country = ?, phone = ?, postal_code = ? where id = ?", user.Name, user.Email, user.LastName, user.Country, user.Phone, user.PostalCode, id)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := result.RowsAffected()

	if err != nil {
		log.Fatal(err)
	}

	if rows != 1 {
		return error2.NewErrorNotFound(fmt.Sprintf("Error: User not found by id %s", id))
	}

	return err
}

//GetUserByEmail
func (r *userRepository) GetUserByEmail(email string) (user *model.User, err error) {

	user = &model.User{}

	var query = "SELECT id, email, name, password FROM users WHERE email = ?"
	row := r.db.QueryRow(query, email)

	if err := row.Scan(&user.ID, &user.Email, &user.Name, &user.Password); err != nil {
		if err == sql.ErrNoRows {
			return nil, error2.NewErrorNotFound(fmt.Sprintf("Error: User not found by email %s", email))
		}
		return nil, err
	}
	return user, nil
}

func (r *userRepository) GetUsers() (users []model.User, err error) {
	users = []model.User{}
	var query = "SELECT id, email, name, password FROM users"
	rows, err := r.db.Query(query)
	defer rows.Close()

	for rows.Next() {
		var user = model.User{}
		err := rows.Scan(&user.ID, &user.Email, &user.Name, &user.Password)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return
}

func (r *userRepository) CreateUser(UserSignUp model.CreateUser) (user *model.User, err error) {

	query := "INSERT INTO users (name, password , email) values  (?, ?, ?)"
	res, err := r.db.Exec(query, UserSignUp.Name, UserSignUp.Password, UserSignUp.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println(err)
		}
		return
	}

	id, err := res.LastInsertId()

	if err != nil {
		return nil, err
	}

	return &model.User{
		ID:       id,
		Name:     UserSignUp.Name,
		Password: UserSignUp.Password,
		Email:    UserSignUp.Email,
	}, nil
}
