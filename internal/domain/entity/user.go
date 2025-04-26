package entity

type User struct {
	ID       int64
	Email    string
	Password string
	Role     string // "admin", "doctor", "patient"
}
