package main

import (
	"Aybolit/internal/adapter/http"
	pgrepo "Aybolit/internal/adapter/repository/postgres"
	"Aybolit/internal/infra/db"
	"Aybolit/internal/usecase/doctor"
	"Aybolit/internal/usecase/patient"
)

func main() {
	pool := db.InitPostgres()

	//Регистрация репозитории
	patientRepo := pgrepo.NewPatientRepo(pool)
	doctorRepo := pgrepo.NewDoctorRepo(pool)
	//Здесь должна быть appoinmentrepo

	//Регистрация UseCase-ов
	//Пациент
	registerPatientUseCase := patient.NewRegisterPatient(patientRepo)
	getterPatientUseCase := patient.NewGetterPatient(patientRepo)
	//Доктор
	createDoctorUseCase := doctor.NewCreateDoctor(doctorRepo)
	getterDoctorUseCase := doctor.NewGetterDoctor(doctorRepo)
	//Записи

	//Регистрация Handler-ов
	patientHandler := http.NewPatientHandler(registerPatientUseCase, getterPatientUseCase)
	doctorHandler := http.NewDoctorHandler(createDoctorUseCase, getterDoctorUseCase)
	//// Здесь должна быть Appointmenthandler
	handler := http.NewHandlers(
		doctorHandler,
		patientHandler,
	)

	r := http.SetupRouter(handler)
	r.Run(":8080")
}
