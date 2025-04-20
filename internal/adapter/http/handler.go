package http

type Handlers struct {
	Patient *PatientHandler
	Doctor  *DoctorHandler
}

func NewHandlers(doctor *DoctorHandler, patient *PatientHandler) *Handlers {
	return &Handlers{Doctor: doctor, Patient: patient}
}
