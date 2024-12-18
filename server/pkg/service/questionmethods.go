package service

import (
	"context"
	"encoding/json"
	"fmt"
	"l33tcode/server/pkg/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (srv *service) SubmitQuestion(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
		return
	}

	var newQuestion models.Question
	if err := json.NewDecoder(c.Request.Body).Decode(&newQuestion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	if err := newQuestion.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		srv.logger.Error("question was invalid", zap.Error(err), zap.Any("question", newQuestion))
		return
	}
	user := c.Params.ByName("user")
	qid, err := srv.questionRepo.SaveQuestion(context.Background(), user, newQuestion)
	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("something went wrong error code(%d)", ErrorCodeDB)})

		srv.logger.Error("Save Question failed", zap.Error(err), zap.Int("errorCode", ErrorCodeDB), zap.Any("question", newQuestion))
		return
	}
	newQuestion.Id = qid
	c.JSON(http.StatusCreated, newQuestion)
}

func (srv *service) GetQuestion(c *gin.Context) {
	if c.Request.Method != http.MethodGet {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
		return
	}

	user := c.Params.ByName("user")
	qid := c.Params.ByName("qid")
	lang := c.Params.ByName("language")

	question, err := srv.questionRepo.GetQuestion(context.Background(), user, qid, lang)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("something went wrong error code(%d)", ErrorCodeDB)})
		srv.logger.Error("Get Question failed", zap.Error(err), zap.Int("errorCode", ErrorCodeDB), zap.String("qid", qid), zap.String("language", lang))
		return
	}

	c.JSON(http.StatusOK, question)
}

func (srv *service) RemoveQuestion(c *gin.Context) {
	if c.Request.Method != http.MethodDelete {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
		return
	}

	user := c.Params.ByName("user")
	qid := c.Params.ByName("qid")

	err := srv.questionRepo.DeleteQuestion(context.Background(), user, qid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("something went wrong error code(%d)", ErrorCodeDB)})
		srv.logger.Error("Delete Question failed", zap.Error(err), zap.Int("errorCode", ErrorCodeDB), zap.String("qid", qid))

		return
	}

	c.JSON(http.StatusOK, "deleted")
}

func (srv *service) ListQuestions(c *gin.Context) {
	if c.Request.Method != http.MethodGet {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
		return
	}

	user := c.Params.ByName("user")
	language := c.Params.ByName("language")

	questions, err := srv.questionRepo.ListQuestions(context.Background(), user, language)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("something went wrong error code(%d)", ErrorCodeDB)})
		srv.logger.Error("List Questions failed", zap.Error(err), zap.Int("errorCode", ErrorCodeDB))
		return
	}

	c.JSON(http.StatusOK, questions)
}

func (srv *service) UpdateQuestion(c *gin.Context) {
	if c.Request.Method != http.MethodPut {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
		return
	}
	var question models.Question
	if err := json.NewDecoder(c.Request.Body).Decode(&question); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	if err := question.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.IsEmpty(question.Id, "id"); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	user := c.Params.ByName("user")

	_, err := srv.questionRepo.SaveQuestion(context.Background(), user, question)
	if err != nil {
		srv.logger.Error("update Question failed", zap.Error(err), zap.Int("errorCode", ErrorCodeDB), zap.Any("qid", question))
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("something went wrong error code(%d)", ErrorCodeDB)})
		return
	}

	c.JSON(http.StatusOK, "OK")
}
