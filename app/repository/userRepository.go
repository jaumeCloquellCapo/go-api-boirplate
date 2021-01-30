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
	FindAll() ([]model.User, error)
	FindById(id int) (user *model.User, err error)
	RemoveById(id int) error
	UpdateById(id int, user model.UpdateUser) error
	FindByEmail(email string) (user *model.User, err error)
	Create(model.CreateUser) (user *model.User, err error)
}

//NewUserRepository
func NewUserRepository(db *provider.DbStore) UserRepositoryInterface {
	return &userRepository{
		db,
	}
}

//FindById
func (r *userRepository) FindById(id int) (user *model.User, err error) {
	user = &model.User{}

	var query = "SELECT id, email, name, postal_code, phone, last_name FROM users WHERE id = ?"
	row := r.db.QueryRow(query, id)

	if err := row.Scan(&user.ID, &user.Email, &user.Name, &user.PostalCode, &user.Phone, &user.LastName); err != nil {
		if err == sql.ErrNoRows {
			return nil, error2.NewErrorNotFound(fmt.Sprintf("Error: User not found by ID %d", id))
		}

		return nil, err
	}

	return user, nil
}

func (r *userRepository) RemoveById(id int) error {

	_, err := r.db.Exec(`DELETE FROM users WHERE id = $1;`, id)
	if err != nil {
		panic(err)
	}

	if err != nil {
		fmt.Print(err.Error())

	}

	return err
}

func (r *userRepository) UpdateById(id int, user model.UpdateUser) error {
	result, err := r.db.Exec("UPDATE users SET name = ?, email = ?, last_name = ?, phone = ?, postal_code = ? where id = ?", user.Name, user.Email, user.LastName, user.Phone, user.PostalCode, id)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := result.RowsAffected()

	if err != nil {
		log.Fatal(err)
	}

	if rows != 1 {
		return error2.NewErrorNotFound(fmt.Sprintf("Error: User not found by id %v", id))
	}

	return err
}

//FindByEmail
func (r *userRepository) FindByEmail(email string) (user *model.User, err error) {

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

func (r *userRepository) FindAll() (users []model.User, err error) {
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

func (r *userRepository) Create(UserSignUp model.CreateUser) (user *model.User, err error) {

	query := "INSERT INTO users (name, password, email, last_name, phone, postal_code) values  (?, ?, ?, ?, ?, ?)"
	res, err := r.db.Exec(query, UserSignUp.Name, UserSignUp.Password, UserSignUp.Email, UserSignUp.LastName, UserSignUp.Phone, UserSignUp.PostalCode)
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
		ID:         int(id),
		Name:       UserSignUp.Name,
		Email:      UserSignUp.Email,
		LastName:   UserSignUp.LastName,
		Phone:      UserSignUp.Phone,
		Country:    UserSignUp.Country,
		PostalCode: UserSignUp.PostalCode,
	}, nil
}
