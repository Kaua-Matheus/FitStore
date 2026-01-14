package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	// Cria uma sequência hash com a senha que foi passada por parâmetro
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14);
	return string(hash), err;
}

func ComparePassword(password string, hash string) bool {
	// Compara uma senha string com o hash e retorna um boleano
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password));
	return err == nil;
}