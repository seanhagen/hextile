package hextile

/**
 * File: layout.go
 * Date: 2021-04-08 23:58:04
 * Creator: Sean Hagen <sean.hagen@gmail.com>
 */

// Layout ...
type Layout struct {
	orientation Orientation
	size        Point
	origin      Point
}

// NewLayout ...
func NewLayout(orientation Orientation, size Point, origin Point) Layout {
	return Layout{orientation, size, origin}
}
