package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nullableocean/golang-todo/internal/models"
	"net/http"
	"strconv"
)

type getAllTasksResponse struct {
	Data []models.Task `json:"data"`
}

// @Summary Get tasks
// @Secure ApiKeyAuth
// @Description "Get all tasks from todo list"
// @Tags Tasks
// @Accept json
// @Produce json
// @Param id path integer true "list id"
// @Success 201 {array} models.Task
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/lists/{id}/tasks [get]
func (h *Handler) getAllTasks(ctx *gin.Context) {
	userId, err := h.getUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	listId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	tasks, err := h.services.TodoTask.GetAll(userId, listId)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, getAllTasksResponse{Data: tasks})
}

// @Summary Get task
// @Secure ApiKeyAuth
// @Description "Get task by id"
// @Tags Tasks
// @Accept json
// @Produce json
// @Param id path integer true "task id"
// @Success 201 {object} models.Task
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/tasks/{id} [get]
func (h *Handler) getTaskById(ctx *gin.Context) {
	userId, err := h.getUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	taskId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	task, err := h.services.TodoTask.GetTaskById(userId, taskId)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, task)
}

// @Summary Create task
// @Secure ApiKeyAuth
// @Description "Create task"
// @Tags Tasks
// @Accept json
// @Produce json
// @Param id path integer true "list id"
// @Param id body models.Task true "task data"
// @Success 201 {int} int "id"
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/lists/{id}/tasks [post]
func (h *Handler) createTask(ctx *gin.Context) {
	var input models.Task

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
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	taskId, err := h.services.TodoTask.Create(userId, listId, input)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, map[string]interface{}{
		"id": taskId,
	})
}

// @Summary Update task
// @Secure ApiKeyAuth
// @Description "Update task"
// @Tags Tasks
// @Accept json
// @Produce json
// @Param id path integer true "task id"
// @Param id body models.TaskUpdateInput true "task update data"
// @Success 201 {string} string "status"
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/tasks/{id} [put]
func (h *Handler) updateTask(ctx *gin.Context) {
	var input models.TaskUpdateInput

	err := ctx.BindJSON(&input)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	userId, err := h.getUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	taskId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.services.TodoTask.Update(userId, taskId, input)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, StatusResponse{"ok"})
}

// @Summary Delete task
// @Secure ApiKeyAuth
// @Description "Delete task"
// @Tags Tasks
// @Accept json
// @Produce json
// @Param id path integer true "task id"
// @Success 201 {string} string "status"
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/tasks/{id} [delete]
func (h *Handler) deleteTask(ctx *gin.Context) {
	userId, err := h.getUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	taskId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.services.TodoTask.Delete(userId, taskId)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, StatusResponse{"ok"})
}
