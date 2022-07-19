package g

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestError(t *testing.T) {
	is := assert.New(t)

	// error
	is.Equal(1, Must(1, nil))
	is.Equal(1, Must1(1, nil))
	{
		v1, v2 := Must2(1, 2, nil)
		is.Equal(1, v1)
		is.Equal(2, v2)
	}
	{
		v1, v2, v3 := Must3(1, 2, 3, nil)
		is.Equal(1, v1)
		is.Equal(2, v2)
		is.Equal(3, v3)
	}
	{
		v1, v2, v3, v4 := Must4(1, 2, 3, 4, nil)
		is.Equal(1, v1)
		is.Equal(2, v2)
		is.Equal(3, v3)
		is.Equal(4, v4)
	}

	// panics
	is.Panics(func() { Must(1, errors.New("")) })
	is.Panics(func() { Must1(1, errors.New("")) })
	is.Panics(func() { Must2(1, 2, errors.New("")) })
	is.Panics(func() { Must3(1, 2, 3, errors.New("")) })
	is.Panics(func() { Must4(1, 2, 3, 4, errors.New("")) })

}
