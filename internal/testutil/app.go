package testutil

import (
	"net/http/httptest"
	"testing"

	"github.com/johnal95/workouts-pwa/internal/app"
	"github.com/johnal95/workouts-pwa/internal/auth"
	"github.com/johnal95/workouts-pwa/internal/config"
	"github.com/johnal95/workouts-pwa/internal/logging"
	"github.com/johnal95/workouts-pwa/internal/routes"
	"github.com/stretchr/testify/require"
)

type TestApp struct {
	T           *testing.T
	App         *app.Application
	Server      *httptest.Server
	AuthService *auth.Service
}

const TestJWTSecret = "TEST_JWT_SECRET"

func NewTestApp(t *testing.T) *TestApp {
	t.Helper()

	dbURL := CreateTestDatabase(t)

	app, err := app.NewApplication(&app.ApplicationOptions{
		DatabaseURL:      dbURL,
		SessionJWTSecret: TestJWTSecret,
		Logger:           logging.NewLogger(config.EnvTest),
	})
	require.NoError(t, err)

	authService := auth.NewService(TestJWTSecret)

	server := httptest.NewServer(routes.SetupRoutesHandler(app))

	t.Cleanup(func() {
		server.Close()
		app.DB.Close()
	})

	return &TestApp{
		T:           t,
		App:         app,
		Server:      server,
		AuthService: authService,
	}
}
