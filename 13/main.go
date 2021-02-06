package main

import (
	"fmt"
	"reflect"
	"time"
)

type test struct {
	e int
	s string
}

type solution = func(string) int

var tests = func() []test {
	return []test{
		{
			e: 10,
			s: "X",
		},
		{
			e: 3,
			s: "III",
		},
		{
			e: 4,
			s: "IV",
		},
		{
			e: 9,
			s: "IX",
		},
		{
			e: 3289,
			s: "MMMCCLXXXIX",
		},
		{
			e: 58,
			s: "LVIII",
		},
		{
			e: 1994,
			s: "MCMXCIV",
		},
	}
}

var solutions = map[string]solution{
	//"initial":   romanToInt,
	//"optimized": romanToInt2,
	"learned": romanToInt3,
}

func main() {
	for name, solution := range solutions {
		run(name, solution, tests())
	}
}

func run(name string, s solution, tests []test) {
	fmt.Println("run", name)
	tt := tests //[len(tests)-1:]
	var d time.Duration
	for i := range tt {
		t := tt[i]
		//fmt.Println(i, t.a, t.d)
		start := time.Now()
		r := s(t.s)
		d += time.Since(start)
		if !reflect.DeepEqual(r, t.e) {
			fmt.Println(" 🛑", "test", i, "got:", r, "expected:", t.e)
		} else {
			//fmt.Println(" 🟢", "got:", r, "expected:", t.e)
		}
	}
	avg := time.Duration(d.Nanoseconds() / int64(len(tt)))
	fmt.Println("  ", name, avg*time.Nanosecond)
}

var bytemans = map[byte]int{
	73: 1,
	86: 5,
	88: 10,
	76: 50,
	67: 100,
	68: 500,
	77: 1000,
}

func romanToInt3(s string) int {
	var n int

	l := len(s) - 1

	var last int
	for i := l; i >= 0; i-- {
		v := bytemans[s[i]]
		if l == i {
			n += v
			last = v
			continue
		}

		if v < last {
			n -= v
		} else {
			n += v
		}
		last = v
	}

	return n
}

var romans = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romanToInt2(s string) int {
	var n int

	l := len(s) - 1

	var last string
	for i := l; i >= 0; i-- {
		c := string(s[i])
		v := romans[c]

		switch c {
		case "I":
			switch last {
			case "V", "X":
				n--
			default:
				n += v
			}

		case "X":
			switch last {
			case "L", "C":
				n -= 10
			default:
				n += v
			}

		case "C":
			switch last {
			case "D", "M":
				n -= 100
			default:
				n += v
			}

		default:
			n += v

		}

		last = c
	}

	return n
}

func romanToInt(s string) int {
	var n int

	l := len(s) - 1

	for i := 0; i <= l; i++ {
		c := string(s[i])

		switch c {
		case "I":
			v, inc := addRomans(s, i, l, "V", "X", 1, 4, 9)
			n += v
			if inc {
				i++
			}

		case "V":
			n += 5

		case "X":
			v, inc := addRomans(s, i, l, "L", "C", 10, 40, 90)
			n += v
			if inc {
				i++
			}

		case "L":
			n += 50

		case "C":
			v, inc := addRomans(s, i, l, "D", "M", 100, 400, 900)
			n += v
			if inc {
				i++
			}

		case "D":
			n += 500

		case "M":
			n += 1000

		}
	}

	return n
}

func addRomans(s string, i, l int, M, T string, Lv, Mv, Tv int) (int, bool) {
	j := i + 1
	if j <= l {
		if nc := string(s[j]); M == nc {
			return Mv, true
		} else if T == nc {
			return Tv, true
		}
	}
	return Lv, false
}
