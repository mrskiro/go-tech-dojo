package characters

import (
	"context"
	"database/sql"

	"github.com/purp1eeeee/go-tech-dojo/domains/character"
)

type Repository interface {
	Get(ctx context.Context) (Characters, error)
}

type repository struct {
	db *sql.DB
}

func NewCharactersRepository(db *sql.DB) Repository {
	return repository{
		db: db,
	}
}

type Character struct {
	ID   uint64
	Name string
}

type Characters []Character

func (r repository) Get(ctx context.Context) (Characters, error) {
	query := "SELECT id, name FROM characters"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	characters := make(Characters, 0)
	for rows.Next() {
		character := Character{}
		err = rows.Scan(&character.ID, &character.Name)
		if err != nil {
			continue
		}
		characters = append(characters, character)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return characters, nil
}

func (cs Characters) ToDomain() character.Characters {
	results := make(character.Characters, len(cs))
	for i, v := range cs {
		results[i] = character.Character{
			ID:   v.ID,
			Name: v.Name,
		}
	}
	return results
}
