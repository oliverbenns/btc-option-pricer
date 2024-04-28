package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetYearsFromDuration(t *testing.T) {
	startDate := time.Date(2018, time.January, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2019, time.July, 2, 0, 0, 0, 0, time.UTC)
	duration := endDate.Sub(startDate)

	result := GetYearsFromDuration(duration)

	// good enough
	assert.Equal(t, 1.4976043805612593, result)
}
