package repository

import (
	error2 "ApiRest/app/error"
	"ApiRest/app/model"
	"ApiRest/internal/storage"
	"database/sql"
	"fmt"
)

// userRepository handles communication with the user store
type userRepository struct {
	db *storage.DbStore
}

//UserRepositoryInterface define the user repository interface methods
type UserRepositoryInterface interface {
	FindById(id int) (user *model.User, err error)
	RemoveById(id int) error
	UpdateById(id int, user model.UpdateUser) error
	Create(model.CreateUser) (user *model.User, err error)
}

// NewUserService implements the user repository interface.
func NewUserRepository(db *storage.DbStore) UserRepositoryInterface {
	return &userRepository{
		db,
	}
}

// FindById implements the method to find a user from the store
func (r *userRepository) FindById(id int) (user *model.User, err error) {
	user = &model.User{}

	var query = "SELECT id, email, name, postal_code, phone, last_name, country FROM users WHERE id = $1"
	row := r.db.QueryRow(query, id)

	if err := row.Scan(&user.ID, &user.Email, &user.Name, &user.PostalCode, &user.Phone, &user.LastName, &user.Country); err != nil {
		if err == sql.ErrNoRows {
			return nil, error2.NewErrorNotFound(fmt.Sprintf("Error: User not found by ID %d", id))
		}

		return nil, err
	}

	return user, nil
}

// RemoveById implements the method to remove a user from the store
func (r *userRepository) RemoveById(id int) error {

	_, err := r.db.Exec(`DELETE FROM users WHERE id = $1;`, id)
	return err
}

// UpdateById implements the method to update a user into the store
func (r *userRepository) UpdateById(id int, user model.UpdateUser) error {
	result, err := r.db.Exec("UPDATE users SET name = $1, email = $2, last_name = $3, phone = $4, postal_code = $5, country = $6 where id = $7", user.Name, user.Email, user.LastName, user.Phone, user.PostalCode, user.Country, id)
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

// Create implements the method to persist a new user
func (r *userRepository) Create(UserSignUp model.CreateUser) (user *model.User, err error) {
	createUserQuery := `INSERT INTO users (name, email, last_name, phone, postal_code, country) 
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id`

	stmt, err := r.db.Prepare(createUserQuery)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var userId int
	err = stmt.QueryRow(UserSignUp.Name, UserSignUp.Email, UserSignUp.LastName, UserSignUp.Phone, UserSignUp.PostalCode, UserSignUp.Country).Scan(&userId)
	if err != nil {
		return nil, err
	}

	return &model.User{
		ID:         userId,
		Name:       UserSignUp.Name,
		Email:      UserSignUp.Email,
		LastName:   UserSignUp.LastName,
		Phone:      UserSignUp.Phone,
		Country:    UserSignUp.Country,
		PostalCode: UserSignUp.PostalCode,
	}, nil
}
