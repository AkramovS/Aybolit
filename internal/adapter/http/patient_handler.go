package http

import (
	"Aybolit/internal/domain/entity"
	"Aybolit/internal/usecase/patient"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type PatientHandler struct {
	registerPatientUseCase patient.RegisterPatientUseCase
	getterPatientUseCase   patient.GetterPatientsUseCase
	getAllPatientsUseCase  patient.GetAllPatientsUseCase
}

func NewPatientHandler(
	registerPatientUseCase patient.RegisterPatientUseCase,
	getterPatientUseCase patient.GetterPatientsUseCase,
	getAllPatientsUseCase patient.GetAllPatientsUseCase,
) *PatientHandler {
	return &PatientHandler{
		registerPatientUseCase: registerPatientUseCase,
		getterPatientUseCase:   getterPatientUseCase,
		getAllPatientsUseCase:  getAllPatientsUseCase,
	}
}

//Создание пациента

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

// Поиск пациента по его ID

// GetByID godoc
// @Summary Получить пациента по ID
// @Description Возвращает данные пациента по его идентификатору
// @Tags patients
// @Accept json
// @Produce json
// @Param id query int true "ID пациента"
// @Success 200 {object} entity.Patient
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/patients/patient [get]
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

// GetByFilters godoc
// @Summary Получить список пациентов по фильтрам
// @Description Возвращает список пациентов, отфильтрованных по заданным параметрам
// @Tags patients
// @Accept json
// @Produce json
// @Param first_name query string false "Имя пациента"
// @Param last_name query string false "Фамилия пациента"
// @Param phone query string false "Телефон пациента"
// @Success 200 {object} []entity.Patient
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/patients/list [get]
func (h *PatientHandler) GetByFilters(c *gin.Context) {
	var filter entity.PatientsQueryParams

	if err := c.ShouldBindQuery(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid filter"})
		return
	}

	patients, err := h.getAllPatientsUseCase.Execute(c, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Patients not found"})
		log.Println("could not get patients:", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"patients": patients})
}
