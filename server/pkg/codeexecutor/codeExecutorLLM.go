package codeexecutor

import (
	"context"
	"fmt"
	"l33tcode/server/pkg/models"
	"net/url"
)

const (
	DefaultModel = "qwen2.5-coder-14b-instruct"
)

type llmCodeExecutor struct {
	selectedModel        string
	systemPromptTemplate string
	URL                  url.URL
}

func NewLLMCodeExecutor(model, host string, port int) models.CodeExecuter {
	ce := &llmCodeExecutor{}
	ce.systemPromptTemplate = `given a code and inputs from the user:
1. check the language the entire code was written if it's not %s, reply to user {"error":"WRONG LANGUAGE"}, stop all executions and do not reply alternative solution!
2. Simulate the code instruction by instruction.
3. Report the trace of the code at the end of each iteration.
4. Think step by step and reply with the output of the function for the given input
5. reply the output in a json {"input": <put the inputs here>, "output": <put the output here>}`

	ce.selectedModel = model
	urla, _ := url.Parse(fmt.Sprintf("%s:%d", host, port))
	ce.URL = *urla
	return ce
}

func (ce *llmCodeExecutor) ExecuteCode(ctx context.Context, user, code string, question *models.Question, language *models.Language) (string, error) {

	return "", fmt.Errorf("not implemented")
}
