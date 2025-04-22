package doctor

import (
	"Aybolit/internal/domain/entity"
	"Aybolit/internal/domain/repository"
)

type getterDoctor struct {
	repo repository.DoctorRepository
}

func NewGetterDoctor(r repository.DoctorRepository) GetterDoctorUseCase {
	return &getterDoctor{repo: r}
}

func (r *getterDoctor) Execute(id int64) (*entity.Doctor, error) {
	doctor, err := r.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return doctor, nil
}
