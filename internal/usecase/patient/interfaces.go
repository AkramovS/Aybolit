package patient

import (
	"Aybolit/internal/domain/entity"
	"context"
)

type RegisterPatientInput struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	BirthDate string `json:"birth_date"`
	Notes     string `json:"notes"`
}

type RegisterPatientUseCase interface {
	Execute(input RegisterPatientInput) error
}

type GetterPatientsUseCase interface {
	Execute(id int64) (*entity.Patient, error)
}

type GetAllPatientsUseCase interface {
	Execute(ctx context.Context, filters entity.PatientsQueryParams) ([]*entity.Patient, error)
}
