package main

func isValid(s string) bool {
	q := []byte{}
	for i := range s {
		c := s[i]

		switch c {
		// ( [ {
		case 40, 91, 123:
			q = append(q, c)

			// )
		case 41:
			if v := pop(&q); 0 == v || v != 40 {
				return false
			}

			// ]
		case 93:
			if v := pop(&q); 0 == v || v != 91 {
				return false
			}

			// }
		case 125:
			if v := pop(&q); 0 == v || v != 123 {
				return false
			}
		}
	}

	return 0 == len(q)
}

func pop(q *[]byte) byte {
	qv := *q
	l := len(qv) - 1

	if -1 == l {
		return 0
	}

	var c byte
	c, *q = qv[l], qv[:l]

	return c
}

func isValid2(s string) bool {
	sl := len(s)
	if 0 != sl%2 {
		return false
	}
	q := []byte{}
	for i := 0; i < sl; i++ {
		c := s[i]

		switch c {
		// ( [ {
		case 40, 91, 123:
			q = append(q, c)

			// )
		case 41:
			l := len(q) - 1
			if -1 == l || q[l] != 40 {
				return false
			}
			q = q[:l]

			// ]
		case 93:
			l := len(q) - 1
			if -1 == l || q[l] != 91 {
				return false
			}
			q = q[:l]

			// }
		case 125:
			l := len(q) - 1
			if -1 == l || q[l] != 123 {
				return false
			}
			q = q[:l]

		}
	}

	return 0 == len(q)
}
