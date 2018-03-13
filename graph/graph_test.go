package graph

import (
	"flag"
	"fmt"
	"strings"
	"testing"
)

func TestGraph_NewGraphFromString(t *testing.T) {
	input := flag.String("input", "AB5,BC4,CD8,DC8,DE6,AD5,CE2,EB3,AE7", "Please Input the Data")
	data := strings.Split(*input, ",")
	if gra, err := NewGraphFromString(data); err == nil {
		fmt.Println(gra.Len())
	}
}

func TestGraph_CalcDistance(t *testing.T) {
	data := []string{"AB5", "BC4", "CD8", "DC8", "DE6", "AD5", "CE2", "EB3", "AE7"}
	if gra, err := NewGraphFromString(data); err == nil {
		fmt.Println(gra.CalcRouteDistance([]string{"A", "D", "C"}))
	}
}

func TestGraph_BFSTraverse(t *testing.T) {
	data := []string{"AB5", "BC4", "CD8", "DC8", "DE6", "AD5", "CE2", "EB3", "AE7"}
	if gra, err := NewGraphFromString(data); err == nil {
		fmt.Println(gra.BFSTraverse("C", "C", 3, false))
	}
}

func TestGraph_ShortestPath(t *testing.T) {
	data := []string{"AB5", "BC4", "CD8", "DC8", "DE6", "AD5", "CE2", "EB3", "AE7"}
	if gra, err := NewGraphFromString(data); err == nil {
		//fmt.Println(gra.DifShortestPath("C","B"))
		fmt.Println(gra.SameShortestPath("B"))
	}
}

func TestGraph_BFSDistLimit(t *testing.T) {
	data := []string{"AB5", "BC4", "CD8", "DC8", "DE6", "AD5", "CE2", "EB3", "AE7"}
	if gra, err := NewGraphFromString(data); err == nil {
		fmt.Println(gra.BFSDistLimit("C", "C", 30))
	}
}
