package repository

import (
	error2 "ApiRest/app/error"
	"ApiRest/app/model"
	"ApiRest/internal/storage"
	"database/sql"
)

// billingRepository handles communication with the user store
type userRepository struct {
	db *storage.DbStore
}

//UserRepositoryInterface define the user repository interface methods
type UserRepositoryInterface interface {
	FindByID(id int) (user *model.User, err error)
	RemoveByID(id int) error
	UpdateByID(id int, user model.UpdateUser) error
	Create(model.CreateUser) (user *model.User, err error)
}

// NewUserRepository implements the user repository interface.
func NewUserRepository(db *storage.DbStore) UserRepositoryInterface {
	return &userRepository{
		db,
	}
}

// FindByID implements the method to find a user from the store
func (r *userRepository) FindByID(id int) (user *model.User, err error) {
	user = &model.User{}

	var query = "SELECT id, cif, name, postal_code, country FROM users WHERE id = $1"
	row := r.db.QueryRow(query, id)

	if err := row.Scan(&user.ID, &user.Cif, &user.Name, &user.PostalCode, &user.Country); err != nil {
		if err == sql.ErrNoRows {
			return nil, error2.ErrNotFound
		}

		return nil, err
	}

	return user, nil
}

// RemoveByID implements the method to remove a user from the store
func (r *userRepository) RemoveByID(id int) error {

	_, err := r.db.Exec(`DELETE FROM users WHERE id = $1;`, id)
	return err
}

// UpdateByID implements the method to update a user into the store
func (r *userRepository) UpdateByID(id int, user model.UpdateUser) error {
	result, err := r.db.Exec("UPDATE users SET name = $1, cif = $2, postal_code = $3, country = $4 where id = $5", user.Name, user.Cif, user.PostalCode, user.Country, id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rows != 1 {
		return error2.ErrNotFound
	}

	return nil
}

// Create implements the method to persist a new user
func (r *userRepository) Create(UserSignUp model.CreateUser) (user *model.User, err error) {
	createUserQuery := `INSERT INTO users (name, cif, postal_code, country) 
		VALUES ($1, $2, $3, $4)
		RETURNING id`

	stmt, err := r.db.Prepare(createUserQuery)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var userID int
	err = stmt.QueryRow(UserSignUp.Name, UserSignUp.Cif, UserSignUp.PostalCode, UserSignUp.Country).Scan(&userID)
	if err != nil {
		return nil, err
	}

	return &model.User{
		ID:         userID,
		Name:       UserSignUp.Name,
		Cif:        UserSignUp.Cif,
		Country:    UserSignUp.Country,
		PostalCode: UserSignUp.PostalCode,
	}, nil
}
