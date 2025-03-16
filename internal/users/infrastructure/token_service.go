package users_infrastructure

import (
	"fmt"
	"time"

	users_domain "github.com/enteresanlikk/go-modular-monolith/internal/users/domain"
	"github.com/golang-jwt/jwt/v5"
)

type tokenService struct {
	config users_domain.TokenConfig
}

func NewTokenService(config users_domain.TokenConfig) users_domain.TokenService {
	return &tokenService{
		config: config,
	}
}

func (s *tokenService) GenerateTokenPair(claims map[string]interface{}) (*users_domain.TokenPair, error) {
	now := time.Now()
	accessTokenExpiry := now.Add(s.config.AccessTokenDuration)
	refreshTokenExpiry := now.Add(s.config.RefreshTokenDuration)

	// Generate access token
	tokenClaims := &users_domain.TokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(accessTokenExpiry),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			Subject:   claims["userId"].(string),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)
	accessTokenString, err := accessToken.SignedString(s.config.AccessTokenSecret)
	if err != nil {
		return nil, fmt.Errorf("failed to sign access token: %w", err)
	}

	// Generate refresh token
	refreshTokenClaims := &users_domain.TokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(refreshTokenExpiry),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			Subject:   claims["userId"].(string),
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	refreshTokenString, err := refreshToken.SignedString(s.config.RefreshTokenSecret)
	if err != nil {
		return nil, fmt.Errorf("failed to sign refresh token: %w", err)
	}

	return &users_domain.TokenPair{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
		ExpiresAt:    accessTokenExpiry.Unix(),
	}, nil
}

func (s *tokenService) ValidateToken(tokenString string, tokenType users_domain.TokenType) (*users_domain.TokenClaims, error) {
	secret := s.config.AccessTokenSecret
	if tokenType == users_domain.RefreshToken {
		secret = s.config.RefreshTokenSecret
	}

	token, err := jwt.ParseWithClaims(tokenString, &users_domain.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	if claims, ok := token.Claims.(*users_domain.TokenClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}

func (s *tokenService) ParseToken(tokenString string) (*users_domain.TokenClaims, error) {
	// Try with access token first
	if claims, err := s.ValidateToken(tokenString, users_domain.AccessToken); err == nil {
		return claims, nil
	}

	// Try with refresh token if access token validation fails
	return s.ValidateToken(tokenString, users_domain.RefreshToken)
}
