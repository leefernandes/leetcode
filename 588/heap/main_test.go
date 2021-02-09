package main

import (
	"fmt"
	"testing"
	"time"
)

type test struct {
	commands []string
	args     [][]string
	want     []interface{}
}

var tests = func() []test {
	return []test{
		{
			commands: []string{},
			args:     [][]string{},
			want:     []interface{}{},
		},
	}
}

func TestDoIt(t *testing.T) {
	start := time.Now()
	for i, tt := range tests() {
		testname := fmt.Sprintf("%d", i)
		t.Run(testname, func(t *testing.T) {
			fmt.Println(tt.commands, tt.args)

			// ["FileSystem","mkdir","ls",  "mkdir"," ls",  "ls",  "ls",  "addContentToFile",   "ls",   "ls",  "ls"]
			// [[],         ["/m"],  ["/m"],["/w"],  ["/"], ["/w"],["/"]  ,["/dycete","emer"],  ["/w"], ["/"], ["/dycete"]]

			fs := Constructor()

			// fs.AddContentToFile("/nope", "hello")

			// fmt.Println("ReadContentFromFile", fs.ReadContentFromFile("/nope"))

			fs.Mkdir("/m")
			fs.Ls("/m")

			fs.Mkdir("/w")
			fs.Ls("/")
			fs.Ls("/w")
			fs.Ls("/")

			fs.AddContentToFile("/dycete", "emer")
			//fmt.Println(fs.Ls("/w"))
			fmt.Println(fs.Ls("/w"))
			fmt.Println(fs.Ls("/"))
			fmt.Println(fs.Ls("/dycete"))
			fmt.Println(fs.Ls("/"))
			fmt.Println(fs.Ls("/"))
			fmt.Println(fs.Ls("/"))
			fmt.Println(fs.Ls("/"))
			//fmt.Println(fs.Ls("/dycete"))

			// fs.Mkdir("/e/f/")
			// fmt.Println("ls ", fs.Ls("/"))

			// nodes := map[string]int{
			// 	"banana": 3, "apple": 2, "pear": 4,
			// }

			// // Create a priority queue, put the items in it, and
			// // establish the priority queue (heap) invariants.
			// sortedNodes := sortedNodes{}
			// i := 0
			// for name := range nodes {
			// 	sortedNodes = append(sortedNodes, &node{
			// 		name:  name,
			// 		index: i,
			// 	})
			// 	i++
			// }

			// heap.Init(&sortedNodes)

			// // Insert a new item and then modify its priority.
			// n := &node{
			// 	name: "orange",
			// }
			// heap.Push(&sortedNodes, n)

			// // Take the items out; they arrive in decreasing priority order.
			// for sortedNodes.Len() > 0 {
			// 	n := heap.Pop(&sortedNodes).(*node)
			// 	fmt.Printf("%s ", n.name)
			// }

			//result := doIt()
			// if result != tt.want {
			// 	t.Errorf("🛑 got %v, want %v", result, tt.want)
			// }
		})
	}

	fmt.Println(time.Since(start))
}
