package repository

import (
	"ApiRest/app/model"
	"ApiRest/internal/storage"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func TestUserRepositoryInit(t *testing.T) {
	type args struct {
		db *storage.DbStore
	}
	tests := []struct {
		name string
		args args
		want UserRepositoryInterface
	}{
		{
			name: "success",
			args: args{
				db: nil,
			},
			want: &userRepository{
				db: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserRepository(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserInit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserRepository_FindByID(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "db")
	defer sqlxDB.Close()

	userPGRepository := NewUserRepository(&storage.DbStore{DB: sqlxDB})

	columns := []string{"id", "email", "name", "postal_code", "phone", "last_name", "country"}
	userID := int(1)
	mockUser := &model.User{
		ID:         userID,
		Name:       "FirstName",
		LastName:   "LastName",
		Email:      "email@gmail.com",
		Country:    "es",
		Phone:      "es",
		PostalCode: "es",
	}

	rows := sqlmock.NewRows(columns).AddRow(
		userID,
		mockUser.Email,
		mockUser.Name,
		mockUser.PostalCode,
		mockUser.Phone,
		mockUser.LastName,
		mockUser.Country,
	)

	mock.ExpectQuery("SELECT id, email, name, postal_code, phone, last_name, country FROM users WHERE id = $1").WithArgs(userID).WillReturnRows(rows)

	foundUser, err := userPGRepository.FindById(mockUser.ID)

	require.NoError(t, err)
	require.NotNil(t, foundUser)
	require.Equal(t, foundUser.ID, userID)
}

func TestUserRepository_FindByID_IncorrectID(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "db")
	defer sqlxDB.Close()

	userPGRepository := NewUserRepository(&storage.DbStore{DB: sqlxDB})

	columns := []string{"id", "email", "name", "postal_code", "phone", "last_name", "country"}
	userID := int(1)
	mockUser := &model.User{
		ID:         userID,
		Name:       "FirstName",
		LastName:   "LastName",
		Email:      "email@gmail.com",
		Country:    "es",
		Phone:      "es",
		PostalCode: "es",
	}

	rows := sqlmock.NewRows(columns).AddRow(
		userID,
		mockUser.Email,
		mockUser.Name,
		mockUser.PostalCode,
		mockUser.Phone,
		mockUser.LastName,
		mockUser.Country,
	)

	mock.ExpectQuery("SELECT id, email, name, postal_code, phone, last_name, country FROM users WHERE id = $1").WithArgs(2).WillReturnRows(rows)

	foundUser, err := userPGRepository.FindById(mockUser.ID)

	require.Error(t, err)
	require.Nil(t, foundUser)
}

func TestUserRepository_Create(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "db")
	defer sqlxDB.Close()

	userPGRepository := NewUserRepository(&storage.DbStore{DB: sqlxDB})

	userID := int(1)
	mockUser := model.CreateUser{
		Name:       "FirstName",
		LastName:   "LastName",
		Email:      "email@gmail.com",
		Country:    "es",
		Phone:      "es",
		PostalCode: "es",
	}

	ep := mock.ExpectPrepare("INSERT INTO users (name, email, last_name, phone, postal_code, country) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id").WillBeClosed()
	ep.ExpectQuery().WithArgs(mockUser.Name, mockUser.Email, mockUser.LastName, mockUser.Phone, mockUser.PostalCode, mockUser.Country).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(userID))

	foundUser, err := userPGRepository.Create(mockUser)
	require.NoError(t, err)
	require.NotNil(t, foundUser)
	require.Equal(t, foundUser.ID, userID)
}
