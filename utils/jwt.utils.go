package utils

import (
	"fmt"
	"test-gonic/config/app_config"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = app_config.SECRET_KEY

func GenerateToken(claims *jwt.MapClaims) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	webToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return webToken, nil
}

func VerifyToken(tokenString string) (*jwt.Token, error) {

	tokenJwt, errJwt := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		_, isValid := t.Method.(*jwt.SigningMethodHMAC)
		if !isValid {
			return nil, fmt.Errorf("unexpected singing method: %v", t.Header["alg"])
		}

		return []byte(secretKey), nil
	})

	if errJwt != nil {
		return nil, errJwt
	}

	return tokenJwt, nil
}

func DecodeToken(tokenString string) (jwt.MapClaims, error) {

	VerifyToken, err := VerifyToken(tokenString)

	if err != nil {
		return nil, err
	}

	claims, isOk := VerifyToken.Claims.(jwt.MapClaims)

	if isOk && VerifyToken.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
