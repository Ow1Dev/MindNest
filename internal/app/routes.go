package app

import (
	"github.com/Ow1Dev/MindNest/internal/handlers"
	"github.com/danielgtaylor/huma/v2"
	"github.com/rs/zerolog"
)

func addRoutes(api huma.API, logger zerolog.Logger) {
	huma.Get(api, "/getActions/{serviceName}", handlers.MakeGetActionsHandler(&logger))
}
