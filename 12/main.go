package main

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
	"time"
)

type test struct {
	num int
	e   string
}

type solution = func(int) string

var tests = func() []test {
	return []test{
		{
			num: 10,
			e:   "X",
		},
		{
			num: 3,
			e:   "III",
		},
		{
			num: 4,
			e:   "IV",
		},
		{
			num: 9,
			e:   "IX",
		},
		{
			num: 3289,
			e:   "MMMCCLXXXIX",
		},
		{
			num: 58,
			e:   "LVIII",
		},
		{
			num: 1994,
			e:   "MCMXCIV",
		},
	}
}

var solutions = map[string]solution{
	"initial":   intToRoman,
	"optimized": intToRoman2,
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
		r := s(t.num)
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

func intToRoman2(num int) string {
	s := new(strings.Builder)

	if num > 9 {
		if num > 99 {
			if num > 999 {
				thousands := num / 1000
				spreadRomans(s, thousands, "M", "", "")
			}

			hundreds := (num % 1000) / 100
			spreadRomans(s, hundreds, "C", "D", "M")
		}

		tens := (num % 100) / 10
		spreadRomans(s, tens, "X", "L", "C")
	}

	singles := num % 10
	spreadRomans(s, singles, "I", "V", "X")

	return s.String()
}

func spreadRomans(s *strings.Builder, place int, L, M, T string) {
	switch {
	case place < 4:
		for i := 0; i < place; i++ {
			s.WriteString(L)
		}

	case 9 > place && place > 4:
		s.WriteString(M)
		d := place - 5
		for i := 0; i < d; i++ {
			s.WriteString(L)
		}

	case 4 == place:
		s.WriteString(L + M)

	case 9 == place:
		s.WriteString(L + T)
	}
}

type roman struct {
	Numeral string
	Num     int
}

var sortedRomans = []roman{
	{
		Num:     1000,
		Numeral: "M",
	},
	{
		Num:     900,
		Numeral: "CM",
	},
	{
		Num:     500,
		Numeral: "D",
	},
	{
		Num:     400,
		Numeral: "CD",
	},
	{
		Num:     100,
		Numeral: "C",
	},
	{
		Num:     90,
		Numeral: "XC",
	},
	{
		Num:     50,
		Numeral: "L",
	},
	{
		Num:     40,
		Numeral: "XL",
	},
	{
		Num:     10,
		Numeral: "X",
	},
	{
		Num:     9,
		Numeral: "IX",
	},
	{
		Num:     5,
		Numeral: "V",
	},
	{
		Num:     4,
		Numeral: "IV",
	},
	{
		Num:     1,
		Numeral: "I",
	},
}

var l int = len(sortedRomans)

func intToRoman(num int) string {
	var numeral string

	places := getPlaces(num)

	for i := range places {
		place := places[i]

		j := sort.Search(l, func(j int) bool {
			return sortedRomans[j].Num <= place
		})

		roman := sortedRomans[j]

		m := place / roman.Num

		for mi := 0; mi < m; mi++ {
			numeral += roman.Numeral
		}

		d := place - (roman.Num * m)

		if d > 0 {
			k := sort.Search(l, func(k int) bool {
				return sortedRomans[k].Num <= d
			})

			r := sortedRomans[k]
			mr := d / r.Num

			for mi := 0; mi < mr; mi++ {
				numeral += r.Numeral
			}
		}
	}

	return numeral
}

func getPlaces(n int) []int {
	if n == 0 {
		return []int{0}
	}

	var p []int
	i := 1

	for n > 0 {
		v := (n % 10) * i
		if v > 0 {
			p = append(p, (n%10)*i)
		}

		n /= 10
		i *= 10
	}

	l := len(p)

	for i := l/2 - 1; i >= 0; i-- {
		j := l - 1 - i
		p[i], p[j] = p[j], p[i]
	}

	return p
}
