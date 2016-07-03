// Package leap provides the handy function IsLeapYear to check if an year is leap
package leap

const testVersion = 2

// IsLeapYear returns true if the year is leap
func IsLeapYear(year int) bool {
	if year%4 == 0 {
		if year%100 == 0 && year%400 != 0 {
			return false
		}
		return true
	}
	return false
}
