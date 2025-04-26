package patient

import (
	"Aybolit/internal/domain/entity"
	"Aybolit/internal/domain/repository"
	"context"
)

type getterByFilter struct {
	repo repository.PatientRepository
}

func NewGetterByFilters(repo repository.PatientRepository) GetAllPatientsUseCase {
	return &getterByFilter{repo}
}

func (r *getterByFilter) Execute(ctx context.Context, filters entity.PatientsQueryParams) ([]*entity.Patient, error) {
	return r.repo.FindByFilters(ctx, filters)

}
