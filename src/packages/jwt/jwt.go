package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"time"
)

type MyCustomClaims struct {
	UserId int `json:"id"`
	jwt.StandardClaims
}

func GenerateToken(secretKey string, id int, timeExpressIn int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyCustomClaims{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Second * time.Duration(timeExpressIn)).Unix(),
			IssuedAt:  time.Now().Local().Unix(),
		},
	})

	key := []byte(secretKey)

	myToken, err := token.SignedString(key)

	if err != nil {
		return "", err
	}

	return myToken, nil

}
func ValidateToken(secretKey string, token string) (int, error) {
	res, err := jwt.ParseWithClaims(token, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return -1, err
	}

	if !res.Valid {
		return -1, err
	}

	claims, ok := res.Claims.(*MyCustomClaims)

	if !ok {
		return -1, errors.New("Payload not valid")
	}

	return claims.UserId, nil

}
