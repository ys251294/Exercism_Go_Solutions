// Package gigasecond provide AddGigaseconds methods to manipulate time
package gigasecond

import "time"

// AddGigasecond add 1000 million seconds to give time.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1e9)
}
