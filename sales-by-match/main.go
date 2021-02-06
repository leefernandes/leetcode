package main

import (
	"fmt"
	"time"
)

type test struct {
	ar []int32
	e  int32
}

func main() {
	tests := []test{
		{
			ar: []int32{1, 2, 1, 1, 2, 3, 2},
			e:  2,
		},
		{
			ar: []int32{1},
			e:  0,
		},
		{
			ar: []int32{4, 4, 4, 4, 4, 5, 5},
			e:  3,
		},
	}

	tt := tests[0:1]
	var d time.Duration
	for i := range tt {
		t := tt[i]
		fmt.Println(i, len(t.ar), t.ar)
		start := time.Now()
		r := sockMerchant(int32(len(t.ar)), t.ar)
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

func sockMerchant(n int32, ar []int32) int32 {
	var pairs int32 = 0
	if n < 2 {
		return pairs
	}
	socks := map[int32]int32{}
	for i := range ar {
		color := ar[i]
		_, ok := socks[color]
		if !ok {
			socks[color] = 0
		}
		socks[color]++
	}

	for _, count := range socks {
		pairs += count / 2
	}

	return pairs
}
