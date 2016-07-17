// Package binary converts binary to decimal
package binary

import "errors"

// ParseBinary first implementation uses the builtin strconv.ParseInt function which does exactly the same... not funny...
// func ParseBinary(binary string) (decimal int, err error) {
// 	v, err := strconv.ParseInt(binary, 2, 32)
// 	return int(v), err
// }
//
// ParseBinary second implementation uses bitwise shift operators to convert a binary string into decimal
func ParseBinary(binary string) (decimal int, err error) {
	l := len(binary)
	for pos, bin := range binary {
		if bin != '0' && bin != '1' {
			return 0, errors.New("invalid binary string")
		}
		if bin == '1' {
			decimal += 1 << uint(l-pos-1)
		}
	}
	return decimal, nil
}
