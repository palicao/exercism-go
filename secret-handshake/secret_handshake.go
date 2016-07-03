// Package secret for the creation of secret handshakes
package secret

var codes = []struct {
	num  int
	word string
}{
	{1, "wink"},
	{2, "double blink"},
	{4, "close your eyes"},
	{8, "jump"},
}

// Handshake returns an array with the handshake
func Handshake(i int) []string {
	handshake := []string{}
	if i <= 0 {
		return handshake
	}
	for _, code := range codes {
		if i&code.num == code.num {
			handshake = append(handshake, code.word)
		}
	}
	if i&16 == 16 {
		handshake = reverse(handshake)
	}
	return handshake
}

// reverse returns a reversed slice
func reverse(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
