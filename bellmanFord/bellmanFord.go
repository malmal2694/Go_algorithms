package bellmanFord

import "fmt"

//Global variables
const UNLIMITED = 10000000

type NodeEdge struct {
	listNodes  []string
	in_degree  map[string]int
	graph      map[string][]string
	weightEdge map[[2]string]int // Weight for each edge
	shortest   map[string]int
	pred       map[string]string
}

func NewGraph(numNode uint) *NodeEdge {
	return &NodeEdge{
		listNodes:  make([]string, 0, numNode), // save list nodes
		in_degree:  make(map[string]int),       // save degree for any nodes
		graph:      make(map[string][]string),  // save graph
		weightEdge: make(map[[2]string]int),    //Save weight of edges
		shortest:   make(map[string]int),
		pred:       make(map[string]string),
	}
}
func (g *NodeEdge) AddNode(node string) {
	g.listNodes = append(g.listNodes, node) //add node to listNodes
	g.in_degree[node] = 0                   // set degree for any node to 0
}
func (g *NodeEdge) AddEdge(from, to string, weight int) {
	g.graph[from] = append(g.graph[from], to)  //Create edge between two nodes
	g.in_degree[to]++                          // Increase in_degree for vertices u
	g.weightEdge[[2]string{from, to}] = weight // Save weight of edge
}
func (g *NodeEdge) findVertexU() string {
	var vertexU string //save output graph sorted
	var lowestNum, numVertexU int = UNLIMITED, 0

	for i := 0; i < len(g.listNodes); i++ {
		if g.shortest[g.listNodes[i]] < lowestNum {
			vertexU = g.listNodes[i] // save vertex u(in_degree == 0) to end of linearOrder
			numVertexU = i
			lowestNum = g.shortest[g.listNodes[i]]
		}
	}
	g.listNodes = removeIndex(g.listNodes, numVertexU)
	return vertexU //return vertices u
}

/*func (g NodeEdge) TopoSort() []string {
	var linearOrder []string
	var vertexU string
	next := g.findVertexU() // make next consisting output findVertexU func

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
}*/
func removeIndex(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func (g *NodeEdge) BellmanFord(start string) {
	for i := 0; i < len(g.listNodes); i++ {
		g.shortest[g.listNodes[i]] = UNLIMITED //Set shortest to unlimited for each vertex v
		//g.pred[g.listNodes[i]] = nil
	}
	g.shortest[start] = 0 //Set shortest[start] = 0

	for i := 1; i < (len(g.listNodes) - 1); i++ { //for i(=0) to n-1
		for j := 0; j < len(g.graph); j++ {
			for k := 0; k < len(g.graph[g.listNodes[j]]); k++ { //For each edge(u, v):
				g.relax(g.listNodes[j], g.graph[g.listNodes[j]][k]) // relax(u, v)
			}
		}
	}
} //End of func BellmanFord
func (g *NodeEdge) relax(u, v string) {
	if g.shortest[u]+g.weightEdge[[2]string{u, v}] < g.shortest[v] { // If shortest[u] + weight(u,v) < shortest[v], do the following
		g.shortest[v] = g.shortest[u] + g.weightEdge[[2]string{u, v}] //Set shortest[v] to shortest[u] + weight(u,v)
		g.pred[v] = u                                                 //Set pred[v] to u
	}
} //End of func relax
func (g NodeEdge) PrintPred() string {
	return fmt.Sprintf("return: %v", g.pred)
}
