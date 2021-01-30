package middleware

import (
	"ApiRest/app/model"
	"ApiRest/provider"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/dgrijalva/jwt-go.v3"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type authMiddleware struct {
	cache *provider.DbCache
}

type AuthMiddlewareInterface interface {
	Handler() gin.HandlerFunc
}

func NewAuthMiddleware(cache *provider.DbCache) AuthMiddlewareInterface {
	return &authMiddleware{
		cache,
	}
}

func (am authMiddleware) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		am.TokenValid(c)
		c.Next()
	}
}

func (am authMiddleware) TokenValid(c *gin.Context) {

	tokenAuth, err := am.ExtractTokenMetadata(c.Request)
	if err != nil {
		log.Print(err.Error())
		//Token either expired or not valid
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Please login first"})
		return
	}

	userID, err := am.FetchAuth(tokenAuth)
	if err != nil {
		//Token does not exists in Redis (User logged out or expired)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Please login first"})
		return
	}

	//To be called from GetUserID()
	c.Set("userID", userID)
}

//FetchAuth ...
func (am authMiddleware) FetchAuth(authD *model.AccessDetails) (int64, error) {
	userid, err := am.cache.Get(provider.REDIS_CTX, authD.AccessUUID).Result()
	if err != nil {
		return 0, err
	}
	userID, _ := strconv.ParseInt(userid, 10, 64)
	return userID, nil
}

//VerifyToken ...
func (am authMiddleware) VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := am.extractToken(r)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	fmt.Print("Token extracted => ", token)

	return token, nil
}

//ExtractToken ...
func (am authMiddleware) extractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func (am authMiddleware) ExtractTokenMetadata(r *http.Request) (AccessDetails *model.AccessDetails, err error) {
	token, err := am.VerifyToken(r)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUUID, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, err
		}
		userID, err := strconv.ParseInt(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			return nil, err
		}
		return &model.AccessDetails{
			AccessUUID: accessUUID,
			UserID:     userID,
		}, nil
	}
	return nil, err
}
