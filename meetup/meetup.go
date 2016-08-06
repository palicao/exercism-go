// Package meetup finds the best day for a meetup
package meetup

import "time"

const testVersion = 3

type WeekSchedule int

const (
	First WeekSchedule = iota
	Second
	Third
	Fourth
	Last
	Teenth
)

// Day finds the right day for the meetup according to the WeekSchedule
func Day(schedule WeekSchedule, weekday time.Weekday, month time.Month, year int) int {
	start := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)

	switch schedule {
	case First:
		return findWeekday(start, weekday, inc)
	case Second:
		return findWeekday(start.AddDate(0, 0, 7), weekday, inc)
	case Third:
		return findWeekday(start.AddDate(0, 0, 14), weekday, inc)
	case Fourth:
		return findWeekday(start.AddDate(0, 0, 21), weekday, inc)
	case Last:
		return findWeekday(start.AddDate(0, 1, -1), weekday, dec)
	case Teenth:
		return findWeekday(start.AddDate(0, 0, 12), weekday, inc)
	}
	return -1
}

// inc adds one day to a given Time
func inc(t time.Time) time.Time {
	return t.AddDate(0, 0, 1)
}

// inc subtracts one day to a given Time
func dec(t time.Time) time.Time {
	return t.AddDate(0, 0, -1)
}

// find a matching weekday in a week
func findWeekday(t time.Time, weekday time.Weekday, f func(time.Time) time.Time) int {
	for i := 0; i < 7; i++ {
		if t.Weekday() == weekday {
			break
		}
		t = f(t)
	}
	return t.Day()
}
