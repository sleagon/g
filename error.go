package g

// Must assert a condition, and if it fails, panic.
func Must[T any](v T, err error) T {
	Must0(err)
	return v
}

// Must0 assert err is nil
func Must0(err error) {
	if err != nil {
		panic(err)
	}
}

// MustN assert n conditions, and if it fails, panic.
func Must1[T any](v T, err error) T {
	Must0(err)
	return v
}

// MustN assert n conditions, and if it fails, panic.
func Must2[T1 any, T2 any](v1 T1, v2 T2, err error) (T1, T2) {
	Must0(err)
	return v1, v2
}

// MustN assert n conditions, and if it fails, panic.
func Must3[T1 any, T2 any, T3 any](v1 T1, v2 T2, v3 T3, err error) (T1, T2, T3) {
	Must0(err)
	return v1, v2, v3
}

// MustN assert n conditions, and if it fails, panic.
func Must4[T1 any, T2 any, T3 any, T4 any](v1 T1, v2 T2, v3 T3, v4 T4, err error) (T1, T2, T3, T4) {
	Must0(err)
	return v1, v2, v3, v4
}
