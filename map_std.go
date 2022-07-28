package g

// Keys returns the keys of the map, order is not guaranteed.
func Keys[K comparable, V any](values map[K]V) []K {
	result := make([]K, 0, len(values))
	for k := range values {
		result = append(result, k)
	}
	return result
}

// Values returns the values of the map, order is not guaranteed.
func Values[K comparable, V any](values map[K]V) []V {
	result := make([]V, 0, len(values))
	for _, v := range values {
		result = append(result, v)
	}
	return result
}

// MapKeys returns the f(key)s of the map, order is not guaranteed.
func MapKeys[K comparable, V any, T any](values map[K]V, f func(K) T) []T {
	result := make([]T, 0, len(values))
	for k := range values {
		result = append(result, f(k))
	}
	return result
}

// MapValues returns the f(value)s of the map, order is not guaranteed.
func MapValues[K comparable, V any, T any](values map[K]V, f func(V) T) []T {
	result := make([]T, 0, len(values))
	for _, v := range values {
		result = append(result, f(v))
	}
	return result
}

// MapEntries returns the f(entries) of the map, order is not guaranteed.
func MapEntries[K comparable, V any, T any](values map[K]V, f func(K, V) T) []T {
	result := make([]T, 0, len(values))
	for k, v := range values {
		result = append(result, f(k, v))
	}
	return result
}

// MaxValue returns the maximum value of the map.
func MaxValue[K comparable, V Numberic](values map[K]V) V {
	first := true
	var result V
	for _, v := range values {
		if first || v > result {
			result = v
		}
		first = false
	}
	return result
}

// MaxKey returns the maximum key of the map.
func MaxKey[K Numberic, V any](values map[K]V) K {
	first := true
	var result K
	for k := range values {
		if first || k > result {
			result = k
		}
		first = false
	}
	return result
}

// MinValue returns the minimum value of the map.
func MinValue[K comparable, V Numberic](values map[K]V) V {
	first := true
	var result V
	for _, v := range values {
		if first || v < result {
			result = v
		}
		first = false
	}
	return result
}

// MinKey returns the minimum key of the map.
func MinKey[K Numberic, V any](values map[K]V) K {
	first := true
	var result K
	for k := range values {
		if first || k < result {
			result = k
		}
		first = false
	}
	return result
}

// MaxValueBy returns the maximum value of the map based on less method.
func MaxValueBy[K comparable, V any](values map[K]V, less func(V, V) bool) V {
	first := true
	var result V
	for _, v := range values {
		if first || less(result, v) {
			result = v
		}
		first = false
	}
	return result
}

// MaxKeyBy returns the maximum key of the map based on less method.
func MaxKeyBy[K comparable, V any](values map[K]V, less func(K, K) bool) K {
	first := true
	var result K
	for k := range values {
		if first || less(result, k) {
			result = k
		}
		first = false
	}
	return result
}

// MinValueBy returns the minimum value of the map based on less method.
func MinValueBy[K comparable, V any](values map[K]V, less func(V, V) bool) V {
	first := true
	var result V
	for _, v := range values {
		if first || less(v, result) {
			result = v
		}
		first = false
	}
	return result
}

// MinKeyBy returns the minimum key of the map based on less method.
func MinKeyBy[K comparable, V any](values map[K]V, less func(K, K) bool) K {
	first := true
	var result K
	for k := range values {
		if first || less(k, result) {
			result = k
		}
		first = false
	}
	return result
}

// PutAll add all elements to exist map, if the key is already exists, the value will be overridden.
func PutAll[K comparable, V any](target map[K]V, elements map[K]V) map[K]V {
	for k, v := range elements {
		target[k] = v
	}
	return target
}
