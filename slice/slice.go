package slice

// FindFirst returns the index of the first element in elems that satisfies the predicate f.
func FindFirst[T any](elems []T, f func(idx int, elem T) bool) (idx int) {
	for i, e := range elems {
		if f(i, e) {
			return i
		}
	}
	return -1
}

// Filter returns a new slice containing all elements in elems that satisfy the predicate f.
func Filter[T any](elems []T, f func(idx int, elem T) bool) (out []T) {
	for i, e := range elems {
		if f(i, e) {
			out = append(out, e)
		}
	}
	return out
}

// ForEach applies the function f to each element in elems, providing the index and element as arguments.
func ForEach[T any](elems []T, f func(idx int, elem T)) {
	for i, e := range elems {
		f(i, e)
	}
}

// Map returns a new slice containing the results of applying the function f to each element in elems.
func Map[T any, R any](elems []T, f func(idx int, elem T) R) (out []R) {
	for i, e := range elems {
		out = append(out, f(i, e))
	}
	return out
}
