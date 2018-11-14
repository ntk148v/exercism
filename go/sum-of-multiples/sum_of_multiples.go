package summultiples

func SumMultiples(limit int, divisors ...int) (sum int) {
	for i := 0; i < limit; i++ {
		for _, v := range divisors {
			if i%v == 0 {
				sum += i
				break
			}
		}
	}
	return
}
