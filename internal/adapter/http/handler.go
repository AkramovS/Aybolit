package http

type Handlers struct {
	Patient     *PatientHandler
	Doctor      *DoctorHandler
	Appointment *AppointmentHandler
}

func NewHandlers(patient *PatientHandler, appointment *AppointmentHandler, doctor *DoctorHandler) *Handlers {
	return &Handlers{Patient: patient, Appointment: appointment, Doctor: doctor}
}
