package main

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
	"time"
)

type test struct {
	nums []int
	want [][]int
}

var tests = func() []test {
	return []test{
		{
			nums: []int{-4, -2, 1, -5, -4, -4, 4, -2, 0, 4, 0, -2, 3, 1, -5, 0},
			want: [][]int{
				{-5, 1, 4}, {-4, 0, 4}, {-4, 1, 3}, {-2, -2, 4}, {-2, 1, 1}, {0, 0, 0},
			},
		},
		{
			nums: []int{0, 1, 1},
			want: [][]int{},
		},
		{
			nums: []int{1, 2, 1, 2, 3, -3, -2, -3, 0},
			want: [][]int{
				{-3, 0, 3},
				{-3, 1, 2},
				{-2, 0, 2},
				{-2, 1, 1},
			},
		},
		{
			nums: []int{-1, 0, 1},
			want: [][]int{
				{-1, 0, 1},
			},
		},
		{
			nums: []int{-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4},
			want: [][]int{
				{-1, 0, 1},
				{-1, -1, 2},
				{-2, -1, 3},
				{-3, -1, 4},
				{-2, 0, 2},
				{-4, 0, 4},
				{-3, 0, 3},
				{-3, 1, 2},
				{-4, 1, 3},
			},
		},
		{
			nums: []int{-1, 0, 1, 2, -1, -4},
			want: [][]int{
				{-1, -1, 2},
				{-1, 0, 1},
			},
		},
		{
			nums: []int{},
			want: [][]int{},
		},
		{
			nums: []int{0},
			want: [][]int{},
		},
	}
}

func TestThreeSum(t *testing.T) {

	start := time.Now()
	for i, tt := range tests() {
		testname := fmt.Sprintf("%d", i)
		t.Run(testname, func(t *testing.T) {
			res := threeSum(tt.nums)
			if !reflect.DeepEqual(res, tt.want) {
				t.Errorf("🛑 test %d got %v, want %v from %v", i, res, tt.want, tt.nums)
			}
		})
	}

	fmt.Println(time.Since(start))
}

func TestThreeSum2(t *testing.T) {
	start := time.Now()
	for i, tt := range tests() {
		testname := fmt.Sprintf("%d", i)
		t.Run(testname, func(t *testing.T) {
			res := threeSum2(tt.nums)
			if len(tt.want) != len(res) {
				t.Errorf("🛑 test %d got %v, want %v from %v", i, res, tt.want, tt.nums)
			} else {
				if !reflect.DeepEqual(res, tt.want) {
					for i := range res {
						r := res[i]
						sort.Ints(r)
						ok := false
						for j := range tt.want {
							w := tt.want[j]
							sort.Ints(w)
							if reflect.DeepEqual(r, w) {
								ok = true
								break
							}
						}
						if !ok {
							t.Errorf("🛑 test %d got %v, want %v from %v", i, res, tt.want, tt.nums)
							return
						}
					}
				}
			}
		})
	}

	fmt.Println(time.Since(start))
}

func TestThreeSum3(t *testing.T) {
	start := time.Now()
	for i, tt := range tests() {
		testname := fmt.Sprintf("%d", i)
		t.Run(testname, func(t *testing.T) {
			res := threeSum3(tt.nums)
			if len(tt.want) != len(res) {
				t.Errorf("🛑 test %d got %v, want %v from %v", i, res, tt.want, tt.nums)
			} else {
				if !reflect.DeepEqual(res, tt.want) {
					for i := range res {
						r := res[i]
						sort.Ints(r)
						ok := false
						for j := range tt.want {
							w := tt.want[j]
							sort.Ints(w)
							if reflect.DeepEqual(r, w) {
								ok = true
								break
							}
						}
						if !ok {
							t.Errorf("🛑 test %d got %v, want %v from %v", i, res, tt.want, tt.nums)
							return
						}
					}
				}
			}
		})
	}

	fmt.Println(time.Since(start))
}
