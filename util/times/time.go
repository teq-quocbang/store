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

func WeekRange(year, week int) (start, end time.Time) {
	start = WeekStart(year, week)
	end = start.AddDate(0, 0, 6)
	return
}

func WeekStart(year, week int) time.Time {
	// Start from the middle of the year:
	t := time.Date(year, 7, 1, 0, 0, 0, 0, time.UTC)

	// Roll back to Monday:
	if wd := t.Weekday(); wd == time.Sunday {
		t = t.AddDate(0, 0, -6)
	} else {
		t = t.AddDate(0, 0, -int(wd)+1)
	}

	// Difference in weeks:
	_, w := t.ISOWeek()
	t = t.AddDate(0, 0, (week-w)*7)

	return t
}
