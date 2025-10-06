package v1

import (
	"net/http"

	"errors"

	"github.com/gin-gonic/gin"
	"github.com/railgodev/denet-test/internal/apperr"
	httpConverter "github.com/railgodev/denet-test/internal/api/v1/converter"
	"github.com/railgodev/denet-test/internal/api/v1/request"
	"github.com/railgodev/denet-test/internal/api/v1/response"
)

func (h *handler) PostReferrer(c *gin.Context) {
	id := c.Param("id")
	role, exists := c.Get("role")
	if !exists || role != "admin" {
		idFromClaims, exists := c.Get("user_id")
		if !exists || idFromClaims != id {
			response.WriteError(c, http.StatusForbidden, apperr.ErrForbidden)
			return
		}
	}
	var req request.PostReferrer
	if err := c.ShouldBindJSON(&req); err != nil {
		response.WriteError(c, http.StatusBadRequest, apperr.ErrInvalidRequestBody)
		h.log.Error(err.Error())
		return
	}
	usecaseReq := httpConverter.ReferrerToRepoModel(&req)
	err := h.uc.Referrer(c, id, usecaseReq)
	if err != nil {
		switch {
		case errors.Is(err, apperr.ErrUserNotFound):
			response.WriteError(c, http.StatusNotFound, err)
		case errors.Is(err, apperr.ErrReferrerNotFound):
			response.WriteError(c, http.StatusBadRequest, err)
		default:
			response.WriteError(c, http.StatusInternalServerError, apperr.ErrGetStatus)
		}
		h.log.Error(err.Error())
		return
	}
	response.Write(c, http.StatusOK, nil)
}
