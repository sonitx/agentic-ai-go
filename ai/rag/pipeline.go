package rag

import (
	"context"
	"main/ai"
	"main/ai/models"
	"main/utils"
)

type RagAgent struct {
}

func NewRagAgent() ai.AgenticInterface {
	return &RagAgent{}
}

func (r *RagAgent) GenerateResponse(ctx context.Context, question string) (*models.ChatResponse, error) {
	utils.ShowInfoLogs("RAG Agent is working")
	return nil, nil
}
