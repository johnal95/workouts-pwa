package user

import "time"

type AuthProvider string

const (
	AuthProviderGoogle = AuthProvider("GOOGLE")
)

type User struct {
	ID           string
	CreatedAt    time.Time
	Email        string
	AuthProvider AuthProvider
	AuthID       string
}
