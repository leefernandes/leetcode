package main

import (
	"fmt"
	"reflect"
	"time"
)

type test struct {
	people [][]int
	e      int
}

type solution = func([][]int) int

var tests = func() []test {
	return []test{
		{
			people: [][]int{
				{1900, 1980},
				{2000, 0},
				{1979, 2015},
				{1950, 1955},
			},
			e: 2010,
		},
	}
}

var solutions = map[string]solution{
	"initial": livingPeople,
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
		r := s(t.people)
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

// return year w/ most living people
func livingPeople(people [][]int) int {
	year := 0
	l := len(people)

	events := make([][]int, l*2)
	for i := range people {
		person := people[i]
		birth := person[0]
		death := person[1]
		events[i] = []int{birth, 1}
		if 0 == death {
			events[i+1] = []int{birth, 0}
		} else {
			events[i+1] = []int{death, -1}
		}
	}

	return year
}
