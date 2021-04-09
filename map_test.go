package hextile

import (
	"fmt"
	"testing"
)

func TestBasic(t *testing.T) {
	tests := []struct {
		a, b, c int
	}{
		{1, 0, 1},
		{1, 1, 2},
	}

	for i, x := range tests {
		tt := x
		t.Run(fmt.Sprintf("test%v", i), func(t *testing.T) {
			c := tt.a + tt.b
			if c != tt.c {
				t.Errorf("something is seriously wrong, %v+%v != %v, got %v instead", tt.a, tt.b, tt.c, c)
			}
		})
	}
}

func TestRectMap(t *testing.T) {
	m := NewRectangleMap(5, 10)

	/*
			   builds the following map:

		{0 0} {1 0} {2 0} {3 0} {4 0} {5 0} {6 0} {7 0} {8 0} {9 0}
		   {0 1} {1 1} {2 1} {3 1} {4 1} {5 1} {6 1} {7 1} {8 1} {9 1}
		{-1 2} {0 2} {1 2} {2 2} {3 2} {4 2} {5 2} {6 2} {7 2} {8 2}
		  {-1 3} {0 3} {1 3} {2 3} {3 3} {4 3} {5 3} {6 3} {7 3} {8 3}
		{-2 4} {-1 4} {0 4} {1 4} {2 4} {3 4} {4 4} {5 4} {6 4} {7 4}

	*/

	tests := []struct {
		q, r                                              int
		right, rightUp, leftUp, left, leftDown, rightDown bool
	}{
		{0, 0, true, false, false, false, false, true},
		{2, 2, true, true, true, true, true, true},
		{-2, 4, true, true, false, false, false, false},
		{9, 0, false, false, false, true, true, true},
		{7, 4, false, true, true, true, false, false},
		{2, 4, true, true, true, true, false, false},
		{4, 0, true, false, false, true, true, true},
	}

	for i, x := range tests {
		tt := x
		t.Run(fmt.Sprintf("test-%v", i), func(t *testing.T) {
			t.Parallel()

			n := map[Direction]bool{
				DirectionRight:     tt.right,
				DirectionRightUp:   tt.rightUp,
				DirectionLeftUp:    tt.leftUp,
				DirectionLeft:      tt.left,
				DirectionLeftDown:  tt.leftDown,
				DirectionRightDown: tt.rightDown,
			}

			h, ok := m.Tile(tt.q, tt.r)
			if !ok {
				t.Fatalf("expected tile at 0,0, got not ok")
			}

			for d, e := range n {
				x := d
				xe := e
				xh := h
				dir := directionName[x]
				nh := xh.Neighbour(x)

				ee := m._isTileAt(nh)
				if ee != xe {
					t.Errorf("tile %v --expected %v, got %v for direction from tile '%v'", xh, xe, ee, dir)
				}
			}
		})
	}
}
