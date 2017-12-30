package geometry

import (
    "math"
    "testing"
)

func TestLLtoRads(t *testing.T) {
    pts := []LatLong{{0.0, 0.0}, {90.0, 0.0}, {-90.0, 0.0}, {0.0, 180.0}}
    expected := []LatLong{{0.0, 0.0}, {math.Pi/2, 0.0}, {-math.Pi/2, 0.0}, {0.0, math.Pi}}
    LLtoRads(pts)
    for i := range(pts) {
        if pts[i] != expected[i] {
            t.Errorf("Saw %v, but wanted %v", pts[i], expected[i])
        }
    }
}

func TestRadsToXYZ(t *testing.T) {
    pts := []LatLong{{90.0, 0.0}, // North Pole: XYZ = (0,0,1)
                     {-90.0, 0.0}, // South Pole: XYZ = (0,0,-1)
                     {0.0, 0.0}, // Intersection of Meridian and Equator: XYZ = (1,0,0)
                     {0.0, 180.0}, // Other side of the globe from there: XYZ = (-1,0,0)
                     {0.0, -180.0}, // And in the opposite direction: XYZ = (-1,0,0)
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
    for i := range(xyzs) {
        if xyzs[i] != expected[i] {
            t.Errorf("Actual result (%v) did not match expected result (%v)", xyzs[i], expected[i])
        }
    }
}
