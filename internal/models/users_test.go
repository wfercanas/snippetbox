package models

import (
	"testing"

	"github.com/wfercanas/snippetbox/internal/assert"
)

func TestUserModelExists(t *testing.T) {
	if testing.Short() {
		t.Skip("models: skipping integration test")
	}

	tests := []struct {
		name     string
		userID   int
		expected bool
	}{
		{
			name:     "Valid ID",
			userID:   1,
			expected: true,
		},
		{
			name:     "Zero ID",
			userID:   0,
			expected: false,
		},
		{
			name:     "Non-Existent ID",
			userID:   2,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := newTestDB(t)
			m := UserModel{DB: db}

			exists, err := m.Exists(tt.userID)
			assert.Equal(t, exists, tt.expected)
			assert.NilError(t, err)
		})
	}
}
