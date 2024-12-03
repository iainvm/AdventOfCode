package slices

// CountSlice returns a map where the index is the value of the slice and the value is the number of times that value appears in the slice
func CountSlice[T comparable](s []T) map[T]int {
	counts := make(map[T]int)
	for _, v := range s {
		counts[v]++
	}
	return counts
}
