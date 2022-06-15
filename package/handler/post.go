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
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
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
// @Param        id   query  int     true "Param ID"
// @Success 200 {object} model.ResponseSuccess
// @Failure 400,404 {object} error.errorResponse
// @Failure 409 {object} error.errorResponseData
// @Failure 500 {object} error.errorResponse
// @Failure default {object} error.errorResponse
// @Router /api/ghost/post/get-post [GET]
func (handler *Handler) getPostID(ctx *gin.Context) {
	logrus := handler.logrus
	paramID := ctx.Query("id")
	if paramID == "" {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, "ID is empty", logrus)
		return
	}
	id, err := strconv.Atoi(paramID)
	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}
	resp, err := handler.services.GetPostByIdWithoutBody(id, logrus)
	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}
	ctx.JSON(http.StatusOK, model.ResponseSuccess{Data: resp, Message: "DONE"})
}

// @Summary Get  Post Body By ID
// @Tags Post
// @Description get post body by id
// @ID get-post-body-id
// @Accept  json
// @Produce  json
// @Param        id   query  int     true "Param ID"
// @Success 200 {object} model.ResponseSuccess
// @Failure 400,404 {object} error.errorResponse
// @Failure 409 {object} error.errorResponseData
// @Failure 500 {object} error.errorResponse
// @Failure default {object} error.errorResponse
// @Router /api/ghost/post/get-body [GET]
func (handler *Handler) getPostBodyID(ctx *gin.Context) {
	logrus := handler.logrus
	paramID := ctx.Query("id")
	if paramID == "" {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, "ID is empty", logrus)
		return
	}
	id, err := strconv.Atoi(paramID)
	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}
	resp, err := handler.services.GetPostBodyById(id, logrus)
	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}
	ctx.JSON(http.StatusOK, model.ResponseSuccess{Data: resp.PostBody, Message: "DONE"})
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
		if err != nil {
			error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
			return
		}
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
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, "NOT UPDATE ", logrus)
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
		if err != nil {
			error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
			return
		}
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
		if err != nil {
			error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
			return
		}
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
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
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
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
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
	if result == 0 {
		if err != nil {
			error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, "DO NOT WORK", logrus)
			return
		}
	}
	ctx.JSON(http.StatusOK, model.ResponseSuccess{Message: "Liked Post", Data: result})
}

// @Summary Commit  Post
// @Tags Post
// @Description Commit post by user
// @ID commit-post-id
// @Accept  json
// @Produce  json
// @Param input body model.CommitPost true "commit info"
// @Success 200 {object} model.ResponseSuccess
// @Failure 400,404 {object} error.errorResponse
// @Failure 409 {object} error.errorResponseData
// @Failure 500 {object} error.errorResponse
// @Failure default {object} error.errorResponse
// @Router /api/post/commit [POST]
//@Security ApiKeyAuth
func (handler *Handler) commitPost(ctx *gin.Context) {
	logrus := handler.logrus

	var input model.CommitPost
	err := ctx.BindJSON(&input)
	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}

	userID, err := getUserId(ctx, logrus)
	if err != nil {
		if err != nil {
			error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
			return
		}
		return
	}

	input.ReaderID = userID
	commitID, err := handler.services.CommitPost(input, logrus)
	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}
	if commitID == 0 {
		if err != nil {
			error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, "DO NOT WORK", logrus)
			return
		}
	}
	ctx.JSON(http.StatusOK, model.ResponseSuccess{Message: " Committed Post Post", Data: commitID})

}

// @Summary View  Post By ID
// @Tags Post
// @Description View post by id
// @ID view-post-id
// @Accept  json
// @Produce  json
// @Param        id   query  int     true "Param ID"
// @Success 200 {object} model.ResponseSuccess
// @Failure 400,404 {object} error.errorResponse
// @Failure 409 {object} error.errorResponseData
// @Failure 500 {object} error.errorResponse
// @Failure default {object} error.errorResponse
// @Router /api/post/view [GET]
func (handler *Handler) viewPost(ctx *gin.Context) {
	logrus := handler.logrus
	userID, err := getUserId(ctx, logrus)
	if err != nil {
		if err != nil {
			error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
			return
		}
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

	result, err := handler.services.ViewPost(userID, postID, logrus)
	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}
	if result == 0 {
		if err != nil {
			error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, "DO NOT WORK", logrus)
			return
		}
	}
	ctx.JSON(http.StatusOK, model.ResponseSuccess{Message: " View Post", Data: result})

}

// @Summary Rated  Post By ID
// @Tags Post
// @Description Rated post by id
// @ID rated-post-id
// @Accept  json
// @Produce  json
// @Param        postID   query  int     true "Param ID"
// @Param        rating   query  int     true "Rating Number"
// @Success 200 {object} model.ResponseSuccess
// @Failure 400,404 {object} error.errorResponse
// @Failure 409 {object} error.errorResponseData
// @Failure 500 {object} error.errorResponse
// @Failure default {object} error.errorResponse
// @Router /api/post/rating [GET]
//@Security ApiKeyAuth
func (handler *Handler) ratedPost(ctx *gin.Context) {
	logrus := handler.logrus
	userID, err := getUserId(ctx, logrus)

	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}

	postIDQuery := ctx.Query("postID")

	if postIDQuery == "" {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, "Param ID is empty", logrus)
		return
	}

	paramRating := ctx.Query("rating")

	if paramRating == "" {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, "Param Rating is empty", logrus)
		return
	}

	postID, err := strconv.Atoi(postIDQuery)
	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}

	userRating, err := strconv.Atoi(paramRating)
	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}

	result, err := handler.services.RatingPost(userID, postID, userRating, logrus)
	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}
	if result == 0 {
		if err != nil {
			error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, "DO NOT WORK", logrus)
			return
		}
	}
	ctx.JSON(http.StatusOK, model.ResponseSuccess{Message: " Rating Post", Data: ""})

}

