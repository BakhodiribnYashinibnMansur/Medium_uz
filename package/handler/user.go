package handler

import (
	"net/http"
	"strconv"

	"github.com/BakhodiribnYashinibnMansur/Medium_uz/model"
	"github.com/BakhodiribnYashinibnMansur/Medium_uz/util/error"

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
	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}
	effectedRowsNum, err := handler.services.VerifyCode(userId, user.Email, code, logrus)
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

// @Summary Send code for  Verification Email
// @Description send code to email for  verification
// @ID send-code-email
// @Tags   Account
// @Accept       json
// @Produce      json
// @Success      200   {object}      model.ResponseSuccess
// @Failure 400,404 {object} error.errorResponse
// @Failure 409 {object} error.errorResponse
// @Failure 500 {object} error.errorResponse
// @Failure default {object} error.errorResponse
// @Router       /api/account/sendcode [GET]
//@Security ApiKeyAuth
func (handler *Handler) sendCodeToEmail(ctx *gin.Context) {
	logrus := handler.logrus

	id, err := getUserId(ctx, logrus)
	if err != nil {
		return
	}
	userId := strconv.Itoa(id)
	user, err := handler.services.GetUserData(userId, logrus)
	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}
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
// @Router   /api/account/upload-image [PATCH]
//@Security ApiKeyAuth
func (handler *Handler) uploadAccountImage(ctx *gin.Context) {
	logrus := handler.logrus
	id, err := getUserId(ctx, logrus)

	if err != nil {
		return
	}

	ctx.Request.ParseMultipartForm(10 << 20)
	file, header, err := ctx.Request.FormFile("file")

	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}

	imageURL, err := handler.services.UploadImage(file, header, logrus)
	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}

	effectedRowsNum, err := handler.services.UpdateAccountImage(id, imageURL, logrus)
	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}

	if effectedRowsNum == 0 {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, "User not found", logrus)
		return
	}
	ctx.JSON(http.StatusOK, model.ResponseSuccess{Message: "Uploaded", Data: imageURL})
}

// @Summary Update Account
// @Description update account give data
// @ID update-account
// @Tags   Account
// @Accept       json
// @Produce      json
// @Param input body model.User false "account info"
// @Success      200   {object}      model.ResponseSuccess
// @Failure 400,404 {object} error.errorResponse
// @Failure 409 {object} error.errorResponse
// @Failure 500 {object} error.errorResponse
// @Failure default {object} error.errorResponse
// @Router       /api/account/update [PUT]
//@Security ApiKeyAuth
func (handler *Handler) updateAccount(ctx *gin.Context) {
	logrus := handler.logrus
	id, err := getUserId(ctx, logrus)
	if err != nil {
		return
	}

	var input model.User
	err = ctx.BindJSON(&input)

	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}

	checkedUser, err := handler.services.CheckDataExistsEmailNickName(input.Email, input.NickName, logrus)

	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}

	if !checkedUser.NickName || !checkedUser.Email {
		error.NewHandlerErrorResponseData(ctx, http.StatusConflict, "email or nickname already exist", checkedUser, logrus)
		return
	}
	logrus.Info("signUp data send for  check user Data to service")
	effectedRowsNum, err := handler.services.UpdateAccount(id, input, logrus)
	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}

	if effectedRowsNum == 0 {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, "User not found", logrus)
		return
	}
	ctx.JSON(http.StatusOK, model.ResponseSuccess{Message: "Updated"})
}

// @Summary Get Account Data
// @Description return account data. if you send id = "" or null return current user data. if send id = number return number user data .
// @ID get-account
// @Tags   Account
// @Accept       json
// @Produce      json
// @Param id query int false "id"
// @Success      200   {object}      model.ResponseSuccess
// @Failure 400,404 {object} error.errorResponse
// @Failure 409 {object} error.errorResponse
// @Failure 500 {object} error.errorResponse
// @Failure default {object} error.errorResponse
// @Router       /api/account/get [GET]
//@Security ApiKeyAuth
func (handler *Handler) getUser(ctx *gin.Context) {
	logrus := handler.logrus
	id := ctx.Query("id")
	authID, err := getUserId(ctx, logrus)
	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}
	userId := strconv.Itoa(authID)

	if id == "" {
		id = userId
	}

	user, err := handler.services.GetUserData(id, logrus)
	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}
	ctx.JSON(http.StatusOK, model.ResponseSuccess{Message: "GOT", Data: user})
}

// @Summary Add Following
// @Tags Account
// @Description Following user
// @ID following-account
// @Accept  json
// @Produce  json
// @Param        followingID   query  int     true "Param ID"
// @Success 200 {object} model.ResponseSuccess
// @Failure 400,404 {object} error.errorResponse
// @Failure 409 {object} error.errorResponseData
// @Failure 500 {object} error.errorResponse
// @Failure default {object} error.errorResponse
// @Router /api/account/following [GET]
// @Security ApiKeyAuth
func (handler *Handler) followingUser(ctx *gin.Context) {

	logrus := handler.logrus
	userID, err := getUserId(ctx, logrus)
	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}
	paramID := ctx.Query("followingID")

	if paramID == "" {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, "Param is empty", logrus)
		return
	}

	followingID, err := strconv.Atoi(paramID)
	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}
	result, err := handler.services.FollowingAccount(userID, followingID, logrus)
	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}
	if result == 0 {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, "NOT UPDATED", logrus)
		return
	}
	ctx.JSON(http.StatusOK, model.ResponseSuccess{Message: "Following User", Data: result})
}

