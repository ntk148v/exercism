package robotname

import (
	"fmt"
	"math/rand"
)

// Robot reprents a robot with its name.
type Robot struct {
	name string
}

// Store all used names.
var usedName []string

// stringInSlice checks n exists in slice s or not.
func stringInSlice(n string, s []string) bool {
	for _, r := range s {
		if r == n {
			return true
		}
	}
	return false
}

// Name defines the random unique name of robot.
func (r *Robot) Name() string {
	if r.name == "" {
		for {
			r.name = fmt.Sprintf("%c%c%03d",
				byte(65+rand.Intn(26)),
				byte(65+rand.Intn(26)),
				rand.Intn(1000))
			if !stringInSlice(r.name, usedName) {
				usedName = append(usedName, r.name)
				break
			}
		}
	}
	return r.name
}

// Reset resets a given robot's name.
func (r *Robot) Reset() {
	*r = Robot{}
}
