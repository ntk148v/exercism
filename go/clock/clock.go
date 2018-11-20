package clock

import "fmt"

type Clock int

// New returns a new Clock.
func New(h, m int) Clock {
	c := Clock((h*60 + m) % 1440)
	if c < 0 {
		c += 1440
	}
	return c
}

// String returns a time of Clock.
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c/60, c%60)
}

// Add returns a added-minutes Clock.
func (c Clock) Add(m int) Clock {
	return New(0, int(c)+m)
}

// Subtract returns a new subtracted-minutes Clock.
func (c Clock) Subtract(m int) Clock {
	return New(0, int(c)-m)
}
