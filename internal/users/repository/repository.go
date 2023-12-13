package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/martinyonatann/go-unit-test/internal/users"
	"github.com/martinyonatann/go-unit-test/internal/users/entities"
)

const (
	QueryCreateUser = `INSERT INTO users (id, name, password, created_at, updated_at) VALUES (:id, :name, :password, :created_at, :updated_at)`
	QueryDetailUser = `SELECT id, name, password, created_at, updated_at FROM users WHERE id = ?;`
)

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) users.Repositories {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, request entities.Users) error {
	_, err := r.db.NamedExecContext(ctx, QueryCreateUser, request)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) Detail(ctx context.Context, userUUID string) (response entities.Users, err error) {
	if err := r.db.GetContext(ctx, &response, QueryDetailUser, userUUID); err != nil {
		return response, err
	}

	return response, nil
}
