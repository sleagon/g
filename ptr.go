package g

// ValueOf returns the value of the given pointer.
func ValueOf[T any](ptr *T) T {
	if ptr == nil {
		var t T
		return t
	}
	return *ptr
}

// Ptr returns the pointer of the given value, Ptr will panic if v is nil.
func Ptr[T any](v T) *T {
	return &v
}
