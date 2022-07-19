package g

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValueOf(t *testing.T) {
	is := assert.New(t)
	// interface{}
	{
		data := "v"
		di := interface{}(data)
		is.True(data == ValueOf(&di))
	}

	// *interface{}
	{
		var pd *string
		di := interface{}(pd)
		pdi := &di
		is.True(pd == ValueOf(pdi))
	}

	// *int
	{
		value := 1
		is.Equal(1, ValueOf(&value))

		var pv *int
		is.True(0 == ValueOf(pv))
	}
}
