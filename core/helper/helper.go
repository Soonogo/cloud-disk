package helper

import (
	"cloud-disk/core/define"
	"crypto/md5"
	"fmt"

	"github.com/golang-jwt/jwt/v4"
)

func Md5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

func GenerateToken(id int, identity string, name string) (string, error) {
	uc := define.UserClaim{
		Id:       id,
		Identity: identity,
		Name:     name,
	}
	// false
	// token := jwt.NewWithClaims(jwt.SigningMethodES256, uc)
	// yse
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	fmt.Println(token, "2")

	tokenString, err := token.SignedString([]byte(define.JwtKey))

	if err != nil {
		return "", err

	}

	fmt.Println(tokenString, "11")
	return tokenString, nil
}
