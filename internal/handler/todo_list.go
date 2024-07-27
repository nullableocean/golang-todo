package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nullableocean/golang-todo/internal/models"
	"net/http"
	"strconv"
)

type getAllListResponse struct {
	Data []models.TodoList `json:"data"`
}

// @Summary Get todo lists
// @Secure ApiKeyAuth
// @Description "Get all todo lists"
// @Tags Lists
// @Accept json
// @Produce json
// @Success 201 {array} models.TodoList
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/lists [get]
func (h *Handler) getAllLists(ctx *gin.Context) {
	userId, err := h.getUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	lists, err := h.services.TodoList.GetAll(userId)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, getAllListResponse{Data: lists})
}

// @Summary Get list
// @Secure ApiKeyAuth
// @Description "Get list by id"
// @Tags Lists
// @Accept json
// @Produce json
// @Param id path integer true "list id"
// @Success 201 {object} models.TodoList
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/lists/{id} [get]
func (h *Handler) getListById(ctx *gin.Context) {
	userId, err := h.getUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	listId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid id")
		return
	}

	list, err := h.services.TodoList.GetListById(userId, listId)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, list)
}

// @Summary Create list
// @Secure ApiKeyAuth
// @Description "Create todo list"
// @Tags Lists
// @Accept json
// @Produce json
// @Param id body models.TodoList true "todo list data"
// @Success 201 {int} integer "id"
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/lists [post]
func (h *Handler) createList(ctx *gin.Context) {
	var input models.TodoList

	err := ctx.BindJSON(&input)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid input data")
		return
	}

	userId, err := h.getUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	listId, err := h.services.TodoList.Create(userId, input)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, map[string]interface{}{
		"id": listId,
	})
}

// @Summary Update list
// @Secure ApiKeyAuth
// @Description "Update list"
// @Tags Lists
// @Accept json
// @Produce json
// @Param id path integer true "list id"
// @Param id body models.TodoListUpdateInput true "list update data"
// @Success 201 {string} string "status"
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/lists [put]
func (h *Handler) updateList(ctx *gin.Context) {
	var input models.TodoListUpdateInput

	err := ctx.BindJSON(&input)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid input data")
		return
	}

	userId, err := h.getUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	listId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid id")
		return
	}

	err = h.services.TodoList.Update(userId, listId, input)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, StatusResponse{"ok"})
}

// @Summary Delete list
// @Secure ApiKeyAuth
// @Description "Delete list"
// @Tags Lists
// @Accept json
// @Produce json
// @Param id path integer true "list id"
// @Success 201 {string} string "status"
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/lists/{id} [delete]
func (h *Handler) deleteList(ctx *gin.Context) {
	userId, err := h.getUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	listId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid id")
		return
	}

	err = h.services.TodoList.Delete(userId, listId)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, StatusResponse{"ok"})

}
