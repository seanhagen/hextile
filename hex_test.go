package hextile

import (
	"fmt"
	"testing"
)

func TestNewHex(t *testing.T) {
	tests := []struct {
		q int
		r int
		s int
	}{
		{1, 0, -1},
		{1, 1, -2},
	}

	for i, x := range tests {
		tt := x
		t.Run(fmt.Sprintf("test%v", i), func(t *testing.T) {
			h := NewHex(tt.q, tt.r)

			if h.q != tt.q {
				t.Errorf("invalid q, expected %v got %v", tt.q, h.q)
			}

			if h.r != tt.r {
				t.Errorf("invalid r, expected %v got %v", tt.r, h.r)
			}

			if h.s != tt.s {
				t.Errorf("invalid s, expected %v got %v", tt.s, h.s)
			}
		})
	}
}
