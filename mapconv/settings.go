package mapconv

type settings struct {
	axis            float64
	flattening      float64
	centralMeridian *float64
	latOfOrigin     float64
	scale           float64
	falseNorthing   float64
	falseEasting    float64
}

type ProjectionType string

const (
	RT90_75_GonV        ProjectionType = "rt90_7.5_gon_v"
	RT90_50_GonV        ProjectionType = "rt90_5.0_gon_v"
	RT90_25_GonV        ProjectionType = "rt90_2.5_gon_v"
	RT90_00_GonV        ProjectionType = "rt90_0.0_gon_v"
	RT90_25_GonO        ProjectionType = "rt90_2.5_gon_o"
	RT90_50_GonO        ProjectionType = "rt90_5.0_gon_o"
	Bessel_RT90_75_GonV ProjectionType = "bessel_rt90_7.5_gon_v"
	Bessel_RT90_50_GonV ProjectionType = "bessel_rt90_5.0_gon_v"
	Bessel_RT90_25_GonV ProjectionType = "bessel_rt90_2.5_gon_v"
	Bessel_RT90_00_GonV ProjectionType = "bessel_rt90_0.0_gon_v"
	Bessel_RT90_25_GonO ProjectionType = "bessel_rt90_2.5_gon_o"
	Bessel_RT90_50_GonO ProjectionType = "bessel_rt90_5.0_gon_o"
	SWEREF_99_TM        ProjectionType = "sweref_99_tm"
	SWEREF_99_1200      ProjectionType = "sweref_99_1200"
	SWEREF_99_1330      ProjectionType = "sweref_99_1330"
	SWEREF_99_1500      ProjectionType = "sweref_99_1500"
	SWEREF_99_1630      ProjectionType = "sweref_99_1630"
	SWEREF_99_1800      ProjectionType = "sweref_99_1800"
	SWEREF_99_1415      ProjectionType = "sweref_99_1415"
	SWEREF_99_1545      ProjectionType = "sweref_99_1545"
	SWEREF_99_1715      ProjectionType = "sweref_99_1715"
	SWEREF_99_1845      ProjectionType = "sweref_99_1845"
	SWEREF_99_2015      ProjectionType = "sweref_99_2015"
	SWEREF_99_2145      ProjectionType = "sweref_99_2145"
	SWEREF_99_2315      ProjectionType = "sweref_99_2315"
)

