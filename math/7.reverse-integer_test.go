package math

import "testing"

func Test_reverse(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"reverse case 1", args{123}, 321},
		{"reverse case 2", args{-123}, -321},
		{"reverse case 3", args{120}, 21},
		{"reverse case 4", args{0}, 0},
		{"reverse case 5", args{1563847412}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.x); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
