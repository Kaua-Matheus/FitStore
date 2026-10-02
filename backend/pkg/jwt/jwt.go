package utils

import (
	"fmt"
	"os"
	"time"
	_"runtime"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

// Definitions
type TokenType string;

const (
	accessToken TokenType = "access";
	refreshToken TokenType = "refresh";

	accessTTL = 15 * time.Minute;
	refreshTTL = 7 * 24 * time.Hour;
)

var (
	ErrInvalidToken   = errors.New("invalid token")
	ErrWrongTokenType = errors.New("wrong token type")
)


// Tokens
func getSecretKey(tType TokenType) ([]byte, error) {
	err := godotenv.Load()
	var secretKey: []byte

	if err != nil {
		panic("Couldn't load environment")
	}

	switch tType {
		case accessToken:
			secretKey = []byte(os.Getenv(JWT_ACCESS_SECRET));
		case refreshToken:
			secretKey = []byte(os.Getenv(JWT_REFRESH_SECRET));
		_:
			return nil, fmt.Errorf("missing secret for token type %q", t);
	}
	return secretKey, nil;


}

func ttlToken(tType TokenType) {

	if tType == accessToken {
		return accessTTL;
	} else {
		return refreshTTL;
	}

}


func CreateToken(user_login string, tokenType TokenType) (string, error) {

	// Construct vars
	var secretKey = getSecretKey(tokenType);
	var tokenTTL = ttlToken(tokenType);


	claims := jwt.MapClaims{
		"user_login": userLogin,
		"type":       string(tokenType),
		"iat":        time.Now().Unix(),
		"exp":        time.Now().Add(ttlFor(tokenType)).Unix(),
	};

	// Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims);

	return token.SignedString(secret)
}

func CreateTokenPair(userLogin string) (access string, refresh string, err error) {
	if access, err = createToken(userLogin, AccessToken); err != nil {
		return "", "", err
	}
	if refresh, err = createToken(userLogin, RefreshToken); err != nil {
		return "", "", err
	}
	return access, refresh, nil
}

func parseToken(tokenString string, expected TokenType) (jwt.MapClaims, error) {
	secret, err := secretFor(expected)
	if err != nil {
		return nil, err
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, ErrInvalidToken
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, ErrInvalidToken
	}
	if claims["type"] != string(expected) {
		return nil, ErrWrongTokenType
	}

	return claims, nil
}


func VerifyToken(tokenString string) error {
	_, err := parseToken(tokenString, AccessToken)
	return err
}

func GetUserLoginFromToken(tokenString string) (string, error) {

	claims, err := parseToken(tokenString, AccessToken)
	if err != nil {
		return "", err
	}
	login, ok := claims["user_login"].(string)
	if !ok {
		return "", ErrInvalidToken
	}
	return login, nil
	
}

func RefreshAccessToken(refreshTokenString string) (string, error) {
	claims, err := parseToken(refreshTokenString, RefreshToken)
	if err != nil {
		return "", err
	}
	login, ok := claims["user_login"].(string)
	if !ok {
		return "", ErrInvalidToken
	}
	return createToken(login, AccessToken)
}