package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

const SecretKey = "+as-0df80asg86das98fg8s9d-0" //私钥

// CreateToken 生成token
func CreateToken(username string) (string, error) {
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Minute * 3600 * 2400).Unix(),
	})
	token, err := at.SignedString([]byte(SecretKey))
	if err != nil {
		return "", err
	}
	return token, nil
}

// ParseToken 解析token
func ParseToken(token string) (string, error) {
	claim, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		return "", err
	}
	return claim.Claims.(jwt.MapClaims)["username"].(string), nil
}
