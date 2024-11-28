package service

import (
	"l33tcode/server/pkg/models"

	"github.com/gin-gonic/gin"
)

//go:generate mockgen -destination ../mocks/mock_services_interfaces_test.go  -source service.go -package service
type QuestionService interface {
	SubmitQuestion(ctx *gin.Context)
	GetQuestion(ctx *gin.Context)
	RemoveQuestion(ctx *gin.Context)
	UpdateQuestion(ctx *gin.Context)
	ListQuestions(ctx *gin.Context)
}

type CodeSubmitterService interface {
	SubmitCode(ctx *gin.Context)
	TestCode(ctx *gin.Context)
	ListCodeExecutors(ctx *gin.Context)
	SetCodeExecutor(ctx *gin.Context)
}

type Service interface {
	QuestionService
	CodeSubmitterService

	ListSupportedLanguages(ctx *gin.Context)
}
type service struct {
	codeExecutorsMap    map[string]models.CodeExecuter
	currentCodeExecutor string
	questionRepo        models.QuestionRepository
	languageRepo        models.LanguageRepository
}

func NewService(questionRepo models.QuestionRepository,
	languageRepo models.LanguageRepository,
	codeExecutorFactory map[string]models.CodeExecuter,
	defaultCodeExecutor string) Service {
	srv := &service{
		currentCodeExecutor: defaultCodeExecutor,
		codeExecutorsMap:    codeExecutorFactory,
		questionRepo:        questionRepo,
		languageRepo:        languageRepo,
	}
	return srv
}
