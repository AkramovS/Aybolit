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
	query := `INSERT INTO patients (full_name, phone, birth_date, notes) 
				VALUES ($1, $2, $3, $4)`
	_, err := r.db.Exec(context.Background(), query, p.FullName, p.Phone, p.BirthDate, p.Notes)
	return err
}

func (r *PatientRepo) GetByID(id int64) (*entity.Patient, error) {
	query := `SELECT id, full_name, phone, birth_date, notes FROM patients WHERE id = $1`
	row := r.db.QueryRow(context.Background(), query, id)

	p := &entity.Patient{}
	err := row.Scan(&p.ID, &p.FullName, &p.Phone, &p.BirthDate, &p.Notes)
	if err != nil {
		return nil, err
	}

	return p, nil
}
