package g

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCondition(t *testing.T) {
	is := assert.New(t)

	is.Equal(1, Ternary(true, 1, 2))
	is.Equal(2, Ternary(false, 1, 2))

	// ternaryF
	is.Equal(1, TernaryF(true, func() int { return 1 }, func() int { return 2 }))
	is.Equal(2, TernaryF(false, func() int { return 1 }, func() int { return 2 }))
}
