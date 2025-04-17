package main

import (
	"Aybolit/internal/adapter/http"
	pgrepo "Aybolit/internal/adapter/repository/postgres"
	"Aybolit/internal/infra/db"
	"Aybolit/internal/usecase/patient"
)

func main() {
	pool := db.NewPostgresConnection()
	patientRepo := pgrepo.NewPatientRepo(pool)
	useCase := patient.NewRegisterPatient(patientRepo)
	handler := http.NewPatientHandler(useCase)

	r := http.SetupRouter(handler)
	r.Run(":8080")
}

//// 1. Разобраться с конфигами
//// 2. инициализировать создание базы и миграции
//// 3. Добавление новой сущности ( например доктор)
//// 4. Добавить новые методы для новых сущностей
//// 5. Пообщаться с чат гпт
