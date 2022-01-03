package projecteuler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNonAbundantSums(t *testing.T) {
	assert.Equal(t, 4179871, NonAbundantSums())
}
