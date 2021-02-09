package main

import (
	"fmt"
	"testing"
	"time"
)

type test struct {
	nums []int
	k    int
	want int
}

var tests = func() []test {
	return []test{
		{
			nums: []int{3, 2, 1, 5, 6, 4},
			k:    2,
			want: 5,
		},
		{
			nums: []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
			k:    4,
			want: 4,
		},
	}
}

func TestFindKthLargest(t *testing.T) {
	start := time.Now()
	for i, tt := range tests() {
		testname := fmt.Sprintf("%d", i)
		t.Run(testname, func(t *testing.T) {
			got := findKthLargest(tt.nums, tt.k)
			ok := got == tt.want
			if !ok {
				t.Errorf("🛑 test %d got %v, want %v from %v %v", i, got, tt.want, tt.nums, tt.k)
			}
		})
	}

	fmt.Println(time.Since(start))
}
