package main

import (
	"Aybolit/internal/adapter/http"
	pgrepo "Aybolit/internal/adapter/repository/postgres"
	"Aybolit/internal/infra/db"
	"Aybolit/internal/usecase/patient"
)

func main() {
	pool := db.InitPostgres()
	patientRepo := pgrepo.NewPatientRepo(pool)
	registerUseCase := patient.NewRegisterPatient(patientRepo)
	getterUseCase := patient.NewGetterPatient(patientRepo)
	handler := http.NewPatientHandler(registerUseCase, getterUseCase)

	r := http.SetupRouter(handler)
	r.Run(":8080")
}
