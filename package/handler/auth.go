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
// @Success 200 {object} model.ResponseSignUp
// @Failure 400,404 {object} error.errorResponse
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
	logrus.Info("signUp data send for  create user to service")
	id, err := handler.services.Authorization.CreateUser(input, logrus)

	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}

	token, err := handler.services.Authorization.GenerateToken(input.FirstName, input.Password, logrus)
	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusInternalServerError, err.Error(), logrus)
		return
	}
	err = handler.services.SendMessageEmail(input.Email, input.FirstName, logrus)

	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}
	ctx.JSON(http.StatusOK, model.ResponseSignUp{Id: id, Token: token})
}

func (handler *Handler) signIn(ctx *gin.Context) {
}

// @Summary  Verification Email
// @Description verification email with code
// @ID verify-email
// @Tags   Auth
// @Accept       json
// @Produce      json
// @Param username query string false "username"
// @Param code query string false "code"
// @Success      200   {object}      model.ResponseSuccess
// @Failure 400,404 {object} error.errorResponse
// @Failure 500 {object} error.errorResponse
// @Failure default {object} error.errorResponse
// @Router       /auth/verify [GET]
func (handler *Handler) verifyEmail(ctx *gin.Context) {
	logrus := handler.logrus
	code := ctx.Query("code")
	username := ctx.Query("username")
	logrus.Infof("DONE: get : %s ,%s", code, username)
	id, err := handler.services.VerifyCode(username, code, logrus)
	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}

	if id == 0 {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, "User not found", logrus)
		return
	}
	ctx.JSON(http.StatusOK, model.ResponseSuccess{Message: "DONE"})
}

// @Summary Resend cod for  Verification Email
// @Description resend code to email for  verification
// @ID resend-code-email
// @Tags   Auth
// @Accept       json
// @Produce      json
// @Param email query string false "email"
// @Param username query string false "username"
// @Success      200   {object}      model.ResponseSuccess
// @Failure 400,404 {object} error.errorResponse
// @Failure 500 {object} error.errorResponse
// @Failure default {object} error.errorResponse
// @Router       /auth/resend [GET]
func (handler *Handler) resendCode(ctx *gin.Context) {
	logrus := handler.logrus
	email := ctx.Query("email")
	username := ctx.Query("username")
	logrus.Infof("DONE: get : %s ,%s", email, username)
	err := handler.services.SendMessageEmail(email, username, logrus)
	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}
	ctx.JSON(http.StatusOK, model.ResponseSuccess{Message: "DONE"})
}
