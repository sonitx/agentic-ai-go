package services

import (
	"context"
	"main/ai"
	"main/ai/general"
	"main/ai/math"
	"main/ai/models"
	"main/ai/rag"
	"main/ai/router"
	"main/utils"
)

type AgentService struct {
	Classifier *router.Classifier
	General    *general.GeneralAgent
	Math       *math.MathAgent
	Rag        *rag.RagAgent
}

func NewAgentService() *AgentService {
	return &AgentService{}
}

func (s *AgentService) GenerateResponse(ctx context.Context, question string) (*models.ChatResponse, error) {
	intent, err := s.Classifier.Classify(ctx, question)
	if err != nil {
		return nil, err
	}

	utils.ShowInfoLogs("Question: %s, Intent: %s", question, intent)
	var agenticAi ai.AgenticInterface
	switch intent {
	case router.IntentKnowledge:
		agenticAi = s.Rag
	case router.IntentMath:
		agenticAi = s.Math
	default:
		agenticAi = s.General
	}

	return agenticAi.GenerateResponse(ctx, question)
}
