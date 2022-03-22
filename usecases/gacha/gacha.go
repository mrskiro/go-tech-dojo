package gacha

import (
	"context"
	"math/rand"
	"sort"

	"github.com/purp1eeeee/go-tech-dojo/domains/character"
	"github.com/purp1eeeee/go-tech-dojo/repositories/character_probability"
	"github.com/purp1eeeee/go-tech-dojo/repositories/characters"
	"github.com/purp1eeeee/go-tech-dojo/repositories/tx"
	"github.com/purp1eeeee/go-tech-dojo/repositories/user_characters"
)

type UseCase interface {
	Drow(ctx context.Context, userID string, times int) (character.Characters, error)
}

type useCase struct {
	txRepo             tx.Repository
	probabilityRepo    character_probability.Repository
	charactersRepo     characters.Repository
	userCharactersRepo user_characters.Repository
}

func NewGachaUseCase(txRepo tx.Repository, probabilityRepo character_probability.Repository, charactersRepo characters.Repository, userCharactersRepo user_characters.Repository) UseCase {
	return useCase{
		txRepo:             txRepo,
		probabilityRepo:    probabilityRepo,
		charactersRepo:     charactersRepo,
		userCharactersRepo: userCharactersRepo,
	}
}

func (u useCase) Drow(ctx context.Context, userID string, times int) (character.Characters, error) {
	characters, err := u.charactersRepo.Get(ctx)
	if err != nil {
		return nil, err
	}
	probabilities, err := u.probabilityRepo.Get(ctx)
	if err != nil {
		return nil, err
	}

	charactersMap := characters.ToDomain().ToMap()

	weights := probabilities.ToDomain().CalculateOneHundred()

	targets := make(character.Characters, 0, len(weights))
	for _, v := range weights {
		c, ok := charactersMap[v.CharacterID]
		if !ok {
			continue
		}
		targets = append(targets, c)
	}

	boundaries := make([]int, len(weights)+1)
	for i := 1; i < len(boundaries); i++ {
		boundaries[i] = boundaries[i-1] + int(weights[i-1].Probability)
	}

	boundaryLast := boundaries[len(boundaries)-1]

	results := make(character.Characters, 0, times)
	for i := 0; i < times; i++ {
		x := rand.Intn(boundaryLast) + 1
		idx := sort.SearchInts(boundaries, x) - 1
		results = append(results, targets[idx])
	}

	err = u.txRepo.Start(ctx, func(ctx context.Context) error {
		for _, v := range results {
			err = u.userCharactersRepo.Create(ctx, userID, user_characters.Character{ID: v.ID})
			if err != nil {
				continue
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return results, nil
}
