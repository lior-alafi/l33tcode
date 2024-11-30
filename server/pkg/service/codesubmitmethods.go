package service

import (
	"context"
	"encoding/json"
	"errors"
	"l33tcode/server/pkg/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (srv *service) SubmitCode(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
		return
	}

	var newCodeSubmit models.CodeSubmitRequest
	if err := json.NewDecoder(c.Request.Body).Decode(&newCodeSubmit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	user := c.Params.ByName("user")
	if err := models.IsEmpty(newCodeSubmit.QID, "qid"); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		srv.logger.Error("selected code executor doesn't exist", zap.Error(err), zap.Any("submittedCode", newCodeSubmit))
		return
	}
	question, err := srv.questionRepo.GetQuestion(context.Background(), user, newCodeSubmit.QID, newCodeSubmit.Language)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		srv.logger.Error("getting question for codesubmit failed", zap.Error(err), zap.Any("submittedCode", newCodeSubmit))
		return
	}

	result, err := srv.codeExecutorsMap[srv.currentCodeExecutor].ExecuteCode(context.Background(), user, newCodeSubmit.Code, &question, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		srv.logger.Error("failed to issue codeExecution with current codesummit", zap.Error(err), zap.Any("submittedCode", newCodeSubmit))
		return
	}

	c.JSON(http.StatusOK, result)

}

func (srv *service) ListCodeExecutors(c *gin.Context) {
	if c.Request.Method != http.MethodGet {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
		return
	}
	codeExecutors := make([]string, 0, len(srv.codeExecutorsMap))
	for k := range srv.codeExecutorsMap {
		codeExecutors = append(codeExecutors, k)
	}
	c.JSON(http.StatusOK, codeExecutors)
}

func (srv *service) SetCodeExecutor(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
		return
	}

	var selectedCodeExecutor models.SelectCodeExecuterRequest
	if err := json.NewDecoder(c.Request.Body).Decode(&selectedCodeExecutor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	if _, ok := srv.codeExecutorsMap[selectedCodeExecutor.Name]; !ok {
		err := errors.New("selected code executor doesn't exist")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		srv.logger.Error("selected code executor doesn't exist", zap.Error(err), zap.Any("selectedCE", selectedCodeExecutor))
		return
	}

	srv.currentCodeExecutor = selectedCodeExecutor.Name
	c.JSON(http.StatusOK, "OK")
}
