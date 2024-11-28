package models

import "context"

//go:generate mockgen -destination ../mocks/mock_repo_interfaces_test.go  -source interfaces.go -package models
type QuestionRepository interface {
	SaveQuestion(ctx context.Context, user, q Question) (string, error)
	GetQuestion(ctx context.Context, user, qid, language string) (Question, error)
	DeleteQuestion(ctx context.Context, user, qid string) error
	ListQuestions(ctx context.Context, user, language string) ([]Question, error)
}

type CodeExecuter interface {
	ExecuteCode(ctx context.Context, user, code string, question *Question, language *Language) (string, error)
	TestCode(ctx context.Context, user, code, inputs string, question *Question, language *Language)
}
type LanguageRepository interface {
	ListSupportedLanguages(ctx context.Context, user string) ([]Language, error)
}
