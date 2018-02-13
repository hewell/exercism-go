// Package gigasecond calculates the moment when someone has lived for 10^9 seconds
package gigasecond

import "time"

// AddGigasecond adds 10^9 seconds to the provided time.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * (time.Duration(1e9)))
}
