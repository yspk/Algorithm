package graph

import (
	"fmt"
	"strconv"
)

type Edge struct {
	tail   string
	head   string
	weight int
}

type Vertex struct {
	id   string
	dist int
	arcs map[string]int // arcs[vertex id] = weight
}

type Graph struct {
	visited  map[string]bool
	vertices map[string]Vertex
}

func NewGraph(vs map[string]Vertex) *Graph {
	g := new(Graph)
	g.visited = make(map[string]bool)
	g.vertices = make(map[string]Vertex)
	for i, v := range vs {
		v.dist = 1000000
		g.vertices[i] = v
	}
	return g
}

func (g *Graph) Len() int    { return len(g.vertices) }
func (g *Graph) visit(v string) { g.visited[v] = true }

func NewGraphFromString(fn []string) (*Graph, error) {
	s := make(map[string]Vertex)
	for _, v := range fn {
		tail := v[0:1]
		head := v[1:2]
		t := v[2:]
		weight, err := strconv.Atoi(t)
		if err != nil {
			fmt.Println("Data err:", err)
			return nil, err
		}

		if val , ok := s[tail]; ok {
			val.arcs[head] = weight
		} else {
			arcs := make(map[string]int)
			arcs[head] = weight
			s[tail] = Vertex{id: tail, arcs: arcs, dist: 0}
		}

	}
	fmt.Println(s)
	return NewGraph(s),nil
}