// @Summary Get Commit
// @Tags Post
// @Description get commits
// @ID get-commits
// @Accept  json
// @Produce  json
// @Param        offset   query  int     false "Offset "
// @Param        limit   query  int     false "Limit "
// @Param        postID   query  int     true "postID "
// @Success 200 {object} model.ResponseSuccess
// @Failure 400,404 {object} error.errorResponse
// @Failure 409 {object} error.errorResponseData
// @Failure 500 {object} error.errorResponse
// @Failure default {object} error.errorResponse
// @Router /api/ghost/post/get-commit [GET]
func (handler *Handler) getCommits(ctx *gin.Context) {
	logrus := handler.logrus
	var pagination model.Pagination
	postIDQuery := ctx.Query("postID")
	if postIDQuery == "" {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, "Offset is empty", logrus)
		return
	}

	postID, err := strconv.Atoi(postIDQuery)
	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}
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
	result, err := handler.services.GetCommitPost(postID, pagination, logrus)
	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}
	ctx.JSON(http.StatusOK, model.ResponseSuccess{Message: "Post Commits", Data: result})

}

// @Summary Get User Posts
// @Tags Post
// @Description Get User Posts
// @ID get-user-post
// @Accept  json
// @Produce  json
// @Param        offset   query  int     false "Offset "
// @Param        limit   query  int     false "Limit "
// @Param       userID   query  int     true "userID "
// @Success 200 {object} model.ResponseSuccess
// @Failure 400,404 {object} error.errorResponse
// @Failure 409 {object} error.errorResponseData
// @Failure 500 {object} error.errorResponse
// @Failure default {object} error.errorResponse
// @Router /api/ghost/post/get-user-post [GET]
func (handler *Handler) getUserPost(ctx *gin.Context) {
	logrus := handler.logrus
	var pagination model.Pagination
	userIDQuery := ctx.Query("userID")
	if userIDQuery == "" {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, "Offset is empty", logrus)
		return
	}

	userID, err := strconv.Atoi(userIDQuery)
	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}
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
	result, err := handler.services.GetUserPost(userID, pagination, logrus)
	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}
	ctx.JSON(http.StatusOK, model.ResponseSuccess{Message: "User Posts Body", Data: result})

}

// @Summary Get Most Viewed
// @Tags Post
// @Description Get Most Viewed
// @ID get-most-viewed-post
// @Accept  json
// @Produce  json
// @Param        offset   query  int     false "Offset "
// @Param        limit   query  int     false "Limit "
// @Success 200 {object} model.ResponseSuccess
// @Failure 400,404 {object} error.errorResponse
// @Failure 409 {object} error.errorResponseData
// @Failure 500 {object} error.errorResponse
// @Failure default {object} error.errorResponse
// @Router /api/ghost/post/get-most-viewed [GET]
func (handler *Handler) getMostViewed(ctx *gin.Context) {
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
	result, err := handler.services.GetMostViewed(pagination, logrus)
	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}
	ctx.JSON(http.StatusOK, model.ResponseSuccess{Message: "User Posts Body", Data: result})

}

// @Summary Get Most Liked
// @Tags Post
// @Description Get Most Liked
// @ID get-most-liked-post
// @Accept  json
// @Produce  json
// @Param        offset   query  int     false "Offset "
// @Param        limit   query  int     false "Limit "
// @Success 200 {object} model.ResponseSuccess
// @Failure 400,404 {object} error.errorResponse
// @Failure 409 {object} error.errorResponseData
// @Failure 500 {object} error.errorResponse
// @Failure default {object} error.errorResponse
// @Router /api/ghost/post/get-most-liked [GET]
func (handler *Handler) getMostLiked(ctx *gin.Context) {
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
	result, err := handler.services.GetMostLiked(pagination, logrus)
	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}
	ctx.JSON(http.StatusOK, model.ResponseSuccess{Message: "User Posts Body", Data: result})

}

// @Summary Get Most Rating
// @Tags Post
// @Description Get Most Rating
// @ID get-most-rated-post
// @Accept  json
// @Produce  json
// @Param        offset   query  int     false "Offset "
// @Param        limit   query  int     false "Limit "
// @Success 200 {object} model.ResponseSuccess
// @Failure 400,404 {object} error.errorResponse
// @Failure 409 {object} error.errorResponseData
// @Failure 500 {object} error.errorResponse
// @Failure default {object} error.errorResponse
// @Router /api/ghost/post/get-most-rated [GET]
func (handler *Handler) getMostRated(ctx *gin.Context) {
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
	result, err := handler.services.GetMostRated(pagination, logrus)
	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}
	ctx.JSON(http.StatusOK, model.ResponseSuccess{Message: "User Posts Body", Data: result})

}

// @Summary Get Resent Post
// @Tags Post
// @Description Get Resent Post
// @ID get-resent-post
// @Accept  json
// @Produce  json
// @Param        offset   query  int     false "Offset "
// @Param        limit   query  int     false "Limit "
// @Success 200 {object} model.ResponseSuccess
// @Failure 400,404 {object} error.errorResponse
// @Failure 409 {object} error.errorResponseData
// @Failure 500 {object} error.errorResponse
// @Failure default {object} error.errorResponse
// @Router /api/ghost/post/resent [GET]
func (handler *Handler) getResentPost(ctx *gin.Context) {
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
	result, err := handler.services.GetResentPost(pagination, logrus)
	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
		return
	}
	ctx.JSON(http.StatusOK, model.ResponseSuccess{Message: "User Posts Body", Data: result})

}
