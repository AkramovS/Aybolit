package patient

import (
	"Aybolit/internal/domain/entity"
	"Aybolit/internal/domain/repository"
)

type getterPatient struct {
	repo repository.PatientRepository
}

func NewGetterPatient(r repository.PatientRepository) GetterPatientsUseCase {
	return &getterPatient{repo: r}
}

func (r *getterPatient) Execute(id int64) (*entity.Patient, error) {
	return r.repo.GetByID(id)
}
