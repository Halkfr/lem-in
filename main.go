package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	if len(os.Args) == 1 {
		fmt.Println("ERROR: invalid data format, missing file txt")
		return
	} else if len(os.Args) > 2 {
		fmt.Println("ERROR: invalid data format, too many arguments")
		return
	}

	s := readFile()
	test := &Graph{}

	sort := &SortPaths{
		from:    0,
		counter: 1,
	}

	if antNum(s, test) <= 0 || antNum(s, test) > 100000 { 
		fmt.Println("ERROR: invalid data format, invalid number for ants")
		return
	}

	if startRoom(s, test) == "" {
		fmt.Println("ERROR: invalid data format, ##start room not found")
		return
	}

	if endRoom(s, test) == "" {
		fmt.Println("ERROR: invalid data format, ##end room not found")
		return
	}

	for _, v := range allRooms(s) {
		test.AddVertex(v)
	}

	for _, v := range theLinks(s) {
		vertex := strings.Split(v, ", ")
		test.AddEdge(vertex[0], vertex[1])
		test.AddEdge(vertex[1], vertex[0])
	}

	test.Print()
	findAllPaths(test)
	sortPaths(test, sort)
	bestPathCombinations(test, sort)
	desidePath(test, sort)
	createAnts(test, sort)

	fmt.Printf("Program executed in %v", time.Since(start))
}
