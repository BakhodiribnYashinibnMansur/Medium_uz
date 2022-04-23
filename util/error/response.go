package error

import (
	"mediumuz/util/logrus"

	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Message string `json:"message"`
}

type statusResponse struct {
	Status string `json:"status"`
}

func NewHandlerErrorResponse(ctx *gin.Context, statusCode int, message string, logrus *logrus.Logger) {
	logrus.Error(message)
	ctx.AbortWithStatusJSON(statusCode, errorResponse{message})
}
