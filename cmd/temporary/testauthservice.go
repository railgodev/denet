// cmd/genjwt/main.go
package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Claims struct {
	UserID    string `json:"user_id,omitempty"`
	Role      string `json:"role,omitempty"`
	TokenType string `json:"token_type,omitempty"`
	jwt.RegisteredClaims
}

// CreateAccessToken создает HS256 токен с TTL
func CreateAccessToken(secret []byte, userID, role string, ttl time.Duration) (string, *Claims, error) {
	now := time.Now().UTC()
	claims := &Claims{
		UserID:    userID,
		Role:      role,
		TokenType: "access",
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(ttl)),
			ID:        uuid.NewString(),
			Issuer:    "your-service",
			Subject:   userID,
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := t.SignedString(secret)
	if err != nil {
		return "", nil, err
	}
	return signed, claims, nil
}

func main() {
	// декодируем Base64 secret в 64 байта
	secret, err := base64.StdEncoding.DecodeString("Qjd8p2I7wrWqQLpEjEm5ud3elczndQUV2nnszACGREDWki/b6uqoYU9yhhw1kNR5TZtmiOpz2nn0IuQE/uqSiQ==")
	if err != nil {
		log.Fatalf("failed to decode JWT_SECRET from base64: %v", err)
	}

	userID := "dd74ba03-5608-4ad2-b251-b0e87843d0a1"
	role := "admin"
	ttl := time.Hour

	tokenStr, claims, err := CreateAccessToken(secret, userID, role, ttl)
	if err != nil {
		log.Fatalf("failed to create token: %v", err)
	}

	// Печатаем только токен (для Authorization header)
	fmt.Println(tokenStr)

	// Печатаем отладочную информацию в stderr
	fmt.Fprintf(os.Stderr, "issued at: %v\nexpires at: %v\njti: %s\n",
		claims.IssuedAt.Time.Format(time.RFC3339),
		claims.ExpiresAt.Time.Format(time.RFC3339),
		claims.ID,
	)
}
