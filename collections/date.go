package collections

import (
	"fmt"
	"time"
)

// DateStringToTime expects input string in the format 2006-01-02
// and returns time object at 00:00
func DateStringToTime(dateStr string) (time.Time, error) {
	layout := "2006-01-02T15:04:05.000Z"
	return time.Parse(layout, fmt.Sprintf("%sT00:00:00.000Z", dateStr))
}

// DateTimeStringToTime expects input string in the format 2006-01-02T15:04:05.000Z
// and returns time object
func DateTimeStringToTime(dateStr string) (time.Time, error) {
	layout := "2006-01-02T15:04:05.000Z"
	return time.Parse(layout, dateStr)
}
