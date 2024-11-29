package service

import (
	"encoding/json"
	"l33tcode/server/pkg/models"
	"net/http"

	"github.com/gin-gonic/gin"
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
}
