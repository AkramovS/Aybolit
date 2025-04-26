package user

type RegisterUserInput struct {
	Email    string `json:"email"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type LoginUserInput struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type RegisterUserUseCase interface {
	Execute(input RegisterUserInput) error
}

type LoginUserUseCase interface {
	Execute(input LoginUserInput) (string, error)
}
