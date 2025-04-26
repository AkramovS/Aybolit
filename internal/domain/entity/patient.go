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

type PatientsQueryParams struct {
	FirstName string `form:"first_name"`
	LastName  string `form:"last_name"`
	Phone     string `form:"phone"`
}
