package main

import (
	"fmt"
	"l33tcode/server/pkg/config"
	"l33tcode/server/pkg/service"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// r.HandleMethodNotAllowed = true
	if err := config.LoadConfigurations("config.json"); err != nil {
		panic("missing configuration")
	}
	srv := service.NewService(nil, nil, nil, "")

	r.POST("/admin/question/submit", srv.SubmitQuestion)
	r.GET("/question/list", srv.ListQuestions)
	r.GET("/question/:qid", srv.GetQuestion)
	r.DELETE("/admin/question/:qid", srv.RemoveQuestion)
	r.PUT("/admin/question/:qid", srv.UpdateQuestion)

	r.GET("/admin/languages/list", srv.ListSupportedLanguages)
	r.GET("/admin/codeexecutor/list", srv.ListCodeExecutors)
	r.POST("/admin/codeexecutor/set", srv.SetCodeExecutor)
	r.Run(fmt.Sprintf(":%d", config.Cfg.Port))
}
