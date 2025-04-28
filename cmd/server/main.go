package main

import (
	"Aybolit/internal/adapter/http"
	"Aybolit/internal/adapter/http/handlers"
	pgrepo "Aybolit/internal/adapter/repository/postgres"
	"Aybolit/internal/infra/db"
	"Aybolit/internal/usecase/appointment"
	"Aybolit/internal/usecase/doctor"
	"Aybolit/internal/usecase/patient"
	"Aybolit/internal/usecase/user"
)

// @title Aybolit API
// @version 1.0
// @description API for Aybolit CRM
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	pool := db.InitPostgres()

	//Регистрация репозитории
	patientRepo := pgrepo.NewPatientRepo(pool)
	doctorRepo := pgrepo.NewDoctorRepo(pool)
	appointmentRepo := pgrepo.NewAppointmentRepo(pool)
	userRepo := pgrepo.NewUserRepo(pool)

	//Регистрация UseCase-ов
	//Пользователь
	registerUserUseCase := user.NewCreateUser(userRepo)
	loginUserUseCase := user.NewLoginUser(userRepo)
	//Пациент
	registerPatientUseCase := patient.NewRegisterPatient(patientRepo)
	getterPatientUseCase := patient.NewGetterPatient(patientRepo)
	getterAllPatientUseCase := patient.NewGetterByFilters(patientRepo)
	//Доктор
	createDoctorUseCase := doctor.NewCreateDoctor(doctorRepo)
	getterDoctorUseCase := doctor.NewGetterDoctor(doctorRepo)
	getterAllDoctorsUseCase := doctor.NewGetterByFilter(doctorRepo)
	//Записи
	adoptionAppointmentUseCase := appointment.NewPurposeDoctor(appointmentRepo)

	//Регистрация Handler-ов
	patientHandler := handlers.NewPatientHandler(registerPatientUseCase, getterPatientUseCase, getterAllPatientUseCase)
	doctorHandler := handlers.NewDoctorHandler(createDoctorUseCase, getterDoctorUseCase, getterAllDoctorsUseCase)
	appointmentHandler := handlers.NewAppointmentHandler(adoptionAppointmentUseCase)
	userHandler := handlers.NewUserHandler(registerUserUseCase, loginUserUseCase)

	handler := handlers.NewHandlers(
		patientHandler,
		appointmentHandler,
		doctorHandler,
		userHandler,
	)

	r := http.SetupRouter(handler)
	r.Run(":8080")

}
