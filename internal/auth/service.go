package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Service struct {
	sessionJWTSecret []byte
	sessionDuration  time.Duration
}

func NewService(sessionJWTSecret string) *Service {
	return &Service{
		sessionJWTSecret: []byte(sessionJWTSecret),
		sessionDuration:  24 * time.Hour,
	}
}

type SessionPayload struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
}

type sessionClaims struct {
	SessionPayload *SessionPayload `json:"sub"`
	jwt.RegisteredClaims
}

func (s *Service) CreateSessionToken(userID, email string) (string, error) {
	claims := sessionClaims{
		SessionPayload: &SessionPayload{
			UserID: userID,
			Email:  email,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.sessionDuration)),
		},
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(s.sessionJWTSecret)
}

func (s *Service) VerifySessionToken(tokenString string) (*SessionPayload, error) {
	claims := &sessionClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("%w", ErrUnexpectedSigningMethod)
		}
		return s.sessionJWTSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, fmt.Errorf("%w", ErrInvalidSessionToken)
	}
	return claims.SessionPayload, nil
}
