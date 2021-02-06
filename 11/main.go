package main

import (
	"fmt"
	"reflect"
	"time"
)

type test struct {
	height []int
	e      int
}

type solution = func([]int) int

var tests = func() []test {
	return []test{
		{
			height: []int{1, 1},
			e:      1,
		},
		{
			height: []int{4, 3, 2, 1, 4},
			e:      4,
		},
		{
			height: []int{30000, 1, 2, 3, 30000},
			e:      4,
		},
		{
			height: []int{1, 2, 1},
			e:      2,
		},
		{
			height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			e:      49,
		},
		{
			height: []int{1, 28, 6, 2, 5231, 2344, 8, 3, 7},
			e:      2344,
		},
		{
			height: []int{1, 28, 2361, 2, 3, 38, 3, 7223},
			e:      11805,
		},
		{
			height: []int{30000, 23328, 0, 63, 2, 5231, 22344, 8, 3, 7, 0},
			e:      134064,
		},
	}
}

var solutions = map[string]solution{
	"initial":  maxArea,
	"optimize": maxArea2,
	"learned":  maxArea3,
	"copied":   maxArea4,
}

func main() {
	for name, solution := range solutions {
		run(name, solution, tests())
	}
}

func run(name string, s solution, tests []test) {
	fmt.Println("run", name)
	tt := tests[len(tests)-1:]
	var d time.Duration
	for i := range tt {
		t := tt[i]
		//fmt.Println(i, t.a, t.d)
		start := time.Now()
		r := s(t.height)
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

func maxArea4(height []int) int {
	var maxArea int
	first := 0
	last := len(height) - 1

	for first < last {
		maxArea = max(maxArea, min(height[first], height[last])*(last-first))

		if height[first] < height[last] {
			first = first + 1
		} else {
			last = last - 1
		}
	}

	return maxArea
}

func max(a, b int) int {
	if a > b || a == b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b || a == b {
		return a
	}
	return b
}

func maxArea3(height []int) int {
	var maxArea int

	l := 0
	r := len(height) - 1
	for l < r {
		d := r - l

		var m int
		left := height[l]
		right := height[r]
		if left < right {
			m = left
			l++
		} else {
			m = right
			r--
		}

		if a := m * d; a > maxArea {
			maxArea = a
		}

	}

	return maxArea
}

func maxArea2(height []int) int {
	var maxArea int

	l := len(height) - 1

	var left int
	for i := range height {
		if height[i] <= left {
			continue
		}

		left = height[i]

		var right int
		for j := l; j > i; j-- {
			if height[j] <= right {
				continue
			}

			right = height[j]

			d := j - i

			var min int
			if left < right {
				min = left
			} else {
				min = right
			}

			area := min * d
			if area > maxArea {
				maxArea = area
			}
		}
	}

	return maxArea
}

func maxArea(height []int) int {
	var maxArea int

	l := len(height) - 1

	for i := range height {
		left := height[i]
		for j := l; j > i; j-- {
			right := height[j]
			d := j - i

			var min int
			if left < right {
				min = left
			} else {
				min = right
			}

			area := min * d
			if area > maxArea {
				maxArea = area
			}
		}
	}

	return maxArea
}
