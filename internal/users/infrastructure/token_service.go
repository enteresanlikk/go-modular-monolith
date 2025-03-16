package usersInfrastructure

import (
	"fmt"
	"time"

	usersDomain "github.com/enteresanlikk/go-modular-monolith/internal/users/domain"
	"github.com/golang-jwt/jwt/v5"
)

type tokenService struct {
	config usersDomain.TokenConfig
}

func NewTokenService(config usersDomain.TokenConfig) usersDomain.TokenService {
	return &tokenService{
		config: config,
	}
}

func (s *tokenService) GenerateTokenPair(claims map[string]interface{}) (*usersDomain.TokenPair, error) {
	now := time.Now()
	accessTokenExpiry := now.Add(s.config.AccessTokenDuration)
	refreshTokenExpiry := now.Add(s.config.RefreshTokenDuration)

	tokenClaims := &usersDomain.TokenClaims{
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

	refreshTokenClaims := &usersDomain.TokenClaims{
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

	return &usersDomain.TokenPair{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
		ExpiresAt:    accessTokenExpiry.Unix(),
	}, nil
}

func (s *tokenService) ValidateToken(tokenString string, tokenType usersDomain.TokenType) (*usersDomain.TokenClaims, error) {
	secret := s.config.AccessTokenSecret
	if tokenType == usersDomain.RefreshToken {
		secret = s.config.RefreshTokenSecret
	}

	token, err := jwt.ParseWithClaims(tokenString, &usersDomain.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	if claims, ok := token.Claims.(*usersDomain.TokenClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}

func (s *tokenService) ParseToken(tokenString string) (*usersDomain.TokenClaims, error) {
	if claims, err := s.ValidateToken(tokenString, usersDomain.AccessToken); err == nil {
		return claims, nil
	}

	return s.ValidateToken(tokenString, usersDomain.RefreshToken)
}
