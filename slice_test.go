package g

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	is := assert.New(t)

	is.ElementsMatch([]int{2, 3, 4}, Map([]int{1, 2, 3}, func(i int, _ int) int { return i + 1 }))
}

func TestMapE(t *testing.T) {
	is := assert.New(t)

	r, e := MapE([]int{1, 2, 3}, func(i int, _ int) (int, error) { return i + 1, nil })
	is.NoError(e)
	is.ElementsMatch([]int{2, 3, 4}, r)

	r, e = MapE([]int{1, 2, 3}, func(i int, _ int) (int, error) {
		if i == 2 {
			return 0, errors.New("error")
		}
		return i + 1, nil
	})

	is.Error(e)
	is.Nil(r)
}

func TestReduce(t *testing.T) {
	is := assert.New(t)

	is.Equal(6, Reduce([]int{1, 2, 3}, func(sum int, v int, _ int) int { return sum + v }))
	is.ElementsMatch(
		[]int{2, 3, 4},
		Values(Reduce([]int{1, 2, 3}, func(all map[int]int, v int, i int) map[int]int {
			if all == nil {
				all = make(map[int]int)
			}
			all[v] = i + 2
			return all
		})),
	)
}

func TestReduceE(t *testing.T) {
	is := assert.New(t)

	r, e := ReduceE([]int{1, 2, 3}, func(sum int, v int, _ int) (int, error) { return sum + v, nil })
	is.NoError(e)
	is.Equal(6, r)

	r, e = ReduceE([]int{1, 2, 3}, func(sum int, v int, _ int) (int, error) {
		if v == 2 {
			return 0, errors.New("error")
		}
		return sum + v, nil
	})
	is.Error(e)
	is.Equal(0, r)
}

func TestPartition(t *testing.T) {
	is := assert.New(t)

	is.ElementsMatch([][]int{[]int{1, 2}, []int{3, 4}}, Partition([]int{1, 2, 3, 4}, 2))
}

func TestGroupBy(t *testing.T) {
	is := assert.New(t)

	is.ElementsMatch(
		[][]int{{1}, {2}, {3}},
		Values(GroupBy([]int{1, 2, 3}, func(v int, _ int) int { return v })),
	)

	is.ElementsMatch(
		[][]int{{1, 2, 3}},
		Values(GroupBy([]int{1, 2, 3}, func(v int, _ int) int { return 1 })),
	)
}

func TestContains(t *testing.T) {
	is := assert.New(t)

	is.True(Contains([]int{1, 2, 3}, 1))
	is.False(Contains([]int{1, 2, 3}, 4))
}

func TestUniq(t *testing.T) {
	is := assert.New(t)

	is.ElementsMatch([]int{1, 2, 3}, Uniq([]int{1, 2, 3, 1, 2, 3}))
}

func TestIndexOf(t *testing.T) {
	is := assert.New(t)

	is.Equal(0, IndexOf([]int{1, 2, 3}, 1))
	is.Equal(-1, IndexOf([]int{1, 2, 3}, 4))
}

func TestLastIndexOf(t *testing.T) {
	is := assert.New(t)

	is.Equal(2, LastIndexOf([]int{1, 2, 3}, 3))
	is.Equal(-1, LastIndexOf([]int{1, 2, 3}, 4))
}

func TestIndexOfBy(t *testing.T) {
	is := assert.New(t)

	is.Equal(0, IndexOfBy([]int{1, 2, 3}, func(v int, _ int) bool { return v == 1 }))
	is.Equal(-1, IndexOfBy([]int{1, 2, 3}, func(v int, _ int) bool { return v == 4 }))
}

func TestLastIndexOfBy(t *testing.T) {
	is := assert.New(t)

	is.Equal(2, LastIndexOfBy([]int{1, 2, 3}, func(v int, _ int) bool { return v == 3 }))
	is.Equal(-1, LastIndexOfBy([]int{1, 2, 3}, func(v int, _ int) bool { return v == 4 }))
}

func TestFilter(t *testing.T) {
	is := assert.New(t)

	is.ElementsMatch([]int{2, 4}, Filter([]int{1, 2, 3, 4}, func(v int, _ int) bool { return v%2 == 0 }))
}

func TestFlat2(t *testing.T) {
	is := assert.New(t)

	is.ElementsMatch([]int{1, 2, 3, 4}, Flat2([][]int{{1, 2}, {3, 4}}))
}

func TestFlat3(t *testing.T) {
	is := assert.New(t)

	is.ElementsMatch([]int{1, 2, 3, 4}, Flat3([][][]int{{{1, 2}, {3, 4}}}))
}

func TestSlice2Map(t *testing.T) {
	is := assert.New(t)

	is.ElementsMatch([]int{1, 2, 3}, Values(Slice2Map([]int{1, 2, 3}, func(v int, idx int) int { return idx })))
	is.ElementsMatch([]int{0, 1, 2}, Keys(Slice2Map([]int{1, 2, 3}, func(v int, idx int) int { return idx })))
}

func TestReverse(t *testing.T) {
	is := assert.New(t)

	is.ElementsMatch([]int{3, 2, 1}, Reverse([]int{1, 2, 3}))
}

func TestAssertSlice(t *testing.T) {
	is := assert.New(t)

	values := []interface{}{1, 2, 3}

	typedValues := AssertSlice[int](values)

	is.ElementsMatch([]int{1, 2, 3}, typedValues)

	anySlice := AnySlice(typedValues)

	is.ElementsMatch(values, anySlice)
}

func TestUniqBy(t *testing.T) {
	is := assert.New(t)

	is.ElementsMatch([]int{1, 2}, UniqBy([]int{1, 2, 3, 4, 5}, func(v int) int { return v % 2 }))
	is.ElementsMatch([]int{1, 2, 3}, UniqBy([]int{1, 2, 3, 2, 2, 1, 2, 3}, func(v int) int { return v * v }))
}
