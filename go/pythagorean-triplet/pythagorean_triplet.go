package pythagorean

type Triplet [3]int

// Range returns a list of all Pythagorean triplets with sides in the
// range min to max inclusive.
func Range(min, max int) (ts []Triplet) {
	for a := min; a <= max; a++ {
		for b := a; b <= max; b++ {
			for c := b; c <= max; c++ {
				if a*a+b*b == c*c {
					ts = append(ts, Triplet{a, b, c})
				}
			}
		}
	}
	return
}

// Sum returns a list of all Pythagorean triplets where the sum a+b+c
// (the perimeter) is equal to p.
func Sum(sum int) (ts []Triplet) {
	max := sum / 2
	rts := Range(1, max)
	for _, rt := range rts {
		if sum == rt[0]+rt[1]+rt[2] {
			ts = append(ts, rt)
		}
	}
	return ts
}
