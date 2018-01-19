package geometry

import (
	"math"
	"testing"
)

func TestDegreesToRadians(t *testing.T) {
	var cases = []struct {
		input   float64 // Input Argument
		desired float64 // Expected Result
	}{
		{0.0, 0.0},
		{90.0, math.Pi / 2},
		{-90.0, -math.Pi / 2},
		{180.0, math.Pi},
	}
	for _, testcase := range cases {
		actual := DegreesToRadians(testcase.input)
		if actual != testcase.desired {
			t.Errorf("Saw %v, but wanted %v", actual, testcase.desired)
		}
	}
}

func TestLLtoRads(t *testing.T) {
	pts := []LatLong{{0.0, 0.0}, {90.0, 0.0}, {-90.0, 0.0}, {0.0, 180.0}}
	expected := []LatLong{{0.0, 0.0}, {math.Pi / 2, 0.0}, {-math.Pi / 2, 0.0}, {0.0, math.Pi}}
	LLtoRads(pts)
	for i := range pts {
		if pts[i] != expected[i] {
			t.Errorf("Saw %v, but wanted %v", pts[i], expected[i])
		}
	}
}

func TestRadsToXYZ(t *testing.T) {
	pts := []LatLong{{90.0, 0.0}, // North Pole: XYZ = (0,0,1)
		{-90.0, 0.0},         // South Pole: XYZ = (0,0,-1)
		{0.0, 0.0},           // Intersection of Meridian and Equator: XYZ = (1,0,0)
		{0.0, 180.0},         // Other side of the globe from there: XYZ = (-1,0,0)
		{0.0, -180.0},        // And in the opposite direction: XYZ = (-1,0,0)
		{37.7949, -122.4017}} // Site of the wreck of the Niantic
	// But RadsToXYZ needs radians, so convert that:
	LLtoRads(pts)

	expected := []XYZ{{0.0, 0.0, 1.0}, // North Pole
		{0.0, 0.0, -1.0}, // South Pole
		{1.0, 0.0, 0.0},
		{-1.0, 0.0, 0.0},
		{-1.0, 0.0, 0.0},
		{-0.4234352546006439, -0.6671834396143197, 0.6128367181115155}}

	xyzs := RadsToXYZ(pts)
	if len(pts) != len(xyzs) {
		t.Errorf("Result length incorrect (%d elements came back; we wanted %d)", len(xyzs), len(pts))
	}
	for i := range xyzs {
		dx := math.Abs(xyzs[i].X - expected[i].X)
		dy := math.Abs(xyzs[i].Y - expected[i].Y)
		dz := math.Abs(xyzs[i].Z - expected[i].Z)
		if dx > 1.0e-15 || dy > 1.0e-15 || dz > 1.0e-15 {
			t.Errorf("Actual result (%v) did not match expected result (%v)", xyzs[i], expected[i])
		}
	}
}

func TestGreatArcAngle(t *testing.T) {
	const earthRadius = 6371000.0 // Earth's radius in meters
	var cases = []struct {
		lat1, long1 float64 // input: latitude and longitude #1
		lat2, long2 float64 // input: latitude and longitude #2
		expected    float64 // expected: the angle between them
	}{
		{0.0, 0.0, math.Pi / 2.0, 0.0, math.Pi / 2.0 * earthRadius},
		{0.6512790002530597, -2.1275756615226147, 0.6512826479911965, -2.1275634791244356, 66.0},
	}

	for _, example := range cases {
		angle := GreatArcAngle(LatLong{example.lat1, example.long1}, LatLong{example.lat2, example.long2})
		distance := math.Abs(angle * earthRadius) // in meters
		if math.Abs(distance-example.expected) > 0.5 {
			t.Errorf("Great Arc Angle for (%g,%g)--(%g,%g) calculated as %g; expected %g",
				example.lat1, example.long1, example.lat2, example.long2, angle, example.expected/earthRadius)
		}
	}
}
