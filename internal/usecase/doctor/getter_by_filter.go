package doctor

import (
	"Aybolit/internal/domain/entity"
	"Aybolit/internal/domain/repository"
	"context"
)

type getterByFilter struct {
	repo repository.DoctorRepository
}

func NewGetterByFilter(repo repository.DoctorRepository) GetAllDoctorsUseCase {
	return &getterByFilter{repo}
}

func (r *getterByFilter) Execute(ctx context.Context, filters entity.DoctorQueryParams) ([]*entity.Doctor, error) {
	return r.repo.FindByFilters(ctx, filters)
}
