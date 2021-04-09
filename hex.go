package hextile

/**
 * File: hex.go
 * Date: 2021-04-08 22:48:30
 * Creator: Sean Hagen <sean.hagen@gmail.com>
 */

import (
	"fmt"
	"math"
)

const nsid = "4c47310a-1e9f-4ffb-8897-cd053e60e0c1"

type Direction int32

const (
	DirectionRight Direction = iota
	DirectionRightUp
	DirectionLeftUp
	DirectionLeft
	DirectionLeftDown
	DirectionRightDown
)

var (
	directionRight     = NewHex(1, 0)
	directionRightUp   = NewHex(1, -1)
	directionLeftUp    = NewHex(0, -1)
	directionLeft      = NewHex(-1, 0)
	directionLeftDown  = NewHex(-1, 1)
	directionRightDown = NewHex(0, 1)
)

var directionToHex = map[Direction]Hex{
	DirectionRight:     directionRight,
	DirectionRightUp:   directionRightUp,
	DirectionLeftUp:    directionLeftUp,
	DirectionLeft:      directionLeft,
	DirectionLeftDown:  directionLeftDown,
	DirectionRightDown: directionRightDown,
}

var directionName = map[Direction]string{
	DirectionRight:     "Right",
	DirectionRightUp:   "Right Up",
	DirectionLeftUp:    "Left Up",
	DirectionLeft:      "Left",
	DirectionLeftDown:  "Left Down",
	DirectionRightDown: "Right Down",
}

// Hex stores info about a hex in axial coordinates
type Hex struct {
	q int
	r int
	s int
}

// NewHex ...
func NewHex(q, r int) Hex {
	return NewHexS(q, r, -q-r)
}

// NewHexS ...
func NewHexS(q, r, s int) Hex {
	return Hex{
		q: q,
		r: r,
		s: s,
	}
}

// Q ...
func (h Hex) Q() int {
	return h.q
}

// R ...
func (h Hex) R() int {
	return h.r
}

// S ...
func (h Hex) S() int {
	return h.s
}

// Eq ...
func (h Hex) Eq(o Hex) bool {
	return h.q == o.q && h.r == o.r && h.s == o.s
}

// Add ...
func (h Hex) Add(o Hex) Hex {
	nq := h.q + o.q
	nr := h.r + o.r
	ns := h.s + o.s
	return Hex{nq, nr, ns}
}

// Subtract ...
func (h Hex) Subtract(o Hex) Hex {
	nq := h.q - o.q
	nr := h.r - o.r
	ns := h.s - o.s
	return Hex{nq, nr, ns}
}

// Multiply ...
func (h Hex) Multiply(k int) Hex {
	nq := h.q * k
	nr := h.r * k
	ns := h.s * k
	return Hex{nq, nr, ns}
}

// Length ...
func (h Hex) Length() int {
	aq := math.Abs(float64(h.q))
	ar := math.Abs(float64(h.r))
	as := math.Abs(float64(h.s))
	d := (aq + ar + as) / 2
	return int(d)
}

// Distance ...
func Distance(a, b Hex) int {
	s := a.Subtract(b)
	return s.Length()
}

// GetDirection ...
func GetDirection(d Direction) Hex {
	if d < 0 || d > 6 {
		return Hex{}
	}

	return directionToHex[d]
}

// GetHexNeighbour ...
func GetHexNeighbour(h Hex, d Direction) Hex {
	return h.Add(GetDirection(d))
}

// Neighbour ...
func (h Hex) Neighbour(d Direction) Hex {
	x := GetDirection(d)
	return h.Add(x)
}

// String ...
func (h Hex) String() string {
	return fmt.Sprintf("{%v %v}", h.q, h.r)
}
