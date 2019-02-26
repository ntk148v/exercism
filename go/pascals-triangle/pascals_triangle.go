package pascal

func Triangle(n int) [][]int {
	if n < 1 {
		return [][]int{}
	}
	t := make([][]int, n)
	r := []int{1}
	t[0] = r
	for i := 1; i < n; i++ {
		prev := r
		r = make([]int, i+1)
		r[0] = 1
		r[i] = 1
		for j := 1; j < i; j++ {
			r[j] = prev[j-1] + prev[j]
		}
		t[i] = r
	}
	return t
}
