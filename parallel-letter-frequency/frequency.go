// Package letter calculates letter frequency using concurrency
package letter

type FreqMap map[rune]int

// Frequency calculates letter frequency in a string
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(texts []string) FreqMap {
	c := make(chan FreqMap)
	for _, s := range texts {
		go func(s string) { c <- Frequency(s) }(s)
	}

	total := FreqMap{}
	for range texts {
		for letter, freq := range <-c {
			total[letter] += freq
		}
	}
	return total
}
