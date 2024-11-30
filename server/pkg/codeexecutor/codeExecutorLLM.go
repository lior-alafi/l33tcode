package codeexecutor

import (
	"context"
	"fmt"
	"l33tcode/server/pkg/models"
	"net/http"
	"net/url"

	"go.uber.org/zap"
)

type llmCodeExecutor struct {
	selectedModel        string
	systemPromptTemplate string
	codeSumbitPrompt     string
	URL                  url.URL
	client               *http.Client
	logger               *zap.Logger
}

func NewLLMCodeExecutor(logger *zap.Logger, model, host string, port int, chatURL, systemPromptTemplate, CodeSubmittingTemplate string) models.CodeExecuter {
	ce := &llmCodeExecutor{logger: logger}
	ce.systemPromptTemplate = systemPromptTemplate
	ce.codeSumbitPrompt = CodeSubmittingTemplate
	ce.selectedModel = model
	urla, _ := url.Parse(fmt.Sprintf("%s:%d/%s", host, port, chatURL))
	ce.URL = *urla
	ce.client = http.DefaultClient
	return ce
}

func (ce *llmCodeExecutor) ExecuteCode(ctx context.Context, user, code string, question *models.Question, language *models.Language) (string, error) {

	return "", fmt.Errorf("not implemented")
}
