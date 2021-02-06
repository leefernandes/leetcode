package main

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

func main() {
	v := -3200000
	expected := 0
	s := strconv.Itoa(v)
	fmt.Println("s:", s)

	start := time.Now()
	r := divide(v)
	fmt.Println(time.Since(start))
	fmt.Println("result:", r == expected, r)

	// start = time.Now()
	// r = stringify(v)
	// fmt.Println(time.Since(start))
	// fmt.Println("result:", r == expected, r)
}

func stringify(x int) int {
	prefix := ""
	if x < 0 {
		prefix = "-"
		x = x * -1
	}

	s := strconv.Itoa(x)
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	v, _ := strconv.Atoi(prefix + string(r))
	return v
}

func divide(x int) int {
	if 0 == x {
		return 0
	}

	if x < math.MinInt32 || x > math.MaxInt32 {
		return 0
	}

	i := 0
	n := 0
	for x != 0 {
		d := x % 10
		n = n*10 + d
		x = x / 10
		i++
	}

	if n < math.MinInt32 || n > math.MaxInt32 {
		return 0
	}

	return n
}
