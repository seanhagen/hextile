package hextile

import (
	"fmt"
	"reflect"
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

func TestHex_FieldAccessor(t *testing.T) {
	type fields struct {
		q int
		r int
	}

	type want struct {
		q int
		r int
		s int
	}

	tests := []struct {
		name   string
		fields fields
		want   want
	}{
		{"basic", fields{0, 0}, want{0, 0, 0}},
		{"small", fields{1, 0}, want{1, 0, -1}},
		{"complex", fields{2, 5}, want{2, 5, -7}},
	}
	for _, x := range tests {
		tt := x
		t.Run(tt.name, func(t *testing.T) {
			h := NewHex(tt.fields.q, tt.fields.r)
			if got := h.Q(); got != tt.want.q {
				t.Errorf("hex: %v -- Hex.Q() = %v, want %v", h, got, tt.want.q)
			}

			if got := h.R(); got != tt.want.r {
				t.Errorf("hex: %v -- Hex.Q() = %v, want %v", h, got, tt.want.r)
			}

			if got := h.S(); got != tt.want.s {
				t.Errorf("hex: %v -- Hex.S() = %v, want %v", h, got, tt.want.s)
			}

		})
	}
}

func TestHex_Eq(t *testing.T) {
	type h struct {
		q, r int
	}
	tests := []struct {
		name   string
		fields h
		want   h
		valid  bool
	}{
		{"invalid", h{0, 0}, h{1, 1}, false},
		{"basic", h{0, 0}, h{0, 0}, true},
		{"positive q", h{1, 0}, h{1, 0}, true},
		{"postitve r", h{0, 1}, h{0, 1}, true},
		{"negative q", h{-1, 0}, h{-1, 0}, true},
		{"negative r", h{0, -1}, h{0, -1}, true},
		{"both pos", h{1, 1}, h{1, 1}, true},
		{"both neg", h{-1, -1}, h{-1, -1}, true},
	}
	for _, x := range tests {
		tt := x
		t.Run(tt.name, func(t *testing.T) {
			h := NewHex(tt.fields.q, tt.fields.r)
			e := NewHex(tt.want.q, tt.want.r)
			if got := h.Eq(e); got != tt.valid {
				t.Errorf("Hex (%v) Eq(%v) = %v, want %v", h, e, got, tt.valid)
			}
		})
	}
}

func TestHex_String(t *testing.T) {
	type fields struct {
		q int
		r int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"basic", fields{0, 0}, "{0 0}"},
		{"another", fields{1, 1}, "{1 1}"},
	}

	for _, x := range tests {
		tt := x
		t.Run(tt.name, func(t *testing.T) {
			h := Hex{
				q: tt.fields.q,
				r: tt.fields.r,
			}
			if got := h.String(); got != tt.want {
				t.Errorf("Hex.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHex_Add(t *testing.T) {
	tests := []struct {
		name  string
		start Hex
		args  Hex
		want  Hex
	}{
		{"0,0 - add 0,0", NewHex(0, 0), NewHex(0, 0), NewHex(0, 0)},
		{"0,0 - add 1,0", NewHex(0, 0), NewHex(1, 0), NewHex(1, 0)},
		{"1,1 - add 1,0", NewHex(1, 1), NewHex(1, 0), NewHex(2, 1)},
		{"1,1 - add 0,1", NewHex(1, 1), NewHex(0, 1), NewHex(1, 2)},
		{"4,4 - add 1,0", NewHex(4, 4), NewHex(1, 0), NewHex(5, 4)},
		{"4,4 - add 0,1", NewHex(4, 4), NewHex(0, 1), NewHex(4, 5)},
	}
	for _, x := range tests {
		tt := x
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.start.Add(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%v.Add(%v) = %v, want %v", tt.start, tt.args, got, tt.want)
			}
		})
	}
}

func TestHex_Subtract(t *testing.T) {
	tests := []struct {
		name  string
		start Hex
		args  Hex
		want  Hex
	}{
		{"0,0 - sub 0,0", NewHex(0, 0), NewHex(0, 0), NewHex(0, 0)},
		{"0,0 - sub 1,0", NewHex(0, 0), NewHex(1, 0), NewHex(-1, 0)},
		{"1,1 - sub 1,0", NewHex(1, 1), NewHex(1, 0), NewHex(0, 1)},
		{"1,1 - sub 0,1", NewHex(1, 1), NewHex(0, 1), NewHex(1, 0)},
		{"4,4 - sub 1,0", NewHex(4, 4), NewHex(1, 0), NewHex(3, 4)},
		{"4,4 - sub 0,1", NewHex(4, 4), NewHex(0, 1), NewHex(4, 3)},
	}
	for _, x := range tests {
		tt := x
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.start.Subtract(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%v.Subtract(%v) = %v, want %v", tt.start, tt.args, got, tt.want)
			}
		})
	}
}

func TestHex_Multiply(t *testing.T) {
	tests := []struct {
		name  string
		start Hex
		args  int
		want  Hex
	}{
		{"0,0 multiply 0", NewHex(0, 0), 0, NewHex(0, 0)},
		{"0,0 multiply 1", NewHex(0, 0), 1, NewHex(0, 0)},
		{"1,1 multiply 0", NewHex(1, 1), 0, NewHex(0, 0)},
		{"1,1 multiply 1", NewHex(1, 1), 1, NewHex(1, 1)},
		{"1,1 multiply 2", NewHex(1, 1), 2, NewHex(2, 2)},
		{"2,2 multiply 0", NewHex(2, 2), 0, NewHex(0, 0)},
		{"2,2 multiply 1", NewHex(2, 2), 1, NewHex(2, 2)},
		{"2,2 multiply 2", NewHex(2, 2), 2, NewHex(4, 4)},
		{"10,10 multiply 1", NewHex(10, 10), 1, NewHex(10, 10)},
		{"10,10 multiply 2", NewHex(10, 10), 2, NewHex(20, 20)},
	}
	for _, x := range tests {
		tt := x
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.start.Multiply(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%v.Multiply(%v) = %v, want %v", tt.start, tt.args, got, tt.want)
			}
		})
	}
}

