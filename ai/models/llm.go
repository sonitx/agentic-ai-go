package models

type ChatMessage struct {
	Role    string // "system" | "user" | "assistant"
	Content string
}

type CompletionParams struct {
	Model       string
	Temperature float32
	MaxTokens   int
	JsonSchema  string // nếu muốn structured output
	Tools       any    // mở rộng: tool calling
}
