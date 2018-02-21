// Package clock time without dates
package clock

import "fmt"

type clock struct {
	hour   int
	minute int
}

// New constructor for clock
func New(h, m int) clock {
	hr, min := overflow(h, m)
	return clock{hr, min}
}

// String stringer for clock, e.g. "04:12"
func (c clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

// Add adds a given number of minutes to a clock
func (c clock) Add(in int) clock {
	hr, min := overflow(c.hour, c.minute+in)
	return clock{hr, min}
}

// Subtract subtracts a given number of minutes from a clock
func (c clock) Subtract(in int) clock {
	hr, min := overflow(c.hour, c.minute-in)
	return clock{hr, min}
}

// overflow handles rollover of hour & minute values
func overflow(h, m int) (int, int) {
	min := m % 60
	hr := (h + m/60) % 24
	// handle negative minutes
	if min < 0 {
		min += 60
		hr -= 1
	}
	// handle negative hours
	if hr < 0 {
		hr += 24
	}
	return hr, min
}
