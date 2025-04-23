package entity

import "time"

type Patient struct {
	ID        int64
	FirstName string
	LastName  string
	Phone     string
	BirthDate time.Time
	Notes     string
}
