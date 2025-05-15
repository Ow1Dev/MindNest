package app

import (
	"net/http"

	"github.com/Ow1Dev/MindNest/internal/middleware"
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"
	"github.com/rs/zerolog"
)

func NewServer(logger *zerolog.Logger) http.Handler {
	mux := http.NewServeMux()
	api := humago.New(mux, huma.DefaultConfig("My API", "1.0.0"))
	addRoutes(api, *logger)
	var handler http.Handler = mux
	handler = logging.NewLoggingMiddleware(logger, handler)
	return handler
}

