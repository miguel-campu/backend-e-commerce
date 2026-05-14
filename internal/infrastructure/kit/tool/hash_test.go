package tool

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashSHA256(t *testing.T) {
	text := "password123"
	expectedHash := "ef92b778bafe771e89245b89ecbc08a44a4e166c06659911881f383d4473e94f" // SHA256 of password123

	t.Run("Generate Hash", func(t *testing.T) {
		hash := HashSHA256(text)
		assert.NotEmpty(t, hash)
		assert.Equal(t, expectedHash, hash)
	})

	t.Run("Check Hash Correct", func(t *testing.T) {
		isValid := CheckHashSHA256(text, expectedHash)
		assert.True(t, isValid)
	})

	t.Run("Check Hash Incorrect", func(t *testing.T) {
		isValid := CheckHashSHA256("wrongpassword", expectedHash)
		assert.False(t, isValid)
	})
}
