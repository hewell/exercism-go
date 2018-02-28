// Package robotname generates name for robots
package robotname

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	capitals = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	names    = make(map[string]bool)
)

type Robot struct {
	name string
}

// init
func init() {
	rand.Seed(time.Now().UnixNano())
}

/* Name returns robot's name
   if the name is empty, assign a new name */
func (robot *Robot) Name() string {
	if robot.name != "" {
		return robot.name
	}

	for {
		name := fmt.Sprintf("%v%d", randomCapitalLetters(2), rand.Intn(900)+100)
		if _, present := names[name]; !present {
			names[name] = true
			robot.name = name
			return name
		}
	}
}

// Reset resets robot name
func (robot *Robot) Reset() {
	robot.name = ""
	delete(names, robot.name)
}

// randomString returns specified number of random capital letters
func randomCapitalLetters(in int) string {
	r := make([]rune, in)
	for i := range r {
		r[i] = capitals[rand.Intn(len(capitals))]
	}
	return string(r)
}
