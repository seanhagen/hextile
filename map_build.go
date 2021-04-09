package hextile

import (
	"math"
)

/**
 * File: map_build.go
 * Date: 2021-04-09 15:30:57
 * Creator: Sean Hagen <sean.hagen@gmail.com>
 */

func newMap() Map {
	return Map{
		data:  map[Hex]interface{}{},
		tiles: map[Hex]bool{},
	}
}

// NewParallelogramLeftMap ...
func NewParallelogramLeftMap(q1, q2, r1, r2 int) Map {
	m := newMap()

	for q := q1; q <= q1; q++ {
		for r := r1; r <= r2; r++ {
			h := NewHex(q, r)
			m.AddHex(h)
		}
	}

	return m
}

// NewParallelogramRightMap ...
func NewParallelogramRightMap(r1, r2, s1, s2 int) Map {
	m := newMap()

	for r := r1; r <= r2; r++ {
		for s := s1; s <= s2; s++ {
			h := NewHexS(-s-r, r, s)
			m.AddHex(h)
		}
	}

	return m
}

func NewTriangleDown(size int) Map {
	m := newMap()

	for q := 0; q <= size; q++ {
		for r := 0; r <= size-q; r++ {
			h := NewHex(q, r)
			m.AddHex(h)
		}
	}

	return m
}

func NewHexagonMap(radius int) Map {
	m := newMap()

	for q := -radius; q <= radius; q++ {
		r1 := max(-radius, -q-radius)
		r2 := min(radius, -q+radius)
		for r := r1; r <= r2; r++ {
			h := NewHex(q, r)
			m.AddHex(h)
		}
	}

	return m
}

// NewRectangleMap ...
func NewRectangleMap(height, width int) Map {
	m := newMap()
	//(s,r)
	for r := 0; r < height; r++ {
		offset := int(math.Floor(float64(r) / 2))
		for q := -offset; q < width-offset; q++ {
			h := NewHex(q, r)
			m.AddHex(h)
		}
	}
	return m
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
