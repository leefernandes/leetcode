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

			// ["FileSystem",    "ls",   "mkdir",     "addContentToFile",      "ls",    "readContentFromFile"]
			// [[],             ["/"],   ["/a/b/c"],   ["/a/b/c/d","hello"],   ["/"],   ["/a/b/c/d"]]

			// ["FileSystem",  "ls",    "mkdir",      "ls"]
			// [[],            ["/"],   ["/a/b/c"],    ["/a/b"]]

			fs := Constructor()

			fmt.Println("ls /", fs.Ls("/"))

			fs.Mkdir("/a/b/c")

			fmt.Println("ls /a/b", fs.Ls("/a/b"))

			fmt.Println("ls /", fs.Ls("/"))

			// fs.AddContentToFile("/a/b/c/d", "hello")

			// fmt.Println("ls /", fs.Ls("/"))

			// fmt.Println("readContentFromFile /a/b/c/d", fs.ReadContentFromFile("/a/b/c/d"))

			//result := doIt()
			// if result != tt.want {
			// 	t.Errorf("🛑 got %v, want %v", result, tt.want)
			// }
		})
	}

	fmt.Println(time.Since(start))
}
