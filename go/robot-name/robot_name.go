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

// Name defines the random unique name of robot.
func (r *Robot) Name() (string, error) {
	if r.name > "" {
		return r.name, nil
	}
	ra := rand.New(rand.NewSource(time.Now().Unix()))
	first := fmt.Sprintf("%c%c%03d",
		byte(65+ra.Intn(26)),
		byte(65+ra.Intn(26)),
		ra.Intn(1000))
	for inUse, ok := r.name == "", true; ok && inUse; inUse, ok = usedNames[r.name] {
		r.name = fmt.Sprintf("%c%c%03d",
			byte(65+ra.Intn(26)),
			byte(65+ra.Intn(26)),
			ra.Intn(1000))
		if r.name == first {
			return "", errors.New("All valid robot names are used")
		}
	}
	usedNames[r.name] = true
	return r.name, nil
}

// Reset resets a given robot's name.
func (r *Robot) Reset() {
	*r = Robot{}
}
