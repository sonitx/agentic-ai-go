package utils

var AppConfig *Config

type Config struct {
	Server struct {
		Port int    `yaml:"port"`
		Mode string `yaml:"mode"`
	} `yaml:"server"`
	LLMProvider LLMProvider `yaml:"llm-providers"`
	AgenticAI   AgenticAI   `yaml:"agentic-ai"`
}

type LLMProvider struct {
	Gemini struct {
		APIKey string `yaml:"apikey"`
	} `yaml:"gemini"`
	Ollama struct {
		ServerAddress string `yaml:"server-address"`
	} `yaml:"ollama"`
}

type AgenticAI struct {
	Router AgenticAIDetail `yaml:"router"`
	RAG    AgenticAIDetail `yaml:"rag"`
	Direct AgenticAIDetail `yaml:"direct"`
	Logic  AgenticAIDetail `yaml:"logic"`
}

type AgenticAIDetail struct {
	ModelType string `yaml:"model-type"`
	ModelName string `yaml:"model-name"`
}
