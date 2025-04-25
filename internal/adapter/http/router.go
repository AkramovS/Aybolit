package http

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(handler *Handlers) *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.POST("/patients", handler.Patient.Register)
		api.GET("/patients/patient", handler.Patient.GetByID)

		api.POST("/doctors", handler.Doctor.Create)
		api.GET("/doctors/doctor", handler.Doctor.GetById)
		api.GET("/doctors/list", handler.Doctor.GetByFilters)

		api.POST("/appointment", handler.Appointment.Adoption)
	}

	return router
}
