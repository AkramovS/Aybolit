package repository

import "Aybolit/internal/domain/entity"

type DoctorRepository interface {
	Create(doctor *entity.Doctor) error
	GetByID(id int64) (*entity.Doctor, error)
}
