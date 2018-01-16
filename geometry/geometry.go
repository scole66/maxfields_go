// The geometry code for the maxfield project.
// It implements all of the geometry calculations for dealing with triangle intersections, calculating great circles,
// and generally dealing with latitude and longitude.
package geometry

import (
	"math"
)

// LatLong is a simple two-value struct, used to represent latitude and longitude. Note that we don't specify radians
// or degrees here; please document your variables to indicate which set of units are being used!
type LatLong struct {
	Latitude  float64
	Longitude float64
}

// XYZ is a three-value struct, used to represent rectilinear coordinates. Units are specified by the user.
type XYZ struct {
	X, Y, Z float64
}

// LLtoRads takes a slice of LatLong pairs with units in degrees and converts them into equivalent LatLong pairs with
// units in radians. This is an in-place conversion.
func LLtoRads(pts []LatLong) {
	const radiansPerDegree = math.Pi / 180.0
	for i := range pts {
		pts[i].Latitude *= radiansPerDegree
		pts[i].Longitude *= radiansPerDegree
	}
}

// RadsToXYZ takes a slice of LatLong pairs (measured in radians), and returns a new slice of XYZ triples after having
// done the conversion to rectilinear coordinates. The initial sphere is the unit sphere, so the range of XYZ will be
// from -1 to 1.
func RadsToXYZ(pts []LatLong) []XYZ {
	rval := make([]XYZ, len(pts))
	for i, pt := range pts {
		// Radius from Axis:
		radius := math.Cos(pt.Latitude)

		// And then the XYZ values. Note that we assume we're doing lat/long on the unit sphere.
		rval[i] = XYZ{
			math.Cos(pt.Longitude) * radius,
			math.Sin(pt.Longitude) * radius,
			math.Sin(pt.Latitude)}
	}
	return rval
}

// GreatArcAngle, given two points on a sphere (as latitude and longitude, in radians), calculate the angle between
// them (in a great-arc sense), in radians.
func GreatArcAngle(x, y LatLong) float64 {
	deltaLongitude := math.Abs(x.Longitude - y.Longitude)
	sind := math.Sin(deltaLongitude)
	cosd := math.Cos(deltaLongitude)
	cosx := math.Cos(x.Latitude)
	cosy := math.Cos(y.Latitude)
	sinx := math.Sin(x.Latitude)
	siny := math.Sin(y.Latitude)

	// From wikipedia, the article on Great-circle distance
	part1 := cosy * sind
	part2 := cosx*siny - sinx*cosy*cosd
	numer := math.Sqrt(part1*part1 + part2*part2)
	denom := sinx*siny + cosx*cosy*cosd
	return math.Atan2(numer, denom)
}
