package times

import (
	"time"
)

const (
	ThreeMonth time.Duration = time.Duration(7776000000000000) // unit nano second
	SixMonth   time.Duration = time.Duration(ThreeMonth * 2)
)

func StringToTime(t string) (time.Time, error) {
	return time.Parse(time.RFC3339, t)
}

func IsLessThan(start time.Time, end time.Time, expectDuration time.Duration) bool {
	s := end.Sub(start)
	return s < expectDuration
}

func IsMoreThan(start time.Time, end time.Time, expectedDuration time.Duration) bool {
	s := end.Sub(start)
	return s > expectedDuration
}
