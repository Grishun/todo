package handler

import (
	"github.com/gin-gonic/gin"
	"log/slog"
)

type err struct {
	Message string `json:"message"`
}

func newErrResp(c *gin.Context, statusCode int, message string) {
	slog.Error(message)
	c.AbortWithStatusJSON(statusCode, err{message})
}
