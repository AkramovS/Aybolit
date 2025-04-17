package patient

import (
	"Aybolit/internal/domain/entity"
	"Aybolit/internal/domain/repository"
)

type registerPatient struct {
	repo repository.PatientRepository
}

func NewRegisterPatient(r repository.PatientRepository) RegisterPatientUseCase {
	return &registerPatient{repo: r}
}

func (r *registerPatient) Execute(input RegisterPatientInput) error {
	patient := &entity.Patient{
		FullName:  input.FullName,
		Phone:     input.Phone,
		BirthDate: input.BirthDate,
		Notes:     input.Notes,
	}

	return r.repo.Create(patient)
}
