package entity

type Doctor struct {
	ID         int64
	FirstName  string
	LastName   string
	Activity   bool
	Experience int
	Phone      string
	Email      string
	Notes      string
}

type DoctorQueryParams struct {
	FirstName  string `form:"first_name"`
	LastName   string `form:"last_name"`
	Active     *bool  `form:"active"`
	Experience *int64 `form:"experience"`
	Phone      string `form:"phone"`
	Email      string `form:"email"`
}
