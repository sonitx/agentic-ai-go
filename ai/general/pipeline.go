package general

import (
	"context"
	"fmt"
	"main/ai"
	"main/ai/llms"
	"main/ai/models"
	"main/utils"
)

type GeneralAgent struct {
}

// GenerateResponse implements ai.AgenticInterface.
func (g *GeneralAgent) GenerateResponse(ctx context.Context, question string) (*models.ChatResponse, error) {
	utils.ShowInfoLogs("General Agent is working")

	agenticConf := utils.AppConfig.AgenticAI
	prompt := fmt.Sprintf(instruction, question)
	res, err := llms.GenerateResponse(ctx, models.ChatRequest{
		ModelType: agenticConf.Direct.ModelType,
		ModelName: agenticConf.Direct.ModelName,
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

func NewGeneralAgent() ai.AgenticInterface {
	return &GeneralAgent{}
}
