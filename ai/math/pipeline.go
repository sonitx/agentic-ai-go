package math

import (
	"context"
	"fmt"
	"main/ai"
	"main/ai/llms"
	"main/ai/models"
	"main/utils"
)

type MathAgent struct {
}

func NewMathAgent() ai.AgenticInterface {
	return &MathAgent{}
}

func (m *MathAgent) GenerateResponse(ctx context.Context, question string) (*models.ChatResponse, error) {
	utils.ShowInfoLogs("Math Agent is working")

	agenticConf := utils.AppConfig.AgenticAI
	prompt := fmt.Sprintf(instruction, question)
	res, err := llms.GenerateResponse(ctx, models.ChatRequest{
		ModelType: agenticConf.Logic.ModelType,
		ModelName: agenticConf.Logic.ModelName,
		Question:  prompt,
		Tools:     nil,
	})
	if err != nil {
		return nil, err
	}
	return &models.ChatResponse{
		Answer:     res.Text(),
		TotalToken: res.Usage.TotalTokens,
	}, nil
}
