package main

import (
	"fmt"
	"testing"
	"time"
)

type test struct {
	A    []int
	B    []int
	want int
}

var tests = func() []test {
	return []test{
		{
			A:    []int{0, 1, 1, 1, 1},
			B:    []int{1, 0, 1, 0, 1},
			want: 2,
		},
		{
			A:    []int{1, 2, 3, 2, 1},
			B:    []int{3, 2, 1, 4, 7},
			want: 3,
		},
		{
			A:    []int{0, 0, 1, 0, 0, 0, 3},
			B:    []int{0, 0, 0, 1, 0, 0, 4},
			want: 5,
		},
		{
			A:    []int{0, 0, 0, 0, 0, 0, 1, 0, 0, 0},
			B:    []int{0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
			want: 9,
		},
	}
}

func TestFindLength2(t *testing.T) {

	start := time.Now()
	for i, tt := range tests() {
		testname := fmt.Sprintf("%d", i)
		t.Run(testname, func(t *testing.T) {
			result := findLength2(tt.A, tt.B)
			if result != tt.want {
				t.Errorf("🛑 test %d got %v, want %v", i, result, tt.want)
			}
		})
	}

	fmt.Println(time.Since(start))
}

func TestFindLength(t *testing.T) {

	start := time.Now()
	for i, tt := range tests() {
		testname := fmt.Sprintf("%d", i)
		t.Run(testname, func(t *testing.T) {
			result := findLength(tt.A, tt.B)
			if result != tt.want {
				t.Errorf("🛑 test %d got %v, want %v", i, result, tt.want)
			}
		})
	}

	fmt.Println(time.Since(start))
}
