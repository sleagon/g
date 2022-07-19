package g

// Ternary is a ternary operator, like C's ?:.
func Ternary[T any](cond bool, v1 T, v2 T) T {
	if cond {
		return v1
	}
	return v2
}

// If is a if-else operator, like C's if-else.
func TernaryF[T any](cond bool, f1 func() T, f2 func() T) T {
	if cond {
		return f1()
	}
	return f2()
}
