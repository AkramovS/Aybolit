package postgres

import (
	"Aybolit/internal/domain/entity"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

type AppointmentRepo struct {
	db *pgxpool.Pool
}

func NewAppointmentRepo(db *pgxpool.Pool) *AppointmentRepo {
	return &AppointmentRepo{db: db}
}

func (r *AppointmentRepo) Create(ctx context.Context, a *entity.Appointment) (int64, error) {
	query := `INSERT INTO appointments(patien_id, doctor_id, start_time, end_time, notes, created_at, updated_at, canceled ) 
values ($1, $2, $3, $4, $5, $6, $7);`
	_, err := r.db.Exec(context.Background(), query, a.PatientID, a.DoctorID, a.StartTime, a.EndTime, a.Notes, a.CreatedAt, a.UpdatedAt, a.Canceled)
	if err != nil {
		return 0, err
	}
	return 0, nil
}

func (r *AppointmentRepo) GetById(ctx context.Context, id int64) (*entity.Appointment, error) {
	query := `SELECT id, patient_id, doctor_id, start_time, end_time, notes, created_at, updated_at FROM appointments WHERE id = $1;`
	row := r.db.QueryRow(context.Background(), query, id)
	p := entity.Appointment{}
	err := row.Scan(&p.ID, &p.PatientID, &p.DoctorID)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *AppointmentRepo) Update(ctx context.Context, a *entity.Appointment) error {
	query := `UPDATE  appointments  SET id, patient_id, doctor_id, updated_at  WHERE id = $1;`
	_, err := r.db.Exec(context.Background(), query, a.ID, a.PatientID, a.DoctorID, a.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (r *AppointmentRepo) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM appointments WHERE id = $1;`
	_, err := r.db.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *AppointmentRepo) ListByDoctor(ctx context.Context, doctorID int64) ([]*entity.Appointment, error) {
	query := `SELECT patient_id, doctor_id, start_time, end_time, notes, created_at, updated_at, canceled FROM appointments WHERE doctor_id = $1;`
	_, err := r.db.Query(context.Background(), query, doctorID)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *AppointmentRepo) ListByPatient(ctx context.Context, patientID int64) ([]*entity.Appointment, error) {
	query := `SELECT patient_id, doctor_id, created_at, updated_at, canceled  FROM appointments WHERE patient_id = $1;`
	_, err := r.db.Query(context.Background(), query, patientID)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
