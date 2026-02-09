package testutil

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/johnal95/workouts-pwa/internal/sqlx"
	"github.com/johnal95/workouts-pwa/migrations"
	"github.com/stretchr/testify/require"
)

const BaseDSN = "postgres://postgres:postgres@127.0.0.1:5432/postgres?sslmode=disable"
const TestUserID = "019c404e-d046-74f2-acda-9e002993367a"
const TestUserEmail = "john.doe@example.com"
const TestUserAuthProvider = "GOOGLE"
const TestUserAuthID = "019c4050-43fe-7db3-b2f9-f8488cb23564"

func CreateTestDatabase(t *testing.T) string {
	t.Helper()

	dbName := "test_" + uuid.NewString()
	testDatabaseURL := fmt.Sprintf("postgres://postgres:postgres@127.0.0.1:5432/%s?sslmode=disable", dbName)

	adminDB, err := sqlx.Open(BaseDSN)
	require.NoError(t, err)

	_, err = adminDB.Exec(fmt.Sprintf(`CREATE DATABASE "%s"`, dbName))
	require.NoError(t, err)

	testDB, err := sqlx.Open(testDatabaseURL)
	require.NoError(t, err)
	defer testDB.Close()

	err = sqlx.Migrate(testDB, migrations.FS, ".")
	require.NoError(t, err)

	_, err = testDB.Exec(`INSERT INTO users (id, email, auth_id, auth_provider) VALUES ($1, $2, $3, $4)`,
		TestUserID, TestUserEmail, TestUserAuthID, TestUserAuthProvider)
	require.NoError(t, err)

	t.Cleanup(func() {
		adminDB.Exec(fmt.Sprintf(`DROP DATABASE "%s" WITH (FORCE)`, dbName))
		adminDB.Close()
	})

	return testDatabaseURL
}
