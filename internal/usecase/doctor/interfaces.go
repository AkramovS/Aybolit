package doctor

import "Aybolit/internal/domain/entity"

type CreateDoctorInput struct {
	Name       string `json:"name"`
	Activity   string `json:"activity"`
	Experience string `json:"experience"`
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
