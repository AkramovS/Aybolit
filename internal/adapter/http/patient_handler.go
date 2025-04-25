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

//Поиск пациента по его имени

func (h *PatientHandler) GetByName(c *gin.Context) {
	nameParam := c.Query("first_name")
	if nameParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "First name is required"})
		return
	}
	_, err := strconv.ParseInt(nameParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid first_name"})
		log.Println("error=", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "patients found"})
}

//Поиск пациента по его фамилии

func (h *PatientHandler) GetByLastName(c *gin.Context) {
	lnameParam := c.Query("last_name")
	if lnameParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Last name is required"})
		return
	}
	_, err := strconv.ParseInt(lnameParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid last_name"})
		log.Println("error=", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "patients found"})
}

//Поиск пациента по его номеру телефона

func (h *PatientHandler) GetByPhoneNumber(c *gin.Context) {
	phoneNumberParam := c.Query("phone_number")
	if phoneNumberParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Phone number is required"})
		return
	}
	_, err := strconv.ParseInt(phoneNumberParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid phone_number"})
		log.Println("error=", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "patients found"})
}
