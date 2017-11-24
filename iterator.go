package iterator

import (
	"io"
)

type Graph struct {
	nodes []Node
}

type Node struct {
	id    int64
	Value int64
}

func (node Node) ID() int64 {
	return node.id
}

func NewGraph(n int64) *Graph {
	nodes := make([]Node, n, n)
	for i := range nodes {
		nodes[i] = Node{id: int64(i), Value: int64(i)}
	}
	return &Graph{nodes: nodes}
}

type Iterator struct {
	index int
	g     *Graph
}

func (g *Graph) NodeSlice() []Node {
	return g.nodes
}

func (g *Graph) Iterator() *Iterator {
	return &Iterator{g: g}
}

func (it *Iterator) Next() (*Node, error) {

	if it.index >= len(it.g.nodes) {
		return nil, io.EOF
	}

	node := it.g.nodes[it.index]
	it.index++
	return &node, nil
}
