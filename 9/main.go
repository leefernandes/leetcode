package main

import (
	"fmt"
	"math"
	"time"
)

type test struct {
	v int
	e bool
}

func main() {
	tests := []test{
		{
			v: 1000021,
			e: false,
		},
		{
			v: -1212,
			e: false,
		},
		{
			v: 1212,
			e: false,
		},
		{
			v: -1444441,
			e: false,
		},
		{
			v: 1444441,
			e: true,
		},
		{
			v: 123321,
			e: true,
		},
		{
			v: 33,
			e: true,
		},
		{
			v: 98789,
			e: true,
		},
		{
			v: 98789,
			e: true,
		},
		{
			v: 9,
			e: true,
		},
		{
			v: -9,
			e: false,
		},
		{
			v: 888888,
			e: true,
		},
		{
			v: 8888888,
			e: true,
		},
		{
			v: 121,
			e: true,
		},
	}

	tt := tests
	var d time.Duration
	for i := range tt {
		t := tt[i]
		start := time.Now()
		r := isPalindrome2(t.v)
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

func isPalindrome3(x int) bool {
	if x < 0 {
		return false
	}

	old := x
	new := 0
	for ; x != 0; x = x / 10 {
		new = new*10 + (x % 10)
	}
	return new == old
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	} else if x < 10 {
		return true
	}

	n := 0
	y := x
	for y != 0 {
		d := y % 10
		n = n*10 + d
		y = y / 10
	}

	return n == x
}

func isPalindrome2(x int) bool {
	if x < 0 {
		return false
	} else if x < 10 {
		return true
	}

	l := int(math.Log10(float64(x))+1) / 2

	for i := 0; i < l; i++ {
		ll := int(math.Log10(float64(x)))
		lp := int(math.Pow10(ll - i))
		l := (x / lp) % 10
		r := x % 10
		if l != r {
			return false
		}
		if 0 == l {
			l = 1
		}
		x = x / 10
	}

	return true
}
