package model

import (
	"encoding/json"
)

// Interval is a value object that represents a time interval.
type Interval string

func (i *Interval) String() string {
	return string(*i)
}

func (i *Interval) UnmarshalJSON(b []byte) error {
	var a string
	err := json.Unmarshal(b, &a)
	if err != nil {
		return err
	}

	*i = Interval(a)
	return nil
}

// Seconds returns the seconds of the interval.
func (i *Interval) Seconds() int {
	m, ok := SupportedIntervals[*i]
	if !ok {
		return parseInterval(*i)
	}

	return m
}

var (
	Interval1s  = Interval("1s")
	Interval1m  = Interval("1m")
	Interval3m  = Interval("3m")
	Interval5m  = Interval("5m")
	Interval15m = Interval("15m")
	Interval30m = Interval("30m")
	Interval1h  = Interval("1h")
	Interval2h  = Interval("2h")
	Interval4h  = Interval("4h")
	Interval6h  = Interval("6h")
	Interval12h = Interval("12h")
	Interval1d  = Interval("1d")
	Interval3d  = Interval("3d")
	Interval1w  = Interval("1w")
	Interval2w  = Interval("2w")
	Interval1mo = Interval("1mo")
)

// IntervalMap is a map of intervals.
type IntervalMap map[Interval]int

// SupportedIntervals is a map of supported intervals.
var SupportedIntervals = IntervalMap{
	Interval1s:  1,
	Interval1m:  1 * 60,
	Interval3m:  3 * 60,
	Interval5m:  5 * 60,
	Interval15m: 15 * 60,
	Interval30m: 30 * 60,
	Interval1h:  60 * 60,
	Interval2h:  60 * 60 * 2,
	Interval4h:  60 * 60 * 4,
	Interval6h:  60 * 60 * 6,
	Interval12h: 60 * 60 * 12,
	Interval1d:  60 * 60 * 24,
	Interval3d:  60 * 60 * 24 * 3,
	Interval1w:  60 * 60 * 24 * 7,
	Interval2w:  60 * 60 * 24 * 14,
	Interval1mo: 60 * 60 * 24 * 30,
}

func parseInterval(v Interval) int {
	t := 0
	index := 0

	for i, rn := range string(v) {
		if rn >= '0' && rn <= '9' {
			t = t*10 + int(rn-'0')
		} else {
			index = i
			break
		}
	}

	switch string(v[index:]) {
	case "s":
		return t
	case "m":
		t *= 60
	case "h":
		t *= 60 * 60
	case "d":
		t *= 60 * 60 * 24
	case "w":
		t *= 60 * 60 * 24 * 7
	case "mo":
		t *= 60 * 60 * 24 * 30
	default:
		panic("unknown interval input: " + v)
	}

	return t
}
