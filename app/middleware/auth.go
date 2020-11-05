package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/dgrijalva/jwt-go.v3"
	"net/http"
	"os"
	"strings"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		TokenValid(c.Request)
		c.Next()
	}
}

//TokenValid ...
func TokenValid(r *http.Request) error {
	token, err := VerifyToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

//VerifyToken ...
func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("APP_KEY")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

//ExtractToken ...
func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}
