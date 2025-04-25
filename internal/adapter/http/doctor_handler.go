package http

import (
	"Aybolit/internal/domain/entity"
	"Aybolit/internal/usecase/doctor"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type DoctorHandler struct {
	createDoctorUseCase  doctor.CreateDoctorUseCase
	getterDoctorUseCase  doctor.GetterDoctorUseCase
	getAllDoctorsUseCase doctor.GetAllDoctorsUseCase
}

func NewDoctorHandler(
	createDoctorUseCase doctor.CreateDoctorUseCase,
	getterDoctorUseCase doctor.GetterDoctorUseCase,
	getAllDoctorsUseCase doctor.GetAllDoctorsUseCase,
) *DoctorHandler {
	return &DoctorHandler{
		createDoctorUseCase:  createDoctorUseCase,
		getterDoctorUseCase:  getterDoctorUseCase,
		getAllDoctorsUseCase: getAllDoctorsUseCase,
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

//Поиск доктора по его ID

func (h *DoctorHandler) GetById(c *gin.Context) {
	idParam := c.Query("id")
	if idParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is invalid"})
		log.Println("could not parse id: ", err)
		return
	}
	doctor, err := h.getterDoctorUseCase.Execute(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not get doctor"})
		log.Println("could not get doctor: ", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"doctor": doctor})
}

// Поиск доктора по его имени

func (h *DoctorHandler) GetByFilters(c *gin.Context) {
	var filter entity.DoctorQueryParams

	if err := c.ShouldBindQuery(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameters"})
		return
	}

	doctors, err := h.getAllDoctorsUseCase.Execute(c, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not get doctors"})
		log.Println("could not get doctors: ", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"doctors": doctors})
}
