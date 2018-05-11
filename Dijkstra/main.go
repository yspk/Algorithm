package main

import (
	"./graph"
	"flag"
	"fmt"
	"strings"
)

func main() {
	input := flag.String("input", "AB5,BC4,CD8,DC8,DE6,AD5,CE2,EB3,AE7", "Please Input the Data")
	data := strings.Split(*input, ",")
	//Data Format
	var gra *graph.Graph
	var err error
	var dist int
	var route []string
	if gra, err = graph.NewGraphFromString(data); err != nil {
		fmt.Println(err)
	}
	// 1.	The distance of the route A-B-C.
	route = strings.Split("A-B-C", "-")
	if dist, err = gra.CalcRouteDistance(route); err != nil {
		fmt.Println("Output #1: ", err)
	} else {
		fmt.Println("Output #1: ", dist)
	}

	// 2.	The distance of the route A-D.
	route = strings.Split("A-D", "-")
	if dist, err = gra.CalcRouteDistance(route); err != nil {
		fmt.Println("Output #2: ", err)
	} else {
		fmt.Println("Output #2: ", dist)
	}

	// 3.	The distance of the route A-D-C.
	route = strings.Split("A-D-C", "-")
	if dist, err = gra.CalcRouteDistance(route); err != nil {
		fmt.Println("Output #3: ", err)
	} else {
		fmt.Println("Output #3: ", dist)
	}

	// 4.	The distance of the route A-E-B-C-D.
	route = strings.Split("A-E-B-C-D", "-")
	if dist, err = gra.CalcRouteDistance(route); err != nil {
		fmt.Println("Output #4: ", err)
	} else {
		fmt.Println("Output #4: ", dist)
	}

	// 5.	The distance of the route A-E-D.
	route = strings.Split("A-E-D", "-")
	if dist, err = gra.CalcRouteDistance(route); err != nil {
		fmt.Println("Output #5: ", err)
	} else {
		fmt.Println("Output #5: ", dist)
	}

	// 6.	The number of trips starting at C and ending at C with a maximum of 3 stops.  In the sample data below, there are two such trips: C-D-C (2 stops). and C-E-B-C (3 stops).
	fmt.Println("Output #6: ", gra.BFSTraverse("C", "C", 3, false))

	// 7.	The number of trips starting at A and ending at C with exactly 4 stops.  In the sample data below, there are three such trips: A to C (via B,C,D); A to C (via D,C,D); and A to C (via D,E,B).
	if gra1, err := graph.NewGraphFromString(data); err == nil {
		fmt.Println("Output #7: ", gra1.BFSTraverse("A", "C", 4, true))
	}

	// 8.	The length of the shortest route (in terms of distance to travel) from A to C.
	fmt.Println("Output #8: ", gra.DifShortestPath("A", "C"))

	// 9.	The length of the shortest route (in terms of distance to travel) from B to B.
	if gra2, err := graph.NewGraphFromString(data); err == nil {
		fmt.Println("Output #9: ", gra2.SameShortestPath("B"))
	}

	// 10.	The number of different routes from C to C with a distance of less than 30.  In the sample data, the trips are: CDC, CEBC, CEBCDC, CDCEBC, CDEBC, CEBCEBC, CEBCEBCEBC.
	if gra3, err := graph.NewGraphFromString(data); err == nil {
		fmt.Println("Output #10: ", gra3.BFSDistLimit("C", "C", 70))
	}
}
