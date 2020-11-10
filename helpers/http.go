package helpers

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func HashAndSalt(pwd []byte) (string, error) {

	// Use GenerateFromPassword to hash & salt pwd
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash), nil
}

func ComparePasswords(password1 string, password2 string) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash1 := []byte(password1)
	byteHash2 := []byte(password2)
	err := bcrypt.CompareHashAndPassword(byteHash1, byteHash2)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
