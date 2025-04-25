package doctor

import (
	"Aybolit/internal/domain/entity"
	"context"
)

type CreateDoctorInput struct {
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Activity   bool   `json:"activity"`
	Experience int    `json:"experience"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	Notes      string `json:"notes"`
}

type CreateDoctorUseCase interface {
	Execute(input CreateDoctorInput) error
}

type GetterDoctorUseCase interface {
	Execute(id int64) (*entity.Doctor, error)
}

type GetAllDoctorsUseCase interface {
	Execute(ctx context.Context, filters entity.DoctorQueryParams) ([]*entity.Doctor, error)
}
