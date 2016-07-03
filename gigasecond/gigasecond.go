// Package gigasecond allows you to celebrate your gigasecond anniversary
package gigasecond

import "time"

const (
	testVersion = 4
	giga        = 1000000000
)

// AddGigasecond adds 10**9 seconds to the provided time
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Duration(giga) * time.Second)
}
