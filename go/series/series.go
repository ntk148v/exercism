package series

// All returns a list of all substrings of s with length n
func All(n int, s string) (r []string) {
	for i := 0; n <= len(s); i++ {
		r = append(r, s[i:n])
		n++
	}
	return
}

// UnsafeFirst returns the first substring of s with length n
func UnsafeFirst(n int, s string) string {
	return s[:n]
}
