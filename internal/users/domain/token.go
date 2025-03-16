package usersDomain

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenType string

const (
	AccessToken  TokenType = "access"
	RefreshToken TokenType = "refresh"
)

type TokenClaims struct {
	jwt.RegisteredClaims
}

type TokenConfig struct {
	AccessTokenDuration  time.Duration
	RefreshTokenDuration time.Duration
	AccessTokenSecret    []byte
	RefreshTokenSecret   []byte
}

type TokenPair struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	ExpiresAt    int64  `json:"expiresAt"`
}

type TokenService interface {
	GenerateTokenPair(claims map[string]interface{}) (*TokenPair, error)
	ValidateToken(tokenString string, tokenType TokenType) (*TokenClaims, error)
	ParseToken(tokenString string) (*TokenClaims, error)
}
