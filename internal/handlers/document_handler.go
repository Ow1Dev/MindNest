package handlers

import (
	"context"
	"fmt"

	"github.com/rs/zerolog"
)

type GetActionsRequest struct {
	serviceName string `path:"serviceName" example:"echo" doc:"The name of the service"`
}

type GetActionsDto struct {
	Body struct {
		Message string `json:"message" example:"Hello, world!" doc:"Greeting message"`
	}
}

func MakeGetActionsHandler(logger *zerolog.Logger) func(ctx context.Context, input *GetActionsRequest) (*GetActionsDto, error) {
	return func(ctx context.Context, input *GetActionsRequest) (*GetActionsDto, error) {
		// Use thing here if needed
		message := fmt.Sprintf("Hello, %s!", input.serviceName)

		return &GetActionsDto{
			Body: struct {
				Message string `json:"message" example:"Hello, world!" doc:"Greeting message"`
			}{
				Message: message,
			},
		}, nil
	}
}
