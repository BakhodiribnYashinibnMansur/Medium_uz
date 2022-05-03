package handler

import (
	"mediumuz/model"
	"mediumuz/util/error"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Create Post
// @Tags Post
// @Description create post
// @ID create-post
// @Accept  json
// @Produce  json
// @Param input body model.Post true "post info"
// @Success 200 {object} model.ResponseSign
// @Failure 400,404 {object} error.errorResponse
// @Failure 409 {object} error.errorResponseData
// @Failure 500 {object} error.errorResponse
// @Failure default {object} error.errorResponse
// @Router /api/post/create [post]
func (handler *Handler) createPost(ctx *gin.Context) {
	logrus := handler.logrus
	var input model.Post
	err := ctx.BindJSON(&input)
	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}
	id, err := getUserId(ctx, logrus)
	if err != nil {
		return
	}
	logrus.Infof("User id : %s", id)
}
