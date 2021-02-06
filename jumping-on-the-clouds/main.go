package main

import (
	"fmt"
	"time"
)

type test struct {
	c []int32
	e int32
}

func main() {
	tests := []test{
		{
			c: []int32{0, 0, 0, 0, 1, 0},
			e: 3,
		},
		{
			c: []int32{0, 0, 1, 0, 0, 1, 0},
			e: 4,
		},
		{
			c: []int32{0, 0, 0},
			e: 1,
		},
		{
			c: []int32{0, 0},
			e: 1,
		},
		{
			c: []int32{0, 1, 0, 1, 0, 1, 0, 1, 0, 0},
			e: 5,
		},
	}

	tt := tests //[0:1]
	var d time.Duration
	for i := range tt {
		t := tt[i]
		fmt.Println(i, t.c)
		start := time.Now()
		r := jumpingOnClouds(t.c)
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

func jumpingOnClouds(c []int32) int32 {
	var jumps int32 = 0
	end := len(c) - 1
	p := 0

	for p < end {
		p2 := p + 2
		if p2 == end || p2 < end && c[p2] != 1 {
			p = p2
		} else {
			p++
		}
		jumps++
	}

	return jumps
}
