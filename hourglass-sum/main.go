package main

import (
	"fmt"
	"math"
	"time"
)

type test struct {
	arr [][]int32
	e   int32
}

type summer = func([][]int32) int32

var tests = []test{
	{
		arr: [][]int32{
			{-9, -9, -9, 1, 1, 1},
			{0, -9, 0, 4, 3, 2},
			{-9, -9, -9, 1, 2, 3},
			{0, 0, 8, 6, 6, 0},
			{0, 0, 0, -2, 0, 0},
			{0, 0, 1, 2, 4, 0},
		},
		e: 28,
	},
	{
		arr: [][]int32{
			{1, 1, 1, 0, 0, 0},
			{0, 1, 0, 0, 0, 0},
			{1, 1, 1, 0, 0, 0},
			{0, 0, 2, 4, 4, 0},
			{0, 0, 0, 2, 0, 0},
			{0, 0, 1, 2, 4, 0},
		},
		e: 19,
	},
	{
		arr: [][]int32{
			{-1, -1, 0, -9, -2, -2},
			{-2, -1, -6, -8, -2, -5},
			{-1, -1, -1, -2, -3, -4},
			{-1, -9, -2, -4, -4, -5},
			{-7, -3, -3, -2, -9, -9},
			{-1, -3, -1, -2, -4, -5},
		},
		e: -6,
	},
	{
		arr: [][]int32{
			{0, -4, -6, 0, -7, -6},
			{-1, -2, -6, -8, -3, -1},
			{-8, -4, -2, -8, -8, -6},
			{-3, -1, -2, -5, -7, -4},
			{-3, -5, -3, -6, -6, -6},
			{-3, -6, 0, -8, -6, -7},
		},
		e: -19,
	},
}

func main() {
	run("mine", hourglassSum)
	run("hackguru", hourglassSum2)
}

func run(name string, impl summer) {
	fmt.Println("run", name)
	tt := tests //[2:]
	var d time.Duration
	for i := range tt {
		t := tt[i]
		fmt.Println(i, t.arr)
		start := time.Now()
		r := impl(t.arr)
		d += time.Since(start)
		if r != t.e {
			fmt.Println(" 🛑", r == t.e, "got:", r, "expected:", t.e)
		} else {
			fmt.Println(" 🟢", r == t.e, "got:", r, "expected:", t.e)
		}
	}
	avg := time.Duration(d.Nanoseconds() / int64(len(tt)))
	fmt.Println(name, avg*time.Nanosecond)
}

func hourglassSum(arr [][]int32) int32 {
	var max int32 = math.MinInt32

	for row := 0; row < 4; row++ {
		for col := 0; col < 4; col++ {
			sum := arr[row][col] + arr[row][col+1] + arr[row][col+2] +
				arr[row+1][col+1] +
				arr[row+2][col] + arr[row+2][col+1] + arr[row+2][col+2]

			if sum > max {
				max = sum
			}
		}
	}

	return max
}

func hourglassSum2(arr [][]int32) int32 {
	max := int32(math.MinInt32)
	j := 0
	i := 0
	for j < 4 {
		tmp := arr[i][j] + arr[i][j+1] + arr[i][j+2] + arr[i+1][j+1] + arr[i+2][j] + arr[i+2][j+1] + arr[i+2][j+2]
		if tmp > max {
			max = tmp
		}
		i++
		if i == 4 {
			i = 0
			j++
		}
	}
	return max
}
