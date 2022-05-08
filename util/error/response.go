package error

import (
	"github.com/BakhodiribnYashinibnMansur/Medium_uz/util/logrus"

	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Message string `json:"message"`
}
type errorResponseData struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type statusResponse struct {
	Status string `json:"status"`
}

func NewHandlerErrorResponse(ctx *gin.Context, statusCode int, message string, logrus *logrus.Logger) {
	logrus.Error(message)
	ctx.AbortWithStatusJSON(statusCode, errorResponse{message})
}

func NewHandlerErrorResponseData(ctx *gin.Context, statusCode int, message string, data interface{}, logrus *logrus.Logger) {
	logrus.Error(message)
	ctx.AbortWithStatusJSON(statusCode, errorResponseData{message, data})
}
