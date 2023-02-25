package main

import (
	"fmt"
	"os"
)

type Graph struct {
	ants         int
	start        string
	end          string
	pathStartEnd bool
	current      *Vertex
	vertices     []*Vertex
	path         []string
	allPaths     [][]string
}

type Vertex struct {
	visited  bool
	name     string
	adjacent []*Vertex
}

type SortPaths struct {
	counter       int
	from          int
	pathsWithAnts []int
	sliceFrom     []int
	searchSlice   [][]string
	search        [][][]string
	result        [][][]string
}

func (g *Graph) AddVertex(key string) {
	if key == g.end {
		g.vertices = append(g.vertices, &Vertex{name: key, visited: true})
	} else if contains(g.vertices, key) {
		// err := fmt.Errorf("Vertex %v not added because it is an existing key", key)
		// fmt.Println(err.Error())
	} else {
		g.vertices = append(g.vertices, &Vertex{name: key, visited: false})
	}
}

func contains(vertexList []*Vertex, key string) bool {
	for _, vertex := range vertexList {
		if key == vertex.name {
			return true
		}
	}
	return false
}

func (g *Graph) getVertex(key string) *Vertex {
	for i, vertex := range g.vertices {
		if vertex.name == key {
			return g.vertices[i]
		}
	}
	return nil
}

func (g *Graph) AddEdge(from, to string) {
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)

	if fromVertex == nil || toVertex == nil {
		// err := fmt.Errorf("invalid edge (%v-->%v)", from, to)
		// fmt.Println(err.Error())
	} else if contains(fromVertex.adjacent, to) {
		// err := fmt.Errorf("existing edge (%v-->%v)", from, to)
		// fmt.Println(err.Error())
	} else {
		fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
	}
}

func connectedGraph(vertices []*Vertex) bool {
	for _, vertex := range vertices {
		if len(vertex.adjacent) == 0 {
			fmt.Println("ERROR: invalid data format, disconnected graph")
			os.Exit(1)
			return false
		}
	}
	return true
}

func (g *Graph) Print() {
	connectedGraph(g.vertices)
	g.current = g.getVertex(g.start)
	g.path = append(g.path, g.start)
}
