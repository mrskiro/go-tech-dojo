package middlewares

import (
	"net/http"

	"github.com/purp1eeeee/go-tech-dojo/ctxlib"
	"github.com/purp1eeeee/go-tech-dojo/handlers/http/utils"
	"github.com/purp1eeeee/go-tech-dojo/repositories/tokens"
)

type Middlewares interface {
	SetOperaterID(h http.Handler) http.Handler
	SetUserIDByToken(h http.Handler) http.Handler
}

type middlewares struct {
	tokensRepo tokens.Repository
}

func NewMiddlewares(tokensRepo tokens.Repository) Middlewares {
	return middlewares{
		tokensRepo: tokensRepo,
	}
}

func (m middlewares) SetOperaterID(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		token := r.Header.Get("x-token")

		ctx = ctxlib.SetContextWithToken(ctx, token)
		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func (m middlewares) SetUserIDByToken(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		token, err := ctxlib.GetTokenFromContext(ctx)
		if err != nil {
			h.ServeHTTP(w, r)
			return
		}
		if token == "" {
			h.ServeHTTP(w, r)
			return
		}

		userID, err := m.tokensRepo.FindUserIDByToken(ctx, token)
		if err != nil {
			utils.RenderInternalServerError(ctx, w, r)
			return
		}
		ctx = ctxlib.SetContextWithUserID(ctx, userID)
		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
