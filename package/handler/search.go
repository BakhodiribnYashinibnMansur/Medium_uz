package handler

import (
	"net/http"
	"strconv"

	"github.com/BakhodiribnYashinibnMansur/Medium_uz/model"
	"github.com/BakhodiribnYashinibnMansur/Medium_uz/util/error"
	"github.com/gin-gonic/gin"
)

// @Summary Search  Post By search text
// @Tags Search
// @Description Search post by search text
// @ID search-universal
// @Accept  json
// @Produce  json
// @Param        offset   query  int     false "Offset "
// @Param        limit   query  int     false "Limit "
// @Param        search   query  string     true "search text"
// @Success 200 {object} model.ResponseSuccess
// @Failure 400,404 {object} error.errorResponse
// @Failure 409 {object} error.errorResponseData
// @Failure 500 {object} error.errorResponse
// @Failure default {object} error.errorResponse
// @Router /api/ghost/search [GET]
func (handler *Handler) searchAll(ctx *gin.Context) {
	logrus := handler.logrus
	search := ctx.Query("search")
	if search == "" {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, "Search text is empty", logrus)
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
	resultSearch, err := handler.services.UniversalSearch(search, pagination, logrus)
	if err != nil {
		error.NewHandlerErrorResponse(ctx, http.StatusBadRequest, err.Error(), logrus)
	}
	ctx.JSON(http.StatusOK, model.ResponseSuccess{Message: "Search Result", Data: resultSearch})
}
