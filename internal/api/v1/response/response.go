package response

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)
// TODO: tags issue
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func WriteError(c *gin.Context, httpStatus int, err error) {
	msg := err.Error()
	c.JSON(httpStatus, Error{
		Code:    httpStatus,
		Message: msg,
	})
}
func Write(c *gin.Context, httpStatus int, obj any) {
	c.JSON(httpStatus, obj)
}

type Leader struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Points   int    `json:"points"`
}

type Status struct {
	ID           uuid.UUID  `json:"id"`
	Username     string     `json:"username"`
	Email        string     `json:"email" binding:"required, email"`
	Points       int        `json:"points" binding:"gte=0"`
	ReferralCode *string    `json:"referral_code"`
	ReferredBy   *uuid.UUID `json:"referred_by"`
	CreatedAt    time.Time  `json:"created_at"`
}
