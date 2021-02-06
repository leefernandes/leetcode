package main

import (
	"fmt"
	"time"
)

type test struct {
	s string
	p string
	e bool
}

func main() {
	tests := []test{
		{
			s: "aasdfasdfasdfasdfas",
			p: "aasdf.*asdf.*asdf.*asdf.*s",
			e: true,
		},
		{
			s: "aaa",
			p: "ab*a*c*a",
			e: true,
		},
		{
			s: "p",
			p: "pp*",
			e: true,
		},
		{
			s: "a",
			p: "ap*",
			e: true,
		},
		{
			s: "pp",
			p: ".pp*",
			e: true,
		},
		{
			s: "pppppp",
			p: ".pp*",
			e: true,
		},
		{
			s: "xpppppp",
			p: "xpp*",
			e: true,
		},
		{
			s: "xpppppp",
			p: "x.pp*",
			e: true,
		},
		{
			s: "xypppppp",
			p: "x.pp*",
			e: true,
		},
		{
			s: "yxpppppp",
			p: "x.pp*",
			e: false,
		},
		{
			s: "xpppppp",
			p: "y.pp*",
			e: false,
		},
		{
			s: "aa",
			p: "aa",
			e: true,
		},
		{
			s: "aa",
			p: "aaa",
			e: false,
		},
		{
			s: "aa",
			p: "a",
			e: false,
		},
		{
			s: "aa",
			p: "a*",
			e: true,
		},
		{
			s: "asdfasd",
			p: ".*x",
			e: false,
		},
		{
			s: "yasdfas",
			p: "y.*",
			e: true,
		},
		{
			s: "yasdfasx",
			p: "y.*x",
			e: true,
		},
		{
			s: "yasdfasy",
			p: "y.*x",
			e: false,
		},
		{
			s: "asdfasdx",
			p: ".*x",
			e: true,
		},
		{
			s: "asdfasd",
			p: ".*.*",
			e: true,
		},
		{
			s: "aaaaaaapple",
			p: "a*p*l.",
			e: true,
		},
		{
			s: "aaaaaaapple",
			p: "a*a.a*ppp*le*",
			e: true,
		},
		{
			s: "aaaaaaapple",
			p: "a*a.a*pppp*le*",
			e: false,
		},
		{
			s: "caaaaandy",
			p: "ca*ndy",
			e: true,
		},
		{
			s: "caaaaandy",
			p: "ca*aaaandy",
			e: true,
		},
		{
			s: "caaaaandy",
			p: "caa*aaaaan*dy",
			e: false,
		},
		{
			s: "caaaaandy",
			p: "caa*aaaa*aa*andy",
			e: false,
		},
		{
			s: "ab",
			p: ".*",
			e: true,
		},
		{
			s: "aab",
			p: "c*a*b",
			e: true,
		},
		{
			s: "mississippi",
			p: "mis*is*p*.",
			e: false,
		},
		{
			s: "mississippi",
			p: "mis*is*i*p*.",
			e: true,
		},
	}

	tt := tests[0:1]
	var d time.Duration
	for i := range tt {
		t := tt[i]
		fmt.Println(i, t.s, t.p)
		start := time.Now()
		r := isMatch2(t.s, t.p)
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

func isMatch2(s string, p string) bool {
	if s == p {
		return true
	}
	pl := len(p)
	sl := len(s)

	lookaheadMap := map[byte][]byte{}

	// string cursor
	j := sl - 1

	// pattern cursor
	i := pl - 1
	for i > -1 {
		pc := p[i]
		fmt.Println(" i:", i, string(pc), "j:", j)

		if 42 == pc {
			n := i - 1
			m := p[n]
			if _, ok := lookaheadMap[m]; !ok {
				lookaheadMap[m] = []byte{}
			}
			// .* matches everything
			if 46 == m {
				for i := 0; i < j; i++ {
					lookaheadMap[m] = append(lookaheadMap[m], s[j])
				}
				j = 0
				i--
				continue
			}

			min := 0
			for n > 0 {
				n--
				v := p[n]
				if v != m {
					break
				}
				min++
			}

			if sl-min < 0 {
				return false
			}

			x := 0
			for x <= j {
				v := s[j-x]
				if v != m {
					if x < min {
						return false
					}
					break
				}
				lookaheadMap[m] = append(lookaheadMap[m], v)
				x++
			}

			i -= min + 2
			j -= x

		} else {
			var lookahead []byte
			if v, ok := lookaheadMap[pc]; !ok {
				if v, ok := lookaheadMap[46]; !ok {
					lookahead = []byte{}
				} else {
					lookahead = v
					fmt.Println("wtf?", lookahead[len(lookahead)-1], string(lookahead[len(lookahead)-1]))
				}
			} else {
				lookahead = v
			}

			if len(lookahead) > 0 {
				fmt.Println("aaaaay!", pc, string(pc), lookahead[len(lookahead)-1], string(lookahead[len(lookahead)-1]))
				fmt.Println(46 == pc)
				fmt.Println(pc == lookahead[len(lookahead)-1])
				fmt.Println(46 == lookahead[len(lookahead)-1])
			}
			// use asertisk overflow
			if j < i && len(lookahead) > 0 && (46 == pc || pc == lookahead[len(lookahead)-1] || 46 == lookahead[len(lookahead)-1]) {
				i--
				continue
			}

			if j < 0 {
				return false
			}

			fmt.Println("    jesus fucking christ", i, string(pc), j, s[j], string(s[j]))
			if 46 != pc && s[j] != pc {
				return false
			}

			i--
			j--
		}
	}

	if i != j {
		return false
	}

	return true
}

// isMatch failing on lookaheads
// {
// 	s: "aaaaaaapple",
// 	p: "a*a.a*ppp*le*",
// 	e: true,
// },
func isMatch(s string, p string) bool {
	if s == p {
		return true
	}
	pl := len(p)
	sl := len(s)

	var lookahead byte
	var lookaheadL int

	// string cursor
	j := 0

	// pattern cursor
	i := 0
	for i < pl {
		pc := p[i]

		switch pc {
		// .
		case 46:
			lookahead = 0
			lookaheadL = 0
			fmt.Println("  wildcard")

		// *
		case 42:
			lookahead = p[i-1]
			fmt.Println("  lookahead", i, string(lookahead)+"*")
			for j < sl {
				n := s[j]
				fmt.Println("    peering", j, string(n))
				if lookahead != 46 && n != lookahead {
					break
				}
				lookaheadL++
				j++
			}

		default:
			// exit if pattern exceeds string length
			if j >= sl {
				// handle wildcard lookahead
				if lookahead == 46 {
					return pc == 46 || pc == s[sl-1]
				}
				return false
			}
			sc := s[j]
			fmt.Println("  character", "i:", i, string(pc), "j:", j, string(sc))
			if pc != 46 && sc != pc {
				// contract lookahead buffer by 1 if sc matches
				fmt.Println("    contract lookahead")
				if 0 == lookaheadL {
					return false
				}
				if pc != lookahead {
					return false
				}
				lookaheadL--
			} else {
				j++
			}
		}
		i++
	}

	if j < sl {
		return false
	}

	return true
}
