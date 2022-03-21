package users

import (
	"context"
	"database/sql"
)

type Repository interface {
	Create(ctx context.Context, user User) error
}

type repository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) Repository {
	return repository{
		db: db,
	}
}

type User struct {
	ID   string
	Name string
}

func (r repository) Create(ctx context.Context, user User) error {
	query := "INSERT INTO users (id, name, created_at, updated_at) VALUES ($1, $2, NOW(), NOW())"
	_, err := r.db.ExecContext(ctx, query, user.ID, user.Name)
	if err != nil {
		return err
	}
	return nil
}
