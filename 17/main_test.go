package main

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
	"time"
)

type test struct {
	digits string
	want   []string
}

var tests = func() []test {
	return []test{
		{
			digits: "9",
			want:   []string{"w", "x", "y", "z"},
		},
		{
			digits: "23",
			want:   []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			digits: "",
			want:   []string{},
		},
		{
			digits: "2",
			want:   []string{"a", "b", "c"},
		},
		{
			digits: "2345",
			want:   []string{"adgj", "adgk", "adgl", "adhj", "adhk", "adhl", "adij", "adik", "adil", "aegj", "aegk", "aegl", "aehj", "aehk", "aehl", "aeij", "aeik", "aeil", "afgj", "afgk", "afgl", "afhj", "afhk", "afhl", "afij", "afik", "afil", "bdgj", "bdgk", "bdgl", "bdhj", "bdhk", "bdhl", "bdij", "bdik", "bdil", "begj", "begk", "begl", "behj", "behk", "behl", "beij", "beik", "beil", "bfgj", "bfgk", "bfgl", "bfhj", "bfhk", "bfhl", "bfij", "bfik", "bfil", "cdgj", "cdgk", "cdgl", "cdhj", "cdhk", "cdhl", "cdij", "cdik", "cdil", "cegj", "cegk", "cegl", "cehj", "cehk", "cehl", "ceij", "ceik", "ceil", "cfgj", "cfgk", "cfgl", "cfhj", "cfhk", "cfhl", "cfij", "cfik", "cfil"},
		},
	}
}

func TestLetterCombinations(t *testing.T) {
	start := time.Now()
	for i, tt := range tests() {
		testname := fmt.Sprintf("%d", i)
		t.Run(testname, func(t *testing.T) {
			got := letterCombinations(tt.digits)
			ok := isItGood(got, tt.want)
			if !ok {
				t.Errorf("🛑 test %d got %v, want %v from %v", i, got, tt.want, tt.digits)
			}
		})
	}

	fmt.Println(time.Since(start))
}

func TestLetterCombinationsQueue(t *testing.T) {
	start := time.Now()
	for i, tt := range tests() {
		testname := fmt.Sprintf("%d", i)
		t.Run(testname, func(t *testing.T) {
			got := letterCombinationsQueue(tt.digits)
			ok := isItGood(got, tt.want)
			if !ok {
				t.Errorf("🛑 test %d got %v, want %v from %v", i, got, tt.want, tt.digits)
			}
		})
	}

	fmt.Println(time.Since(start))
}

func isItGood(got []string, want []string) bool {
	if len(got) != len(want) {
		return false
	} else {
		sort.Strings(got)
		sort.Strings(want)
		if !reflect.DeepEqual(got, want) {
			return false
		}
	}
	return true
}
