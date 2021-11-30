package trees

import "testing"

func TestGraph(t *testing.T) {
	e := &GraphNode{data: 5, adjacencies: nil}
	d := &GraphNode{data: 4, adjacencies: nil}
	c := &GraphNode{data: 3, adjacencies: []*GraphNode{d, e}}
	b := &GraphNode{data: 2, adjacencies: nil}
	a := &GraphNode{data: 1, adjacencies: []*GraphNode{b, c}}

	G := &Graph{nodes: []*GraphNode{a, b, c, d, e}}

	bfs := G.breadthFirstSearch(4)
	if bfs == nil || bfs.data != 4 {
		t.Error("failed bfs")
	}

	dfs := G.depthFirstSearch(4)
	if dfs == nil || bfs.data != 4 {
		t.Error("failed dfs")
	}
}
