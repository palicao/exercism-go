// Package clock handles times without dates.
package clock

import "fmt"

const testVersion = 4

// Clock is the struct representing our clock
type Clock struct {
	hour   int
	minute int
}

// New instanciates a new clock
func New(hour int, minute int) Clock {
	newMinute := minute % 60
	if newMinute < 0 {
		newMinute = 60 + newMinute
		hour--
	}
	newHour := (hour + (minute / 60)) % 24
	if newHour < 0 {
		newHour = 24 + newHour
	}
	return Clock{hour: newHour, minute: newMinute}
}

// String returns a string representation of the clock
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

// Add adds (or subtracts) minutes
func (c Clock) Add(minutes int) Clock {
	return New(c.hour, c.minute+minutes)
}
