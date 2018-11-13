package diffsquares

// SquareOfSum calculates square of the sum of the first num
// natural numbers.
// https://en.wikipedia.org/wiki/Triangular_number
func SquareOfSum(num int) int {
	return num * num * (num*num + 2*num + 1) / 4
}

// SumOfSquares calculates the sum of the squares of the first
// num natural numbers.
// https://en.wikipedia.org/wiki/Square_pyramidal_number
func SumOfSquares(num int) int {
	return num * (num + 1) * (2*num + 1) / 6
}

// Difference returns the difference between the square of the sum
// and the sum of the squares of the first num natural numbers.
func Difference(num int) int {
	return SquareOfSum(num) - SumOfSquares(num)
}
