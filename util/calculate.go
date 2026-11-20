package util

import (
	"fmt"
	"math"
	"time"
)

func CalculateDistance(latSrc float64, longSrc float64, latDest float64, longDest float64) float64 {
	if latSrc == 0 || longSrc == 0 || latDest == 0 || longDest == 0 {
		return 0
	}

	latSrcRad := latSrc * math.Pi / 180
	latDestRad := latDest * math.Pi / 180
	deltaLat := (latDest - latSrc) * math.Pi / 180
	deltaLong := (longDest - longSrc) * math.Pi / 180

	a := math.Sin(deltaLat/2)*math.Sin(deltaLat/2) +
		math.Cos(latSrcRad)*math.Cos(latDestRad)*
			math.Sin(deltaLong/2)*math.Sin(deltaLong/2)
	if a > 1 {
		a = 1
	}

	return 6371.0088 * 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
}

func CalculateDateToWIB(date string, timeFormat string, timeZone string) (dateTime time.Time, err error) {
	loc, err := time.LoadLocation("Local")
	if err != nil {
		return
	}

	dateTime, err = time.ParseInLocation(timeFormat, date, loc)
	if err != nil {
		return
	}

	switch timeZone {
	case "WIB":
		return dateTime, nil
	case "WITA":
		return dateTime.Add(-60 * time.Minute), nil
	case "WIT":
		return dateTime.Add(-120 * time.Minute), nil
	default:
		return dateTime, fmt.Errorf("invalid time zone")
	}
}

// calculate time difference between start time and end time, return in seconds
func TimestampDiff(startDate string, endDate string, format string) (time.Duration, error) {
	startTime, err := time.Parse(format, startDate)
	if err != nil {
		return 0, fmt.Errorf("failed to parse start date: %v", err)
	}

	endTime, err := time.Parse(format, endDate)
	if err != nil {
		return 0, fmt.Errorf("failed to parse end date: %v", err)
	}

	duration := endTime.Sub(startTime)
	return duration, nil
}
