package v1

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/railgodev/denet-test/internal/api/v1/request"
	"github.com/railgodev/denet-test/internal/api/v1/response"
	"github.com/railgodev/denet-test/internal/apperr"
)

func (h *handler) PostTaskComplete(c *gin.Context) {
	id := c.Param("id")
	role, exists := c.Get("role")
	if !exists || role != "admin" {
		idFromClaims, exists := c.Get("user_id")
		if !exists || idFromClaims != id {
			response.WriteError(c, http.StatusForbidden, apperr.ErrForbidden)
			return
		}
	}
	idFromClaims, exists := c.Get("user_id")
	if !exists || idFromClaims != id {
		response.WriteError(c, http.StatusForbidden, apperr.ErrForbidden)
		return
	}
	var req request.PostTaskComplete
	if err := c.ShouldBindJSON(&req); err != nil {
		response.WriteError(c, http.StatusBadRequest, apperr.ErrInvalidRequestBody)
		h.log.Error(err.Error())
		return
	}
	err := h.uc.TaskCompletion(c, id, req.TaskType)
	if err != nil {
		switch {
		case errors.Is(err, apperr.ErrUserNotFound):
			response.WriteError(c, http.StatusNotFound, err)
		case errors.Is(err, apperr.ErrTaskNotFound):
			response.WriteError(c, http.StatusBadRequest, err)
		default:
			response.WriteError(c, http.StatusInternalServerError, apperr.ErrGetStatus)
		}
		h.log.Error(err.Error())
		return
	}
	response.Write(c, http.StatusOK, nil)
}
