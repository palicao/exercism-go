// Package robotname manages the creation of unique names for robots
package robotname

import (
	"math/rand"
	"time"
	"fmt"
)

type Robot string

var registry = make(map[*Robot]bool)

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
	r := rand.Intn('Z' - 'A') + 'A'
	return string(r)
}

func randomNumber() string {
	return fmt.Sprintf("%03d", rand.Intn(1000))
}

func (r *Robot) Reset() {
	registry[r] = false
}
