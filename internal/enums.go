package gscan

import "time"

type Interval int

const (
	Hour        Interval = 0
	Day         Interval = 1
	Week        Interval = 2
	Month       Interval = 3
	ThreeMonths Interval = 4
	SixMonths   Interval = 5
	Year        Interval = 6
)

func (i Interval) String() string {
	switch i {
	case Hour:
		return "hour"
	case Day:
		return "day"
	case Week:
		return "week"
	case Month:
		return "month"
	case ThreeMonths:
		return "three-months"
	case SixMonths:
		return "six-months"
	case Year:
		return "year"
	}
	return "nil"
}

func (i Interval) GetStart() int64 {
	unixNow := time.Now().Unix()
	// Durations after a week are calculated using the yearly duration
	// of 8736h
	switch i {
	case Hour:
		hourDiff, _ := time.ParseDuration("1h")
		return unixNow - int64(hourDiff.Seconds())
	case Day:
		dayDiff, _ := time.ParseDuration("24h")
		return unixNow - int64(dayDiff.Seconds())
	case Week:
		weekDiff, _ := time.ParseDuration("168h")
		return unixNow - int64(weekDiff.Seconds())
	case Month:
		monthDiff, _ := time.ParseDuration("811h")
		return unixNow - int64(monthDiff.Seconds())
	case ThreeMonths:
		threeMonthDiff, _ := time.ParseDuration("2434h")
		return unixNow - int64(threeMonthDiff.Seconds())
	case SixMonths:
		sixMonthDiff, _ := time.ParseDuration("4368h")
		return unixNow - int64(sixMonthDiff.Seconds())
	case Year:
		yearDiff, _ := time.ParseDuration("8736h")
		return unixNow - int64(yearDiff.Seconds())
	}
	return unixNow
}

func (i Interval) GetEnd() int64 {
	return time.Now().Unix()
}
