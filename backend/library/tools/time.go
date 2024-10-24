package tools

import (
	"fmt"
	"time"
)

func ConvertToUnix(timeStr string) (int64, error) {
	layout := time.RFC3339Nano

	// Parse the time string
	parsedTime, err := time.Parse(layout, timeStr)
	if err != nil {
		return 0, fmt.Errorf("error parsing time: %v", err)
	}

	// Convert to Unix time
	return parsedTime.Unix(), nil
}

func ConvertToTime(timeStr string) (time.Time, error) {
	layout := time.RFC3339Nano

	// Parse the time string
	parsedTime, err := time.Parse(layout, timeStr)
	if err != nil {
		return time.Time{}, fmt.Errorf("error parsing time: %v", err)
	}

	return parsedTime, nil
}
