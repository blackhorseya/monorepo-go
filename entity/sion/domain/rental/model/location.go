package model

import (
	"math"
)

// DistanceTo will calculate the distance between two locations.
func (x *Location) DistanceTo(other *Location, unit string) float64 {
	radlat1 := math.Pi * x.Latitude / 180
	radlat2 := math.Pi * other.Latitude / 180

	theta := x.Longitude - other.Longitude
	radtheta := math.Pi * theta / 180

	dist := math.Sin(radlat1)*math.Sin(radlat2) + math.Cos(radlat1)*math.Cos(radlat2)*math.Cos(radtheta)

	if dist > 1 {
		dist = 1
	}

	dist = math.Acos(dist)
	dist = dist * 180 / math.Pi
	dist = dist * 60 * 1.1515

	if unit == "K" {
		dist = dist * 1.609344
	}

	if unit == "N" {
		dist = dist * 0.8684
	}

	return dist
}
