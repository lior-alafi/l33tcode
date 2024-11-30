package models

type LLMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
type LLMCodeExecutionRequest struct {
	Model       string      `json:"model"`
	Stream      bool        `json:"stream"`
	Temperature float32     `json:"temperature"`
	MaxTokens   int         `json:"max_tokens"`
	Messages    []LLMessage `json:"messages"`
}
type LLMCodeExecutionResponse struct {
	Error  string         `json:"error,omitempty"`
	Inputs map[string]any `json:"input,omitempty"`
	Output map[string]any `json:"output,omitempty"`
}
