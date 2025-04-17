package patient

type RegisterPatientInput struct {
	FullName  string `json:"full_name"`
	Phone     string `json:"phone"`
	BirthDate string `json:"birth_date"`
	Notes     string `json:"notes"`
}

type RegisterPatientUseCase interface {
	Execute(input RegisterPatientInput) error
}
