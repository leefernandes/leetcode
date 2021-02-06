package main

import (
	"fmt"
	"math"
	"time"
)

type test struct {
	v string
	e int
}

func main() {
	tests := []test{
		{
			v: "-922337203685477580948323",
			e: -2147483648,
		},
		{
			v: "9223372036854775808323443",
			e: 2147483647,
		},
		{
			v: "   +0 123",
			e: 0,
		},
		{
			v: "  -93010 ok",
			e: -93010,
		},
		{
			v: "   -42",
			e: -42,
		},
		{
			v: "4193 with words",
			e: 4193,
		},
		{
			v: "words and 987",
			e: 0,
		},
		{
			v: "-91283472332",
			e: -2147483648,
		},
		{
			v: "21474836460",
			e: 2147483647,
		},
	}

	tt := tests
	for i := range tt {
		t := tt[i]
		start := time.Now()
		fmt.Println("v:", t.v)
		r := myAtoi(t.v)
		fmt.Println(" ", time.Since(start))
		if r != t.e {
			fmt.Println(" 🛑", r == t.e, "got:", r, "expected:", t.e)
		} else {
			fmt.Println(" 🟢", r == t.e, "got:", r, "expected:", t.e)
		}
	}
}

var m map[byte]int = map[byte]int{
	57: 9,
	56: 8,
	55: 7,
	54: 6,
	53: 5,
	52: 4,
	51: 3,
	50: 2,
	49: 1,
	48: 0,
}

var min32 int = math.MinInt32 * -1

func myAtoi(s string) int {
	v := parse(s)
	return int(v)
}

func parse(s string) int {
	j := 0
	var v int = 0
	var sign int = 1

	for i := range s {
		c := s[i]

		switch c {
		// whitespace
		case 32:
			if 0 == j {
				continue
			} else {
				return v * sign
			}

		// -
		case 45:
			if 0 == j {
				j = i + 1
				sign = -1
				continue
			} else {
				return v * sign
			}

		// +
		case 43:
			if 0 == j {
				j = i + 1
				continue
			} else {
				return v * sign
			}

		default:
			if d, ok := m[c]; !ok {
				return v * sign
			} else {
				w := v*10 + d
				if 1 == sign && (w >= math.MaxInt32 || w < v) {
					return math.MaxInt32
				} else if -1 == sign && (w >= min32 || w < v) {
					return math.MinInt32
				}
				v = w
				j++
			}
		}
	}

	return v * sign
}
