package http

import (
	"github.com/purp1eeeee/go-tech-dojo/oapi"
	"github.com/purp1eeeee/go-tech-dojo/usecases/character"
	"github.com/purp1eeeee/go-tech-dojo/usecases/gacha"
	"github.com/purp1eeeee/go-tech-dojo/usecases/user"
)

type Handlers struct {
	userUseCase      user.UseCase
	gachaUseCase     gacha.UseCase
	characterUseCase character.UseCase
}

func NewHandlers(userUseCase user.UseCase, gachaUseCase gacha.UseCase, characterUseCase character.UseCase) oapi.ServerInterface {
	return Handlers{
		userUseCase:      userUseCase,
		gachaUseCase:     gachaUseCase,
		characterUseCase: characterUseCase,
	}
}
