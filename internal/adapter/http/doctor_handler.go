package http

import (
	"Aybolit/internal/usecase/doctor"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DoctorHandler struct {
	createDoctorUseCase doctor.CreateDoctorUseCase
	getterDoctorUseCase doctor.GetterDoctorUseCase
}

func NewDoctorHandler(
	createDoctorUseCase doctor.CreateDoctorUseCase,
	getterDoctorUseCase doctor.GetterDoctorUseCase,
) *DoctorHandler {
	return &DoctorHandler{
		createDoctorUseCase: createDoctorUseCase,
		getterDoctorUseCase: getterDoctorUseCase,
	}
}

func (h *DoctorHandler) Create(c *gin.Context) {
	var input doctor.CreateDoctorInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.createDoctorUseCase.Execute(input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create doctor"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "doctor created successfully"})
}

//TODO: имплементировать GET by ID
