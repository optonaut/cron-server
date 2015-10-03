package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(secret string, personID UUID) (string, error) {
	signkey := []byte(secret)
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims["id"] = personID
	token.Claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	return token.SignedString(signkey)
}
