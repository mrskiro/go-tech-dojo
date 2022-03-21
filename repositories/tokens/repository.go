package tokens

import (
	"context"
	"database/sql"
)

type Repository interface {
	Create(ctx context.Context, userID, token string) error
	FindUserIDByToken(ctx context.Context, token string) (string, error)
}

type repository struct {
	db *sql.DB
}

func NewTokenRepository(db *sql.DB) Repository {
	return repository{
		db: db,
	}
}

func (r repository) Create(ctx context.Context, userID, token string) error {
	query := "INSERT INTO tokens (user_id, token, created_at) VALUES ($1, $2, NOW())"
	_, err := r.db.ExecContext(ctx, query, userID, token)
	if err != nil {
		return err
	}
	return nil
}

func (r repository) FindUserIDByToken(ctx context.Context, token string) (string, error) {
	query := "SELECT user_id FROM tokens WHERE token = $1"
	prep, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return "", err
	}
	defer prep.Close()
	var userID string
	err = prep.QueryRow(token).Scan(&userID)
	if err != nil {
		return "", err
	}
	return userID, nil
}
