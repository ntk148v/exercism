package diffsquares

// SquareOfSum calculates square of the sum of the first num
// natural numbers.
func SquareOfSum(num int) int {
	sum := 0
	for i := 1; i <= num; i++ {
		sum += i
	}
	return sum * sum
}

// SumOfSquares calculates the sum of the squares of the first
// num natural numbers.
func SumOfSquares(num int) int {
	sum := 0
	for i := 1; i <= num; i++ {
		sum += i * i
	}
	return sum
}

// Difference returns the difference between the square of the sum
// and the sum of the squares of the first num natural numbers.
func Difference(num int) int {
	return SquareOfSum(num) - SumOfSquares(num)
}
