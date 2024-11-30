package service

import (
	"l33tcode/server/pkg/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//go:generate mockgen -destination ../mocks/mock_services_interfaces_test.go  -source service.go -package mocks
type QuestionService interface {
	SubmitQuestion(ctx *gin.Context)
	GetQuestion(ctx *gin.Context)
	RemoveQuestion(ctx *gin.Context)
	UpdateQuestion(ctx *gin.Context)
	ListQuestions(ctx *gin.Context)
}

type CodeSubmitterService interface {
	SubmitCode(ctx *gin.Context)
	ListCodeExecutors(ctx *gin.Context)
	SetCodeExecutor(ctx *gin.Context)
}

type Service interface {
	QuestionService
	CodeSubmitterService

	ListSupportedLanguages(ctx *gin.Context)
	GetLanguage(ctx *gin.Context)
}
type service struct {
	codeExecutorsMap    map[string]models.CodeExecuter
	currentCodeExecutor string
	questionRepo        models.QuestionRepository
	logger              *zap.Logger
	languageRepo        models.LanguageRepository
}

func NewService(logger *zap.Logger, questionRepo models.QuestionRepository,
	languageRepo models.LanguageRepository,
	codeExecutorFactory map[string]models.CodeExecuter,
	defaultCodeExecutor string) Service {
	srv := &service{
		currentCodeExecutor: defaultCodeExecutor,
		codeExecutorsMap:    codeExecutorFactory,
		questionRepo:        questionRepo,
		languageRepo:        languageRepo,
		logger:              logger,
	}
	return srv
}
