package repository

import (
	error2 "ApiRest/app/error"
	"ApiRest/app/model"
	"ApiRest/database"
	"github.com/go-redis/redis/v8"
	"github.com/twinj/uuid"
	"gopkg.in/dgrijalva/jwt-go.v3"
	"os"
	"strconv"
	"time"
)

type authRepository struct {
	redis *database.DbCache
}

//AuthRepositoryInterface ...
type AuthRepositoryInterface interface {
	CreateToken(user model.User) (td model.TokenDetails, err error)
	CreateAuth(user model.User, td model.TokenDetails) error
	GetAuth(AccessUUID string) (int64, error)
	DeleteAuth(AccessUUID string) error
}

//NewAuthRepository ...
func NewAuthRepository(db *database.DbCache) AuthRepositoryInterface {
	return &authRepository{
		db,
	}
}

//CreateToken ...
func (ar authRepository) CreateToken(user model.User) (td model.TokenDetails, err error) {
	td.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	td.AccessUUID = uuid.NewV4().String()

	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUUID = uuid.NewV4().String()

	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUUID
	atClaims["user_id"] = user.ID
	atClaims["exp"] = td.AtExpires

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	// A Util function to generate jwt_token which can be used in the request header
	td.AccessToken, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return model.TokenDetails{}, err
	}
	//Creating Refresh Token
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUUID
	rtClaims["user_id"] = user.ID
	rtClaims["exp"] = td.RtExpires

	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)

	td.RefreshToken, err = rt.SignedString([]byte(os.Getenv("REFRESH_SECRET")))

	if err != nil {
		return model.TokenDetails{}, err
	}

	return td, nil
}

//CreateAuth ...
func (ar authRepository) CreateAuth(user model.User, token model.TokenDetails) error {

	at := time.Unix(token.AtExpires, 0) //converting Unix to UTC(to Time object)
	rt := time.Unix(token.RtExpires, 0)
	now := time.Now()

	errAccess := ar.redis.Set(database.RedisCtx, token.AccessUUID, strconv.Itoa(int(user.ID)), at.Sub(now)).Err()
	if errAccess != nil {
		return errAccess
	}

	errRefresh := ar.redis.Set(database.RedisCtx, token.RefreshUUID, strconv.Itoa(int(user.ID)), rt.Sub(now)).Err()
	if errRefresh != nil {
		return errRefresh
	}

	return nil
}

//GetAuth ...
func (ar authRepository) GetAuth(AccessUUID string) (int64, error) {

	userid, err := ar.redis.Get(database.RedisCtx, AccessUUID).Result()

	if err != nil {
		return 0, err
	}

	userID, _ := strconv.ParseInt(userid, 10, 64)

	return userID, nil
}

//DeleteAuth ...
func (ar authRepository) DeleteAuth(AccessUUID string) error {

	_, err := ar.redis.Del(database.RedisCtx, AccessUUID).Result()

	if err != nil {
		if err == redis.Nil {
			return error2.NewErrorNotFound("Redis not found")
		}
		return err
	}

	return nil
}
