package robotname

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Robot reprents a robot with its name.
type Robot struct {
	name string
}

// usedName keeps track of which names are in use.
var usedNames = make(map[string]bool)

// count tracks number of given names
var count int

// max number of available names
var max = 676000

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// Name defines the random unique name of robot.
func (r *Robot) Name() (string, error) {
	if r.name > "" {
		return r.name, nil
	}
	for r.name == "" || usedNames[r.name] {
		if count == max {
			return "", errors.New("All valid robot names are used")
		}
		r.name = fmt.Sprintf("%c%c%03d",
			'A'+byte(rand.Intn(26)),
			'A'+byte(rand.Intn(26)),
			rand.Intn(1000))
		count++
	}
	usedNames[r.name] = true
	return r.name, nil
}

// Reset resets a given robot's name.
func (r *Robot) Reset() {
	*r = Robot{}
}
