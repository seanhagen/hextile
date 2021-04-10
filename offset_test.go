package hextile

import (
	"reflect"
	"testing"
)

func TestQOffsetFromHex(t *testing.T) {
	type args struct {
		o OffsetType
		h Hex
	}
	tests := []struct {
		name  string
		args  args
		want  OffsetCoord
		want1 bool
	}{
		{"invalid offset type", args{OffsetType(3), NewHex(0, 0)}, OffsetCoord{0, 0}, false},

		{"basic,even", args{Even, NewHex(0, 0)}, OffsetCoord{0, 0}, true},
		{"basic, odd", args{Odd, NewHex(0, 0)}, OffsetCoord{0, 0}, true},

		{"positive q, even", args{Even, NewHex(2, 0)}, OffsetCoord{2, 1}, true},
		{"positive qr, even", args{Even, NewHex(2, 1)}, OffsetCoord{2, 2}, true},

		{"positive q, odd", args{Odd, NewHex(2, 0)}, OffsetCoord{2, 1}, true},
		{"positive qr, odd", args{Odd, NewHex(2, 1)}, OffsetCoord{2, 2}, true},

		{"negative q, even", args{Even, NewHex(-2, 0)}, OffsetCoord{-2, -1}, true},
		{"negative qr, even", args{Even, NewHex(-2, -1)}, OffsetCoord{-2, -2}, true},

		{"negative q, odd", args{Odd, NewHex(-2, 0)}, OffsetCoord{-2, -1}, true},
		{"negative qr, odd", args{Odd, NewHex(-2, -1)}, OffsetCoord{-2, -2}, true},
	}
	for _, x := range tests {
		tt := x
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := QOffsetFromHex(tt.args.o, tt.args.h)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QOffsetFromHex() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("QOffsetFromHex() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestROffsetFromHex(t *testing.T) {
	type args struct {
		o OffsetType
		h Hex
	}
	tests := []struct {
		name  string
		args  args
		want  OffsetCoord
		want1 bool
	}{
		{"invalid offset type", args{OffsetType(3), NewHex(0, 0)}, OffsetCoord{0, 0}, false},

		{"basic,even", args{Even, NewHex(0, 0)}, OffsetCoord{0, 0}, true},
		{"basic, odd", args{Odd, NewHex(0, 0)}, OffsetCoord{0, 0}, true},

		{"positive q, even", args{Even, NewHex(2, 0)}, OffsetCoord{2, 0}, true},
		{"positive qr, even", args{Even, NewHex(2, 1)}, OffsetCoord{3, 1}, true},

		{"positive q, odd", args{Odd, NewHex(2, 0)}, OffsetCoord{2, 0}, true},
		{"positive qr, odd", args{Odd, NewHex(2, 1)}, OffsetCoord{2, 1}, true},

		{"negative q, even", args{Even, NewHex(-2, 0)}, OffsetCoord{-2, 0}, true},
		{"negative qr, even", args{Even, NewHex(-2, -1)}, OffsetCoord{-2, -1}, true},

		{"negative q, odd", args{Odd, NewHex(-2, 0)}, OffsetCoord{-2, 0}, true},
		{"negative qr, odd", args{Odd, NewHex(-2, -1)}, OffsetCoord{-3, -1}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ROffsetFromHex(tt.args.o, tt.args.h)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ROffsetFromHex() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ROffsetFromHex() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestQOffsetToHex(t *testing.T) {
	type args struct {
		o  OffsetType
		oc OffsetCoord
	}
	tests := []struct {
		name  string
		args  args
		want  Hex
		want1 bool
	}{
		{"invalid offset type", args{OffsetType(3), OffsetCoord{0, 0}}, NewHex(0, 0), false},

		{"basic,even", args{Even, OffsetCoord{0, 0}}, NewHex(0, 0), true},
		{"basic, odd", args{Odd, OffsetCoord{0, 0}}, NewHex(0, 0), true},

		{"positive q, even", args{Even, OffsetCoord{2, 1}}, NewHex(2, 2), true},
		{"positive qr, even", args{Even, OffsetCoord{2, 2}}, NewHex(2, 3), true},

		{"positive q, odd", args{Odd, OffsetCoord{2, 1}}, NewHex(2, 2), true},
		{"positive qr, odd", args{Odd, OffsetCoord{2, 2}}, NewHex(2, 3), true},

		{"negative q, even", args{Even, OffsetCoord{-2, -1}}, NewHex(-2, -2), true},
		{"negative qr, even", args{Even, OffsetCoord{-2, -2}}, NewHex(-2, -3), true},

		{"negative q, odd", args{Odd, OffsetCoord{-2, -1}}, NewHex(-2, -2), true},
		{"negative qr, odd", args{Odd, OffsetCoord{-2, -2}}, NewHex(-2, -3), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := QOffsetToHex(tt.args.o, tt.args.oc)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QOffsetToHex(%v,%v) got = %v, want %v", tt.args.o, tt.args.oc, got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("QOffsetToHex(%v,%v) got1 = %v, want %v", tt.args.o, tt.args.oc, got1, tt.want1)
			}
		})
	}
}

func TestROffsetToHex(t *testing.T) {
	type args struct {
		o  OffsetType
		oc OffsetCoord
	}
	tests := []struct {
		name  string
		args  args
		want  Hex
		want1 bool
	}{
		{"invalid offset type", args{OffsetType(3), OffsetCoord{0, 0}}, NewHex(0, 0), false},

		{"basic,even", args{Even, OffsetCoord{0, 0}}, NewHex(0, 0), true},
		{"basic, odd", args{Odd, OffsetCoord{0, 0}}, NewHex(0, 0), true},

		{"positive q, even", args{Even, OffsetCoord{2, 0}}, NewHex(2, 0), true},
		{"positive qr, even", args{Even, OffsetCoord{3, 1}}, NewHex(4, 1), true},

		{"positive q, odd", args{Odd, OffsetCoord{2, 0}}, NewHex(2, 0), true},
		{"positive qr, odd", args{Odd, OffsetCoord{2, 1}}, NewHex(2, 1), true},

		{"negative q, even", args{Even, OffsetCoord{-2, 0}}, NewHex(-2, 0), true},
		{"negative qr, even", args{Even, OffsetCoord{-2, -1}}, NewHex(-2, -1), true},

		{"negative q, odd", args{Odd, OffsetCoord{-2, 0}}, NewHex(-2, 0), true},
		{"negative qr, odd", args{Odd, OffsetCoord{-3, -1}}, NewHex(-4, -1), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ROffsetToHex(tt.args.o, tt.args.oc)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ROffsetToHex(%v,%v) got = %v, want %v", tt.args.o, tt.args.oc, got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ROffsetToHex(%v,%v) got1 = %v, want %v", tt.args.o, tt.args.oc, got1, tt.want1)
			}
		})
	}
}
