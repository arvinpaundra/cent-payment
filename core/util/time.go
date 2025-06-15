package util

import (
	"fmt"
	"time"
)

func StringToTime(layout, timeStr string) (time.Time, error) {
	t, err := time.Parse(layout, timeStr)
	if err != nil {
		return time.Time{}, fmt.Errorf("unable to parse %s", timeStr)
	}

	return t, nil
}
