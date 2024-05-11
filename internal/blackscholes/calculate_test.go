package blackscholes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateCall(t *testing.T) {
	t.Run("should calculate call", func(t *testing.T) {
		props := &CalculateProps{
			StrikePrice:     100,
			UnderlyingPrice: 100,
			TimeToExp:       1,
			RiskFreeRate:    0.05,
			Volatility:      0.2,
		}

		result := CalculateCall(props)

		assert.Equal(t, 10.450583572185568, result)
	})
}
