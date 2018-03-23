package main

import "fmt"

import "topoSort"

func main() {
	a := topoSort.NewGraph(9)

	a.AddNode("1")
	a.AddNode("2")
	a.AddNode("3")
	a.AddNode("4")
	a.AddNode("hello")
	a.AddNode("hi")
	a.AddNode("how")
	a.AddNode("sure")
	a.AddNode("but")

	a.AddEdge("1", "4")
	a.AddEdge("1", "3")
	a.AddEdge("1", "hello")
	a.AddEdge("2", "hi")
	a.AddEdge("hi", "how")
	a.AddEdge("3", "how")
	a.AddEdge("how", "sure")
	a.AddEdge("hi", "sure")

	fmt.Println(a.TopoSort())
}
