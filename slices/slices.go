package slices

// Foreach iterates through the given slice
func Foreach[T any](s []T, f func(int, T)) {
	for i, v := range s {
		f(i, v)
	}
}

// Filter filters the given slice using the given filter func
func Filter[T any](s []T, f func(T) bool) []T {
	filtered := []T{}

	for _, v := range s {
		if f(v) {
			filtered = append(filtered, v)
		}
	}

	return filtered
}

// Map maps a value
func Map[T, O any](s []T, mf func(T) O) []O {
	m := []O{}

	for _, v := range s {
		m = append(m, mf(v))
	}

	return m
}

// MapSlice maps slice of values
func MapSlice[T, O any](s []T, mf func(T) []O) []O {
	m := []O{}

	for _, v := range s {
		m = append(m, mf(v)...)
	}

	return m
}

// First returns the first or false
func First[T any](s []T, ff func(T) bool) (T, bool) {
	for _, v := range s {
		if ff(v) {
			return v, true
		}
	}

	t := new(T)
	return *t, false
}

// Contains returns true if elemens exists
func Contains[T any](s []T, f func(T) bool) bool {
	_, ok := First(s, f)
	return ok
}
