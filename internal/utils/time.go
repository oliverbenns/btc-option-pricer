package utils

import "time"

func GetYearsFromDuration(duration time.Duration) float64 {
	yearDuration := time.Duration(365.25 * 24 * time.Hour)

	years := float64(duration) / float64(yearDuration)

	return years
}
