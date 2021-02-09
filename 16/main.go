package main

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	closest := 0
	distance := math.MaxInt64

	length := len(nums)

	if 3 == length {
		s := 0
		for i := 0; i < length; i++ {
			s += nums[i]
		}
		return s
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

			if v == target {
				return v
			}

			d := target - v
			if d < 0 {
				d *= -1
			}

			if d < distance {
				closest = v
				distance = d
			}

			if v < target {
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

	return closest
}
