// +build !windows

package qpc

import (
	"time"
)

// Now returns ticks
func Now() Count {
	return Count(time.Now().UnixNano())
}

func (a Count) Sub(b Count) Count { return a - b }

func (count Count) Nanoseconds() int64 {
	return int64(count)
}

func (count Count) Duration() time.Duration {
	return time.Duration(count) * time.Nanosecond
}
