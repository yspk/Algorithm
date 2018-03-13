package graph

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

type Edge struct {
	tail   string
	head   string
	weight int
}

type Vertex struct {
	id     string
	dist   int
	routes [][]string
	arcs   map[string]int // arcs[vertex id] = weight
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
	//fmt.Println(s)
	return NewGraph(s), nil
}

func (g *Graph) CalcRouteDistance(route []string) (int, error) {
	var dist int
	var s Vertex
	l := len(route)
	if l < 2 {
		err := errors.New("NO SUCH ROUTE")
		return 0, err
	}

	for k, v := range route {
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

func (g *Graph) BFSTraverse(src, dest string, limit int, exact bool) int {
	var v, c Vertex
	var count int
	v = g.vertices[src]
	h := make(Queue, len(v.arcs))
	for id, _ := range v.arcs {
		v = g.vertices[id]
		var route []string
		route = append(route, src)
		v.routes = append(v.routes, route)
		g.vertices[id] = v
		h.Push(v)
	}

	for !h.IsEmpty() {
		v = h.Pop()
		src = v.id
		for w, _ := range v.arcs {
			c = g.vertices[w]
			var pushed bool
			for _, s := range v.routes {
				var route []string
				var rep bool
				route = append(route, s...)
				route = append(route, src)
				for _, m := range c.routes {
					if reflect.DeepEqual(m, route) {
						rep = true
					}
				}
				if !rep {
					c.routes = append(c.routes, route)
					l := len(route)
					if l < limit {
						pushed = true
					}
					if w == dest {
						if exact {
							if len(route) == limit {
								count++
								//fmt.Println(append(route, w))
							}
						} else {
							count++
							//fmt.Println(append(route, w))
						}
					}
				}
			}
			g.vertices[w] = c
			if pushed {
				h.Push(c)
			}
		}

	}
	return count
}

func (g *Graph) DifShortestPath(src, dest string) int {
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

func (g *Graph) SameShortestPath(src string) int {
	v := g.vertices[src]
	var minimum int
	var dists []int
	for id, y := range v.arcs {
		// update the vertices being pointed to with the distance.
		if id == src {
			minimum = y
		}
		dist := g.DifShortestPath(id, src)
		//fmt.Println(id, src, dist)
		dists = append(dists, y+dist)
	}
	l := len(dists)
	if l == 0 {
		//fmt.Println("NO SUCH ROUTE")
	} else if l == 1 {
		if minimum > dists[0] || minimum == 0 {
			minimum = dists[0]
		}
	} else {
		minimum = dists[0]
		for _, v := range dists[1:] {
			if v < minimum {
				minimum = v
			}
		}
	}
	return minimum
}

func (g *Graph) BFSDistLimit(src, dest string, dist int) int {
	var v, c Vertex
	var count int
	v = g.vertices[src]
	h := make(Queue, len(v.arcs))
	for id, _ := range v.arcs {
		v = g.vertices[id]
		var route []string
		route = append(route, src)
		v.routes = append(v.routes, route)
		g.vertices[id] = v
		h.Push(v)
	}

	for !h.IsEmpty() {
		v = h.Pop()
		src = v.id
		for w, _ := range v.arcs {
			c = g.vertices[w]
			var pushed bool
			for _, s := range v.routes {
				var route []string
				var rep bool
				route = append(route, s...)
				route = append(route, src)
				for _, m := range c.routes {
					if reflect.DeepEqual(m, route) {
						rep = true
					}
				}
				if !rep {
					c.routes = append(c.routes, route)
					//l := len(route)
					//if l < limit {
					//	pushed = true
					//}
					route = append(route, w)
					d, _ := g.CalcRouteDistance(route)
					if d < dist {
						pushed = true
					}
					if w == dest {
						d, _ := g.CalcRouteDistance(route)
						if d < dist {
							count++
							//fmt.Println(route)
						}
					}
				}
			}
			g.vertices[w] = c
			if pushed {
				h.Push(c)
			}
		}

	}
	return count
}
