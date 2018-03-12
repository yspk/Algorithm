package graph

import (
	"testing"
	"flag"
	"strings"
	"fmt"
)

func TestGraph_NewGraphFromString(t *testing.T) {
	input := flag.String("input","AB5,BC4,CD8,DC8,DE6,AD5,CE2,EB3,AE7","Please Input the Data")
	data := strings.Split(*input,",")
	if gra,err := NewGraphFromString(data);err == nil {
		fmt.Println(gra.Len())
	}
}

func TestGraph_CalcDistance(t *testing.T) {
	data := []string{"AB5","BC4","CD8","DC8","DE6","AD5","CE2","EB3","AE7"}
	if gra,err := NewGraphFromString(data);err == nil {
		fmt.Println(gra.CalcDistance("A-E-D"))
	}
}

