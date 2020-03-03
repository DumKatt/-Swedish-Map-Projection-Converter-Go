package mapconv

import (
	"errors"
	"math"
)

// Conversion from geodetic coordinates to grid coordinates.
func (s settings) GeodeticToGrid(cordinate GeodeticCoordinate) (RT90Cordinate, error) {
	if cordinate.Validate() {
		return RT90Cordinate{}, errors.New("GeodeticCoordinates is not valid")
	}
	if s.centralMeridian == nil {
		return RT90Cordinate{}, errors.New("centeral meridian is not set")
	}
	// Prepare ellipsoid-based stuff.
	e2 := s.flattening * (2.0 - s.flattening)
	n := s.flattening / (2.0 - s.flattening)
	aRoof := s.axis / (1.0 + n) * (1.0 + n*n/4.0 + n*n*n*n/64.0)
	A := e2
	B := (5.0*e2*e2 - e2*e2*e2) / 6.0
	C := (104.0*e2*e2*e2 - 45.0*e2*e2*e2*e2) / 120.0
	D := (1237.0 * e2 * e2 * e2 * e2) / 1260.0
	beta1 := n/2.0 - 2.0*n*n/3.0 + 5.0*n*n*n/16.0 + 41.0*n*n*n*n/180.0
	beta2 := 13.0*n*n/48.0 - 3.0*n*n*n/5.0 + 557.0*n*n*n*n/1440.0
	beta3 := 61.0*n*n*n/240.0 - 103.0*n*n*n*n/140.0
	beta4 := 49561.0 * n * n * n * n / 161280.0

	// Convert.
	degToRad := math.Pi / 180.0
	phi := cordinate.Latitude * degToRad
	lambda := cordinate.Longitude * degToRad
	lambdaZero := *s.centralMeridian * degToRad

	phiStar := phi - math.Sin(phi)*math.Cos(phi)*(A+
		B*math.Pow(math.Sin(phi), 2)+
		C*math.Pow(math.Sin(phi), 4)+
		D*math.Pow(math.Sin(phi), 6))
	deltaLambda := lambda - lambdaZero
	xiPrim := math.Atan(math.Tan(phiStar) / math.Cos(deltaLambda))
	etaPrim := math.Atanh(math.Cos(phiStar) * math.Sin(deltaLambda))
	x := s.scale*aRoof*(xiPrim+
		beta1*math.Sin(2.0*xiPrim)*math.Cosh(2.0*etaPrim)+
		beta2*math.Sin(4.0*xiPrim)*math.Cosh(4.0*etaPrim)+
		beta3*math.Sin(6.0*xiPrim)*math.Cosh(6.0*etaPrim)+
		beta4*math.Sin(8.0*xiPrim)*math.Cosh(8.0*etaPrim)) + s.falseNorthing
	y := s.scale*aRoof*(etaPrim+
		beta1*math.Cos(2.0*xiPrim)*math.Sinh(2.0*etaPrim)+
		beta2*math.Cos(4.0*xiPrim)*math.Sinh(4.0*etaPrim)+
		beta3*math.Cos(6.0*xiPrim)*math.Sinh(6.0*etaPrim)+
		beta4*math.Cos(8.0*xiPrim)*math.Sinh(8.0*etaPrim)) + s.falseEasting

	return RT90Cordinate{
		X: math.Round(x*1000.0) / 1000.0,
		Y: math.Round(y*1000.0) / 1000.0,
	}, nil
}

// Conversion from grid coordinates to geodetic coordinates.
func (s settings) GridToGeodetic(coordinate RT90Cordinate) (GeodeticCoordinate, error) {
	if s.centralMeridian == nil {
		return GeodeticCoordinate{}, errors.New("centeral meridian is not set")
	}

	// Prepare ellipsoid-based stuff.
	e2 := s.flattening * (2.0 - s.flattening)
	n := s.flattening / (2.0 - s.flattening)
	aRoof := s.axis / (1.0 + n) * (1.0 + n*n/4.0 + n*n*n*n/64.0)
	delta1 := n/2.0 - 2.0*n*n/3.0 + 37.0*n*n*n/96.0 - n*n*n*n/360.0
	delta2 := n*n/48.0 + n*n*n/15.0 - 437.0*n*n*n*n/1440.0
	delta3 := 17.0*n*n*n/480.0 - 37*n*n*n*n/840.0
	delta4 := 4397.0 * n * n * n * n / 161280.0

	Astar := e2 + e2*e2 + e2*e2*e2 + e2*e2*e2*e2
	Bstar := -(7.0*e2*e2 + 17.0*e2*e2*e2 + 30.0*e2*e2*e2*e2) / 6.0
	Cstar := (224.0*e2*e2*e2 + 889.0*e2*e2*e2*e2) / 120.0
	Dstar := -(4279.0 * e2 * e2 * e2 * e2) / 1260.0

	// Convert.
	degToRad := math.Pi / 180
	lambdaZero := *s.centralMeridian * degToRad
	xi := (coordinate.X - s.falseNorthing) / (s.scale * aRoof)
	eta := (coordinate.Y - s.falseEasting) / (s.scale * aRoof)
	xiPrim := (xi -
		delta1*math.Sin(2.0*xi)*math.Cosh(2.0*eta) -
		delta2*math.Sin(4.0*xi)*math.Cosh(4.0*eta) -
		delta3*math.Sin(6.0*xi)*math.Cosh(6.0*eta) -
		delta4*math.Sin(8.0*xi)*math.Cosh(8.0*eta))

	etaPrim := (eta -
		delta1*math.Cos(2.0*xi)*math.Sinh(2.0*eta) -
		delta2*math.Cos(4.0*xi)*math.Sinh(4.0*eta) -
		delta3*math.Cos(6.0*xi)*math.Sinh(6.0*eta) -
		delta4*math.Cos(8.0*xi)*math.Sinh(8.0*eta))

	phiStar := math.Asin(math.Sin(xiPrim) / math.Cosh(etaPrim))
	deltaLambda := math.Atan(math.Sinh(etaPrim) / math.Cos(xiPrim))
	lonRadian := lambdaZero + deltaLambda
	latRadian := (phiStar + math.Sin(phiStar)*math.Cos(phiStar)*
		(Astar+
			Bstar*math.Pow(math.Sin(phiStar), 2)+
			Cstar*math.Pow(math.Sin(phiStar), 4)+
			Dstar*math.Pow(math.Sin(phiStar), 6)))

	return GeodeticCoordinate{
		Latitude:  latRadian * 180.0 / math.Pi,
		Longitude: lonRadian * 180.0 / math.Pi,
	}, nil
}
