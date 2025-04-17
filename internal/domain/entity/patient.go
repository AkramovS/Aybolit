package entity

import "time"

type Patient struct {
	ID        int64
	FullName  string
	Phone     string
	BirthDate time.Time
	Notes     string
}
