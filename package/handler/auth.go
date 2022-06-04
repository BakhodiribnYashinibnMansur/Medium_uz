package handler

import (
	"net/http"

	"github.com/BakhodiribnYashinibnMansur/Medium_uz/model"
	"github.com/BakhodiribnYashinibnMansur/Medium_uz/util/error"

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
	err := handler.services.CheckDataExistsEmailPassword(input.Email, input.Password, logrus)
	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}
	token, err := handler.services.Authorization.GenerateToken(input.Email, logrus)
	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusInternalServerError, err.Error(), logrus)
		return
	}
	ctx.JSON(http.StatusOK, model.ResponseSign{Token: token})
}

// @Summary Check Email for Recovery Password
// @Description Check Email for Recovery Password
// @ID recovery-check-email
// @Tags   Auth
// @Accept       json
// @Produce      json
// @Param email query string false "email"
// @Success      200   {object}      model.ResponseSuccess
// @Failure 400,404 {object} error.errorResponse
// @Failure 409 {object} error.errorResponse
// @Failure 500 {object} error.errorResponse
// @Failure default {object} error.errorResponse
// @Router       /auth/recovery-check [GET]
func (handler *Handler) recoveryCheckEmail(ctx *gin.Context) {
	logrus := handler.logrus
	userEmail := ctx.Query("email")
	if userEmail == "" {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, "Email is empty", logrus)
		return
	}
	checkedUser, err := handler.services.CheckDataExistsEmailNickName(userEmail, "", logrus)
	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}
	if checkedUser.Email {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, "Email Not Found", logrus)
		return
	}
	ctx.JSON(http.StatusOK, model.ResponseSuccess{Message: "DONE", Data: checkedUser.Email})
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
// @Router       /auth/recovery-send [GET]
func (handler *Handler) recoverySendEmail(ctx *gin.Context) {
	logrus := handler.logrus
	userEmail := ctx.Query("email")
	if userEmail == "" {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, "Email is empty", logrus)
		return
	}
	logrus.Infof(userEmail)
	err := handler.services.SendMessageEmail(userEmail, "", logrus)
	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}
	ctx.JSON(http.StatusOK, model.ResponseSuccess{Message: "DONE"})
}

// @Summary Check code for  Recovery Password
// @Description check code for  Recovery Password
// @ID recovery-check-code-email
// @Tags   Auth
// @Accept       json
// @Produce      json
// @Param code query string false "code"
// @Param email query string false "email"
// @Success      200   {object}      model.ResponseSuccess
// @Failure 400,404 {object} error.errorResponse
// @Failure 409 {object} error.errorResponse
// @Failure 500 {object} error.errorResponse
// @Failure default {object} error.errorResponse
// @Router       /auth/recovery-verify [GET]
func (handler *Handler) recoveryCheckEmailCode(ctx *gin.Context) {
	logrus := handler.logrus
	code := ctx.Query("code")
	if code == "" {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, "Code is empty", logrus)
		return
	}
	userEmail := ctx.Query("email")
	if userEmail == "" {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, "Email is empty", logrus)
		return
	}
	err := handler.services.RecoveryCheckEmailCode(userEmail, code, logrus)
	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}

	ctx.JSON(http.StatusOK, model.ResponseSuccess{Message: "VERIFIED"})
}

// @Summary   Recovery Password
// @Description   Recovery Password
// @ID recovery-password
// @Tags   Auth
// @Accept       json
// @Produce      json
// @Param password query string false "password"
// @Param email query string false "email"
// @Success      200   {object}      model.ResponseSuccess
// @Failure 400,404 {object} error.errorResponse
// @Failure 409 {object} error.errorResponse
// @Failure 500 {object} error.errorResponse
// @Failure default {object} error.errorResponse
// @Router       /auth/recovery-password [GET]
func (handler *Handler) recoveryPassword(ctx *gin.Context) {
	logrus := handler.logrus
	password := ctx.Query("password")
	if password == "" {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, "Password is empty", logrus)
		return
	}
	email := ctx.Query("email")
	if email == "" {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, "Email is empty", logrus)
		return
	}
	effectedRowsNum, err := handler.services.UpdateAccountPassword(email, password, logrus)
	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}

	if effectedRowsNum == 0 {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, "User not found", logrus)
		return
	}
	ctx.JSON(http.StatusOK, model.ResponseSuccess{Message: "Updated Account Password"})
}
