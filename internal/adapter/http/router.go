package http

import (
	_ "Aybolit/docs"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Aybolit API
// @version 1.0
// @description API for Aybolit CRM
// @host localhost:8080
// @BasePath /

func SetupRouter(handler *Handlers) *gin.Engine {
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api")
	{

		api.GET("/protected", AuthMiddleware("admin", "doctor"), AuthMiddleware("admin", "doctor"), func(c *gin.Context) {
			userID := c.GetInt64("userID")
			role := c.GetString("role")
			c.JSON(200, gin.H{
				"message": fmt.Sprintf("User ID: %d, Role: %s", userID, role),
			})
		})

		api.POST("/patients", handler.Patient.Register)
		api.GET("/patients/patient", handler.Patient.GetByID)
		api.GET("/patients/list", handler.Patient.GetByFilters)

		api.POST("/doctors", handler.Doctor.Create)
		api.GET("/doctors/doctor", handler.Doctor.GetById)
		api.GET("/doctors/list", handler.Doctor.GetByFilters)
		//TODO Disactive doctor by id
		//TODO Getappointments by filter
		//TODO get appointment by id
		//TODO get patient by filters
		//TODO Update appointments

		api.POST("/appointment", handler.Appointment.Adoption)
	}

	return router
}
