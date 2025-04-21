package http

type Handlers struct {
	Patient     *PatientHandler
	Doctor      *DoctorHandler
	Appointment *AppointmentHandler
	//Appoinment handler
}

// Regenerate constructor
func NewHandlers(doctor *DoctorHandler, patient *PatientHandler) *Handlers {
	return &Handlers{Doctor: doctor, Patient: patient}
}
