package handler

import (
	"fmt"
	"mediumuz/model"
	"mediumuz/util/error"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary  Verification Email
// @Description verification email with code
// @ID verify-email
// @Tags   Account
// @Accept       json
// @Produce      json
// @Param code query string false "code"
// @Success      200   {object}      model.ResponseSuccess
// @Failure 400,404 {object} error.errorResponse
// @Failure 409 {object} error.errorResponse
// @Failure 500 {object} error.errorResponse
// @Failure default {object} error.errorResponse
// @Router       /api/account/verify [GET]
//@Security ApiKeyAuth
func (handler *Handler) verifyEmail(ctx *gin.Context) {
	logrus := handler.logrus
	code := ctx.Query("code")

	id, err := getUserId(ctx, logrus)
	if err != nil {
		return
	}
	userId := strconv.Itoa(id)
	user, err := handler.services.GetUserData(userId, logrus)

	effectedRowsNum, err := handler.services.VerifyCode(userId, user.FirstName, code, logrus)
	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}

	if effectedRowsNum == 0 {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, "User not found", logrus)
		return
	}
	ctx.JSON(http.StatusOK, model.ResponseSuccess{Message: "DONE"})
}

// @Summary Resend cod for  Verification Email
// @Description resend code to email for  verification
// @ID resend-code-email
// @Tags   Account
// @Accept       json
// @Produce      json
// @Success      200   {object}      model.ResponseSuccess
// @Failure 400,404 {object} error.errorResponse
// @Failure 409 {object} error.errorResponse
// @Failure 500 {object} error.errorResponse
// @Failure default {object} error.errorResponse
// @Router       /api/account/resend [GET]
//@Security ApiKeyAuth
func (handler *Handler) resendCodeToEmail(ctx *gin.Context) {
	logrus := handler.logrus

	id, err := getUserId(ctx, logrus)
	if err != nil {
		return
	}
	userId := strconv.Itoa(id)
	user, err := handler.services.GetUserData(userId, logrus)
	logrus.Infof(user.Email)
	err = handler.services.SendMessageEmail(user.Email, user.FirstName, logrus)
	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}
	ctx.JSON(http.StatusOK, model.ResponseSuccess{Message: "DONE"})
}

// @Summary Upload Account Image
// @Description Upload Account Image
// @ID upload-image
// @Tags   Account
// @Accept       json
// @Produce      json
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
// @Router   /api/account/upload-image [POST]
//@Security ApiKeyAuth
func (handler *Handler) uploadAccountImage(ctx *gin.Context) {
	logrus := handler.logrus
	id, err := getUserId(ctx, logrus)
	if err != nil {
		return
	}
	userId := strconv.Itoa(id)
	ctx.Request.ParseMultipartForm(10 << 20)
	file, header, err := ctx.Request.FormFile("file")

	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
		return
	}
	user, err := handler.services.GetUserData(userId, logrus)
	filePath, err := handler.services.UploadAccountImage(file, header, user, logrus)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	effectedRowsNum, err := handler.services.UpdateAccountImage(id, filePath, logrus)
	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}

	if effectedRowsNum == 0 {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, "User not found", logrus)
		return
	}
	ctx.JSON(http.StatusOK, model.ResponseSuccess{Message: "Uploaded", Data: filePath})
}

func (handler *Handler) updateAccount(ctx *gin.Context) {

}

func (handler *Handler) getUser(ctx *gin.Context) {

}

func (handler *Handler) searchUser(ctx *gin.Context) {

}
