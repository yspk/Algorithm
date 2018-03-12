package graph

import (
	"testing"
	"flag"
	"strings"
	"fmt"
)

func TestNewGraphFromString(t *testing.T) {
	input := flag.String("input","AB5,BC4,CD8,DC8,DE6,AD5,CE2,EB3,AE7","Please Input the Data")
	data := strings.Split(*input,",")
	if graph,err := NewGraphFromString(data);err == nil {
		fmt.Println(graph.Len())
	}
}