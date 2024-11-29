package codeexecutor

import (
	"context"
	"fmt"
	"l33tcode/server/pkg/config"
	"l33tcode/server/pkg/models"
	"net/http"
	"net/url"
)

type llmCodeExecutor struct {
	selectedModel        string
	systemPromptTemplate string
	URL                  url.URL
	client               *http.Client
}

func NewLLMCodeExecutor(model, host string, port int) models.CodeExecuter {
	ce := &llmCodeExecutor{}
	ce.systemPromptTemplate = config.Cfg.LLMConfiguration.SystemPromptTemplate

	ce.selectedModel = model
	urla, _ := url.Parse(fmt.Sprintf("%s:%d", host, port))
	ce.URL = *urla
	ce.client = http.DefaultClient
	return ce
}

func (ce *llmCodeExecutor) ExecuteCode(ctx context.Context, user, code string, question *models.Question, language *models.Language) (string, error) {

	return "", fmt.Errorf("not implemented")
}
