package v1

import (
	"net/http"

	"errors"

	"github.com/gin-gonic/gin"
	"github.com/railgodev/denet-test/internal/apperr"

	"github.com/railgodev/denet-test/internal/api/v1/response"
)

func (h *handler) GetStatus(c *gin.Context) {
	id := c.Param("id")
	role, exists := c.Get("role")
	if !exists || role != "admin" {
		idFromClaims, exists := c.Get("user_id")
		if !exists || idFromClaims != id {
			response.WriteError(c, http.StatusForbidden, apperr.ErrForbidden)
			h.log.Debug(id, role)

			return
		}
	}

	status, err := h.uc.Status(c.Request.Context(), id)
	if err != nil {
		switch {
		case errors.Is(err, apperr.ErrInvalidIDFormat):
			response.WriteError(c, http.StatusBadRequest, err)
		case errors.Is(err, apperr.ErrUserNotFound):
			response.WriteError(c, http.StatusNotFound, err)
		default:
			response.WriteError(c, http.StatusInternalServerError, apperr.ErrGetStatus)
		}
		h.log.Error(err.Error())
		return
	}
	response.Write(c, http.StatusOK, status)
}
