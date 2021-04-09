package hextile

/**
 * File: map.go
 * Date: 2021-04-08 23:05:32
 * Creator: Sean Hagen <sean.hagen@gmail.com>
 */

// Map stores any additional information about each hex on a grid
type Map struct {
	data  map[Hex]interface{}
	tiles map[Hex]bool
}

// TileAt ...
func (m Map) IsTileAt(q, r int) bool {
	h := NewHex(q, r)
	return m._isTileAt(h)
}

// _tileAt ...
func (m Map) _isTileAt(h Hex) bool {
	_, ok := m.tiles[h]
	return ok
}

// Tile ...
func (m Map) Tile(q, r int) (Hex, bool) {
	h := NewHex(q, r)
	_, ok := m.tiles[h]
	return h, ok
}

// TileNeighbours ...
func (m Map) TileNeighbours(h Hex) []Hex {
	n := []Hex{}
	if !m._isTileAt(h) {
		return n
	}

	for d := range directionToHex {
		nh := h.Neighbour(d)
		if m._isTileAt(nh) {
			n = append(n, nh)
		}
	}

	return n
}

// AddHex ...
func (m Map) AddHex(h Hex) {
	m.tiles[h] = true
}
