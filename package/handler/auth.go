package handler

import "github.com/gin-gonic/gin"

func (handler *Handler) signUp(c *gin.Context) {
	logrus := handler.logrus
	logrus.Info("Signing up")
}

func (h *Handler) signIn(ctx *gin.Context) {
}
