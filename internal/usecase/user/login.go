package user

import (
	"Aybolit/internal/domain/repository"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

type loginUser struct {
	repo repository.UserRepository
}

func NewLoginUser(r repository.UserRepository) LoginUserUseCase {
	return &loginUser{repo: r}
}

func (u *loginUser) Execute(input LoginUserInput) (string, error) {
	//Find user by login
	user, err := u.repo.FindByLogin(input.Login)
	if err != nil {
		return "", err
	}
	//Compare password
	ok := comparePasswords(input.Password, []byte(user.Password))
	if ok == false {
		return "", errors.New("wrong password")
	}
	// Generic key and send
	token, err := generateJWT(user.Login, user.Role)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return token, nil
}

func comparePasswords(clientPass string, hashPass []byte) bool {
	err := bcrypt.CompareHashAndPassword(hashPass, []byte(clientPass))
	if err != nil {
		return false
	}
	return true
}

func generateJWT(login, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"login": login,
		"role":  role,
		"exp":   time.Now().Add(time.Hour * 1).Unix(), // 1 час жизни
	})

	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
