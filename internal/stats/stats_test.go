package stats

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStdDev(t *testing.T) {
	t.Run("should calculate standard deviation", func(t *testing.T) {
		set := []float64{1, 2, 3, 4, 5}
		result := StdDev(set)

		assert.Equal(t, 1.4142135623730951, result)
	})
}

func TestVariance(t *testing.T) {
	t.Run("should calculate variance", func(t *testing.T) {
		set := []float64{1, 2, 3, 4, 5}
		result := variance(set)

		assert.Equal(t, 2.0, result)
	})

	t.Run("should calculate variance with empty set", func(t *testing.T) {
		set := []float64{}
		result := variance(set)

		assert.Equal(t, 0.0, result)
	})
}

func TestSquare(t *testing.T) {
	t.Run("should calculate square", func(t *testing.T) {
		set := []float64{1, 2, 3, 4, 5}
		result := square(set)

		assert.EqualValues(t, []float64{1, 4, 9, 16, 25}, result)
	})
}

func TestMean(t *testing.T) {
	t.Run("should calculate mean", func(t *testing.T) {
		set := []float64{1, 2, 3, 4, 5}
		result := mean(set)

		assert.Equal(t, 3.0, result)
	})

	t.Run("should calculate mean with empty set", func(t *testing.T) {
		set := []float64{}
		result := mean(set)

		assert.Equal(t, 0.0, result)
	})
}

func TestSum(t *testing.T) {
	t.Run("should calculate sum", func(t *testing.T) {
		set := []float64{1, 2, 3, 4, 5}
		result := sum(set)

		assert.Equal(t, 15.0, result)
	})
}

func TestMinus(t *testing.T) {
	t.Run("should calculate minus", func(t *testing.T) {
		set := []float64{1, 2, 3, 4, 5}
		result := minus(set, 2)

		assert.EqualValues(t, []float64{-1, 0, 1, 2, 3}, result)
	})
}
