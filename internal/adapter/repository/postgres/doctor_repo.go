package postgres

import (
	"Aybolit/internal/domain/entity"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

type DoctorRepo struct {
	db *pgxpool.Pool
}

func NewDoctorRepo(db *pgxpool.Pool) *DoctorRepo {
	return &DoctorRepo{db: db}
}

func (r *DoctorRepo) Create(p *entity.Doctor) error {
	query := `INSERT INTO doctors (first_name, last_name, activity, experience, phone, email, notes) 
				VALUES ($1, $2, $3, $4, $5, $6);`
	_, err := r.db.Exec(context.Background(), query, p.FirstName, p.LastName, p.Activity, p.Experience, p.Phone, p.Email, p.Notes)
	return err
}

func (r *DoctorRepo) GetByID(id int64) (*entity.Doctor, error) {
	query := `SELECT id, first_name, last_name, activity, experience, phone, email, notes FROM doctors WHERE id = $1`
	row := r.db.QueryRow(context.Background(), query, id)

	p := &entity.Doctor{}
	err := row.Scan(&p.ID, &p.FirstName, &p.LastName, &p.Activity, &p.Experience, &p.Phone, &p.Email, &p.Notes)
	if err != nil {
		return nil, err
	}

	return p, nil
}
