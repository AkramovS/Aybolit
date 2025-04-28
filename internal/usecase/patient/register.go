package patient

import (
	"Aybolit/internal/domain/entity"
	"Aybolit/internal/domain/repository"
	"log"
	"time"
)

type registerPatient struct {
	repo repository.PatientRepository
}

func NewRegisterPatient(r repository.PatientRepository) RegisterPatientUseCase {
	return &registerPatient{repo: r}
}

func (r *registerPatient) Execute(input RegisterPatientInput) error {
	layout := "2006-01-02"
	timeBirthDate, err := time.Parse(layout, input.BirthDate)
	if err != nil {
		log.Println("Error parsing time:", err)
		return err
	}

	patient := &entity.Patient{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Phone:     input.Phone,
		BirthDate: timeBirthDate,
		Notes:     input.Notes,
	}

	return r.repo.Create(patient)
}
