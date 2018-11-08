// Package triangle determines a kind of triangle
package triangle

import "math"

type Kind string

const (
	NaT Kind = "Nat" // not a triangle
	Equ Kind = "Equ" // equilateral
	Iso Kind = "Iso" // isosceles
	Sca Kind = "Sca" // scalene
)

// KindFromSides determines if a triangle is squilateral,
// isosceles or scalene
func KindFromSides(a, b, c float64) Kind {
	if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) {
		return NaT
	}
	if (a+b < c) || (b+c < a) || (c+a < b) {
		return NaT
	}
	if (a <= 0) || (b <= 0) || (c <= 0) {
		return NaT
	}
	if math.IsInf(a, 1) || math.IsInf(b, 1) || math.IsInf(c, 1) {
		return NaT
	}
	if (a == b) && (b == c) {
		return Equ
	}
	if (a == b) || (b == c) || (c == a) {
		return Iso
	}
	return Sca
}
