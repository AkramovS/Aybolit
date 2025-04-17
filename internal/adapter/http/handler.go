package http

import (
	"Aybolit/internal/usecase/patient"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type PatientHandler struct {
	registerPatientUseCase patient.RegisterPatientUseCase
	getterPatientUseCase   patient.GetterPatientUseCase
}

func NewPatientHandler(
	registerPatientUseCase patient.RegisterPatientUseCase,
	getterPatientUseCase patient.GetterPatientUseCase,
) *PatientHandler {
	return &PatientHandler{
		registerPatientUseCase: registerPatientUseCase,
		getterPatientUseCase:   getterPatientUseCase,
	}
}

func (h *PatientHandler) Register(c *gin.Context) {
	var input patient.RegisterPatientInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.registerPatientUseCase.Execute(input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not register patient"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "patient registered successfully"})
}

func (h *PatientHandler) GetByID(c *gin.Context) {
	idParam := c.Query("id")
	if idParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		log.Println("error=", err)
		return
	}

	patient, err := h.getterPatientUseCase.Execute(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		log.Println("error=", err)
		return
	}

	c.JSON(http.StatusOK, patient)
}
