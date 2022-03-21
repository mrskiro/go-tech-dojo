package character_probability

import (
	"context"
	"database/sql"

	"github.com/purp1eeeee/go-tech-dojo/domains/character/probability"
)

type Repository interface {
	Get(ctx context.Context) (CharacterProbabilities, error)
}

type repository struct {
	db *sql.DB
}

func NewProbabilityRepository(db *sql.DB) Repository {
	return repository{
		db: db,
	}
}

type CharacterProbability struct {
	ID          uint64
	Probability uint64
	CharacterID uint64
}

type CharacterProbabilities []CharacterProbability

func (r repository) Get(ctx context.Context) (CharacterProbabilities, error) {
	query := "SELECT id, probability, character_id FROM character_probability WHERE deleted_at IS NULL"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	probabilities := make(CharacterProbabilities, 0)
	for rows.Next() {
		probability := CharacterProbability{}
		err = rows.Scan(&probability.ID, &probability.Probability, &probability.CharacterID)
		if err != nil {
			continue
		}
		probabilities = append(probabilities, probability)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return probabilities, nil
}

func (ps CharacterProbabilities) ToDomain() probability.CharacterProbabilities {
	results := make(probability.CharacterProbabilities, len(ps))
	for i, v := range ps {
		results[i] = probability.CharacterProbability{
			ID:          v.ID,
			Probability: v.Probability,
			CharacterID: v.CharacterID,
		}
	}
	return results
}
