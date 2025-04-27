package user

import (
	"Aybolit/internal/domain/repository"
	local_jwt "Aybolit/pkg/jwt"
	"errors"
	"log"
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
	ok := local_jwt.ComparePasswords(input.Password, []byte(user.Password))
	if ok == false {
		return "", errors.New("wrong password")
	}
	// Generic key and send
	token, err := local_jwt.GenerateJWT(user.Login, user.Role)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return token, nil
}
