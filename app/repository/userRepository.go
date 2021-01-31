package repository

import (
	error2 "ApiRest/app/error"
	"ApiRest/app/model"
	"ApiRest/internal/storage"
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
)

type userRepository struct {
	db *storage.DbStore
}

//UserRepositoryInterface ...
type UserRepositoryInterface interface {
	FindAll() ([]model.User, error)
	FindById(id int) (user *model.User, err error)
	RemoveById(id int) error
	UpdateById(id int, user model.UpdateUser) error
	FindByEmail(email string) (user *model.User, err error)
	Create(model.CreateUser) (user *model.User, err error)
}

//NewUserRepository ...
func NewUserRepository(db *storage.DbStore) UserRepositoryInterface {
	return &userRepository{
		db,
	}
}

//FindById ...
func (r *userRepository) FindById(id int) (user *model.User, err error) {
	user = &model.User{}

	var query = "SELECT id, email, name, postal_code, phone, last_name, country FROM users WHERE id = ?"
	row := r.db.QueryRow(query, id)

	if err := row.Scan(&user.ID, &user.Email, &user.Name, &user.PostalCode, &user.Phone, &user.LastName, &user.Country); err != nil {
		if err == sql.ErrNoRows {
			return nil, error2.NewErrorNotFound(fmt.Sprintf("Error: User not found by ID %d", id))
		}

		return nil, err
	}

	return user, nil
}

func (r *userRepository) RemoveById(id int) error {

	_, err := r.db.Exec(`DELETE FROM users WHERE id = $1;`, id)
	return err
}

//UpdateById ...
func (r *userRepository) UpdateById(id int, user model.UpdateUser) error {
	result, err := r.db.Exec("UPDATE users SET name = ?, email = ?, last_name = ?, phone = ?, postal_code = ?, country = ? where id = ?", user.Name, user.Email, user.LastName, user.Phone, user.PostalCode, user.Country, id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rows != 1 {
		return error2.NewErrorNotFound(fmt.Sprintf("UpdateById: User not found by id %v", id))
	}

	return nil
}

//FindByEmail
func (r *userRepository) FindByEmail(email string) (user *model.User, err error) {

	user = &model.User{}

	var query = "SELECT id, email, name, password FROM users WHERE email = ?"
	row := r.db.QueryRow(query, email)

	if err := row.Scan(&user.ID, &user.Email, &user.Name, &user.Password); err != nil {
		if err == sql.ErrNoRows {
			return nil, error2.NewErrorNotFound(fmt.Sprintf("FindByEmail: User not found by email %s", email))
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

	query := "INSERT INTO users (name, password, email, last_name, phone, postal_code, country) values  (?, ?, ?, ?, ?, ?, ?)"
	res, err := r.db.Exec(query, UserSignUp.Name, UserSignUp.Password, UserSignUp.Email, UserSignUp.LastName, UserSignUp.Phone, UserSignUp.PostalCode, UserSignUp.Country)
	if err != nil {
		if me, ok := err.(*mysql.MySQLError); ok {
			if me.Number == 1062 {
				return nil, error2.NewErrorAlreadyExist(fmt.Sprintf("User by email %s already exist", UserSignUp.Email))
			}
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
