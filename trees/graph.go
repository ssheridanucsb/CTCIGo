package trees

import "errors"

type GraphNode struct {
	data        int
	adjacencies []*GraphNode
}

type Graph struct {
	nodes []*GraphNode
}

func dfs(root *GraphNode, key int, vm map[*GraphNode]int) *GraphNode {
	if root == nil {
		return nil
	}
	if root.data == key {
		return root
	}
	vm[root]++
	for _, n := range root.adjacencies {
		_, ok := vm[n]
		if !ok {
			resp := dfs(n, key, vm)
			if resp != nil {
				return resp
			}
		}
	}
	return nil
}

func (g *Graph) depthFirstSearch(key int) *GraphNode {
	visited := make(map[*GraphNode]int)
	//this only works if the graph is connected
	n := dfs(g.nodes[0], key, visited)
	return n
}

func (g *Graph) breadthFirstSearch(key int) *GraphNode {
	q := NewNodeQueue()
	vm := make(map[*GraphNode]int)

	root := g.nodes[0]
	vm[root]++
	q.add(root)

	for !q.isEmpty() {
		r, _ := q.remove()
		if r.data == key {
			return r
		}

		for _, n := range r.adjacencies {
			_, ok := vm[n]
			if !ok {
				vm[n]++
				q.add(n)
			}
		}
	}
	return nil
}

// Queue for graph nodes
type NodeQNode struct {
	graphNode *GraphNode
	next      *NodeQNode
}
type NodeQ struct {
	first *NodeQNode
	last  *NodeQNode
}

func NewNodeQueue() *NodeQ {
	return &NodeQ{first: nil, last: nil}
}

func (q *NodeQ) add(gn *GraphNode) {
	t := &NodeQNode{graphNode: gn, next: nil}
	if q.last != nil {
		q.last.next = t
	}
	q.last = t
	if q.first == nil {
		q.first = q.last
	}

}

func (q *NodeQ) remove() (*GraphNode, error) {
	if q.first == nil {
		return nil, errors.New("empty queue")
	}

	d := q.first.graphNode
	q.first = q.first.next
	if q.first == nil {
		q.last = nil
	}
	return d, nil
}

func (q *NodeQ) isEmpty() bool {
	return q.first == nil
}
