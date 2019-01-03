package accumulate

// Accumulate will apply a function of string manipulation to every string in a slice of strings
func Accumulate(slice []string, f func(string) string) []string {
	var slice2 []string
	// We can also do the below iteration by not rebuilding the slice but replacing the existing entries in the slice (1 by 1) with the transformed entries
	for _, s := range slice {
		slice2 = append(slice2, f(s))
	}
	return slice2
}
