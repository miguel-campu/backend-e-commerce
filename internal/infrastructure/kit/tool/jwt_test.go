package tool

import (
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/jnates/crud_golang/internal/domain/model"
	"github.com/stretchr/testify/assert"
)

func TestGenerateToken(t *testing.T) {
	user := &model.User{
		ID:   uuid.New(),
		Role: "admin",
	}

	t.Run("Generate Token with Env Secret", func(t *testing.T) {
		os.Setenv("JWT_SECRET", "test_secret")
		defer os.Unsetenv("JWT_SECRET")

		token, err := GenerateToken(user)
		assert.NoError(t, err)
		assert.NotEmpty(t, token)
	})

	t.Run("Generate Token with Fallback Secret", func(t *testing.T) {
		os.Unsetenv("JWT_SECRET")

		token, err := GenerateToken(user)
		assert.NoError(t, err)
		assert.NotEmpty(t, token)
	})
}
