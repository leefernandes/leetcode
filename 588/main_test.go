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

			doIt()
			//result := doIt()
			// if result != tt.want {
			// 	t.Errorf("🛑 got %v, want %v", result, tt.want)
			// }
		})
	}

	fmt.Println(time.Since(start))
}
