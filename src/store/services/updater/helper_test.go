package updater

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestPercent(t *testing.T) {
	if num, percent, ok := percent("4%"); ok && num > 0 {
		assert.True(t, percent)
		assert.Equal(t, num, float64(4))
	}

	if num, percent, ok := percent("2.5%"); ok && num > 0 {
		assert.True(t, percent)
		assert.Equal(t, num, float64(2.5))
	}

	if num, percent, ok := percent("2.5"); ok && num > 0 {
		assert.False(t, percent)
	}
}