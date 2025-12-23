package tokens

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/johnal95/workouts-pwa/cmd/config"
)

type SessionPayload struct {
	UserID string `json:"user_id"`
}

type sessionClaims struct {
	SessionPayload *SessionPayload `json:"sub"`
	jwt.RegisteredClaims
}

func CreateToken(userID string) (string, error) {
	sessionJWTSecret := []byte(config.GetSessionJWTSecret())

	claims := sessionClaims{
		SessionPayload: &SessionPayload{
			UserID: userID,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(sessionJWTSecret)
}

func VerifySessionToken(tokenString string) (*SessionPayload, error) {
	sessionJWTSecret := []byte(config.GetSessionJWTSecret())

	claims := &sessionClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return sessionJWTSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	return claims.SessionPayload, nil
}
