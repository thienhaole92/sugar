package sugar

// Map returns a new array with the results of passed through the mapper function on each element.
func Map[I, O any](array []I, mapper func(I) O) []O {
	result := make([]O, len(array))

	for i, v := range array {
		result[i] = mapper(v)
	}

	return result
}
