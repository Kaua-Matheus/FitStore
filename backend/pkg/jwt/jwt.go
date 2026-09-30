package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func getEnv() string {
	err := godotenv.Load()
	if err != nil {
		panic("Couldn't load environment")
	} else {
		key := os.Getenv("JWT_SECRET")
		if key == "" {
			panic("JWT_SECRET environment variable is required")
		}
		return key
	}
}

func CreateToken(user_login string) (string, error) {

	var secretKey = []byte(getEnv())

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user_login": user_login,
			"exp":        time.Now().Add(time.Minute * 15).Unix(), // Timestamp to expire the jwt
		},
	)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) error {

	var secretKey = []byte(getEnv())

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}

func GetUserLoginFromToken(tokenString string) (string, error) {

	var secretKey = []byte(getEnv())

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return "", nil
	}

	if !token.Valid {
		return "", fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("invalid claims")
	}

	user_login, ok := claims["user_login"].(string)
	if !ok {
		return "", fmt.Errorf("user_login not found in token")
	}

	return user_login, nil
}
