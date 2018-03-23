package main

import "fmt"

import "dagShortestPaths"

func main() {
	a := dagShortestPaths.NewGraph(6)

	a.AddNode("r")
	a.AddNode("s")
	a.AddNode("t")
	a.AddNode("x")
	a.AddNode("y")
	a.AddNode("z")

	a.AddEdge("r", "s", 5)
	a.AddEdge("r", "t", 3)
	a.AddEdge("s", "t", 2)
	a.AddEdge("s", "x", 6)
	a.AddEdge("t", "x", 7)
	a.AddEdge("t", "y", 4)
	a.AddEdge("t", "z", 2)
	a.AddEdge("x", "y", -1)
	a.AddEdge("x", "z", 1)
	a.AddEdge("y", "z", -2)

	a.DagShortestPaths("s")

	fmt.Printf("%s", a.PrintPred())
}
