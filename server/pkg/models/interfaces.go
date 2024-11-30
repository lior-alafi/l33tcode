package models

import "context"

//go:generate mockgen -destination ../mocks/mock_repo_interfaces.go  -source interfaces.go -package mocks
type QuestionRepository interface {
	SaveQuestion(ctx context.Context, user string, q Question) (string, error)
	GetQuestion(ctx context.Context, user, qid, language string) (Question, error)
	DeleteQuestion(ctx context.Context, user, qid string) error
	ListQuestions(ctx context.Context, user, language string) ([]Question, error)
}

type CodeExecuter interface {
	ExecuteCode(ctx context.Context, user, code string, question *Question, language *Language) (string, error)
}
type LanguageRepository interface {
	ListSupportedLanguages(ctx context.Context, user string) ([]Language, error)
	GetSupportedLanguagesFromList(ctx context.Context, languages []string) ([]string, error)
	GetLanguage(ctx context.Context, lang string) (Language, error)
}
