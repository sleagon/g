package g

// Max returns the maximum value in the given slice.
func Max[T Numberic](values []T) T {
	var result T
	if len(values) == 0 {
		return result
	}

	result = values[0]
	for i := 1; i < len(values); i++ {
		if values[i] > result {
			result = values[i]
		}
	}
	return result
}

// Min returns the minimum value in the given slice.
func Min[T Numberic](values []T) T {
	var result T
	if len(values) == 0 {
		return result
	}

	result = values[0]
	for i := 1; i < len(values); i++ {
		if values[i] < result {
			result = values[i]
		}
	}
	return result
}

// MaxBy returns the maximum value in the given slice, using the provided less function to compare values.
func MaxBy[T any](values []T, less func(T, T) bool) T {
	var result T
	if len(values) == 0 {
		return result
	}

	result = values[0]
	for i := 1; i < len(values); i++ {
		if less(result, values[i]) {
			result = values[i]
		}
	}
	return result
}

// MinBy returns the minimum value in the given slice, using the provided less function to compare values.
func MinBy[T any](values []T, less func(T, T) bool) T {
	var result T
	if len(values) == 0 {
		return result
	}

	result = values[0]
	for i := 1; i < len(values); i++ {
		if less(values[i], result) {
			result = values[i]
		}
	}
	return result
}
