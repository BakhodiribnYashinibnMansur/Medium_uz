package handler

import (
	"mediumuz/model"
	"mediumuz/util/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler *Handler) signUp(ctx *gin.Context) {
	logrus := handler.logrus
	var input model.SingUpUserJson
	err := ctx.BindJSON(&input)
	if err != nil {
		errors.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}
	logrus.Info("signUp data send for  create user to service")
	id, err := handler.services.Authorization.CreateUser(input, logrus)

	if err != nil {
		errors.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) signIn(ctx *gin.Context) {
}
