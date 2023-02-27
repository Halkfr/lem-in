package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func desidePath(g *Graph, s *SortPaths) {

	var bestValue int
	var min int
	var max int
	var count int
	var bestOrder int
	var sliceInt []int
	var sliceOfSlicesInt [][]int
	allAnts := g.ants

	if len(s.result) != 0 {
		for i := range s.result {
			sliceInt = nil
			for _, combination := range s.result[i] {
				sliceInt = append(sliceInt, len(combination))
			}
			sliceOfSlicesInt = append(sliceOfSlicesInt, sliceInt)
		}
		for _, numbers := range sliceOfSlicesInt {
			for ant := 0; ant < allAnts; {

				min = numbers[0]
				max = numbers[0]
				minOrder := 0

				for j, num := range numbers {
					if num < min {
						min = num
						minOrder = j
					}
				}
				numbers[minOrder]++
				ant++

				for _, num := range numbers {
					if num > max {
						max = num
					}
				}
			}
			count++
			if bestValue == 0 || bestValue > max {
				bestValue = max
				bestOrder = count - 1
			} else {
				break
			}
		}
		s.result = s.result[bestOrder : bestOrder+1]
		s.pathsWithAnts = sliceOfSlicesInt[bestOrder]

		// fmt.Printf("Best path for %v ants: %s. ", allAnts, s.result)
		// fmt.Println("Here is why:", sliceOfSlicesInt[bestOrder])

	} else { // for start-end graph
		s.result = [][][]string{{{g.end}}}
		s.pathsWithAnts = []int{allAnts + 1}
	}
}
func createAnts(g *Graph, s *SortPaths) {

	var sliceAntsInPaths []int
	var antsNames []string
	orderOfTunnel := map[string]int{}
	visitCount := map[string]int{}
	antFinished := map[string]bool{}
	allAnts := g.ants

	data, _ := ioutil.ReadFile(os.Args[1])
	fmt.Println(string(data))
	fmt.Println()
	
	for order, antQueue := range s.pathsWithAnts {
		antsInPaths := antQueue - len(s.result[0][order])
		sliceAntsInPaths = append(sliceAntsInPaths, antsInPaths)
	}
	for ant := 1; ant <= allAnts; {
		for i := range sliceAntsInPaths {
			if sliceAntsInPaths[i]-1 >= 0 {
				antName := "L" + strconv.Itoa(ant)
				orderOfTunnel[antName] = i
				antsNames = append(antsNames, antName)
				sliceAntsInPaths[i] = sliceAntsInPaths[i] - 1
				ant++
			}
		}
	}

	for _, combo := range s.result {

		var outputString string
		room := 0

		for i := len(combo); i <= len(antsNames); {
			for room = 0; room < len(combo); {
				for _, ants := range antsNames[:i] {
					if antFinished[ants] != true { // check if first paths are visited

						if visitCount[ants] == 0 { // if new ant is added

							if room != orderOfTunnel[ants] { // if it goes in the wrong tunnel
								room = 0
								break
							}
							room++
							if room == len(combo) {
								room = 0
							}
						}

						outputString = outputString + ants + "-" + combo[orderOfTunnel[ants]][visitCount[ants]] + " "
						visitCount[ants]++

						if len(combo[orderOfTunnel[ants]]) == visitCount[ants] { // if ant finished
							antFinished[ants] = true
						}
					}

				}
				fmt.Println(outputString)
				outputString = ""

				i += len(combo)
				if i > len(antsNames) {
					i = len(antsNames)
				}

				if len(antFinished) == len(antsNames) { // exit
					return
				}
			}
		}
	}


}
