package ai

import (
	"context"
	"main/ai/models"
)

type AgenticInterface interface {
	GenerateResponse(ctx context.Context, prompt string) (*models.ChatResponse, error)
}
