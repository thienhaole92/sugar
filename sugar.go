package sugar

// Map returns a new array with the results of passed through the mapper function on each element.
func Map[T, R any](array []T, mapper func(T) R) []R {
	result := make([]R, len(array))

	for i, v := range array {
		result[i] = mapper(v)
	}
	return result
}
