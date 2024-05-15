package blackscholes

import (
	"math"
)

type CalculateProps struct {
	StrikePrice     float64
	UnderlyingPrice float64
	TimeToExp       float64 // years
	RiskFreeRate    float64
	Volatility      float64
}

func CalculateCall(props *CalculateProps) float64 {
	a := props.UnderlyingPrice * normDist(d1(props))
	b := props.StrikePrice * math.Exp(-props.RiskFreeRate*props.TimeToExp) * normDist(d2(props))

	return a - b
}

func CalculatePut(props *CalculateProps) float64 {
	a := props.StrikePrice * math.Exp(-props.RiskFreeRate*props.TimeToExp) * normDist(-d2(props))
	b := props.UnderlyingPrice * normDist(-d1(props))

	return a - b
}

func d1(props *CalculateProps) float64 {
	a := math.Log(props.UnderlyingPrice / props.StrikePrice)
	b := (props.RiskFreeRate + math.Pow(props.Volatility, 2)/2) * props.TimeToExp
	c := props.Volatility * math.Sqrt(props.TimeToExp)

	return (a + b) / c
}

func d2(props *CalculateProps) float64 {
	a := d1(props)
	b := props.Volatility * math.Sqrt(props.TimeToExp)
	return a - b
}

func normDist(x float64) float64 {
	return 0.5 * (1 + math.Erf(x/math.Sqrt2))
}