func (s settings) SetProjectionType(projection ProjectionType) {
	// RT90 parameters, GRS 80 ellipsoid.
	switch projection {
	case "rt90_7.5_gon_v":
		s.setGRS80()
		*s.centralMeridian = 11.0 + 18.375/60.0
		s.scale = 1.000006000000
		s.falseNorthing = -667.282
		s.falseEasting = 1500025.141
	case "rt90_5.0_gon_v":
		s.setGRS80()
		*s.centralMeridian = 13.0 + 33.376/60.0
		s.scale = 1.000005800000
		s.falseNorthing = -667.130
		s.falseEasting = 1500044.695
	case "rt90_2.5_gon_v":
		s.setGRS80()
		*s.centralMeridian = 15.0 + 48.0/60.0 + 22.624306/3600.0
		s.scale = 1.00000561024
		s.falseNorthing = -667.711
		s.falseEasting = 1500064.274
	case "rt90_0.0_gon_v":
		s.setGRS80()
		*s.centralMeridian = 18.0 + 3.378/60.0
		s.scale = 1.000005400000
		s.falseNorthing = -668.844
		s.falseEasting = 1500083.521
	case "rt90_2.5_gon_o":
		s.setGRS80()
		*s.centralMeridian = 20.0 + 18.379/60.0
		s.scale = 1.000005200000
		s.falseNorthing = -670.706
		s.falseEasting = 1500102.765
	case "rt90_5.0_gon_o":
		s.setGRS80()
		*s.centralMeridian = 22.0 + 33.380/60.0
		s.scale = 1.000004900000
		s.falseNorthing = -672.557
		s.falseEasting = 1500121.846
	// RT90 parameters, Bessel 1841 ellipsoid.
	case "bessel_rt90_7.5_gon_v":
		s.setBessel()
		*s.centralMeridian = 11.0 + 18.0/60.0 + 29.8/3600.0
	case "bessel_rt90_5.0_gon_v":
		s.setBessel()
		*s.centralMeridian = 13.0 + 33.0/60.0 + 29.8/3600.0
	case "bessel_rt90_2.5_gon_v":
		s.setBessel()
		*s.centralMeridian = 15.0 + 48.0/60.0 + 29.8/3600.0
	case "bessel_rt90_0.0_gon_v":
		s.setBessel()
		*s.centralMeridian = 18.0 + 3.0/60.0 + 29.8/3600.0
	case "bessel_rt90_2.5_gon_o":
		s.setBessel()
		*s.centralMeridian = 20.0 + 18.0/60.0 + 29.8/3600.0
	case "bessel_rt90_5.0_gon_o":
		s.setBessel()
		*s.centralMeridian = 22.0 + 33.0/60.0 + 29.8/3600.0

	// SWEREF99TM and SWEREF99ddmm  parameters.
	case "sweref_99_tm":
		s.setSWEREF99()
		*s.centralMeridian = 15.00
		s.latOfOrigin = 0.0
		s.scale = 0.9996
		s.falseNorthing = 0.0
		s.falseEasting = 500000.0
	case "sweref_99_1200":
		s.setSWEREF99()
		*s.centralMeridian = 12.00
	case "sweref_99_1330":
		s.setSWEREF99()
		*s.centralMeridian = 13.50
	case "sweref_99_1500":
		s.setSWEREF99()
		*s.centralMeridian = 15.00
	case "sweref_99_1630":
		s.setSWEREF99()
		*s.centralMeridian = 16.50
	case "sweref_99_1800":
		s.setSWEREF99()
		*s.centralMeridian = 18.00
	case "sweref_99_1415":
		s.setSWEREF99()
		*s.centralMeridian = 14.25
	case "sweref_99_1545":
		s.setSWEREF99()
		*s.centralMeridian = 15.75
	case "sweref_99_1715":
		s.setSWEREF99()
		*s.centralMeridian = 17.25
	case "sweref_99_1845":
		s.setSWEREF99()
		*s.centralMeridian = 18.75
	case "sweref_99_2015":
		s.setSWEREF99()
		*s.centralMeridian = 20.25
	case "sweref_99_2145":
		s.setSWEREF99()
		*s.centralMeridian = 21.75
	case "sweref_99_2315":
		s.setSWEREF99()
		*s.centralMeridian = 23.25

	// Test-case:11
	//	Lat: 66 0'0", lon: 24 0'0".
	//	X:1135809.413803 Y:555304.016555.
	case "test_case":
		s.axis = 6378137.0
		s.flattening = 1.0 / 298.257222101
		*s.centralMeridian = 13.0 + 35.0/60.0 + 7.692000/3600.0
		s.latOfOrigin = 0.0
		s.scale = 1.000002540000
		s.falseNorthing = -6226307.8640
		s.falseEasting = 84182.8790
		// Not a valid projection.
	default:
		s.centralMeridian = nil
	}
}

// Sets of default parameters.
func (s settings) setGRS80() {
	s.axis = 6378137.0                 // GRS 80.
	s.flattening = 1.0 / 298.257222101 // GRS 80.
	s.centralMeridian = nil
	s.latOfOrigin = 0.0
}

func (s settings) setBessel() {
	s.axis = 6377397.155             // Bessel 1841.
	s.flattening = 1.0 / 299.1528128 // Bessel 1841.
	s.centralMeridian = nil
	s.latOfOrigin = 0.0
	s.scale = 1.0
	s.falseNorthing = 0.0
	s.falseEasting = 1500000.0
}

func (s settings) setSWEREF99() {
	s.axis = 6378137.0                 // GRS 80.
	s.flattening = 1.0 / 298.257222101 // GRS 80.
	s.centralMeridian = nil
	s.latOfOrigin = 0.0
	s.scale = 1.0
	s.falseNorthing = 0.0
	s.falseEasting = 150000.0
}