func TestHex_Length(t *testing.T) {
	tests := []struct {
		name string
		q, r int
		want int
	}{
		{"basic", 0, 0, 0},
		{"{1 1} length", 1, 1, 2},
		{"{2 2} length", 2, 2, 4},
		{"{1 2} length", 1, 2, 3},
	}
	for _, x := range tests {
		tt := x
		t.Run(tt.name, func(t *testing.T) {
			h := NewHex(tt.q, tt.r)
			if got := h.Length(); got != tt.want {
				t.Errorf("%v.Length() = %v, want %v", h, got, tt.want)
			}
		})
	}
}

func TestDistance(t *testing.T) {
	type args struct {
		a Hex
		b Hex
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"basic - same tile", args{NewHex(0, 0), NewHex(0, 0)}, 0},
		{"basic - neighbour", args{NewHex(0, 0), NewHex(0, 1)}, 1},
		{"two away", args{NewHex(0, 0), NewHex(2, -1)}, 2},
		{"three away", args{NewHex(0, 0), NewHex(3, -2)}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Distance(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Distance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetDirection(t *testing.T) {
	tests := []struct {
		name string
		args Direction
		want Hex
	}{
		{"invalid", Direction(99), NewHex(0, 0)},
		{"right", DirectionRight, NewHex(1, 0)},
		{"right up", DirectionRightUp, NewHex(1, -1)},
		{"left up", DirectionLeftUp, NewHex(0, -1)},
		{"left", DirectionLeft, NewHex(-1, 0)},
		{"left down", DirectionLeftDown, NewHex(-1, 1)},
		{"right down", DirectionRightDown, NewHex(0, 1)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetDirection(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDirection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetHexNeighbour(t *testing.T) {
	type args struct {
		h Hex
		d Direction
	}
	tests := []struct {
		name string
		args args
		want Hex
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetHexNeighbour(tt.args.h, tt.args.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetHexNeighbour() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHex_Neighbour(t *testing.T) {
	type fields struct {
		q int
		r int
		s int
	}
	type args struct {
		d Direction
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Hex
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := Hex{
				q: tt.fields.q,
				r: tt.fields.r,
				s: tt.fields.s,
			}
			if got := h.Neighbour(tt.args.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Hex.Neighbour() = %v, want %v", got, tt.want)
			}
		})
	}
}
