package http

import (
	_ "Aybolit/docs"
	"Aybolit/internal/adapter/http/handlers"
	"Aybolit/internal/adapter/http/middleware"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Aybolit API
// @version 1.0
// @description API for Aybolit CRM
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func SetupRouter(handler *handlers.Handlers) *gin.Engine {
	router := gin.Default()
	router.Use(middleware.LoggerMiddleware())
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.POST("/register", handler.User.Register)
	router.POST("/login", handler.User.Login)

	api := router.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
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
