package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authHeader = "Authorization"
	userCtx    = "user_id"
)

func (h *Handler) identifyUser(ctx *gin.Context) {
	header := ctx.GetHeader(authHeader)
	if header == "" {
		newErrorResponse(ctx, http.StatusBadRequest, "empty authorization header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid authorization header")
		return
	}

	tokenString := headerParts[1]
	userId, err := h.services.Authorization.ParseToken(tokenString)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.Set(userCtx, userId)
}

func (h *Handler) getUserId(ctx *gin.Context) (int, error) {
	userId, ok := ctx.Get(userCtx)
	if !ok {
		errorM := "invalid request. Need valid token"
		newErrorResponse(ctx, http.StatusBadRequest, errorM)
		return 0, errors.New(errorM)
	}

	id, ok := userId.(int)
	if !ok {
		errorM := "invalid user_id"
		newErrorResponse(ctx, http.StatusBadRequest, errorM)
		return 0, errors.New(errorM)
	}

	return id, nil
}
