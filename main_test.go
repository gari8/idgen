package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateID(t *testing.T) {
	t.Run("use default-length", func(t *testing.T) {
		id, err := generateID(defaultLength)
		assert.NoError(t, err)
		assert.Equal(t, defaultLength, len(id))
	})

	t.Run("should be error", func(t *testing.T) {
		id, err := generateID(0)
		assert.Error(t, err)
		assert.Empty(t, id)
	})
}
