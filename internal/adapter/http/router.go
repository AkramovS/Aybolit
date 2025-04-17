package http

import "github.com/gin-gonic/gin"

func SetupRouter(patientHandler *PatientHandler) *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.POST("/patients", patientHandler.Register)
		api.GET("/patients/patient", patientHandler.GetByID)
		//TODO : implement route for get patient by id
		// 1. Опеределить какой тип запроса должен быть.
		// 2. Где в запросе должно отправлятться ID.
	}

	return router
}
