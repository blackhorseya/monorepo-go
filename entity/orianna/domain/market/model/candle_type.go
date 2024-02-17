package model

// CandleType is an enum that represents the type of candle.
type CandleType int

const (
	CandleTypeUnknown CandleType = iota
	CandleTypeLongUp
	CandleTypeLongDown
)

var candleTypeMap = map[CandleType]string{
	CandleTypeUnknown:  "unknown",
	CandleTypeLongUp:   "long_up",
	CandleTypeLongDown: "long_down",
}

// String returns the string representation of the candle type.
func (x *CandleType) String() string {
	return candleTypeMap[*x]
}
