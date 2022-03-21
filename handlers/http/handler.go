package http

import (
	"github.com/purp1eeeee/go-tech-dojo/oapi"
	"github.com/purp1eeeee/go-tech-dojo/usecases/user"
)

type Handlers struct {
	userUseCase user.UseCase
}

func NewHandlers(userUseCase user.UseCase) oapi.ServerInterface {
	return Handlers{
		userUseCase: userUseCase,
	}
}
