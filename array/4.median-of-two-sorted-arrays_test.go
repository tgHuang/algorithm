package array

import "testing"

func Test_findMedianSortedArrays(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		// {"findMedianSortedArrays case 1", args{[]int{1, 3}, []int{2}}, 2.0},
		// {"findMedianSortedArrays case 2", args{[]int{1, 2}, []int{3, 4}}, 2.50},
		// {"findMedianSortedArrays case 3", args{[]int{0, 0}, []int{0, 0}}, 0.00},
		// {"findMedianSortedArrays case 4", args{[]int{}, []int{1}}, 1.00},
		// {"findMedianSortedArrays case 5", args{[]int{2}, []int{}}, 2.00},
		// {"findMedianSortedArrays case 5", args{[]int{3}, []int{-2, -1}}, -1.00},
		// {"findMedianSortedArrays case 5", args{[]int{1, 3}, []int{2, 7}}, 2.500},
		{"findMedianSortedArrays case 5", args{[]int{2, 2, 4, 4}, []int{2, 2, 4, 4}}, 3.00},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMedianSortedArrays(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("findMedianSortedArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
