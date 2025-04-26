package http

import (
	"Aybolit/internal/usecase/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
	createUserUseCase user.RegisterUserUseCase
	loginUserUseCase  user.LoginUserUseCase
}

func NewUserHandler(
	createUserUseCase user.RegisterUserUseCase,
	loginUserUseCase user.LoginUserUseCase,
) *UserHandler {
	return &UserHandler{
		createUserUseCase: createUserUseCase,
		loginUserUseCase:  loginUserUseCase,
	}
}

// Register
// @Summary      Регистрация пользователя
// @Description  Создает нового пользователя с указанной ролью (admin, doctor, patient)
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        input body user.RegisterUserInput true "Данные для регистрации"
// @Success      201 {object} map[string]string "Пользователь успешно создан"
// @Failure      400 {object} map[string]string "Неверный формат данных"
// @Failure      500 {object} map[string]string "Ошибка сервера при создании пользователя"
// @Router       /api/register [post]
func (h *UserHandler) Register(c *gin.Context) {
	var input user.RegisterUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.createUserUseCase.Execute(input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully!"})
}

// Login godoc
// @Summary      Авторизация пользователя
// @Description  Вход по логину и паролю. Возвращает JWT токен при успешной авторизации.
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        input body user.LoginUserInput true "Данные для входа"
// @Success      200 {object} map[string]string "Пример: {\"token\": \"your.jwt.token\"}"
// @Failure      400 {object} map[string]string "Ошибка валидации"
// @Failure      500 {object} map[string]string "Внутренняя ошибка"
// @Router       /api/login [post]
func (h *UserHandler) Login(c *gin.Context) {
	var input user.LoginUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.loginUserUseCase.Execute(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
