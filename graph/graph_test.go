package graph_test

import (
	"fmt"
	"github.com/vistrcm/coursera-algorithms-pi-go/graph"
)

func ExampleCC() {
	g := setupGraph()

	cc := graph.NewCC(&g)
	fmt.Printf("cc.Count() = %d\n", cc.Count())
	fmt.Printf("cc.Id(0) = %d\n", cc.Id(0))
	fmt.Printf("cc.Id(9) = %d\n", cc.Id(9))
	// Output: cc.Count() = 3
	// cc.Id(0) = 0
	// cc.Id(9) = 2

}

func setupGraph() graph.Graph {
	g := graph.NewGraph(13)
	g.AddEdge(0, 5)
	g.AddEdge(4, 3)
	g.AddEdge(0, 1)
	g.AddEdge(9, 12)
	g.AddEdge(6, 4)
	g.AddEdge(5, 4)
	g.AddEdge(0, 2)
	g.AddEdge(11, 12)
	g.AddEdge(9, 10)
	g.AddEdge(0, 6)
	g.AddEdge(7, 8)
	g.AddEdge(9, 11)
	g.AddEdge(5, 3)
	return g
}
