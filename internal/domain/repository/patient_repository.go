package repository

import "Aybolit/internal/domain/entity"

type PatientRepository interface {
	Create(patient *entity.Patient) error
	GetByID(id int64) (*entity.Patient, error)
}
