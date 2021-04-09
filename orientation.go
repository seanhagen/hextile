package hextile

import "math"

/**
 * File: orientation.go
 * Date: 2021-04-08 23:55:10
 * Creator: Sean Hagen <sean.hagen@gmail.com>
 */

// Orientation ...
type Orientation struct {
	f0, f1, f2, f3 float64
	b0, b1, b2, b3 float64
	// in multiples of 60°
	startAngle float64
}

// LayoutPointy
var LayoutPointy = Orientation{
	math.Sqrt(3.0), math.Sqrt(3.0) / 2.0, 0.0, 3.0 / 2.0,
	math.Sqrt(3.0) / 3.0, -1.0 / 3.0, 0.0, 2.0 / 3.0,
	0.5,
}

var LayoutFlat = Orientation{
	3.0 / 2.0, 0.0, math.Sqrt(3.0) / 2.0, math.Sqrt(3.0),
	2.0 / 3.0, 0.0, -1.0 / 3.0, math.Sqrt(3.0) / 3.0,
	0.0,
}
