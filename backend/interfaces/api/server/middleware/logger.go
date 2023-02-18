package middleware

import (
	"net/http"

	"github.com/yoshihiro-shu/draft-backend/internal/pkg/logger"
	"go.uber.org/zap"
)

func LoggerMiddleware(logger logger.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.Info("request info",
				zap.String("method", r.Method),
				zap.String("url", r.URL.Path))
			next.ServeHTTP(w, r)
		})
	}
}
