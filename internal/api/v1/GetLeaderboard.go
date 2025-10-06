package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/railgodev/denet-test/internal/apperr"
	"github.com/railgodev/denet-test/internal/api/v1/response"
)

func (h *handler) GetLeaderboard(c *gin.Context) {
	leaders, err := h.uc.Leaderboard(c.Request.Context())
	if err != nil {
		switch {
		default:
			response.WriteError(c, http.StatusInternalServerError, apperr.ErrGetLeaderboard)
		}
		h.log.Error("failed to get leaderboard", err)
		return
	}
	response.Write(c, http.StatusOK, leaders)
}