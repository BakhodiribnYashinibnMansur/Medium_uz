package handler

import (
	"mediumuz/model"
	"mediumuz/util/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary SignUp
// @Tags auth
// @Description create account
// @ID create-account
// @Accept  json
// @Produce  json
// @Param input body model.SingUpUserJson true "account info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errors.errorResponse
// @Failure 500 {object} errors.errorResponse
// @Failure default {object} errors.errorResponse
// @Router /auth/sign-up [post]
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
