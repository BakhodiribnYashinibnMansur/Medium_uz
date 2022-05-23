package handler

import (
	"net/http"
	"strconv"

	"github.com/BakhodiribnYashinibnMansur/Medium_uz/model"
	"github.com/BakhodiribnYashinibnMansur/Medium_uz/util/error"

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
// @Router /api/ghost/post/get/{id} [GET]
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
	resp, err := handler.services.GetPostById(id, logrus)
	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}
	ctx.JSON(http.StatusOK, model.ResponseSuccess{Data: resp, Message: "DONE"})

}

// @Summary Upload Post Image
// @Description Upload Post Image
// @ID upload-image-post
// @Tags   Post
// @Accept       json
// @Produce      json
// @Produce application/octet-stream
// @Produce image/png
// @Produce image/jpeg
// @Produce image/jpg
// @Param file formData file true "file"
// @Param id query int false "id"
// @Accept multipart/form-data
// @Success      200   {object}      model.ResponseSuccess
// @Failure 400,404 {object} error.errorResponse
// @Failure 409 {object} error.errorResponse
// @Failure 500 {object} error.errorResponse
// @Failure default {object} error.errorResponse
// @Router   /api/post/upload-image [PATCH]
//@Security ApiKeyAuth
func (handler *Handler) uploadImagePost(ctx *gin.Context) {
	logrus := handler.logrus
	userID, err := getUserId(ctx, logrus)
	if err != nil {
		return
	}
	id := ctx.Query("id")
	postID, err := strconv.Atoi(id)
	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
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

	effectedRowsNum, err := handler.services.UpdatePostImage(userID, postID, imageURL, logrus)
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

// @Summary Update  Post By ID
// @Tags Post
// @Description Update post by id
// @ID update-post-id
// @Accept  json
// @Produce  json
// @Param        id   query  int     true "Param ID"
// @Param input body model.Post true "post info"
// @Success 200 {object} model.ResponseSuccess
// @Failure 400,404 {object} error.errorResponse
// @Failure 409 {object} error.errorResponseData
// @Failure 500 {object} error.errorResponse
// @Failure default {object} error.errorResponse
// @Router /api/post/update [PUT]
//@Security ApiKeyAuth
func (handler *Handler) updatePost(ctx *gin.Context) {
	logrus := handler.logrus
	userID, err := getUserId(ctx, logrus)
	if err != nil {
		return
	}
	paramID := ctx.Query("id")
	if paramID == "" {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, "Param is empty", logrus)
		return
	}
	postID, err := strconv.Atoi(paramID)
	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}
	var input model.Post
	err = ctx.BindJSON(&input)

	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}

	result, err := handler.services.UpdatePost(userID, postID, input, logrus)
	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}
	if result == 0 {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, "NOT UPDATED", logrus)
		return
	}

	ctx.JSON(http.StatusOK, model.ResponseSuccess{Message: "Uploaded", Data: postID})
}

// @Summary Delete  Post By ID
// @Tags Post
// @Description Delete post by id
// @ID delete-post-id
// @Accept  json
// @Produce  json
// @Param        id   query  int     true "Param ID"
// @Success 200 {object} model.ResponseSuccess
// @Failure 400,404 {object} error.errorResponse
// @Failure 409 {object} error.errorResponseData
// @Failure 500 {object} error.errorResponse
// @Failure default {object} error.errorResponse
// @Router /api/post/delete [DELETE]
//@Security ApiKeyAuth
func (handler *Handler) deletePost(ctx *gin.Context) {

	logrus := handler.logrus
	userID, err := getUserId(ctx, logrus)
	if err != nil {
		return
	}
	paramID := ctx.Query("id")
	if paramID == "" {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, "Param is empty", logrus)
		return
	}
	postID, err := strconv.Atoi(paramID)
	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}
	deletePostResult, deletePostUserResult, err := handler.services.DeletePost(userID, postID, logrus)
	if err != nil {
		if err != nil {
			error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
			return
		}
	}
	if deletePostResult == 0 || deletePostUserResult == 0 {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, "Not Deleted", logrus)
		return
	}
	ctx.JSON(http.StatusOK, model.ResponseSuccess{Message: "Deleted", Data: postID})
}

// @Summary Like  Post By ID
// @Tags Post
// @Description Like post by id
// @ID like-post-id
// @Accept  json
// @Produce  json
// @Param        id   query  int     true "Param ID"
// @Success 200 {object} model.ResponseSuccess
// @Failure 400,404 {object} error.errorResponse
// @Failure 409 {object} error.errorResponseData
// @Failure 500 {object} error.errorResponse
// @Failure default {object} error.errorResponse
// @Router /api/post/like [GET]
//@Security ApiKeyAuth
func (handler *Handler) likePost(ctx *gin.Context) {

	logrus := handler.logrus
	userID, err := getUserId(ctx, logrus)
	if err != nil {
		return
	}
	paramID := ctx.Query("id")
	if paramID == "" {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, "Param is empty", logrus)
		return
	}
	postID, err := strconv.Atoi(paramID)
	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}

	result, err := handler.services.LikePost(userID, postID, logrus)
	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}
	ctx.JSON(http.StatusOK, model.ResponseSuccess{Message: "Liked Post", Data: result})
}
