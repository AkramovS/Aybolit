package entity

import "time"

type Appointment struct {
	ID        int64
	PatientID string
	DoctorID  string
	StartTime time.Time
	EndTime   time.Time
	Notes     string
	CreatedAt time.Time
	UpdatedAt time.Time
	Canceled  bool
}

func (a Appointment) Error() string {
	//TODO implement me
	panic("implement me")
}
