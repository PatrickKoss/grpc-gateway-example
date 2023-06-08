package middleware

import (
	"net/http"
	"time"

	"github.com/patrickkoss/grpc-gateway-example/internal/adapter/logging"
)

// responseWriter is a minimal wrapper for http.ResponseWriter that allows the
// written HTTP status code to be captured for logging.
type responseWriter struct {
	http.ResponseWriter
	status      int
	wroteHeader bool
}

func wrapResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{ResponseWriter: w}
}

func (rw *responseWriter) Status() int {
	return rw.status
}

func (rw *responseWriter) WriteHeader(code int) {
	if rw.wroteHeader {
		return
	}

	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
	rw.wroteHeader = true
}

// LoggingMiddleware logs the incoming HTTP request & its duration.
func LoggingMiddleware(logger logging.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					logger.Error(
						"failed to handle request logging",
					)
				}
			}()

			start := time.Now()
			wrapped := wrapResponseWriter(w)

			next.ServeHTTP(wrapped, r)

			status := wrapped.status
			if status == 0 {
				status = http.StatusOK
			}

			logger.Info(
				"request handled",
				logging.Int("status", status),
				logging.String("method", r.Method),
				logging.String("path", r.URL.EscapedPath()),
				logging.Float64("duration", time.Since(start).Seconds()),
			)

			wrapped.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
			wrapped.Header().Set("Content-Type", "application/json")
		}

		return http.HandlerFunc(fn)
	}
}
