package main

import (
	"fmt"
	"testing"
	"time"
)

type test struct {
	s    string
	want bool
}

var tests = func() []test {
	return []test{
		{
			s:    "()[]{}",
			want: true,
		},
		{
			s:    "{[]}",
			want: true,
		},
		{
			s:    "()",
			want: true,
		},
		{
			s:    "(]",
			want: false,
		},
		{
			s:    "([)]",
			want: false,
		},
		{
			s:    "{[()]}",
			want: true,
		},
		{
			s:    "()[",
			want: false,
		},
		{
			s:    "(ok){kewl}fine[sire]",
			want: true,
		},
		{
			s:    "({[yeps]})",
			want: true,
		},
		{
			s:    "({[})nogood)",
			want: false,
		},
	}
}

func TestIsValid(t *testing.T) {

	start := time.Now()
	for i, tt := range tests() {
		testname := fmt.Sprintf("%d", i)
		t.Run(testname, func(t *testing.T) {
			result := isValid(tt.s)
			if result != tt.want {
				t.Errorf("🛑 test %d got %v, want %v - %v", i, result, tt.want, tt.s)
			}
		})
	}

	fmt.Println(time.Since(start))
}

func TestIsValid2(t *testing.T) {

	start := time.Now()
	for i, tt := range tests() {
		testname := fmt.Sprintf("%d", i)
		t.Run(testname, func(t *testing.T) {
			result := isValid2(tt.s)
			if result != tt.want {
				t.Errorf("🛑 test %d got %v, want %v - %v", i, result, tt.want, tt.s)
			}
		})
	}

	fmt.Println(time.Since(start))
}
