package users

import (
	"context"
	"database/sql"
)

type Repository interface {
	Create(ctx context.Context, user User) error
	FindByID(ctx context.Context, userID string) (User, error)
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

func (r repository) FindByID(ctx context.Context, userID string) (User, error) {
	u := User{}
	query := "SELECT id, name FROM users WHERE id = $1"
	prep, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return User{}, err
	}
	defer prep.Close()
	err = prep.QueryRow(userID).Scan(&u.ID, &u.Name)
	if err != nil {
		return User{}, err
	}
	return u, nil
}
