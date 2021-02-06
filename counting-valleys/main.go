package main

import (
	"fmt"
	"time"
)

type test struct {
	path string
	e    int32
}

func main() {
	tests := []test{
		{
			path: "DDUUUUDD",
			e:    1,
		},
		{
			path: "UDDDUDUU",
			e:    1,
		},
		{
			path: "DUDUDUDUDU",
			e:    5,
		},
		{
			path: "DUDDUDUDUDU",
			e:    1,
		},
		{
			path: "DDUDDUD",
			e:    0,
		},
	}

	tt := tests //[0:1]
	var d time.Duration
	for i := range tt {
		t := tt[i]
		fmt.Println(i, len(t.path), t.path)
		start := time.Now()
		r := countingValleys(int32(len(t.path)), t.path)
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

func countingValleys(steps int32, path string) int32 {
	var valleys int32 = 0
	altitude := 0

	for i := range path {
		c := path[i]
		switch c {
		// D
		case 68:
			altitude--

		// U
		case 85:
			if -1 == altitude {
				valleys++
			}
			altitude++
		}
	}

	return valleys
}
