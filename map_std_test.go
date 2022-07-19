package g

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeys(t *testing.T) {
	is := assert.New(t)

	keys := Keys(map[string]int{"a": 1, "b": 2})
	is.ElementsMatch([]string{"a", "b"}, keys)
}

func TestValues(t *testing.T) {
	is := assert.New(t)

	values := Values(map[string]int{"a": 1, "b": 2})
	is.ElementsMatch([]int{1, 2}, values)
}

func TestMapKeys(t *testing.T) {
	is := assert.New(t)

	keys := MapKeys(map[string]int{"a": 1, "b": 2}, func(v string) string { return v + "_" })
	is.ElementsMatch([]string{"a_", "b_"}, keys)
}

func TestMapValues(t *testing.T) {
	is := assert.New(t)

	values := MapValues(map[string]int{"a": 1, "b": 2}, func(v int) int { return v * 2 })
	is.ElementsMatch([]int{2, 4}, values)
}

func TestMapEntries(t *testing.T) {
	is := assert.New(t)

	entries := MapEntries(map[string]int{"a": 1, "b": 2},
		func(k string, v int) string { return fmt.Sprintf("%s_%d", k, v*2) })
	is.ElementsMatch([]string{"a_2", "b_4"}, entries)
}

func TestMaxKey(t *testing.T) {
	is := assert.New(t)

	is.Equal("b", MaxKey(map[string]int{"a": 1, "b": 2}))
}

func TestMinKey(t *testing.T) {
	is := assert.New(t)

	is.Equal("a", MinKey(map[string]int{"a": 1, "b": 2}))
}

func TestMaxValue(t *testing.T) {
	is := assert.New(t)

	is.Equal(2, MaxValue(map[string]int{"a": 1, "b": 2}))
}

func TestMinValue(t *testing.T) {
	is := assert.New(t)

	is.Equal(1, MinValue(map[string]int{"a": 1, "b": 2}))
}

func TestMaxKeyBy(t *testing.T) {
	is := assert.New(t)

	is.Equal("aa", MaxKeyBy(map[string]int{"a": 1, "aa": 2}, func(v1 string, v2 string) bool { return len(v1) < len(v2) }))
}

func TestMinKeyBy(t *testing.T) {
	is := assert.New(t)

	is.Equal("a", MinKeyBy(map[string]int{"a": 1, "aa": 2}, func(v1 string, v2 string) bool { return len(v1) < len(v2) }))
}

func TestMaxValueBy(t *testing.T) {
	is := assert.New(t)

	is.Equal(2, MaxValueBy(map[string]int{"a": 1, "aa": 2}, func(v1 int, v2 int) bool { return v1 < v2 }))
}

func TestMinValueBy(t *testing.T) {
	is := assert.New(t)

	is.Equal(1, MinValueBy(map[string]int{"a": 1, "aa": 2}, func(v1 int, v2 int) bool { return v1 < v2 }))
}
