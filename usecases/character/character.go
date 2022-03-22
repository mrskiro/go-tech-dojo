package character

import (
	"context"

	"github.com/purp1eeeee/go-tech-dojo/domains/character"
	"github.com/purp1eeeee/go-tech-dojo/repositories/characters"
	"github.com/purp1eeeee/go-tech-dojo/repositories/tx"
	"github.com/purp1eeeee/go-tech-dojo/repositories/user_characters"
)

type UseCase interface {
	GetList(ctx context.Context, userID string) (character.Characters, error)
}

type useCase struct {
	txRepo             tx.Repository
	charactersRepo     characters.Repository
	userCharactersRepo user_characters.Repository
}

func NewCharacterUseCase(txRepo tx.Repository, charactersRepo characters.Repository, userCharactersRepo user_characters.Repository) UseCase {
	return useCase{
		txRepo:             txRepo,
		charactersRepo:     charactersRepo,
		userCharactersRepo: userCharactersRepo,
	}
}

func (u useCase) GetList(ctx context.Context, userID string) (character.Characters, error) {
	ids, err := u.userCharactersRepo.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	characters := make(character.Characters, 0, len(ids))
	for _, id := range ids {
		c, err := u.charactersRepo.GetByID(ctx, id)
		if err != nil {
			continue
		}
		characters = append(characters, character.Character{ID: c.ID, Name: c.Name})
	}

	return characters, nil
}
