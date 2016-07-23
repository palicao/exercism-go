// Package robotname manages the creation of unique names for robots
package robotname

import (
	"fmt"
	"math/rand"
	"time"
)

type Robot string

var registry = make(map[*Robot]bool)

// Name returns an unique robot name
func (r *Robot) Name() string {
	if registry[r] {
		return string(*r)
	}
	rand.Seed(time.Now().UnixNano())
	*r = Robot(randomLetter() + randomLetter() + randomNumber())
	registry[r] = true
	return string(*r)
}

func randomLetter() string {
	r := rand.Intn('Z'-'A') + 'A'
	return string(r)
}

func randomNumber() string {
	return fmt.Sprintf("%03d", rand.Intn(1000))
}

// Reset reset a robot's name
func (r *Robot) Reset() {
	registry[r] = false
}
