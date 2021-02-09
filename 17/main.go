package main

var letters = map[byte][]string{
	50: {"a", "b", "c"},
	51: {"d", "e", "f"},
	52: {"g", "h", "i"},
	53: {"j", "k", "l"},
	54: {"m", "n", "o"},
	55: {"p", "q", "r", "s"},
	56: {"t", "u", "v"},
	57: {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	combos := []string{}

	l := len(digits)

	if 0 == l {
		return combos
	}

	start := letters[digits[0]]

	for i := range start {
		q := []string{start[i]}
		for j := 1; j < l; j++ {
			jv := letters[digits[j]]
			var next []string
			for k := range jv {
				for l := range q {
					next = append(next, q[l]+jv[k])
				}
			}
			q = next
		}
		combos = append(combos, q...)
	}

	return combos
}

func letterCombinationsQueue(digits string) []string {
	combos := []string{}

	l := len(digits)

	if 0 == l {
		return combos
	}

	q := []string{""}
	for len(q) > 0 {
		var s string
		s, q = q[0], q[1:]
		if len(s) == l {
			combos = append(combos, s)
		} else {
			next := letters[digits[len(s)]]
			for i := range next {
				q = append(q, s+next[i])
			}
		}
	}

	return combos
}
