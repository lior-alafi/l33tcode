package repositories

import (
	"context"
	"fmt"
	"l33tcode/server/pkg/models"
)

type elasticLanguageRepository struct {
}

func NewElasticLanguageRepository() models.LanguageRepository {
	return &elasticLanguageRepository{}
}

func (r *elasticLanguageRepository) ListSupportedLanguages(ctx context.Context, user string) ([]models.Language, error) {
	return nil, fmt.Errorf("not implemented")
}
func (r *elasticLanguageRepository) GetSupportedLanguagesFromList(ctx context.Context, languages []string) ([]string, error) {
	return nil, fmt.Errorf("not implemented")
}
func (r *elasticLanguageRepository) GetLanguage(ctx context.Context, lang string) (models.Language, error) {
	return models.Language{}, fmt.Errorf("not implemented")
}
