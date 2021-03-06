package main

import "fmt"

// Node is a single node in a graph list
type Node struct {
	name  string
	value int
}

type graph map[string][]Node

func (g graph) addVertexToGraph(vtx string) {
	if g[vtx] != nil {
		fmt.Println("\n-- Vertex already exists. --")
		return
	}
	g[vtx] = make([]Node, 0)
}

func (g graph) addEdgeToGraph(fromVtx, toVtx string, edgeValue int) {
	if g[fromVtx] == nil { // check if initial vertex exists
		fmt.Println("\n-- Initial vertex " + fromVtx + " does not exist. --")
		return
	}
	for i := range g[fromVtx] { // check if edge already exists
		if g[fromVtx][i].name == toVtx {
			fmt.Println("\n-- Edge from " + fromVtx + " to " + toVtx + " already exists. --")
			return
		}
	}
	if g[toVtx] == nil { // create new destination vertext if it does not exists
		g[toVtx] = make([]Node, 0)
		fmt.Println("\n-- Destination vertex " + toVtx + " created. --")
	}

	g[fromVtx] = append(g[fromVtx], Node{name: toVtx, value: edgeValue})
	return
}

func (g graph) removeVertexFromGraph(vtx string) {
	length := len(g[vtx]) - 1
	for length != -1 { // loop in reverse in order to avoid index out of range
		g.removeEdgeFromGraph(vtx, g[vtx][length].name)
		length--
	}

	for i := range g {
		if i != vtx {
			length := len(g[i]) - 1
			for length != -1 { // loop in reverse in order to avoid index out of range
				if g[i][length].name == vtx {
					g.removeEdgeFromGraph(i, vtx)
				}
				length--
			}
		}
	}

	delete(g, vtx)
}

func (g graph) removeEdgeFromGraph(fromVtx, toVtx string) {
	if g[fromVtx] == nil {
		fmt.Println("\n-- Vertex " + fromVtx + " does not exists. --")
		return
	} else if g[toVtx] == nil {
		fmt.Println("\n-- Vertex " + toVtx + " does not exists. --")
		return
	}

	for i := range g[fromVtx] {
		if g[fromVtx][i].name == toVtx {
			if i == 0 {
				g[fromVtx] = g[fromVtx][1:len(g[fromVtx])]
			} else if i == (len(g[fromVtx]) - 1) {
				g[fromVtx] = g[fromVtx][0:(len(g[fromVtx]) - 1)]
			} else {
				initial := g[fromVtx][0:i]
				final := g[fromVtx][i+1 : len(g[fromVtx])]
				g[fromVtx] = append(initial, final...)
			}
			break
		}
	}
}

func main() {
	i := 0
	g := make(graph)
	for i == 0 {
		fmt.Println("\n1. ADD A VERTEX")
		fmt.Println("2. ADD AN EDGE")
		fmt.Println("3. REMOVE VERTEX")
		fmt.Println("4. REMOVE AN EDGE")
		fmt.Println("5. DISPLAY USING DFS")
		fmt.Println("6. DISPLAY USING BFS")
		fmt.Println("7. SIMPLE DISPLAY")
		fmt.Println("8. EXIT")
		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scanf("%d", &choice)
		switch choice {
		case 1:
			g.addVertex()
		case 2:
			g.addEdge()
		case 3:
			g.removeVertex()
		case 4:
			g.removeEdge()
		case 5:
			result, weights := g.displayDFS()
			fmt.Println(result)
			fmt.Println(weights)
		case 6:
			result, weights := g.displayBFS()
			fmt.Println(result)
			fmt.Println(weights)
		case 7:
			g.simpleDisplay()
		case 8:
			i = 1
		default:
			fmt.Println("Command not recognized.")
		}
	}
}

func (g graph) addVertex() {
	var vtxName string
	fmt.Print("Enter the name of vertex: ")
	fmt.Scanf("%s", &vtxName)
	g.addVertexToGraph(vtxName)
}

func (g graph) addEdge() {
	var fromVtx, toVtx string
	var edgeValue int
	fmt.Print("Enter the initial vertex name: ")
	fmt.Scanf("%s", &fromVtx)
	fmt.Print("Enter the destination vertex name: ")
	fmt.Scanf("%s", &toVtx)
	fmt.Print("Enter the weight of edge: ")
	fmt.Scanf("%d", &edgeValue)
	g.addEdgeToGraph(fromVtx, toVtx, edgeValue)
}

func (g graph) removeVertex() {
	var vtxName string
	fmt.Print("Enter the name of vertex: ")
	fmt.Scanf("%s", &vtxName)
	g.removeVertexFromGraph(vtxName)
}

func (g graph) removeEdge() {
	var fromVtx, toVtx string
	fmt.Print("Enter the initial vertex name: ")
	fmt.Scanf("%s", &fromVtx)
	fmt.Print("Enter the destination vertex name: ")
	fmt.Scanf("%s", &toVtx)
	g.removeEdgeFromGraph(fromVtx, toVtx)
}

func (g graph) displayDFS() ([]string, []int) {
	result, weights := g.DFS()
	return result, weights
}

func (g graph) displayBFS() ([]string, []int) {
	result, weights := g.BFS()
	return result, weights
}

func (g graph) simpleDisplay() {
	fmt.Println("")
	for i := range g {
		fmt.Print(i, " => ")
		for j := range g[i] {
			fmt.Print(g[i][j])
		}
		fmt.Println("")
	}
}

// addVertexToGraph("a")
// addVertexToGraph("b")
// addVertexToGraph("c")
// addVertexToGraph("d")
// addVertexToGraph("e")
// addVertexToGraph("f")
// addEdgeToGraph("a", "b", 18)
// addEdgeToGraph("a", "c", 20)
// addEdgeToGraph("b", "d", 4)
// addEdgeToGraph("c", "e", 10)
// addEdgeToGraph("d", "e", 15)
// addEdgeToGraph("d", "f", 23)
// addEdgeToGraph("e", "f", 7)
// addEdgeToGraph("d", "a", 5)
// addEdgeToGraph("e", "a", 17)
// addEdgeToGraph("b", "e", 21)
