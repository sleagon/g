package g

// Map is a map operator, like js/java's map.
func Map[F any, T any](values []F, f func(F, int) T) []T {
	result := make([]T, 0, len(values))
	for i, v := range values {
		result = append(result, f(v, i))
	}
	return result
}

// Reduce is a reduce operator, like js's reduce.
func Reduce[F any, T any](values []F, f func(T, F, int) T) T {
	var result T
	for i, v := range values {
		result = f(result, v, i)
	}
	return result
}

// Partition is a partition operator, like java's partition.
func Partition[T any](values []T, size int) [][]T {

	partitions := (len(values) + size - 1) / size

	result := make([][]T, partitions)
	for i := 0; i < partitions; i++ {
		end := (i + 1) * size
		if end > len(values) {
			end = len(values)
		}
		result[i] = values[i*size : end]
	}
	return result
}

// GroupBy is a group by operator, like guava's group by.
func GroupBy[T any, K comparable](values []T, f func(T, int) K) map[K][]T {
	result := make(map[K][]T)
	for i, v := range values {
		key := f(v, i)
		result[key] = append(result[key], v)
	}
	return result
}

// Contains check v is contained by values.
func Contains[T comparable](values []T, v T) bool {
	for _, vv := range values {
		if vv == v {
			return true
		}
	}
	return false
}

// Uniq remove duplicated elements from values
func Uniq[T comparable](values []T) []T {
	result := make([]T, 0, len(values))
	for _, v := range values {
		if !Contains(result, v) {
			result = append(result, v)
		}
	}
	return result
}

// UniqBy remove duplicated elements from values, based on given func.
func UniqBy[V any, T comparable](values []V, f func(V) T) []V {
	result := make([]V, 0, len(values))
	for _, v := range values {
		idx := 0
		for _, vv := range result {
			if f(v) == f(vv) {
				break
			}
			idx++
		}
		if idx == len(result) {
			result = append(result, v)
		}
	}
	return result
}

// IndexOf find the index of v in values, return -1 if element not found.
func IndexOf[T comparable](values []T, v T) int {
	for i, vv := range values {
		if vv == v {
			return i
		}
	}
	return -1
}

// LastIndexOf find the last index of v in values, return -1 if element not found.
func LastIndexOf[T comparable](values []T, v T) int {
	for i := len(values) - 1; i >= 0; i-- {
		if values[i] == v {
			return i
		}
	}
	return -1
}

// IndexOfBy find the index of v in values, return -1 if element not found.
func IndexOfBy[T comparable](values []T, f func(T, int) bool) int {
	for i, v := range values {
		if f(v, i) {
			return i
		}
	}
	return -1
}

// LastIndexOfBy find the last index of v in values, return -1 if element not found.
func LastIndexOfBy[T comparable](values []T, f func(T, int) bool) int {
	for i := len(values) - 1; i >= 0; i-- {
		if f(values[i], i) {
			return i
		}
	}
	return -1
}

// Filter is a filter operator, like java's filter.
func Filter[T any](values []T, f func(T, int) bool) []T {
	result := make([]T, 0, len(values))
	for i, v := range values {
		if f(v, i) {
			result = append(result, v)
		}
	}
	return result
}

// Flat2 flatten 2-dimensional array to 1-dimensional array.
func Flat2[T any](values [][]T) []T {
	result := make([]T, 0, len(values))
	for _, v := range values {
		result = append(result, v...)
	}
	return result
}

// Flat3 flatten 3-dimensional array to 1-dimensional array.
func Flat3[T any](values [][][]T) []T {
	result := make([]T, 0, len(values))
	for _, v := range values {
		result = append(result, Flat2(v)...)
	}
	return result
}

// Slice2Map transfer slice to a map base on given func.
func Slice2Map[K comparable, V any](values []V, f func(V, int) K) map[K]V {
	result := make(map[K]V)
	for i, v := range values {
		result[f(v, i)] = v
	}
	return result
}

// Reverse reverse the order of elements in values.
func Reverse[T any](values []T) []T {
	for i, j := 0, len(values)-1; i < j; i, j = i+1, j-1 {
		values[i], values[j] = values[j], values[i]
	}
	return values
}

// AssertSlice assert interface list to target type, panic if elements of values are not T.
func AssertSlice[T any](values []any) []T {
	result := make([]T, len(values))
	for i, v := range values {
		result[i] = v.(T)
	}

	return result
}

// AnySlice convert []T to []interface{}/[]any.
func AnySlice[T any](values []T) []any {
	result := make([]any, len(values))
	for i, v := range values {
		result[i] = v
	}

	return result
}

// CopySlice copy slice.
func CopySlice[T any](values []T) []T {
	result := make([]T, len(values))
	copy(result, values)
	return result
}

func Repeat[T any](v T, n int) []T {
	result := make([]T, n)
	for i := 0; i < n; i++ {
		result[i] = v
	}

	return result
}
