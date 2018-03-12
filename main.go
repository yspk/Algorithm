package main

import (
	"flag"
	"fmt"
	"strings"
	"./graph"
)

func main()  {
	//Ask
	input := flag.String("input","AB5,BC4,CD8,DC8,DE6,AD5,CE2,EB3,AE7","Please Input the Data")
	data := strings.Split(*input,",")
	if gra,err := graph.NewGraphFromString(data);err == nil {
		fmt.Println(gra.Len())
	}

	//Compute

	//Answer
}

func DataFormat()  {

	//Data Format"
}

func Compute()  {
	// Distance

	// Step Limit

	// Shortest Route

	// Distance Limit
}

func Answer()  {
	// Println
}