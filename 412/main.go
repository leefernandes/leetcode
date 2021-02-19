package main

import "fmt"

func fizzBuzz(n int) []string {
	res := make([]string, n)

	for i := 1; i <= n; i++ {
		var s string
		if 0 == i%3 {
			s += "Fizz"
		}
		if 0 == i%5 {
			s += "Buzz"
		}

		if "" == s {
			s = fmt.Sprintf("%d", i)
		}

		res[i-1] = s
	}

	return res
}
