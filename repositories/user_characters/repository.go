package user_characters

import (
	"context"
	"database/sql"
)

type Repository interface {
	Create(ctx context.Context, userID string, character Character) error
	GetByUserID(ctx context.Context, userID string) ([]uint64, error)
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
	ID     uint64
	UserID string
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

func (r repository) GetByUserID(ctx context.Context, userID string) ([]uint64, error) {
	query := "SELECT character_id FROM user_characters WHERE user_id = $1 AND deleted_at IS NULL"
	prep, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer prep.Close()
	rows, err := prep.QueryContext(ctx, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	characterIDs := make([]uint64, 0)
	for rows.Next() {
		var characterID uint64
		err = rows.Scan(&characterID)
		if err != nil {
			continue
		}
		characterIDs = append(characterIDs, characterID)
	}

	return characterIDs, nil
}
