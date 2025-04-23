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
		api.GET("/doctors/doctor", handler.Doctor.Get)

		api.POST("/appointment", handler.Appointment.Adoption)
		api.GET("/doctors?full_name", handler.Doctor.Get)
	}

	return router
}

//TODO 1.нарисуй весь путь от начала до конца. (в Exceldraw)
//TODO 2.Сделать возможность получать список докторов и выбирают доктора.
