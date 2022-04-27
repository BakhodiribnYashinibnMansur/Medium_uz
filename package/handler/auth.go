package handler

import (
	"mediumuz/model"
	"mediumuz/util/error"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary SignUp
// @Tags Auth
// @Description create account
// @ID create-account
// @Accept  json
// @Produce  json
// @Param input body model.User true "account info"
// @Success 200 {object} model.ResponseSign
// @Failure 400,404 {object} error.errorResponse
// @Failure 409 {object} error.errorResponse
// @Failure 500 {object} error.errorResponse
// @Failure default {object} error.errorResponse
// @Router /auth/sign-up [post]
func (handler *Handler) signUp(ctx *gin.Context) {
	logrus := handler.logrus
	var input model.User
	err := ctx.BindJSON(&input)
	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}
	logrus.Info("signUp data send for  check user Data to service")

	count, err := handler.services.CheckDataExists(input.FirstName, logrus)
	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}
	if count != 0 {
		error.NewHandlerErrorResponse(ctx, http.StatusConflict, "firstName already exist", logrus)
		return
	}
	logrus.Info("signUp data send for  create user to service")
	id, err := handler.services.Authorization.CreateUser(input, logrus)
	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}

	token, err := handler.services.Authorization.GenerateToken(input.FirstName, logrus)
	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusInternalServerError, err.Error(), logrus)
		return
	}
	err = handler.services.SendMessageEmail(input.Email, input.FirstName, logrus)

	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}
	ctx.JSON(http.StatusOK, model.ResponseSign{Id: id, Token: token})
}

// @Summary SignIn
// @Tags Auth
// @Description login account
// @ID login-account
// @Accept  json
// @Produce  json
// @Param input body model.SignInInput true "account info"
// @Success 200 {object} model.ResponseSign
// @Failure 400,404 {object} error.errorResponse
// @Failure 409 {object} error.errorResponse
// @Failure 500 {object} error.errorResponse
// @Failure default {object} error.errorResponse
// @Router /auth/sign-in [post]
func (handler *Handler) signIn(ctx *gin.Context) {
	logrus := handler.logrus
	var input model.SignInInput

	if err := ctx.BindJSON(&input); err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}

	token, err := handler.services.Authorization.GenerateToken(input.Username, logrus)
	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusInternalServerError, err.Error(), logrus)
		return
	}
	ctx.JSON(http.StatusOK, model.ResponseSign{Token: token})
}

// // @Summary Resend code for  Recovery Password
// // @Description Resend code for  Recovery Password
// // @ID recovery-password
// // @Tags   Account
// // @Accept       json
// // @Produce      json
// // @Param code query string false "code"
// // @Param password query string false "password"
// // @Success      200   {object}      model.ResponseSuccess
// // @Failure 400,404 {object} error.errorResponse
// // @Failure 409 {object} error.errorResponse
// // @Failure 500 {object} error.errorResponse
// // @Failure default {object} error.errorResponse
// // @Router       /api/account/resend [GET]

// func (handler *Handler) recoveryPassword(ctx *gin.Context) {
// 	logrus := handler.logrus

// }
