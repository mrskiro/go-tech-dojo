package ctxlib

import (
	"context"
	"errors"
)

type ContextKey struct {
	Name string
}

var (
	tokenCtxKey = &ContextKey{"Token"}
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
