package utils

import (
	"go-starterkit-project/config"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func CreateToken(id string) (string, int64, error) {
	// Recreate token
	exp := time.Now().Add(time.Hour * 72).Unix()
	claims := jwt.MapClaims{
		"id":  id,
		"exp": exp,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, errToken := token.SignedString([]byte(config.Config("JWT_SECRET")))
	if errToken != nil {
		return "", 0, errToken
	}

	return t, exp, nil
}
