package http

import "github.com/gin-gonic/gin"

func SetupRouter(patientHandler *PatientHandler) *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.POST("/patients", patientHandler.Register)
	}

	return router
}
