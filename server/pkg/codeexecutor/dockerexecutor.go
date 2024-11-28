package codeexecutor

import (
	"context"
	"fmt"
	"l33tcode/server/pkg/models"
)

type dockerCodeExecutor struct {
}

func NewDockerCodeExecutor(model, host string, port int) models.CodeExecuter {
	ce := &dockerCodeExecutor{}

	return ce
}

func (ce *dockerCodeExecutor) ExecuteCode(ctx context.Context, user, code string, question *models.Question, language *models.Language) (string, error) {

	return "", fmt.Errorf("not implemented")
}
