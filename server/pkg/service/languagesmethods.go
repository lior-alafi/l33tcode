package service

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (srv *service) ListSupportedLanguages(c *gin.Context) {
	if c.Request.Method != http.MethodGet {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
		return
	}

	user := c.Params.ByName("user")

	languages, err := srv.languageRepo.ListSupportedLanguages(context.Background(), user)
	if err != nil {
		srv.logger.Error("list language failed", zap.Error(err), zap.Int("errorCode", ErrorCodeDB), zap.String("user", user))
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("something went wrong error code(%d)", ErrorCodeDB)})
		return
	}

	c.JSON(http.StatusOK, languages)
}

func (srv *service) GetLanguage(c *gin.Context) {
	if c.Request.Method != http.MethodGet {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
		return
	}

	lang := c.Params.ByName("language")
	language, err := srv.languageRepo.GetLanguage(context.Background(), lang)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("something went wrong error code(%d)", ErrorCodeDB)})
		srv.logger.Error("Get language failed", zap.Error(err), zap.Int("errorCode", ErrorCodeDB), zap.String("language", lang))
		return
	}

	c.JSON(http.StatusOK, language)
}
