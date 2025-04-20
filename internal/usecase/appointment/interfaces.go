package appointment

import (
	"Aybolit/internal/domain/entity"
	"context"
	"time"
)

type AppointmentInput struct {
	PatientID string    `json:"patient_id"`
	DoctorID  string    `json:"doctor_id"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Notes     string    `json:"notes"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Canceled  bool      `json:"canceled"`
}
type AdoptionAppointmentUsecase interface {
	Execute(ctx context.Context, input AppointmentInput) (int64, error)
}

type GetterAppointmentUsecase interface {
	Execute(id int64) (*entity.Appointment, error)
}
