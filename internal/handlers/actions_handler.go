package handlers

import (
	"net/http"

	"github.com/Ow1Dev/MindNest/pkg/httputils"
	"github.com/rs/zerolog"
)


func AddRoutes(
	mux *http.ServeMux,
	logger *zerolog.Logger,
) http.Handler {
	mux.Handle("POST /exec-actions/{service}/{action}", execActionsPost(logger))
	return mux
}

func execActionsPost(logger *zerolog.Logger) http.Handler {
	type Request struct {
		service string `query:"service"`
		action string `query:"action"`
	}

	type execActionsDto struct {
		Message string `json:"message"`
	}

	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			service := r.PathValue("service")
			action := r.PathValue("action")

			logger.Info().Msg("Received request for get-actions")
			result := execActionsDto{
				Message: "Execute '" + action + "' inside '" + service + "'!",
			}
			httputils.Encode(w, r, http.StatusOK, result)
		},
	)
}
