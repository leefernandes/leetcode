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
	want   []int
}

func tmp() {
	fmt.Println(0 % -1)
}

var tests = func() []test {
	return []test{
		{
			nums:   []int{2, 7, 11, 15},
			target: 9,
			want:   []int{0, 1},
		},
		{
			nums:   []int{3, 2, 4},
			target: 6,
			want:   []int{1, 2},
		},
		{
			nums:   []int{3, 3},
			target: 6,
			want:   []int{0, 1},
		},
		{
			nums:   []int{3, 2, 3},
			target: 6,
			want:   []int{0, 2},
		},
	}
}

func TestTwoSum(t *testing.T) {

	start := time.Now()
	for i, tt := range tests() {
		testname := fmt.Sprintf("%d", i)
		t.Run(testname, func(t *testing.T) {
			res := twoSum(tt.nums, tt.target)
			if !reflect.DeepEqual(res, tt.want) {
				t.Errorf("🛑 test %d got %v, want %v from %v", i, res, tt.want, tt.nums)
			}
		})
	}

	fmt.Println(time.Since(start))
}

func TestTwoSumHash(t *testing.T) {

	start := time.Now()
	for i, tt := range tests()[1:2] {
		testname := fmt.Sprintf("%d", i)
		t.Run(testname, func(t *testing.T) {
			res := twoSumHash(tt.nums, tt.target)
			if !reflect.DeepEqual(res, tt.want) {
				t.Errorf("🛑 test %d got %v, want %v from %v", i, res, tt.want, tt.nums)
			}
		})
	}

	fmt.Println(time.Since(start))
}
