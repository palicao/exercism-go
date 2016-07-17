// Package brackets ensures that brackets and braces are matched and nested correctly
package brackets

const testVersion = 3

var couples = map[rune]rune{
	'{': '}',
	'[': ']',
	'(': ')',
}

// Bracket checks if parentheses are matched and nested correctly in a string
func Bracket(s string) (bool, error) {
	queue := make([]rune, 0)
	for _, l := range s {

		if !isParentheses(l) {
			continue
		}

		if isOpen(l) {
			queue = append(queue, couples[l])
		} else {
			if len(queue) == 0 || string(l) != string(queue[len(queue)-1:]) {
				return false, nil
			}
			queue = queue[:len(queue)-1]
		}

	}
	return len(queue) == 0, nil
}

func isParentheses(r rune) bool {
	for k, v := range couples {
		if r == k || r == v {
			return true
		}
	}
	return false
}

func isOpen(r rune) bool {
	for k := range couples {
		if r == k {
			return true
		}
	}
	return false
}
