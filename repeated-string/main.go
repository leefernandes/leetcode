package main

import (
	"fmt"
	"time"
)

type test struct {
	s string
	n int64
	e int64
}

func main() {
	tests := []test{
		{
			s: "ababa",
			n: 3,
			e: 2,
		},
		{
			s: "aba",
			n: 10,
			e: 7,
		},
		{
			s: "a",
			n: 1000000000000,
			e: 1000000000000,
		},
		{
			s: "abcac",
			n: 10,
			e: 4,
		},
		{
			s: "x",
			n: 970770,
			e: 0,
		},
	}

	tt := tests //[0:1]
	var d time.Duration
	for i := range tt {
		t := tt[i]
		fmt.Println(i, t.s, t.n)
		start := time.Now()
		r := repeatedString(t.s, t.n)
		d += time.Since(start)
		if r != t.e {
			fmt.Println(" 🛑", r == t.e, "got:", r, "expected:", t.e)
		} else {
			fmt.Println(" 🟢", r == t.e, "got:", r, "expected:", t.e)
		}
	}
	avg := time.Duration(d.Nanoseconds() / int64(len(tt)))
	fmt.Println(avg * time.Nanosecond)
}

func repeatedString(s string, n int64) int64 {
	l := len(s)
	if "a" == s {
		return n
	}

	var count int64 = 0

	if l >= int(n) {
		var i int64
		for i = 0; i < n; i++ {
			c := s[i]
			if 97 == c {
				count++
			}
		}

		return count
	}

	d := int(n) - l
	r := d / l

	for i := 0; i < l; i++ {
		c := s[i]
		if 97 == c {
			count++
		}
	}

	count += int64(r) * count

	rollover := d - (r * l)
	for i := 0; i < rollover; i++ {
		c := s[i]
		if 97 == c {
			count++
		}
	}

	return count
}
