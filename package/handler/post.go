package handler

import (
	"mediumuz/model"
	"mediumuz/util/error"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Create Post
// @Tags Post
// @Description create post
// @ID create-post
// @Accept  json
// @Produce  json
// @Param input body model.Post true "post info"
// @Success 200 {object} model.ResponseSuccess
// @Failure 400,404 {object} error.errorResponse
// @Failure 409 {object} error.errorResponseData
// @Failure 500 {object} error.errorResponse
// @Failure default {object} error.errorResponse
// @Router /api/post/create [post]
//@Security ApiKeyAuth
func (handler *Handler) createPost(ctx *gin.Context) {
	logrus := handler.logrus
	var input model.Post
	err := ctx.BindJSON(&input)
	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}

	userId, err := getUserId(ctx, logrus)
	if err != nil {
		return
	}

	postId, err := handler.services.CreatePost(userId, input, logrus)
	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}
	ctx.JSON(http.StatusOK, model.ResponseSuccess{Data: postId, Message: "DONE"})
}

// @Summary Get  Post By ID
// @Tags Post
// @Description get post by id
// @ID get-post-id
// @Accept  json
// @Produce  json
// @Param        id   path  int     true "Param ID"
// @Success 200 {object} model.ResponseSuccess
// @Failure 400,404 {object} error.errorResponse
// @Failure 409 {object} error.errorResponseData
// @Failure 500 {object} error.errorResponse
// @Failure default {object} error.errorResponse
// @Router /api/post/get/{id} [GET]
//@Security ApiKeyAuth
func (handler *Handler) getPostID(ctx *gin.Context) {
	logrus := handler.logrus
	paramID := ctx.Param("id")
	if paramID == "" {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, "Param is empty", logrus)
		return
	}
	id, err := strconv.Atoi(paramID)
	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}
	check, err := handler.services.CheckPostId(id, logrus)
	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}
	if check == 0 {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, "ID Not Fount", logrus)
		return
	}
	resp, err := handler.services.GetPostById(id, logrus)
	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}
	ctx.JSON(http.StatusOK, model.ResponseSuccess{Data: resp, Message: "DONE"})

}
