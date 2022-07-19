package g

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseS(t *testing.T) {
	is := assert.New(t)
	is.Equal("srever", ReverseS("revers"))
}
