package auth

import (
	"encoding/base64"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/railgodev/denet-test/internal/api/v1/response"
)

type Claims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

// take out errors into variables
var (
	ErrTokenRequired        = errors.New("token required")
	ErrInvalidAuthHeader    = errors.New("invalid auth header")
	ErrInvalidToken         = errors.New("invalid token")
	ErrInvalidTokenClaims   = errors.New("invalid token claims")
	ErrTokenExpired         = errors.New("token expired")
	ErrInvalidSigningMethod = errors.New("invalid signing method")
)

func NewHMACMiddleware(secret []byte, log *slog.Logger) gin.HandlerFunc {

	return func(c *gin.Context) {
		auth := strings.TrimSpace(c.GetHeader("Authorization"))
		if auth == "" {
			response.WriteError(c, http.StatusUnauthorized, ErrTokenRequired)
			return
		}
		if !strings.HasPrefix(auth, "Bearer ") {
			response.WriteError(c, http.StatusUnauthorized, ErrInvalidAuthHeader)
			return
		}
		tokenStr := strings.TrimSpace(strings.TrimPrefix(auth, "Bearer "))
		SECRET, _ :=base64.StdEncoding.DecodeString(string(secret))
		token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(t *jwt.Token) (interface{}, error) {
			// Validate alg is HS256
			if t.Method.Alg() != jwt.SigningMethodHS256.Alg() {
				return nil, ErrInvalidSigningMethod
			}
			return SECRET, nil
		})
		log.Debug("parsed token", "token", token, "err", err)
		if err != nil {
			response.WriteError(c, http.StatusUnauthorized, fmt.Errorf(ErrInvalidToken.Error()+": %w", err))
			return
		}

		claims, ok := token.Claims.(*Claims)
		if !ok || !token.Valid {
			response.WriteError(c, http.StatusUnauthorized, ErrInvalidTokenClaims)
			return
		}

		// optional: check expiry explicitly
		if claims.ExpiresAt != nil && time.Now().After(claims.ExpiresAt.Time) {
			response.WriteError(c, http.StatusUnauthorized, ErrTokenExpired)
			return
		}

		// put claims into context
		c.Set("user_id", claims.UserID)
		c.Set("role", claims.Role)
		c.Next()
	}
}
