package ctxlib

import (
	"context"
	"errors"
)

type ContextKey struct {
	Name string
}

var (
	tokenCtxKey  = &ContextKey{"Token"}
	userIDCtxKet = &ContextKey{"UserID"}
)

func SetContextWithToken(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, tokenCtxKey, token)
}

func GetTokenFromContext(ctx context.Context) (string, error) {
	val := ctx.Value(tokenCtxKey)
	token, ok := val.(string)
	if !ok {
		return "", errors.New("not found token")
	}
	return token, nil
}

// WIP
func SetContextWithUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, userIDCtxKet, userID)
}

func GetUserIDFromContext(ctx context.Context) (string, error) {
	val := ctx.Value(tokenCtxKey)
	userID, ok := val.(string)
	if !ok {
		return "", errors.New("not found userID")
	}
	return userID, nil
}
