// Package clock time without dates
package clock

import "fmt"

type clock struct {
	hour, minute int
}

// New constructor for clock
func New(h, m int) clock {
	return *((&clock{h, m}).overflow())
}

// String stringer for clock, e.g. "04:12"
func (c clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

// Add adds a given number of minutes to a clock
func (c clock) Add(in int) clock {
	c.minute += in
	return *((&c).overflow())
}

// Subtract subtracts a given number of minutes from a clock
func (c clock) Subtract(in int) clock {
	c.minute -= in
	return *((&c).overflow())
}

// overflow handles rollover of hour & minute values
func (c *clock) overflow() *clock {
	min := c.minute % 60
	hr := (c.hour + c.minute/60) % 24
	// handle negative minutes
	if min < 0 {
		min += 60
		hr -= 1
	}
	// handle negative hours
	if hr < 0 {
		hr += 24
	}
	c.hour = hr
	c.minute = min
	return c
}
