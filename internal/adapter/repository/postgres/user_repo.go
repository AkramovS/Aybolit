package postgres

import (
	"Aybolit/internal/domain/entity"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

type UserRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) Create(u *entity.User) error {
	query := `INSERT INTO users ( email, login, password, role)
          VALUES ( $1, $2, $3 ,$4)`
	_, err := r.db.Exec(context.Background(), query, u.Email, u.Login, u.Password, u.Role)
	return err
}

func (r *UserRepo) FindByLogin(login string) (*entity.User, error) {
	query := `SELECT email, login, password, role FROM users WHERE login = $1`
	row := r.db.QueryRow(context.Background(), query, login)
	u := entity.User{}
	err := row.Scan(&u.Email, &u.Login, &u.Password, &u.Role)
	if err != nil {
		return nil, err
	}

	return &u, nil
}
