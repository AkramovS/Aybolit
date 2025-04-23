package doctor

import (
	"Aybolit/internal/domain/entity"
	"Aybolit/internal/domain/repository"
)

type createDoctor struct {
	repo repository.DoctorRepository
}

func NewCreateDoctor(r repository.DoctorRepository) CreateDoctorUseCase {
	return &createDoctor{repo: r}
}

func (r *createDoctor) Execute(input CreateDoctorInput) error {
	doctor := &entity.Doctor{
		FirstName:  input.FirstName,
		LastName:   input.LastName,
		Activity:   input.Activity,
		Experience: input.Experience,
		Phone:      input.Phone,
		Email:      input.Email,
		Notes:      input.Notes,
	}

	return r.repo.Create(doctor)
}
