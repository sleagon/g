package g

// ValueOf returns the value of the given pointer.
func ValueOf[T any](ptr *T) T {
	if ptr == nil {
		var t T
		return t
	}
	return *ptr
}
