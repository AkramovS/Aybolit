package postgres

import (
	"Aybolit/internal/domain/entity"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

type PatientRepo struct {
	db *pgxpool.Pool
}

func NewPatientRepo(db *pgxpool.Pool) *PatientRepo {
	return &PatientRepo{db: db}
}

func (r *PatientRepo) Create(p *entity.Patient) error {
	query := `INSERT INTO patients (first_name, last_name, phone, birth_date, notes) 
				VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.Exec(context.Background(), query, p.FirstName, p.LastName, p.Phone, p.BirthDate, p.Notes)
	return err
}

func (r *PatientRepo) GetByID(id int64) (*entity.Patient, error) {
	query := `SELECT id, first_name, last_name, phone, birth_date, notes FROM patients WHERE id = $1`
	row := r.db.QueryRow(context.Background(), query, id)

	p := &entity.Patient{}
	err := row.Scan(&p.ID, &p.FirstName, &p.LastName, &p.Phone, &p.BirthDate, &p.Notes)
	if err != nil {
		return nil, err
	}

	return p, nil
}
