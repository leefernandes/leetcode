package main

import (
	"fmt"
	"strconv"
	"time"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// [9,9,9,9,9,9,9]
// [9,9,9,9]

func main() {
	l1 := makeList([]int{
		//1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 9, 1,
		9, 9, 9, 9, 9, 9, 9,
	})

	l2 := makeList([]int{
		//8, 7, 2, 1,
		9, 9, 9, 9,
	})

	start := time.Now()
	v := addTwoNumbers(l1, l2)
	fmt.Println(time.Since(start))

	b := []byte{}

	for v.Next != nil {
		b = append(b, strconv.Itoa(v.Val)...)
		v = v.Next
	}
	b = append(b, strconv.Itoa(v.Val)...)

	fmt.Println(string(b))

	fmt.Printf("%v", v)
}

func makeList(values []int) *ListNode {
	n := len(values)

	l := make([]*ListNode, n)

	for i := range values {
		if i > 0 {
			l[i] = &ListNode{
				Val:  values[i],
				Next: l[i-1],
			}
		} else {
			l[i] = &ListNode{Val: values[i]}
		}
	}

	return l[n-1]
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var p ListNode
	var f *ListNode

	c := 0

	for l1 != nil || l2 != nil || c != 0 {
		v := c

		if l1 != nil {
			v += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			v += l2.Val
			l2 = l2.Next
		}

		if v >= 10 {
			v = v - 10
			c = 1
		} else {
			c = 0
		}

		n := ListNode{
			Val: v,
		}

		if f == nil {
			f = &n
		} else {
			p.Next = &n
		}

		p = n
	}

	return f
}
