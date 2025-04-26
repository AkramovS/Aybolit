package main

import (
	"Aybolit/internal/adapter/http"
	pgrepo "Aybolit/internal/adapter/repository/postgres"
	"Aybolit/internal/infra/auth"
	"Aybolit/internal/infra/db"
	"Aybolit/internal/usecase/appointment"
	"Aybolit/internal/usecase/doctor"
	"Aybolit/internal/usecase/patient"
	"Aybolit/pkg/config"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	pool := db.InitPostgres()

	//Регистрация репозитории
	patientRepo := pgrepo.NewPatientRepo(pool)
	doctorRepo := pgrepo.NewDoctorRepo(pool)
	appointmentRepo := pgrepo.NewAppointmentRepo(pool)

	//Регистрация UseCase-ов
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
	patientHandler := http.NewPatientHandler(registerPatientUseCase, getterPatientUseCase, getterAllPatientUseCase)
	doctorHandler := http.NewDoctorHandler(createDoctorUseCase, getterDoctorUseCase, getterAllDoctorsUseCase)
	appointmentHandler := http.NewAppointmentHandler(adoptionAppointmentUseCase)

	handler := http.NewHandlers(
		patientHandler,
		appointmentHandler,
		doctorHandler,
	)

	r := http.SetupRouter(handler)
	r.Run(":8080")

	cfg := config.Load()

	// Инициализация JWT
	auth.InitJWT(cfg.JWTSecret)

	// Инициализация Gin сервера
	router := gin.Default()

	// Запуск
	addr := ":" + cfg.Port
	log.Printf("Server running on %s", addr)
	if err := router.Run(addr); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
