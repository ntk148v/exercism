package robotname

import (
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
func (r *Robot) Name() string {
	rand.Seed(time.Now().UTC().UnixNano())
	for inUse, ok := r.name == "", true; ok && inUse; inUse, ok = usedNames[r.name] {
		r.name = fmt.Sprintf("%c%c%03d",
			byte(65+rand.Intn(26)),
			byte(65+rand.Intn(26)),
			rand.Intn(1000))
	}
	usedNames[r.name] = true
	return r.name
}

// Reset resets a given robot's name.
func (r *Robot) Reset() {
	*r = Robot{}
}
