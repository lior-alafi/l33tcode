package service

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (srv *service) ListSupportedLanguages(c *gin.Context) {
	if c.Request.Method != http.MethodGet {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
		return
	}

	user := c.Params.ByName("user")

	languages, err := srv.languageRepo.ListSupportedLanguages(context.Background(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("something went wrong error code(%d)", ErrorCodeDB)})
		return
	}

	c.JSON(http.StatusOK, languages)
}
