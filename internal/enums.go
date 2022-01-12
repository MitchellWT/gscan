package gscan

import (
	"strings"
	"time"
)

// TODO: Seperate all enums (and function groups) into their
// own file

type Interval int
type UndefinedIntervalError struct{}

func (uete UndefinedIntervalError) Error() string {
	return "Error: Undefined interval provided!"
}

const (
	IntervalUndefined Interval = -1
	All               Interval = 0
	Hour              Interval = 1
	Day               Interval = 2
	Week              Interval = 3
	Month             Interval = 4
	ThreeMonths       Interval = 5
	SixMonths         Interval = 6
	Year              Interval = 7
)

func ToInterval(s string) (Interval, error) {
	s = strings.ToLower(s)
	switch s {
	case "hour":
		return Hour, nil
	case "day":
		return Day, nil
	case "week":
		return Week, nil
	case "month":
		return Month, nil
	case "three-months":
		return ThreeMonths, nil
	case "six-months":
		return SixMonths, nil
	case "year":
		return Year, nil
	case "all":
		return All, nil
	default:
		return IntervalUndefined, UndefinedIntervalError{}
	}
}

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
	case All:
		return "all"
	default:
		return "nil"
	}
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
	case All:
		return 0
	}
	return unixNow
}

func (i Interval) GetEnd() int64 {
	return time.Now().Unix()
}

type ExportType int
type UndefinedExportTypeError struct{}

func (uete UndefinedExportTypeError) Error() string {
	return "Error: Undefined export type provided!"
}

const (
	ExportTypeUndefined ExportType = -1
	Raw                 ExportType = 0
	Total               ExportType = 1
)

func ToExportType(s string) (ExportType, error) {
	s = strings.ToLower(s)
	switch s {
	case "total":
		return Total, nil
	case "raw":
		return Raw, nil
	default:
		return ExportTypeUndefined, UndefinedExportTypeError{}
	}
}

func (et ExportType) String() string {
	switch et {
	case Total:
		return "total"
	case Raw:
		return "raw"
	default:
		return "nil"
	}
}
