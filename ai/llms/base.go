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

const (
	ProviderGemini = "gemini"
	ProviderOllama = "ollama"
)

func GenerateResponse(ctx context.Context, req models.ChatRequest) (*ai.ModelResponse, error) {
	g, err := getProvider(ctx, req.ModelType)
	if err != nil {
		return nil, err
	}

	opts := []ai.GenerateOption{
		ai.WithPrompt(req.Question),
		ai.WithModelName(req.ModelName),
	}

	// define tools
	toolsRef := make([]ai.ToolRef, len(req.Tools))
	for i, item := range req.Tools {
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

func getProvider(ctx context.Context, providerName string) (*genkit.Genkit, error) {
	var g *genkit.Genkit
	providerConfig := utils.AppConfig.LLMProvider

	switch providerName {
	case ProviderGemini:
		g = genkit.Init(ctx, genkit.WithPlugins(&googlegenai.GoogleAI{
			APIKey: providerConfig.Gemini.APIKey,
		}))
	case ProviderOllama:
		g = genkit.Init(ctx, genkit.WithPlugins(
			&ollama.Ollama{
				ServerAddress: providerConfig.Ollama.ServerAddress,
			},
		))
	default:
		return nil, errors.New("provider not found")
	}
	return g, nil
}
