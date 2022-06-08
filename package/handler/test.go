package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler *Handler) testHttpsHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "SUCCESSFUL TESTED")
}
