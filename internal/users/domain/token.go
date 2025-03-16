package usersDomain

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// TokenType represents the type of token (access or refresh)
type TokenType string

const (
	AccessToken  TokenType = "access"
	RefreshToken TokenType = "refresh"
)

// TokenClaims represents the custom claims for our tokens
type TokenClaims struct {
	jwt.RegisteredClaims
}

// TokenConfig represents token configuration
type TokenConfig struct {
	AccessTokenDuration  time.Duration
	RefreshTokenDuration time.Duration
	AccessTokenSecret    []byte
	RefreshTokenSecret   []byte
}

// TokenPair represents a pair of access and refresh tokens
type TokenPair struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	ExpiresAt    int64  `json:"expiresAt"`
}

// TokenService defines the interface for token operations
type TokenService interface {
	GenerateTokenPair(claims map[string]interface{}) (*TokenPair, error)
	ValidateToken(tokenString string, tokenType TokenType) (*TokenClaims, error)
	ParseToken(tokenString string) (*TokenClaims, error)
}
