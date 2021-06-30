package math

import "testing"

func Test_romanToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"romanToInt case 1", args{"IV"}, 4},
		{"romanToInt case 2", args{"VI"}, 6},
		{"romanToInt case 3", args{"III"}, 3},
		{"romanToInt case 5", args{"IX"}, 9},
		{"romanToInt case 6", args{"LVIII"}, 58},
		{"romanToInt case 7", args{"MCMXCIV"}, 1994},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := romanToInt(tt.args.s); got != tt.want {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
