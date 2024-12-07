package codeexecutor

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
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
	for _, tst := range question.SupportedLanguagges[0].Tests {
		jsonData, err := ce.PrepareMessage(ctx, user, code, &tst, language)
		if err != nil {
			ce.logger.Error("ExecuteCode::Error preparing llm code execution request", zap.Error(err), zap.Any("test", tst), zap.String("user", user), zap.String("qid", question.Id))
			return "", err
		}

		body, err := ce.httpClientAux(jsonData, tst, user, question)
		if err != nil {
			return "", err
		}
		var llmresp models.LLMCodeExecutionResponse
		if err := json.Unmarshal(body, &llmresp); err != nil {
			ce.logger.Error("ExecuteCode::fail to unmarshal response", zap.Error(err), zap.Any("test", tst), zap.String("user", user), zap.String("qid", question.Id))
			return "", err
		}

		if err := models.IsEmpty(llmresp.Error, "llmresponse.error"); err != nil {
			return models.CodeExecutionResponsError, nil
		}

		if fmt.Sprintf("%v", llmresp.Output) != tst.Expected {
			return models.CodeExecutionResponsFail, nil
		}
	}

	return models.CodeExecutionResponsPass, nil
}

func (ce *llmCodeExecutor) httpClientAux(jsonData []byte, tst models.Test, user string, question *models.Question) ([]byte, error) {
	req, err := http.NewRequest(http.MethodPost, ce.URL.String(), bytes.NewBuffer(jsonData))
	if err != nil {
		ce.logger.Error("ExecuteCode::Error creating request", zap.Error(err), zap.Any("test", tst), zap.String("user", user), zap.String("qid", question.Id))
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := ce.client.Do(req)
	if err != nil {
		ce.logger.Error("ExecuteCode::Error Post request", zap.Error(err), zap.Any("test", tst), zap.String("user", user), zap.String("qid", question.Id))
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		ce.logger.Error("ExecuteCode::read Response", zap.Error(err), zap.Any("test", tst), zap.String("user", user), zap.String("qid", question.Id))
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		err := errors.New("bad status code")
		ce.logger.Error("ExecuteCode::bad status code", zap.Error(err), zap.Any("test", tst), zap.String("user", user), zap.String("qid", question.Id))
		return nil, err
	}
	return body, nil
}
func (ce *llmCodeExecutor) PrepareMessage(ctx context.Context, user, code string, test *models.Test, language *models.Language) ([]byte, error) {
	codexecPayload := models.LLMCodeExecutionRequest{Model: ce.selectedModel,
		Stream:      false,
		Temperature: 0.7,
		MaxTokens:   -1,
		Messages: []models.LLMessage{
			{
				Role:    "system",
				Content: fmt.Sprintf(ce.systemPromptTemplate, language.Name),
			},
			{
				Role:    "user",
				Content: fmt.Sprintf(ce.codeSumbitPrompt, code, test.Inputs),
			},
		},
	}

	return json.Marshal(codexecPayload)
}
