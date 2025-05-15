package app

import (
	"net/http"

	"github.com/Ow1Dev/MindNest/internal/handlers"
	"github.com/Ow1Dev/MindNest/internal/middleware"
	"github.com/rs/zerolog"
)

func NewServer(logger *zerolog.Logger) http.Handler {
	mux := http.NewServeMux()
	handlers.AddRoutes(mux, logger) // Route directly to handler
	var handler http.Handler = mux
	handler = logging.NewLoggingMiddleware(logger, handler)
	return handler
}

