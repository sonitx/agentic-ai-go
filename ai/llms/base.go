package llms

import (
	"context"
	"errors"
	"main/ai/models"
	"main/utils"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
	"github.com/firebase/genkit/go/plugins/googlegenai"
	"github.com/firebase/genkit/go/plugins/ollama"
)

func GenerateResponse(ctx context.Context, prompt string, tools []models.AITool) (*ai.ModelResponse, error) {
	g, modelName := getModel(ctx, utils.AppConfig.ModelConfig)
	if g == nil {
		return nil, errors.New("model not found")
	}

	opts := []ai.GenerateOption{
		ai.WithPrompt(prompt),
		ai.WithModelName(modelName),
	}

	// define tools
	toolsRef := make([]ai.ToolRef, len(tools))
	for i, item := range tools {
		toolsRef[i] = genkit.DefineTool(g, item.Name, item.Description, item.Function)
	}

	if len(toolsRef) > 0 {
		opts = append(opts, ai.WithTools(toolsRef...))
	}

	response, err := genkit.Generate(ctx, g, opts...)
	if err != nil {
		utils.ShowErrorLogs(err)
		return nil, err
	}
	return response, nil
}

func getModel(ctx context.Context, agenticConfig utils.AgenticAI, llmProvider utils.LLMProvider) (*genkit.Genkit, string) {
	var g *genkit.Genkit
	var modelName string

	if agenticConfig.Router.ModelType == "gemini" {
		g = genkit.Init(ctx, genkit.WithPlugins(&googlegenai.GoogleAI{
			APIKey: llmProvider.Gemini.APIKey,
		}))
		modelName = agenticConfig.Router.ModelName
	} else if llmProvider.Ollama.Enable {
		g = genkit.Init(ctx, genkit.WithPlugins(
			&ollama.Ollama{
				ServerAddress: llmProvider.Ollama.ServerAddress,
			},
		))
		modelName = agenticConfig.Router.ModelName
	}
	return g, modelName
}
