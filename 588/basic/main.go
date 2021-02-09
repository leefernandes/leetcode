package main

import (
	"sort"
	"strings"
)

type node struct {
	content string
	isDir   bool
	name    string
	ls      []string
}

type FileSystem struct {
	mem map[string]*node
}

func Constructor() FileSystem {
	return FileSystem{
		mem: map[string]*node{
			"/": &node{
				name:  "/",
				isDir: true,
				ls:    []string{},
			},
		},
	}
}

func (fs *FileSystem) Ls(path string) []string {
	n := fs.mem[path]
	if n.isDir {
		sort.Strings(n.ls)
		return n.ls
	}

	return []string{n.name}
}

func (fs *FileSystem) Mkdir(path string) {
	fs.Mknode(path, true)
}

func (fs *FileSystem) Mknode(path string, isDir bool) {
	parts := strings.Split(path, "/")
	l := len(parts) - 1

	var lastNew string
	for i := l; i >= 0; i-- {
		name := parts[i]
		var key string = strings.Join(parts[0:i+1], "/")
		if "" == key {
			name = "/"
			key = "/"
		}

		if n, exists := fs.mem[key]; exists {
			if lastNew != "" {
				n.ls = append(n.ls, lastNew)
			}

			lastNew = ""

			continue
		}

		n := node{
			isDir: isDir,
			name:  name,
		}

		if isDir {
			n.ls = []string{}
		}

		if lastNew != "" {
			n.ls = append(n.ls, lastNew)
		}

		fs.mem[key] = &n

		lastNew = name
	}
}

func (fs *FileSystem) AddContentToFile(filePath string, content string) {
	file, exist := fs.mem[filePath]
	if !exist {
		fs.Mknode(filePath, false)
		file = fs.mem[filePath]
	}

	file.content += content
}

func (fs *FileSystem) ReadContentFromFile(filePath string) string {
	return fs.mem[filePath].content
}
