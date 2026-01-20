package utils

import (
	"fmt"
	"time"
	"os"

	"github.com/joho/godotenv"
	"github.com/golang-jwt/jwt/v5"
)

func getEnv() (string) {
	err := godotenv.Load(); if err != nil {
		panic("Couldn't load environment")
	} else {
		key := os.Getenv("JWT_SECRET")
		if key == "" {
			panic("JWT_SECRET environment variable is required")
		}
		return key;
	}
}

var secretKey = []byte(getEnv())

func CreateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(), // Pode ser alterado para diminuir ou aumentar o tempo do token
		},
	)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) error {
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
