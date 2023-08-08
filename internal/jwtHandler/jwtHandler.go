package jwthandler

import (
	"fmt"
	"os"
	"smart-recommendation/internal/errorhandler"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JwtHandler struct {
	Secret string
}

func GetJwt() *JwtHandler {
	return &JwtHandler{}
}

func (obj *JwtHandler) GetToken(lifepan time.Duration) (token string, err errorhandler.ErrorData) {
	exp := time.Now().Add(lifepan)
	claims := jwt.MapClaims{
		"exp": exp.Unix(),
		"iat": time.Now().Unix(),
		"nbf": time.Now().Unix(),
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, errData := jwtToken.SignedString([]byte(obj.Secret))
	if errData != nil {
		err = errorhandler.FailedGetToken(errData)
		return
	}

	return
}

func (obj *JwtHandler) ValidationToken(token string) (isValid bool, claim jwt.MapClaims, err errorhandler.ErrorData) {
	jwttoken, errData := jwt.Parse(token, func(jwttoken *jwt.Token) (interface{}, error) {
		if _, ok := jwttoken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}
		obj.Secret = os.Getenv("JWT_SECRET")
		return []byte(obj.Secret), nil
	})

	if errData != nil {
		err = errorhandler.FailedvalidationToken(fmt.Errorf("Token is Invalid"))
		return
	}

	if !jwttoken.Valid {
		err = errorhandler.FailedvalidationToken(fmt.Errorf("Token is Invalid"))
		return
	}

	claim, ok := jwttoken.Claims.(jwt.MapClaims)
	if !ok {
		err = errorhandler.FailedvalidationToken(fmt.Errorf("cannot assert jwt payload to MapClaims"))
		return
	}

	isValid = true
	return

}
