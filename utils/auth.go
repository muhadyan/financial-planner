package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/muhadyan/financial-planner/config"
	"github.com/muhadyan/financial-planner/model"
)

func GenerateJWT(claims model.JWTPayload) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(config.GetConfig().JWTSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
