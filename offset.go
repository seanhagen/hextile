package hextile

/**
 * File: offset.go
 * Date: 2021-04-08 23:04:19
 * Creator: Sean Hagen <sean.hagen@gmail.com>
 */

// OffsetType is the type used to determine which way to slide hex tiles,
// depending on orientation
type OffsetType int32

const (
	// Even is an argument for the Hex <-> OffsetCoord functions
	Even OffsetType = 1
	// Odd is an argument for the Hex <-> OffsetCoord functions
	Odd = -1
)

// OffsetCoord records the coordinates in 'offset' notation, -- ie, column and row
//
// From the reference article on implementation:
//
//   There are four offset types: odd-r, even-r, odd-q, even-q. The “r”
//   types are used with with pointy top hexagons and the “q” types are
//   used with flat top. Whether it’s even or odd can be encoded as an
//   offset direction +1 or -1. For pointy top, the offset direction tells
//   us whether to slide alternate rows right or left. For flat top, the
//   offset direction tells us whether to slide alternate columns up or down.
type OffsetCoord struct {
	Col int
	Row int
}

// QOffsetFromHex returns the offset coordinates for a given hex tile when in
// "q" orientation -- that is, flat top.
func QOffsetFromHex(o OffsetType, h Hex) (OffsetCoord, bool) {
	if o != Even && o != Odd {
		return OffsetCoord{}, false
	}

	col := h.q

	//int row = h.r + int((h.q + offset * (h.q & 1)) / 2);
	row := coordTransform(o, h.q, h.r)
	return OffsetCoord{col, row}, true
}

// ROffsetFromHex returns the offset coordinates for a given hex tile when in
// "r" orientation -- that is, pointy top.
func ROffsetFromHex(o OffsetType, h Hex) (OffsetCoord, bool) {
	if o != Even && o != Odd {
		return OffsetCoord{}, false
	}
	col := coordTransform(o, h.r, h.q)
	row := h.r

	return OffsetCoord{col, row}, true
}

// QOffsetToHex returns the Hex coordinates for a set of offset coordinates
// when in "q" orientation -- flat top.
func QOffsetToHex(o OffsetType, oc OffsetCoord) (Hex, bool) {
	if o != Even && o != Odd {
		return Hex{}, false
	}
	q := oc.Col
	r := coordTransform(o, oc.Col, oc.Row)
	return NewHex(q, r), true
}

// ROffsetToHex returns the Hex coordinates for a set of offset coordinates
// when in "r" orientation -- pointy top.
func ROffsetToHex(o OffsetType, oc OffsetCoord) (Hex, bool) {
	if o != Even && o != Odd {
		return Hex{}, false
	}
	q := coordTransform(o, oc.Row, oc.Col)
	r := oc.Row
	return NewHex(q, r), true
}

func coordTransform(o OffsetType, a, b int) int {
	// this produces:
	// output = b + int((a + offset * (a & 1)) / 2);

	fq := float64(a)
	i := float64(a & 1)
	j := fq + float64(o)*i
	c := int(j / 2)

	return b + c
}
