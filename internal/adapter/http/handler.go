package http

import (
	"Aybolit/internal/usecase/patient"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PatientHandler struct {
	useCase patient.RegisterPatientUseCase
}

func NewPatientHandler(u patient.RegisterPatientUseCase) *PatientHandler {
	return &PatientHandler{useCase: u}
}

func (h *PatientHandler) Register(c *gin.Context) {
	var input patient.RegisterPatientInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.useCase.Execute(input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not register patient"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "patient registered successfully"})
}
