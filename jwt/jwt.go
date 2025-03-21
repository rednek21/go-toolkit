package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

type Claims struct {
	Username    string   `json:"username"`
	Roles       []string `json:"roles"`
	Permissions []string `json:"permissions"`
	jwt.RegisteredClaims
}

type Manager struct {
	AccessSecret  string
	RefreshSecret string
	AccessTTL     time.Duration
	RefreshTTL    time.Duration
}

func NewManager(AccessSecret, RefreshSecret string, AccessTTL, RefreshTTL time.Duration) *Manager {
	return &Manager{
		AccessSecret:  AccessSecret,
		RefreshSecret: RefreshSecret,
		AccessTTL:     AccessTTL,
		RefreshTTL:    RefreshTTL,
	}
}

func (m *Manager) GenerateTokenPair(username string, roles, permissions []string) (string, string, error) {
	now := time.Now()

	accessClaims := Claims{
		Username:    username,
		Roles:       roles,
		Permissions: permissions,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(m.AccessTTL)),
			ID:        uuid.New().String(),
		},
	}
	accessToken, err := m.generateToken(m.AccessSecret, accessClaims)
	if err != nil {
		return "", "", err
	}

	refreshClaims := Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(m.RefreshTTL)),
			ID:        uuid.New().String(),
		},
	}
	refreshToken, err := m.generateToken(m.RefreshSecret, refreshClaims)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (m *Manager) ParseToken(tokenString string, isRefresh bool) (*Claims, error) {
	secret := []byte(m.AccessSecret)
	if isRefresh {
		secret = []byte(m.RefreshSecret)
	}

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return secret, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

func (m *Manager) generateToken(secret string, claims Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
