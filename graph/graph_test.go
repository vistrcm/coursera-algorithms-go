package graph_test

import (
	"fmt"
	"github.com/vistrcm/coursera-algorithms-go/graph"
)

func printGraph(g graph.Interface) {
	for v := 0; v < g.V(); v++ {
		for _, w := range g.Adj(v) {
			fmt.Printf("%d -> %d\n", v, w)
		}
	}
}

func ExampleCC_print() {
	g := setupGraph()
	printGraph(&g)

	// Unordered output: 0 -> 5
	// 4 -> 3
	// 0 -> 1
	// 9 -> 12
	// 6 -> 4
	// 5 -> 4
	// 0 -> 2
	// 11 -> 12
	// 9 -> 10
	// 0 -> 6
	// 7 -> 8
	// 9 -> 11
	// 5 -> 3
	// 5 -> 0
	// 3 -> 4
	// 1 -> 0
	// 12 -> 9
	// 4 -> 6
	// 4 -> 5
	// 2 -> 0
	// 12 -> 11
	// 10 -> 9
	// 6 -> 0
	// 8 -> 7
	// 11 -> 9
	// 3 -> 5

}
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
