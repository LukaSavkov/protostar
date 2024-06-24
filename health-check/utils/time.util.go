package utils

import (
	"time"
)

func FormatTime(millis int64) string {

	// Convert milliseconds to seconds and nanoseconds
	seconds := millis / 1000
	nanoseconds := (millis % 1000) * 1e6

	// Create a time.Time structure
	t := time.Unix(seconds, nanoseconds)

	// Format the time as a string
	formattedTime := t.Format(time.RFC3339Nano)

	// Return the formatted time string
	return formattedTime
}
