package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {

	n := 1
	s := "AB"
	start := time.Now()
	v := convert(s, n)
	fmt.Println(time.Since(start))
	expected := "AB"
	fmt.Println("result:", v == expected, v)
}

func convert(s string, numRows int) string {
	if 1 == numRows {
		return s
	}

	n := len(s)
	row := 0
	col := 0

	j := 0
	e := numRows - 1
	c := numRows + numRows - 2

	v := [][]byte{}

	for i := 0; i < n; i++ {
		b := s[i]

		if j < numRows {
			row = j
		} else {
			col++
			row = e - (j - e)
		}

		if len(v) < row+1 {
			v = append(v, []byte{b})
		} else {
			v[row] = append(v[row], b)
		}

		j++
		if j == c {
			col++
			j = 0
		}
	}

	var sb strings.Builder
	for i := range v {
		sb.Write(v[i])
	}

	return sb.String()
}
