package stats

import "math"

// StdDev calculates the population standard deviation.
func StdDev(data []float64) float64 {
	v := variance(data)
	return math.Sqrt(v)
}

// variance calculates the population variance.
func variance(data []float64) float64 {
	if len(data) == 0 {
		return 0
	}

	avg := mean(data)
	values := minus(data, avg)
	values = square(values)
	total := sum(values)

	return total / float64(len(values))
}

func square(data []float64) []float64 {
	var squared []float64
	for _, value := range data {
		squared = append(squared, value*value)
	}
	return squared
}

func mean(data []float64) float64 {
	if len(data) == 0 {
		return 0
	}

	total := sum(data)

	return total / float64(len(data))
}

func sum(data []float64) float64 {
	var sum float64
	for _, value := range data {
		sum += value
	}
	return sum
}

func minus(data []float64, value float64) []float64 {
	var minus []float64
	for _, val := range data {
		minus = append(minus, val-value)
	}
	return minus
}
