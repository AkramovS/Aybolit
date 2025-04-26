package entity

type User struct {
	ID       int64
	Email    string
	Login    string
	Password string
	Role     string // "admin", "doctor", "patient"
}
