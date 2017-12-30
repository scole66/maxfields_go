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
    Latitude float64
    Longitude float64
}

// XYZ is a three-value struct, used to represent rectilinear coordinates. Units are specified by the user.
type XYZ struct {
    X, Y, Z float64
}

// LLtoRads takes a slice of LatLong pairs with units in degrees and converts them into equivalent LatLong pairs with
// units in radians. This is an in-place conversion.
func LLtoRads(pts []LatLong) {
    factor := math.Pi / 180.0
    for i := range pts {
        pts[i].Latitude *= factor
        pts[i].Longitude *= factor
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
