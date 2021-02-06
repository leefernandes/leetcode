package main

import (
	"fmt"
	"testing"
	"time"
)

type test struct {
	strs []string
	want string
}

var tests = func() []test {
	return []test{
		{
			strs: []string{"flower", "flow", "flight"},
			want: "fl",
		},
		{
			strs: []string{"goo", "goober", "goopy", "goobers", "goop"},
			want: "goo",
		},
		{
			strs: []string{"dog", "doggie", "car"},
			want: "",
		},
		{
			strs: []string{},
			want: "",
		},
		{
			strs: []string{"O"},
			want: "O",
		},
		{
			strs: []string{"Higher", "Higgle", ""},
			want: "",
		},
	}
}

func TestLongestCommonPrefix(t *testing.T) {

	start := time.Now()
	for i, tt := range tests() {
		testname := fmt.Sprintf("%d", i)
		t.Run(testname, func(t *testing.T) {
			result := longestCommonPrefix(tt.strs)
			if result != tt.want {
				t.Errorf("🛑 got %v, want %v", result, tt.want)
			}
		})
	}

	fmt.Println(time.Since(start))
}

func TestLongestCommonPrefix2(t *testing.T) {

	start := time.Now()
	for i, tt := range tests() {
		testname := fmt.Sprintf("%d", i)
		t.Run(testname, func(t *testing.T) {
			result := longestCommonPrefix2(tt.strs)
			if result != tt.want {
				t.Errorf("🛑 got %v, want %v", result, tt.want)
			}
		})
	}

	fmt.Println(time.Since(start))
}

func TestLongestCommonPrefix3(t *testing.T) {

	start := time.Now()
	for i, tt := range tests() {
		testname := fmt.Sprintf("%d", i)
		t.Run(testname, func(t *testing.T) {
			result := longestCommonPrefix3(tt.strs)
			if result != tt.want {
				t.Errorf("🛑 got %v, want %v", result, tt.want)
			}
		})
	}

	fmt.Println(time.Since(start))
}