// @Summary Add Follower
// @Tags Account
// @Description Follower user
// @ID follower-account
// @Accept  json
// @Produce  json
// @Param        followerID   query  int     true "Param ID"
// @Success 200 {object} model.ResponseSuccess
// @Failure 400,404 {object} error.errorResponse
// @Failure 409 {object} error.errorResponseData
// @Failure 500 {object} error.errorResponse
// @Failure default {object} error.errorResponse
// @Router /api/account/follower [GET]
// @Security ApiKeyAuth
func (handler *Handler) followerUser(ctx *gin.Context) {

	logrus := handler.logrus
	userID, err := getUserId(ctx, logrus)
	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}
	paramID := ctx.Query("followerID")

	if paramID == "" {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, "Param is empty", logrus)
		return
	}

	followerID, err := strconv.Atoi(paramID)
	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}
	result, err := handler.services.FollowerAccount(userID, followerID, logrus)
	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}
	if result == 0 {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, "NOT UPDATED", logrus)
		return
	}
	ctx.JSON(http.StatusOK, model.ResponseSuccess{Message: "Follower User", Data: result})
}

// @Summary Get Account Follower
// @Tags Account
// @Description Follower user
// @ID get-follower-account
// @Accept  json
// @Produce  json
// @Param        offset   query  int     false "Offset "
// @Param        limit   query  int     false "Limit "
// @Success 200 {object} model.ResponseSuccess
// @Failure 400,404 {object} error.errorResponse
// @Failure 409 {object} error.errorResponseData
// @Failure 500 {object} error.errorResponse
// @Failure default {object} error.errorResponse
// @Router /api/account/get-followers [GET]
// @Security ApiKeyAuth
func (handler *Handler) getFollowers(ctx *gin.Context) {

	logrus := handler.logrus
	userID, err := getUserId(ctx, logrus)
	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}
	var pagination model.Pagination
	offsetQuery := ctx.DefaultQuery("offset", "0")
	if offsetQuery == "" {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, "Param is empty", logrus)
		return
	}

	offset, err := strconv.Atoi(offsetQuery)
	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}
	limitQuery := ctx.DefaultQuery("limit", "10")

	if limitQuery == "" {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, "Param is empty", logrus)
		return
	}

	limit, err := strconv.Atoi(limitQuery)
	pagination.Offset = offset
	pagination.Limit = limit
	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}
	result, err := handler.services.GetFollowers(userID, pagination, logrus)
	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}
	ctx.JSON(http.StatusOK, model.ResponseSuccess{Message: "Follower User", Data: result})
}

// @Summary Get Account Interesting post
// @Tags Account
// @Description Get Account Interesting post
// @ID get-user-interesting
// @Accept  json
// @Produce  json
// @Param        offset   query  int     false "Offset "
// @Param        limit   query  int     false "Limit "
// @Param        tag   query  string     false "Tag "
// @Success 200 {object} model.ResponseSuccess
// @Failure 400,404 {object} error.errorResponse
// @Failure 409 {object} error.errorResponseData
// @Failure 500 {object} error.errorResponse
// @Failure default {object} error.errorResponse
// @Router /api/account/user-interesting [GET]
// @Security ApiKeyAuth
func (handler *Handler) getUserInterestingPost(ctx *gin.Context) {
	logrus := handler.logrus

	var pagination model.Pagination
	offsetQuery := ctx.DefaultQuery("offset", "0")
	if offsetQuery == "" {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, "Offset is empty", logrus)
		return
	}

	offset, err := strconv.Atoi(offsetQuery)
	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}
	limitQuery := ctx.DefaultQuery("limit", "10")

	if limitQuery == "" {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, "Limit is empty", logrus)
		return
	}

	limit, err := strconv.Atoi(limitQuery)
	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}
	pagination.Offset = offset
	pagination.Limit = limit

	tagQuery := ctx.Query("tag")

	if tagQuery == "" {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, "Tag is empty", logrus)
		return
	}
	result, err := handler.services.GetUserInterestingPost(tagQuery, pagination, logrus)
	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}
	ctx.JSON(http.StatusOK, model.ResponseSuccess{Message: "Post DONE User", Data: result})
}

// @Summary Get Account Following
// @Tags Account
// @Description Followings user
// @ID get-following-account
// @Accept  json
// @Produce  json
// @Param        offset   query  int     false "Offset "
// @Param        limit   query  int     false "Limit "
// @Success 200 {object} model.ResponseSuccess
// @Failure 400,404 {object} error.errorResponse
// @Failure 409 {object} error.errorResponseData
// @Failure 500 {object} error.errorResponse
// @Failure default {object} error.errorResponse
// @Router /api/account/get-followings [GET]
// @Security ApiKeyAuth
func (handler *Handler) getFollowings(ctx *gin.Context) {
	logrus := handler.logrus
	userID, err := getUserId(ctx, logrus)
	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}
	var pagination model.Pagination
	offsetQuery := ctx.DefaultQuery("offset", "0")
	if offsetQuery == "" {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, "Param is empty", logrus)
		return
	}

	offset, err := strconv.Atoi(offsetQuery)
	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}
	limitQuery := ctx.DefaultQuery("limit", "10")

	if limitQuery == "" {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, "Param is empty", logrus)
		return
	}

	limit, err := strconv.Atoi(limitQuery)
	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}
	pagination.Offset = offset
	pagination.Limit = limit
	result, err := handler.services.GetFollowings(userID, pagination, logrus)
	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}
	ctx.JSON(http.StatusOK, model.ResponseSuccess{Message: "Following User", Data: result})
}
