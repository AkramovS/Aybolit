package repository

import "Aybolit/internal/domain/entity"

type UserRepository interface {
	Create(user *entity.User) error
	FindByLogin(login string) (*entity.User, error)
}
