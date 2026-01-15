package router

import (
	"context"
	"fmt"
	"main/ai/llms"
	"main/ai/models"
	"main/utils"
	"strings"
)

const (
	IntentKnowledge string = "KNOWLEDGE"
	IntentMath      string = "MATH"
	IntentGeneral   string = "GENERAL"
)

type Classifier struct {
}

func NewClassifier() *Classifier {
	return &Classifier{}
}

func (c *Classifier) Classify(ctx context.Context, question string) (string, error) {
	agenticConf := utils.AppConfig.AgenticAI
	prompt := fmt.Sprintf(instruction, question)
	res, err := llms.GenerateResponse(ctx, models.ChatRequest{
		ModelType: agenticConf.Router.ModelType,
		ModelName: agenticConf.Router.ModelName,
		Question:  prompt,
		Tools:     nil,
	})
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(res.Text()), nil
}
