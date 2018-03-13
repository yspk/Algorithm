package graph

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Edge struct {
	tail   string
	head   string
	weight int
}

type Vertex struct {
	id    string
	dist  int
	routes [][]string
	arcs  map[string]int // arcs[vertex id] = weight
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

func (g *Graph) Len() int       { return len(g.vertices) }
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

		if val, ok := s[tail]; ok {
			val.arcs[head] = weight
		} else {
			arcs := make(map[string]int)
			arcs[head] = weight
			s[tail] = Vertex{id: tail, arcs: arcs, dist: 0}
		}

	}
	fmt.Println(s)
	return NewGraph(s), nil
}

func (g *Graph) CalcRouteDistance(route string) (int, error) {
	var dist int
	var s Vertex
	dails := strings.Split(route, "-")
	l := len(dails)
	if l < 2 {
		err := errors.New("NO SUCH ROUTE")
		return 0, err
	}

	for k, v := range dails {
		if k == 0 {
			s = g.vertices[v]
		} else {

			if val, ok := s.arcs[v]; ok {
				dist += val
			} else {
				err := errors.New("NO SUCH ROUTE")
				return 0, err
			}
			s = g.vertices[v]
		}

	}
	return dist, nil
}

func (g *Graph) BFSTraverse(src, dest string, limit int,exact bool) int  {
	var v,c Vertex
	var count int

	v = g.vertices[src]
	h := make(Queue, len(v.arcs))
	for id, _ := range v.arcs {
		v = g.vertices[id]
		//v.dist = y
		g.vertices[id] = v
		var route []string
		route = append(route,src)
		v.routes = append(v.routes, route)
		h.Push(v)
	}

	for !h.IsEmpty() {
		v = h.Pop()
		src = v.id
		for w, _ := range v.arcs {
			c = g.vertices[w]
			var l int
			//var rep bool
			for _,s := range v.routes {
				var route []string
				route = append(route,s...)
				route = append(route,src)
				l = len(route)
				fmt.Println(append(route,w))
				c.routes = append(c.routes,route)
			}

			if l < limit {
				h.Push(c)
			}
			//if w == dest {
			//	route := append(c.route,w)
			//	fmt.Println(route)
			//	if exact {
			//		if len(route) == limit {
			//			count ++
			//		}
			//	}else {
			//		count ++
			//	}
			//}
		}

	}
	return count
}

func (g *Graph) ShortestPath(src, dest string) (x int) {
	g.visit(src)
	v := g.vertices[src]
	h := make(Queue, len(v.arcs))
	// initialize the heap with out edges from src
	for id, y := range v.arcs {
		v := g.vertices[id]
		// update the vertices being pointed to with the distance.
		v.dist = y
		g.vertices[id] = v
		h.Push(v)
	}
	for src != dest {
		if h.IsEmpty() {
			return 1000000
		}
		v = h.Pop()
		src = v.id
		if g.visited[src] {
			continue
		}
		g.visit(src)
		for w, d := range v.arcs {
			if g.visited[w] {
				continue
			}
			c := g.vertices[w]
			distance := d + v.dist
			if distance < c.dist {
				c.dist = distance
				g.vertices[w] = c
			}
			h.Push(c)
		}
	}
	v = g.vertices[dest]
	return v.dist
}
