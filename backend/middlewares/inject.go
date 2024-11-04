package middlewares

import (
	"context"
	"net/http"

	"github.com/ariefsn/gotik/constant"
)

func Inject() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Inject Writer
			ctx := context.WithValue(r.Context(), constant.WriterCtxKey, w)
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		})
	}
}
