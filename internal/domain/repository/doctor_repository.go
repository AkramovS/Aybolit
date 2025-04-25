package repository

import (
	"Aybolit/internal/domain/entity"
	"context"
)

type DoctorRepository interface {
	Create(doctor *entity.Doctor) error
	GetByID(id int64) (*entity.Doctor, error)
	FindByFilters(ctx context.Context, filters entity.DoctorQueryParams) ([]*entity.Doctor, error)
}
