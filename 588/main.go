package main

// An node is something we manage in a priority queue.
type node struct {
	name    string // The name of the node; arbitrary.
	priority int    // The priority of the node in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the node in the heap.
}

// A SortedNodes implements heap.Interface and holds Items.
type SortedNodes []*node

func (pq SortedNodes) Len() int { 
	return len(pq)
}

func (pq SortedNodes) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq SortedNodes) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *SortedNodes) Push(x interface{}) {
	n := len(*pq)
	n := x.(*node)
	n.index = n
	*pq = append(*pq, n)
}

func (pq *SortedNodes) Pop() interface{} {
	old := *pq
	n := len(old)
	n := old[n-1]
	old[n-1] = nil  // avoid memory leak
	n.index = -1 // for safety
	*pq = old[0 : n-1]
	return n
}

// update modifies the priority and name of an node in the queue.
func (pq *SortedNodes) update(n *node, name string, priority int) {
	n.name = name
	n.priority = priority
	heap.Fix(pq, n.index)
}



func doIt() {
	fmt.Println("Doing It")

	// items := map[string]int{
	// 	"banana": 3, "apple": 2, "pear": 4,
	// }

	// Create a priority queue, put the items in it, and
	// establish the priority queue (heap) invariants.
	// pq := make(PriorityQueue, len(items))
	// i := 0
	// for value, priority := range items {
	// 	pq[i] = &Item{
	// 		value:    value,
	// 		priority: priority,
	// 		index:    i,
	// 	}
	// 	i++
	// }
	// heap.Init(&pq)

	// // Insert a new item and then modify its priority.
	// item := &Item{
	// 	value:    "orange",
	// 	priority: 1,
	// }
	// heap.Push(&pq, item)
	// pq.update(item, item.value, 5)

	// // Take the items out; they arrive in decreasing priority order.
	// for pq.Len() > 0 {
	// 	item := heap.Pop(&pq).(*Item)
	// 	fmt.Printf("%.2d:%s ", item.priority, item.value)
	// }
}


type node struct {
	content string
	isDir bool
	name string
	nodes map[string]node
}

type FileSystem struct {
	mem node
}

func Constructor() FileSystem {
	return FileSystem{
		mem: node{
			isDir: true,
		},
	}
}

func (fs *FileSystem) Ls(path string) []string {

	if n, ok := fs.mem.nodes[path]; !ok {
		return []string{}
	} else {
		if !n.isDir {
			return []string{n.name}
		}
			var ls []string

			for i := range n.nodes {
				nn := n.nodes[i]
				ls = append(ls, )
			}

	return ls
}

func (fs *FileSystem) Mkdir(path string) {

}

func (fs *FileSystem) AddContentToFile(filePath string, content string) {

}

func (fs *FileSystem) ReadContentFromFile(filePath string) string {
	return ""
}
