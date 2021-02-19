package main

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

type test struct {
	n    int
	want []string
}

var tests = func() []test {
	return []test{
		{
			n: 15,
			want: []string{
				"1",
				"2",
				"Fizz",
				"4",
				"Buzz",
				"Fizz",
				"7",
				"8",
				"Fizz",
				"Buzz",
				"11",
				"Fizz",
				"13",
				"14",
				"FizzBuzz",
			},
		},
	}
}

func TestFizzBuzz(t *testing.T) {
	start := time.Now()
	for i, tt := range tests() {
		testname := fmt.Sprintf("%d", i)
		t.Run(testname, func(t *testing.T) {
			got := fizzBuzz(tt.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("🛑 test %d got %v, want %v - %v", i, got, tt.want, tt.n)
			}
		})
	}

	fmt.Println(time.Since(start))
}
