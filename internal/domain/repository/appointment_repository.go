package repository

import (
	"Aybolit/internal/domain/entity"
	"context"
)

type AppointmentRepository interface {
	Create(ctx context.Context, a *entity.Appointment) (int64, error)
	GetById(ctx context.Context, id int64) (*entity.Appointment, error)
	Update(ctx context.Context, a *entity.Appointment) error
	Delete(ctx context.Context, id int64) error
	ListByDoctor(ctx context.Context, doctorID int64) ([]*entity.Appointment, error)
	ListByPatient(ctx context.Context, patientID int64) ([]*entity.Appointment, error)
}
