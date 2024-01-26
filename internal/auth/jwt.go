package auth

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

const SecretKey = "dcbfa8b48a95ef606ddcc52e5e4ff7390375817b50858e852b1ec254064ab134"

func ValidateToken(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		return false, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if exp, ok := claims["exp"].(float64); ok {
			return time.Now().Unix() < int64(exp), nil
		}
	}

	return false, err
}

func GenerateToken(clientId string) (string, error) {
	expirationTime := time.Now().Add(2 * time.Hour)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"client_id": clientId,
		"exp":       expirationTime.Unix(),
	})

	tokenString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
