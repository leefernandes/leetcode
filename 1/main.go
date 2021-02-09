package main

func twoSum(nums []int, target int) []int {
	l := len(nums)
	for i := 0; i < l; i++ {
		a := nums[i]
		start := i + 1
		for j := start; j < l; j++ {
			b := nums[j]
			if a+b == target {
				return []int{i, j}
			}
		}
	}

	return []int{}
}

func twoSumHash(nums []int, target int) []int {
	l := len(nums)

	hash := make(map[int]int, l)

	for i := 0; i < l; i++ {
		hash[nums[i]] = i
	}

	for i := 1; i < l; i++ {
		n := target - nums[i]
		if j, ok := hash[n]; ok && j != i {
			return []int{i, j}
		}
	}

	return []int{}
}
