package repositories

import (
	"context"
	"fmt"
	"l33tcode/server/pkg/models"
)

type elasticQuestionsRepository struct {
}

func NewElasticQuestionsRepository() models.QuestionRepository {
	return &elasticQuestionsRepository{}
}
func (r *elasticQuestionsRepository) SaveQuestion(ctx context.Context, user string, q models.Question) (string, error) {
	return "", fmt.Errorf("not implemented")
}
func (r *elasticQuestionsRepository) GetQuestion(ctx context.Context, user, qid, language string) (models.Question, error) {

	return models.Question{}, fmt.Errorf("not implemented")
}
func (r *elasticQuestionsRepository) DeleteQuestion(ctx context.Context, user, qid string) error {
	return fmt.Errorf("not implemented")
}
func (r *elasticQuestionsRepository) ListQuestions(ctx context.Context, user, language string) ([]models.Question, error) {
	return nil, fmt.Errorf("not implemented")
}
