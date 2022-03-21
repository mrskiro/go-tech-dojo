package middlewares

import (
	"net/http"

	"github.com/purp1eeeee/go-tech-dojo/ctxlib"
)

type Middlewares interface {
	SetOperaterID(h http.Handler) http.Handler
}

type middlewares struct{}

func NewMiddlewares() Middlewares {
	return middlewares{}
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
