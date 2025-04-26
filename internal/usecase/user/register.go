package user

import (
	"Aybolit/internal/domain/entity"
	"Aybolit/internal/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

type createUser struct {
	repo repository.UserRepository
}

func NewCreateUser(repo repository.UserRepository) RegisterUserUseCase {
	return &createUser{repo: repo}
}

func (u *createUser) Execute(input RegisterUserInput) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &entity.User{
		Email:    input.Email,
		Login:    input.Login,
		Password: string(hashedPassword),
		Role:     input.Role,
	}

	return u.repo.Create(user)
}
