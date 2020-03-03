package mapconv

type GeodeticCoordinate struct {
	Latitude  float64
	Longitude float64
}

func (cord GeodeticCoordinate) Validate() bool {
	if cord.Latitude < -90 || cord.Latitude > 90 {
		return false
	}
	if cord.Longitude < -180 || cord.Longitude > 180 {
		return false
	}
	return true
}

type RT90Cordinate struct {
	X float64
	Y float64
}
