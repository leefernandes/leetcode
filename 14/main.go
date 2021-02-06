package main

import (
	"math"
)

func longestCommonPrefix(strs []string) string {
	var longest string
	var longestLen int

	l := len(strs)

	combos := map[string]int{}

	for i := 0; i < l; i++ {
		w := strs[i]

		var letters string
		for j := range w {
			letters += string(w[j])

			if _, ok := combos[letters]; ok {
				combos[letters]++
			} else {
				combos[letters] = 1
			}
		}
	}

	for k, n := range combos {
		kL := len(k)
		if n == l && kL > longestLen {
			longestLen = kL
			longest = k
		}
	}

	return longest
}

func longestCommonPrefix2(strs []string) string {
	var longest []byte
	var longestL int

	l := len(strs)

	if 0 == l {
		return ""
	} else if 1 == l {
		return strs[0]
	}

	for i := 0; i < l; i++ {
		wi := strs[i]
		wil := len(wi)

	JLoop:
		for j := i + 1; j < l; j++ {
			wj := strs[j]
			wjl := len(wj)

			var max int
			if 0 != longestL {
				max = longestL
				if max > wjl {
					max = wjl
				}
			} else if wil > wjl {
				max = wjl
			} else {
				max = wil
			}

			var kLongest []byte

			for k := 0; k < max; k++ {
				x := wi[k]
				if x != wj[k] {
					if 0 == k && 0 == i {
						return ""
					}
					longest = kLongest
					longestL = k
					continue JLoop
				}
				kLongest = append(kLongest, x)
			}
			longest = kLongest
		}
	}

	return string(longest)
}

func longestCommonPrefix3(strs []string) string {
	var longest []byte

	l := len(strs)

	if 0 == l {
		return ""
	} else if 1 == l {
		return strs[0]
	}

	var min int = math.MaxInt64

	for i := 0; i < l; i++ {
		w := strs[i]
		if wl := len(w); wl < min {
			min = wl
		}
	}

	var last byte
	for j := 0; j < min; j++ {
		for i := 0; i < l; i++ {
			c := strs[i][j]
			if last == 0 {
				last = c
			} else if c != last {
				return string(longest)
			}
		}
		longest = append(longest, last)
		last = 0
	}

	return string(longest)
}
