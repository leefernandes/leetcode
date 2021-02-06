package main

import (
	"fmt"
	"math/big"
	"strconv"
	"time"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1Ints := []int{
		1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 9, 1,
	}

	n := len(l1Ints)

	l1 := make([]*ListNode, n)

	for i := range l1Ints {
		if i > 0 {
			l1[i] = &ListNode{
				Val:  l1Ints[i],
				Next: l1[i-1],
			}
		} else {
			l1[i] = &ListNode{Val: l1Ints[i]}

		}
	}

	//a := ListNode{Val: 2}
	//b := ListNode{Val: 4}
	//c := ListNode{Val: 3}
	//a.Next = &b
	//b.Next = &c

	d := ListNode{Val: 5}
	e := ListNode{Val: 6}
	f := ListNode{Val: 4}
	d.Next = &e
	e.Next = &f

	start := time.Now()
	v := addTwoNumbers(l1[n-1], &d)
	fmt.Println(time.Since(start))
	fmt.Printf("%v", v)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	l1V := listValue(l1)
	l2V := listValue(l2)

	sum := big.NewInt(0)
	sum.Add(l1V, l2V)

	s := sum.String()

	n := len(s)

	l := make([]*ListNode, n)
	for i := range s {
		v, err := strconv.Atoi(string(s[i]))
		if err != nil {
			panic(err)
		}
		if i > 0 {
			l[i] = &ListNode{
				Val:  v,
				Next: l[i-1],
			}
		} else {
			l[i] = &ListNode{
				Val: v,
			}
		}
	}

	return l[n-1]
}

func listValue(l *ListNode) *big.Int {
	s := ""
	for l.Next != nil {
		s = fmt.Sprintf("%v%s", l.Val, s)
		l = l.Next

	}
	s = fmt.Sprintf("%v%s", l.Val, s)
	i := new(big.Int)
	i.SetString(s, 0)
	return i
}
