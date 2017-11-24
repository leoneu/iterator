package iterator_test

import (
	"io"
	"testing"

	"github.com/leoneu/iterator"
)

const N = 100000

func BenchmarkIterator(b *testing.B) {

	g := iterator.NewGraph(N)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		iter := g.Iterator()
		for j := 0; ; j++ {
			node, err := iter.Next()
			if err == io.EOF {
				break
			}
			// do something
			node.Value = node.Value + 1

			if j != int(node.ID()) {
				b.Fatalf("want %d, got %d", j, node.ID())
			}
		}
	}
}

func BenchmarkSlice(b *testing.B) {

	g := iterator.NewGraph(N)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		nodes := g.NodeSlice()
		for j, node := range nodes {

			// do something
			node.Value = node.Value + 1

			if j != int(node.ID()) {
				b.Fatalf("want %d, got %d", j, node.ID())
			}
		}
	}
}
