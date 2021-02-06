package main

import (
	"fmt"
	"reflect"
	"time"
)

type test struct {
	nums []int
	k    int
	e    []int
}

type solution = func([]int, int)

var tests = func() []test {
	return []test{
		{
			nums: []int{1, 2, 3, 4, 5, 6, 7},
			k:    3,
			e:    []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			nums: []int{-1, -100, 3, 99},
			k:    2,
			e:    []int{3, 99, -1, -100},
		},
	}
}

var solutions = map[string]solution{
	"initial": rotate,
}

func main() {
	for name, solution := range solutions {
		run(name, solution, tests())
	}
}

func run(name string, s solution, tests []test) {
	fmt.Println("run", name)
	tt := tests //[2:]
	var d time.Duration
	for i := range tt {
		t := tt[i]
		//fmt.Println(i, t.a, t.d)
		start := time.Now()
		s(t.nums, t.k)
		d += time.Since(start)
		if !reflect.DeepEqual(t.nums, t.e) {
			fmt.Println(" 🛑", "test", i, "got:", t.nums, "expected:", t.e)
		} else {
			//fmt.Println(" 🟢", "got:", r, "expected:", t.e)
		}
	}
	avg := time.Duration(d.Nanoseconds() / int64(len(tt)))
	fmt.Println("  ", name, avg*time.Nanosecond)
}

func rotate(nums []int, k int) {
	l := len(nums)

	k %= l

	if 0 == k {
		return
	}

	// reverse down to k
	for i := k/2 - 1; i >= 0; i-- {
		x := l - k + i
		j := l - 1 - i
		nums[x], nums[j] = nums[j], nums[x]
	}

	// reverse down from k
	for i := (l-k)/2 - 1; i >= 0; i-- {
		j := l - k - 1 - i
		nums[i], nums[j] = nums[j], nums[i]
	}

	// // reverse all
	for i := l/2 - 1; i >= 0; i-- {
		j := l - 1 - i
		nums[i], nums[j] = nums[j], nums[i]
	}

	return
}
