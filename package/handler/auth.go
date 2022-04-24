package handler

import (
	"fmt"
	"io"
	"log"
	"mediumuz/model"
	"mediumuz/util/error"
	"net/http"
	"os"

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
// @Failure 409 {object} error.errorResponse
// @Failure 500 {object} error.errorResponse
// @Failure default {object} error.errorResponse
// @Router       /auth/verify [GET]
func (handler *Handler) verifyEmail(ctx *gin.Context) {
	logrus := handler.logrus
	code := ctx.Query("code")
	username := ctx.Query("username")
	logrus.Infof("DONE: get : %s ,%s", code, username)
	count, err := handler.services.CheckDataExists(username, logrus)
	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}
	if count == 0 {
		error.NewHandlerErrorResponse(ctx, http.StatusConflict, "user data not exist", logrus)
		return
	}
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
// @Failure 409 {object} error.errorResponse
// @Failure 500 {object} error.errorResponse
// @Failure default {object} error.errorResponse
// @Router       /auth/resend [GET]
func (handler *Handler) resendCodeToEmail(ctx *gin.Context) {
	logrus := handler.logrus
	email := ctx.Query("email")
	username := ctx.Query("username")
	logrus.Infof("DONE: get : %s ,%s", email, username)
	count, err := handler.services.CheckDataExists(username, logrus)
	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}
	if count == 0 {
		error.NewHandlerErrorResponse(ctx, http.StatusConflict, "user data not exist", logrus)
		return
	}
	err = handler.services.SendMessageEmail(email, username, logrus)
	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}
	ctx.JSON(http.StatusOK, model.ResponseSuccess{Message: "DONE"})
}

// @Summary Upload Account Image
// @Description Upload Account Image
// @ID upload-image
// @Tags   Auth
// @Accept       json
// @Produce      json
// @Produce text/plain
// @Produce application/octet-stream
// @Produce image/png
// @Produce image/jpeg
// @Produce image/jpg
// @Param file formData file true "file"
// @Accept multipart/form-data
// @Success      200   {object}      model.ResponseSuccess
// @Failure 400,404 {object} error.errorResponse
// @Failure 409 {object} error.errorResponse
// @Failure 500 {object} error.errorResponse
// @Failure default {object} error.errorResponse
// @Router       /auth/sign-up [PATCH]
func (handler *Handler) uploadAccountImage(ctx *gin.Context) {

	file, header, err := ctx.Request.FormFile("file")
	if err != nil {
		ctx.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
		return
	}
	filename := header.Filename
	out, err := os.Create("assest/public" + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}
	filepath := "http://localhost:8080/assest/public/" + filename
	ctx.JSON(http.StatusOK, gin.H{"filepath": filepath})
}
