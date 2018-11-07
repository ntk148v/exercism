/*Package leap does leap year check*/
package leap

// IsLeapYear check a given year whether leap year or not.
func IsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}
