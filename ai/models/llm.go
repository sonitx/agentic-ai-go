package models

import "github.com/firebase/genkit/go/ai"

type ChatRequest struct {
	ModelType string
	ModelName string
	Question  string
	Tools     []AITool
}

type ChatResponse struct {
	Answer     string
	TotalToken int
}

type AITool struct {
	Name        string
	Description string
	Function    func(ctx *ai.ToolContext, input any) (string, error)
}
