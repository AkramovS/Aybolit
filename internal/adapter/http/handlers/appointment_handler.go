package handlers

import (
	"Aybolit/internal/usecase/appointment"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AppointmentHandler struct {
	adoptionAppointmentUseCase appointment.AdoptionAppointmentUseCase
}

func NewAppointmentHandler(adoptionAppointmentUseCase appointment.AdoptionAppointmentUseCase) *AppointmentHandler {
	return &AppointmentHandler{adoptionAppointmentUseCase: adoptionAppointmentUseCase}

}

func (h *AppointmentHandler) Adoption(c *gin.Context) {
	var input appointment.AppointmentInput
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := h.adoptionAppointmentUseCase.Execute(c, input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not adoption"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "adoption added successfully"})
}
