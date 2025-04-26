package repository

import (
	"Aybolit/internal/domain/entity"
	"context"
)

type PatientRepository interface {
	Create(patient *entity.Patient) error
	GetByID(id int64) (*entity.Patient, error)
	FindByFilters(ctx context.Context, filters entity.PatientsQueryParams) ([]*entity.Patient, error)
}
