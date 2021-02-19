package main

func findLength(A []int, B []int) int {
	max := 0
	lenA := len(A)
	lenB := len(B)
	for i := range A {
		if lenA-i <= max {
			break
		}
		a := A[i]
		for j := range B {
			if lenB-j <= max {
				break
			}
			b := B[j]
			if a == b {
				if v := streak(A, B, lenA, lenB, i, j); v > max {
					max = v
				}
			}
		}
	}
	return max
}

func streak(A, B []int, lenA, lenB, i, j int) int {
	max := 1

	i++
	j++

	if i == lenA || j == lenB {
		return max
	}

	for A[i] == B[j] {
		max++
		i++
		j++

		if i == lenA || j == lenB {
			return max
		}
	}

	return max
}

func findLength2(A []int, B []int) int {
	max := 0
	lenA := len(A)
	lenB := len(B)

	if lenA < lenB {
		A, B = B, A
	}

	memo := make([]int, lenB)

	for i := 0; i < lenA; i++ {
		//fmt.Println("i:", i)

		next := make([]int, lenB)
		for j := 0; j < lenB; j++ {
			//fmt.Println("  j:", j, memo)

			var carryover int
			if i > 0 && j > 0 {
				carryover = memo[j-1]
			}

			var val int
			if A[i] == B[j] {
				val = carryover + 1
			} else {
				val = 0
			}

			next[j] = val

			if val > max {
				max = val
			}
		}

		memo = next
	}

	return max
}
