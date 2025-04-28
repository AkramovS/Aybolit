package postgres

import (
	"Aybolit/internal/domain/entity"
	"context"
	"fmt"
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
				VALUES ($1, $2, $3, $4, $5, $6, $7);`
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

func (r *DoctorRepo) FindByFilters(ctx context.Context, filters entity.DoctorQueryParams) ([]*entity.Doctor, error) {
	query := `SELECT id, first_name, last_name, activity, experience, phone, email, notes FROM doctors WHERE 1 = 1`
	args := []interface{}{}
	argIndex := 1

	if filters.FirstName != "" {
		query += fmt.Sprintf(" AND first_name = $%d", argIndex)
		args = append(args, filters.FirstName)
		argIndex++
	}

	if filters.LastName != "" {
		query += fmt.Sprintf(" AND last_name = $%d", argIndex)
		args = append(args, filters.LastName)
		argIndex++
	}

	if filters.Experience != nil {
		query += fmt.Sprintf(" AND experience = $%d", argIndex)
		args = append(args, filters.Experience)
		argIndex++
	}

	if filters.Active != nil {
		query += fmt.Sprintf(" AND activity = $%d", argIndex)
		args = append(args, filters.Active)
		argIndex++
	}

	if filters.Phone != "" {
		query += fmt.Sprintf(" AND phone = $%d", argIndex)
		args = append(args, filters.Phone)
		argIndex++
	}

	if filters.Email != "" {
		query += fmt.Sprintf(" AND email = $%d", argIndex)
		args = append(args, filters.Email)
	}

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var doctors []*entity.Doctor
	for rows.Next() {
		d := &entity.Doctor{}
		err := rows.Scan(&d.ID, &d.FirstName, &d.LastName, &d.Activity, &d.Experience, &d.Phone, &d.Email, &d.Notes)
		if err != nil {
			return nil, err
		}
		doctors = append(doctors, d)
	}

	return doctors, nil
}
