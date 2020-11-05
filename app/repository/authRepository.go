package repository

import (
	"ApiRest/app/model"
	"ApiRest/provider"
	"github.com/go-redis/redis/v8"
	"github.com/twinj/uuid"
	"gopkg.in/dgrijalva/jwt-go.v3"
	"os"
	"strconv"
	"time"
)

type authRepository struct {
	db *redis.Client
}

type AuthRepository interface {
	CreateToken(user model.User) (td *model.TokenDetails, err error)
	CreateAuth(user model.User, td *model.TokenDetails) error
}

func NewAuthRepository(db *redis.Client) AuthRepository {
	return &authRepository{
		db,
	}
}

func (a authRepository) CreateToken(user model.User) (td *model.TokenDetails, err error) {
	td.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	td.UUID = uuid.NewV4().String()

	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["access_uuid"] = td.UUID
	atClaims["user_id"] = user.ID
	atClaims["exp"] = td.AtExpires

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.Token, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	return
}

//CreateAuth ...
func (a authRepository) CreateAuth(user model.User, token *model.TokenDetails) error {
	at := time.Unix(token.AtExpires, 0) //converting Unix to UTC(to Time object)
	now := time.Now()
	errAccess := a.db.Set(provider.REDIS_CTX, token.UUID, strconv.Itoa(int(user.ID)), at.Sub(now)).Err()
	if errAccess != nil {
		return errAccess
	}

	return nil
}

//////////
