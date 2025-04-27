package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Claims struct {
	Login string `json:"login"`
	Role  string `json:"role"`
	jwt.RegisteredClaims
}

func ComparePasswords(clientPass string, hashPass []byte) bool {
	err := bcrypt.CompareHashAndPassword(hashPass, []byte(clientPass))
	if err != nil {
		return false
	}
	return true
}

func GenerateJWT(login string, role string) (string, error) {
	claims := Claims{
		Login: login,
		Role:  role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
