package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	matches := [][]int{}
	l := len(nums)

	hash := map[int][]map[int]int{}

	for i := 0; i < l; i++ {
		ni := nums[i]
		for j := i + 1; j < l; j++ {
			nj := nums[j]
			n := ni + nj

			if _, ok := hash[n]; !ok {
				hash[n] = []map[int]int{}
			}
			hash[n] = append(hash[n], map[int]int{
				i: ni,
				j: nj,
			})

		}
	}

	dedupe := map[string]int{}
	for i := 0; i < l; i++ {
		ni := nums[i]
		n := 0 - ni
		if v, ok := hash[n]; ok {
			for k := range v {
				combo := v[k]

				if _, dupe := combo[i]; dupe {
					continue
				}

				adds := make([]int, 3)
				j2 := 0
				for _, j := range combo {
					adds[j2] = j
					j2++
				}

				adds[2] = ni
				sort.Ints(adds)

				var key string
				for j := range adds {
					key += fmt.Sprintf("%d", adds[j])
				}

				if _, dupe := dedupe[key]; dupe {
					continue
				}

				dedupe[key] = 1

				matches = append(matches, adds)
			}
		}
	}

	return matches
}

// 1,2,3,-3,-3,0

// -3,-3,0,1,1,2,2,3
type key struct {
	a, b, c int
}

func threeSum2(nums []int) [][]int {
	matches := [][]int{}
	l := len(nums)
	if l < 3 {
		return matches
	}

	sort.Ints(nums)

	if 3 == l {
		s := 0
		for i := 0; i < l; i++ {
			s += nums[i]
		}
		if 0 != s {
			return matches
		}
		return append(matches, nums)
	}

	u := map[key]int{}

	for i := 0; i < l-2; i++ {
		h := map[int]int{}
		in := nums[i]
		if i > 0 && nums[i-1] == in {
			continue
		}

		for j := i + 1; j < l; j++ {
			jn := nums[j]
			q := 0 - in - jn
			if _, ok := h[q]; ok {
				k := key{in, q, jn}
				if _, dupe := u[k]; dupe {
					continue
				}
				u[k] = 1

				matches = append(matches, []int{
					in, q, jn,
				})
			} else {
				h[jn] = j
			}
		}
	}

	return matches
}

func threeSum3(nums []int) [][]int {
	matches := [][]int{}
	length := len(nums)
	if length < 3 {
		return matches
	}

	if 3 == length {
		s := 0
		for i := 0; i < length; i++ {
			s += nums[i]
		}
		if 0 != s {
			return matches
		}
		return append(matches, nums)
	}

	sort.Ints(nums)

	for i := 0; i < length-2; i++ {
		iv := nums[i]
		if i > 0 && nums[i-1] == iv {
			continue
		}
		l := i + 1
		r := length - 1
		for l < r {
			lv := nums[l]
			rv := nums[r]

			v := iv + lv + rv

			if 0 == v {
				matches = append(matches, []int{iv, lv, rv})
				l++
				r--

				for l < r && nums[l] == lv {
					l++
				}
				for r > l && nums[r] == rv {
					r--
				}

			} else if v < 0 {
				l++
				for l < r && nums[l] == lv {
					l++
				}
			} else {
				r--
				for r > l && nums[r] == rv {
					r--
				}
			}
		}
	}

	return matches
}
