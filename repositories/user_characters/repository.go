package user_characters

import (
	"context"
	"database/sql"
)

type Repository interface {
	Create(ctx context.Context, userID string, character Character) error
}

type repository struct {
	db *sql.DB
}

func NewUserCharactersRepository(db *sql.DB) Repository {
	return repository{
		db: db,
	}
}

type Character struct {
	ID   uint64
	Name string
}

func (r repository) Create(ctx context.Context, userID string, character Character) error {
	query := "INSERT INTO user_characters (character_id, user_id, created_at, updated_at) VALUES ($1, $2, NOW(), NOW())"
	prep, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer prep.Close()

	err = prep.QueryRow(character.ID, userID).Err()
	if err != nil {
		return err
	}
	return nil
}
