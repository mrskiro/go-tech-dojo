package user

import (
	"context"

	"github.com/purp1eeeee/go-tech-dojo/domains/token"
	"github.com/purp1eeeee/go-tech-dojo/domains/user"
	"github.com/purp1eeeee/go-tech-dojo/repositories/tokens"
	"github.com/purp1eeeee/go-tech-dojo/repositories/tx"
	"github.com/purp1eeeee/go-tech-dojo/repositories/users"
)

type UseCase interface {
	Register(ctx context.Context, name string) (string, error)
	GetMe(ctx context.Context, token string) (user.User, error)
	Update(ctx context.Context, user users.User) error
}

type useCase struct {
	txRepo     tx.Repository
	tokensRepo tokens.Repository
	usersRepo  users.Repository
}

func NewUserUseCase(txRepo tx.Repository, tokensRepo tokens.Repository, usersRepo users.Repository) UseCase {
	return useCase{
		txRepo:     txRepo,
		tokensRepo: tokensRepo,
		usersRepo:  usersRepo,
	}
}

func (u useCase) Register(ctx context.Context, name string) (string, error) {
	userID := user.GenID()
	token, err := token.GenToken(10)
	if err != nil {
		return "", err
	}
	err = u.txRepo.Start(ctx, func(ctx context.Context) error {
		err := u.tokensRepo.Create(ctx, userID, token.String())
		if err != nil {
			return err
		}
		err = u.usersRepo.Create(ctx, users.User{ID: userID, Name: name})
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return "", err
	}
	return token.String(), nil
}

func (u useCase) GetMe(ctx context.Context, token string) (user.User, error) {
	userID, err := u.tokensRepo.FindUserIDByToken(ctx, token)
	if err != nil {
		return user.User{}, err
	}
	me, err := u.usersRepo.FindByID(ctx, userID)
	if err != nil {
		return user.User{}, err
	}

	return user.User{
		ID:   user.ID(me.ID),
		Name: me.Name,
	}, nil
}

func (u useCase) Update(ctx context.Context, user users.User) error {
	err := u.txRepo.Start(ctx, func(ctx context.Context) error {
		u.usersRepo.DeleteByID(ctx, user.ID)
		u.usersRepo.Create(ctx, users.User{ID: user.ID, Name: user.Name})
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
