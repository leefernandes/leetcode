package main

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

type test struct {
	nums   []int
	target int
	want   int
}

var tests = func() []test {
	return []test{
		{
			nums:   []int{-1, 2, 1, -4},
			target: 1,
			want:   2,
		},
		{
			nums:   []int{1, 1, -1, -1, 3},
			target: 3,
			want:   3,
		},
	}
}

func TestThreeSumClosest(t *testing.T) {
	start := time.Now()
	for i, tt := range tests()[1:2] {
		testname := fmt.Sprintf("%d", i)
		t.Run(testname, func(t *testing.T) {
			res := threeSumClosest(tt.nums, tt.target)
			if !reflect.DeepEqual(res, tt.want) {
				t.Errorf("🛑 test %d got %v, want %v from %v", i, res, tt.want, tt.nums)
			}
		})
	}

	fmt.Println(time.Since(start))
}
