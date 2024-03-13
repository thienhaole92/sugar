package sugar

// Find returns the first element in an array that passes a given test and corresponding index.
// Index of -1 is returned if no element passes the test.
func Find[T any](array []T, fn func(T) bool) (T, int) {
	for i, v := range array {
		if fn(v) {
			return v, i
		}
	}
	var res T
	return res, -1
}
