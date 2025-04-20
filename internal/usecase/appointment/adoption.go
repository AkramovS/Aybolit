package appointment

import (
	"Aybolit/internal/domain/entity"
	"Aybolit/internal/domain/repository"
	"context"
)

type purposeAppointment struct {
	repo repository.AppointmentRepository
}

func NewPurposeDoctor(r repository.AppointmentRepository) AdoptionAppointmentUsecase {
	return &purposeAppointment{repo: r}
}

func (r *purposeAppointment) Execute(ctx context.Context, input AppointmentInput) (int64, error) {
	appointment := &entity.Appointment{
		PatientID: input.PatientID,
		DoctorID:  input.DoctorID,
		StartTime: input.StartTime,
		EndTime:   input.EndTime,
		Notes:     input.Notes,
		CreatedAt: input.CreatedAt,
		UpdatedAt: input.UpdatedAt,
		Canceled:  input.Canceled,
	}
	id, err := r.repo.Create(ctx, appointment)
	if err != nil {
		return 0, err
	}

	return id, nil
}
