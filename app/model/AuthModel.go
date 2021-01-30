package model

import (
	"gopkg.in/dgrijalva/jwt-go.v3"
)

type UserClaims struct {
	UserId int64 `json:"user_id"`
	jwt.StandardClaims
}

//TokenDetails ...
type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUUID   string
	RefreshUUID  string
	AtExpires    int64
	RtExpires    int64
}

//AccessDetails ...
type AccessDetails struct {
	AccessUUID string
	UserID     int64
}

//Token ...
type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
