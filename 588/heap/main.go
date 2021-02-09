package main

import (
	"container/heap"
	"strings"
)

// An node is something we manage in sortedNodes.
type node struct {
	content string
	isDir   bool
	index   int
	name    string
	nodes   sortedNodes
	hash    map[string]*node
}

func newNode(name string, isDir bool) *node {
	if !isDir {
		return &node{
			name:  name,
			isDir: isDir,
		}
	}

	return &node{
		name:  name,
		isDir: true,
		nodes: newSortedNodes(),
		hash:  map[string]*node{},
	}
}

func (n *node) Add(name string, isDir bool) *node {
	if child, exists := n.hash[name]; exists {
		return child
	}
	child := newNode(name, isDir)

	heap.Push(&n.nodes, child)
	n.hash[child.name] = child

	return child
}

func (n *node) Mknode(path []string, isDir bool) *node {
	child := n.Add(path[0], isDir)
	if len(path) > 1 {
		return child.Mknode(path[1:], isDir)
	}

	return child
}

func (n *node) Walk(path []string) *node {
	if !n.isDir {
		return n
	}

	if 0 == len(path) {
		return n
	}

	next := path[0]
	nn, ok := n.hash[next]
	if !ok {
		return nil
	}

	return nn.Walk(path[1:])
}

func newSortedNodes() sortedNodes {
	var nodes sortedNodes
	heap.Init(&nodes)
	return nodes
}

// A sortedNodes implements heap.Interface and holds nodes
type sortedNodes []*node

func (sn sortedNodes) Len() int {
	return len(sn)
}

func (sn sortedNodes) Less(i, j int) bool {
	return sn[i].name < sn[j].name
}

func (sn sortedNodes) Swap(i, j int) {
	sn[i], sn[j] = sn[j], sn[i]
	sn[i].index = i
	sn[j].index = j
}

func (sn *sortedNodes) Push(x interface{}) {
	l := len(*sn)
	n := x.(*node)
	n.index = l
	*sn = append(*sn, n)
}

func (sn *sortedNodes) Pop() interface{} {
	old := *sn
	l := len(old)
	n := old[l-1]
	old[l-1] = nil // avoid memory leak
	n.index = -1   // for safety
	*sn = old[0 : l-1]
	return n
}

func (sn *sortedNodes) Ls() []*node {
	l := sn.Len()
	ls := make([]*node, l)

	i := 0
	for sn.Len() > 0 {
		n := heap.Pop(sn).(*node)
		ls[i] = n
		i++
	}

	*sn = ls

	return ls
}

type FileSystem struct {
	mem *node
}

func Constructor() FileSystem {
	return FileSystem{
		mem: &node{
			name:  "/",
			isDir: true,
			nodes: newSortedNodes(),
			hash:  map[string]*node{},
		},
	}
}

func (fs *FileSystem) Ls(path string) []string {
	l := fs.mem.nodes.Len()
	if 0 == l {
		return []string{}
	}

	var tail *node
	if "/" == path {
		tail = fs.mem
	} else {
		path = strings.TrimPrefix(path, "/")
		list := strings.Split(path, "/")
		tail = fs.mem.Walk(list)
	}

	if !tail.isDir {
		return []string{
			tail.name,
		}
	}

	l = tail.nodes.Len()
	contents := make([]string, l)
	nodes := tail.nodes.Ls()

	for i := range nodes {
		contents[i] = nodes[i].name
	}

	return contents
}

func (fs *FileSystem) Mkdir(path string) {
	path = strings.TrimPrefix(path, "/")
	fs.mem.Mknode(strings.Split(path, "/"), true)
}

func (fs *FileSystem) AddContentToFile(filePath string, content string) {
	filePath = strings.TrimPrefix(filePath, "/")
	list := strings.Split(filePath, "/")
	file := fs.mem.Walk(list)
	if nil == file {
		file = fs.mem.Mknode(list, false)
	}
	file.content += content
}

func (fs *FileSystem) ReadContentFromFile(filePath string) string {
	filePath = strings.TrimPrefix(filePath, "/")
	list := strings.Split(filePath, "/")
	file := fs.mem.Walk(list)
	return file.content
}
