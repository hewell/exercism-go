// Package gigasecond
// Calculate the moment when someone has lived for 10^9 seconds
package gigasecond

// import path for the time package from the standard library
import "time"

// Add 10^9 seconds to the provided time
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second*(time.Duration(1e9)));
}
