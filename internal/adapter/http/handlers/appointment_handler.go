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

// Adoption godoc
// @Summary Create new appointment (Adoption)
// @Description Registers a new appointment between a patient and a doctor
// @Tags Appointments
// @Accept json
// @Produce json
// @Param appointment body appointment.AppointmentInput true "Appointment details"
// @Success 201 {object} map[string]string "adoption added successfully"
// @Failure 400 {object} map[string]string "bad request"
// @Failure 500 {object} map[string]string "could not adoption"
// @Security BearerAuth
// @Router /api/appointment [post]
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
