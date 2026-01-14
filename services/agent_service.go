package services

import (
	"context"
	"errors"
	"main/models"
)

type AgentService struct {
}

func NewAgentService() *AgentService {
	return &AgentService{}
}

func (s *AgentService) GenerateResponse(ctx context.Context, agentKey string, prompt string) (*models.ChatResponse, error) {
	return nil, errors.New("agent not found")
}
