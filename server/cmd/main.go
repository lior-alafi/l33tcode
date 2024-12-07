package main

import (
	"fmt"
	"l33tcode/server/pkg/codeexecutor"
	"l33tcode/server/pkg/config"
	"l33tcode/server/pkg/models"
	"l33tcode/server/pkg/repositories"
	"l33tcode/server/pkg/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	r := gin.Default()

	if err := config.LoadConfigurations("config.json"); err != nil {
		panic("missing configuration")
	}
	llmce := codeexecutor.NewLLMCodeExecutor(logger,
		config.Cfg.LLMConfiguration.Model,
		config.Cfg.LLMConfiguration.Host,
		config.Cfg.LLMConfiguration.Port,
		config.Cfg.LLMConfiguration.ChatURL,
		config.Cfg.LLMConfiguration.SystemPromptTemplate,
		config.Cfg.LLMConfiguration.SubmitPattern,
	)
	codeExecFactory := map[string]models.CodeExecuter{
		"llm": llmce,
	}

	esLangsRepo := repositories.NewElasticLanguageRepository()
	esQuestionsRepo := repositories.NewElasticQuestionsRepository()
	srv := service.NewService(logger, esQuestionsRepo, esLangsRepo, codeExecFactory, "llm")

	r.POST("/admin/question/submit", srv.SubmitQuestion)
	r.GET("/question/list", srv.ListQuestions)
	r.GET("/question/:qid", srv.GetQuestion)
	r.DELETE("/admin/question/:qid", srv.RemoveQuestion)
	r.PUT("/admin/question/:qid", srv.UpdateQuestion)

	r.GET("/admin/languages/list", srv.ListSupportedLanguages)
	r.GET("/admin/codeexecutor/list", srv.ListCodeExecutors)
	r.POST("/admin/codeexecutor/set", srv.SetCodeExecutor)

	r.POST("/code/submit")

	r.Run(fmt.Sprintf(":%d", config.Cfg.Port))
}
