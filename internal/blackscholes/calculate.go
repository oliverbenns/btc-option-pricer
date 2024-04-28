package blackscholes

type CalculateProps struct {
	StrikePrice     float64
	UnderlyingPrice float64
	TimeToExp       float64 // years
	RiskFreeRate    float64
	Volatility      float64
}

func Calculate(props *CalculateProps) float64 {
	return 0.0
}
