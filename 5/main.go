package main

import (
	"fmt"
)

func main() {
	//s := "abaraxbagessega"
	//s := "zloxoly"
	//s := "zllllzlboooob"
	//s := "xanthaxanaranagrammargan"
	//s := "xmeufheabbajckdalf"
	s := "abb"
	v := longestPalindrome(s)
	fmt.Println("longest", v)
}

type result struct {
	b []byte
	l int
}

func longestPalindrome(s string) string {
	return mo(s)

	//return ltr(s)
}

// mo loops from the middle outward
// left & right finding the first largest
// odd or even palindrome
func mo(s string) string {
	if "" == s {
		return ""
	}

	fmt.Println("s:", s)
	b := []byte(s)
	l := len(b)
	m := l / 2
	even := false
	if 0 == l%2 {
		fmt.Println("even")
		even = true
	}

	var win result

	dmax := m + 1
	for d := 0; d < dmax; d++ {

		var lodd result
		var rodd result
		var leven result
		var reven result

		if even {
			fmt.Println("possible remaining", (m-d)*2, l)
			if (m-d)*2 <= win.l {
				return string(win.b)
			}
			lm := m - d

			fmt.Println("left even")
			leven = findEven(b, lm, l)
			fmt.Println("  ", leven.l, string(leven.b))
			if leven.l == l {
				fmt.Println("  🎂")
				return s
			}

			rm := m + d + 1

			fmt.Println("right even")
			reven = findEven(b, rm-1, l)
			fmt.Println("  ", reven.l, string(reven.b))
			if reven.l == l {
				fmt.Println("  🎂")
				return s
			}

			fmt.Println("left odd")
			lodd = findOdd(b, lm, l)
			fmt.Println("  ", lodd.l, string(lodd.b))

			fmt.Println("right odd")
			rodd = findOdd(b, rm-1, l)
			fmt.Println("  ", rodd.l, string(rodd.b))

		} else {
			fmt.Println("possible remaining", (dmax-d)*2, l)
			if (dmax-d)*2 <= win.l {
				return string(win.b)
			}

			lm := m - d

			fmt.Println("left odd")
			lodd = findOdd(b, lm, l)
			fmt.Println("  ", lodd.l, string(lodd.b))
			if lodd.l == l {
				fmt.Println("  🎂")
				return s
			}

			rm := m + d + 1

			fmt.Println("right odd")
			rodd = findOdd(b, rm, l)
			fmt.Println("  ", rodd.l, string(rodd.b))
			if rodd.l == l {
				fmt.Println("  🎂")
				return s
			}

			fmt.Println("left even")
			leven = findEven(b, lm, l)
			fmt.Println("  ", leven.l, string(leven.b))

			fmt.Println("right even")
			reven = findEven(b, rm-1, l)
			fmt.Println("  ", reven.l, string(reven.b))
		}

		// figure max remaining size from here

		if lodd.l > win.l {
			win = lodd
		}

		if rodd.l > win.l {
			win = rodd
		}

		if leven.l > win.l {
			win = leven
		}

		if reven.l > win.l {
			win = reven
		}
	}

	return string(win.b)
}

// odd checks for odd length palindrome
func findOdd(b []byte, m int, l int) result {
	if 0 == m {
		return result{
			b: b[0:1],
			l: 1,
		}
	}
	if m == l {
		v := b[l-1 : l]
		return result{
			b: v,
			l: len(v),
		}
	}
	max := m + 1
	for d := 1; d < max; d++ {
		xi := m - d
		x := b[xi : xi+1]
		yi := m + d
		y := b[yi : yi+1]
		fmt.Println("  m:", m, string(b[m]))
		fmt.Println("  x:", xi, string(x[0]))
		fmt.Println("  y:", yi, string(y[0]))
		fmt.Println("  range:", string(b[xi:yi+1]))
		if x[0] != y[0] {
			v := b[m+1-d : m+d]
			return result{
				b: v,
				l: len(v),
			}
		}
	}
	v := b[m-m : m+(m)+1]
	return result{
		b: v,
		l: len(v),
	}
}

// even checks for even length palindrome
func findEven(b []byte, m int, l int) result {
	if 0 == m {
		return result{
			b: b[0:1],
			l: 1,
		}
	}
	if m == l {
		v := b[l-1 : l]
		return result{
			b: v,
			l: len(v),
		}
	}
	max := m + 1
	fmt.Println("max?", max, len(b))
	for d := 1; d < max; d++ {
		xi := m - d
		x := b[xi : xi+1]
		yi := m + d - 1
		y := b[yi : yi+1]
		fmt.Println("  m:", m, string(b[m]))
		fmt.Println("  x:", xi, string(x[0]))
		fmt.Println("  y:", yi, string(y[0]))
		fmt.Println("  range:", string(b[xi:yi+1]))
		if x[0] != y[0] {
			v := b[m+1-d : m+d-1]
			return result{
				b: v,
				l: len(v),
			}
		}
	}
	v := b[m-m : m+(m)]
	return result{
		b: v,
		l: len(v),
	}
}

// @todo cannot pick first because favors first left match instead of longest
func ltr(s string) string {
	b := []byte(s)
	l := len(b)
	e := l - 1
	for i := range b {
		if i == e {
			return s[0:1]
		}
		bb := b[i:l]
		if m := rtl(bb); m != nil {
			return string(m)
		}
	}

	return s[0:1]
}

func rtl(b []byte) []byte {
	l := len(b)
	for i := l; i > 1; i-- {
		var x []byte
		var y []byte
		m := i / 2
		var t string
		if 0 == i%2 {
			// even
			x = b[0:m]
			y = b[m:i]
			t = "even"

		} else {
			// odd
			x = b[0:m]
			y = b[m+1 : i]
			t = "odd"
		}
		fmt.Println(t, string(b), string(x), string(y))
		if ok := match(x, y, m); ok {
			fmt.Println("!!  match", string(b[0:i]))
			return b[0:i]
		}
	}
	return nil
}

func match(a []byte, b []byte, m int) bool {
	//fmt.Println("a:", string(a), "b:", string(b))
	for i := 0; i < m; i++ {
		aa := a[i]
		bb := b[m-i-1]
		if aa != bb {
			return false
		}
	}
	return true
}
