package topoSort

type NodeEdge struct {
	listNodes []string
	in_degree map[string]int
	graph     map[string][]string
}

func NewGraph(numNode uint) *NodeEdge {
	return &NodeEdge{
		listNodes: make([]string, 0, numNode), // save list nodes
		in_degree: make(map[string]int),       // save degree for any nodes
		graph:     make(map[string][]string),  // save graph
	}
}
func (g *NodeEdge) AddNode(node string) {
	g.listNodes = append(g.listNodes, node) //add node to listNodes
	g.in_degree[node] = 0                   // set degree for any node to 0
}
func (g *NodeEdge) AddEdge(from, to string) {
	g.graph[from] = append(g.graph[from], to) //create edge between two nodes
	g.in_degree[to]++                         // increase in_degree for vertices u
}
func (g *NodeEdge) findVerticesU() []string {
	var listVerticesU []string //save output graph sorted

	for i := 0; i < len(g.listNodes); i++ {
		if g.in_degree[g.listNodes[i]] == 0 {
			listVerticesU = append(listVerticesU, g.listNodes[i]) // save vertex u(in_degree == 0) to end of linearOrder
		}
	}
	return listVerticesU //return vertices u
}
func (g NodeEdge) TopoSort() []string {
	var linearOrder []string
	var vertexU string
	next := g.findVerticesU() // make next consisting output findVerticesU func

	for len(next) > 0 {
		vertexU = next[0]
		next = removeIndex(next, 0)
		linearOrder = append(linearOrder, vertexU) //append vertexU to end of linearOrder
		for j := 0; j < len(g.graph[vertexU]); j++ {
			g.in_degree[g.graph[vertexU][j]]--         //decrement in_degree vertexV(vertices adjacent vertex u)
			if g.in_degree[g.graph[vertexU][j]] == 0 { //if in_degree[v](v: a vertex adjacent vertex u) is equal to 0, do the following
				next = append(next, g.graph[vertexU][j]) //ad vertex v to end of next
			}
		}
	}
	return linearOrder //return graph sorted
}
func removeIndex(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}
