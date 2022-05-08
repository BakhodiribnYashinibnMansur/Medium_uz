package handler

import (
	"github.com/BakhodiribnYashinibnMansur/Medium_uz/model"
	"github.com/BakhodiribnYashinibnMansur/Medium_uz/util/error"
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
// @Failure 409 {object} error.errorResponseData
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

	checkedUser, err := handler.services.CheckDataExistsEmailNickName(input.Email, input.NickName, logrus)
	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}
	if !checkedUser.NickName || !checkedUser.Email {
		error.NewHandlerErrorResponseData(ctx, http.StatusConflict, "email or nickname already exist", checkedUser, logrus)
		return
	}
	logrus.Info("signUp data send for  create user to service")
	id, err := handler.services.Authorization.CreateUser(input, logrus)
	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}

	token, err := handler.services.Authorization.GenerateToken(input.Email, logrus)
	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusInternalServerError, err.Error(), logrus)
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
	checkedUser, err := handler.services.CheckDataExistsEmailNickName(input.Email, "", logrus)
	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}
	if checkedUser.Email {
		error.NewHandlerErrorResponse(ctx, http.StatusConflict, "email  not  exist", logrus)
		return
	}
	token, err := handler.services.Authorization.GenerateToken(input.Email, logrus)
	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusInternalServerError, err.Error(), logrus)
		return
	}
	ctx.JSON(http.StatusOK, model.ResponseSign{Token: token})
}

// @Summary Send code for  Recovery Password
// @Description send code for  Recovery Password
// @ID recovery-password-send-code
// @Tags   Auth
// @Accept       json
// @Produce      json
// @Param email query string false "email"
// @Success      200   {object}      model.ResponseSuccess
// @Failure 400,404 {object} error.errorResponse
// @Failure 409 {object} error.errorResponse
// @Failure 500 {object} error.errorResponse
// @Failure default {object} error.errorResponse
// @Router       /api/account/recovery [GET]
func (handler *Handler) recoveryForMessageToEmail(ctx *gin.Context) {
	// logrus := handler.logrus
}

// @Summary Check code for  Recovery Password
// @Description check code for  Recovery Password
// @ID recovery-code-email
// @Tags   Auth
// @Accept       json
// @Produce      json
// @Param code query string false "code"
// @Success      200   {object}      model.ResponseSuccess
// @Failure 400,404 {object} error.errorResponse
// @Failure 409 {object} error.errorResponse
// @Failure 500 {object} error.errorResponse
// @Failure default {object} error.errorResponse
// @Router       /api/account/recovery-verify [GET]
func (handler *Handler) recoveryCheckEmailCode(ctx *gin.Context) {
	// logrus := handler.logrus

}

// @Summary   Recovery Password
// @Description   Recovery Password
// @ID recovery-password
// @Tags   Auth
// @Accept       json
// @Produce      json
// @Param password query string false "password"
// @Success      200   {object}      model.ResponseSuccess
// @Failure 400,404 {object} error.errorResponse
// @Failure 409 {object} error.errorResponse
// @Failure 500 {object} error.errorResponse
// @Failure default {object} error.errorResponse
// @Router       /api/account/recovery-password [GET]
func (handler *Handler) recoveryPassword(ctx *gin.Context) {
	// logrus := handler.logrus

}
