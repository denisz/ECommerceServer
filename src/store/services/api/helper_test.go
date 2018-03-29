package api

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"math"
)


func TestInBetween(t *testing.T) {
	assert.True(t, InBetween(2, 1, 3))
	assert.True(t, InBetween(2, 1, math.MaxInt32))
}
