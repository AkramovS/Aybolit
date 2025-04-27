package handlers

type Handlers struct {
	Patient     *PatientHandler
	Doctor      *DoctorHandler
	Appointment *AppointmentHandler
	User        *UserHandler
}

func NewHandlers(
	patient *PatientHandler,
	appointment *AppointmentHandler,
	doctor *DoctorHandler,
	user *UserHandler,
) *Handlers {
	return &Handlers{Patient: patient, Appointment: appointment, Doctor: doctor, User: user}
}
