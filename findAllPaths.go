package main

import (
	"sort"
)

func findAllPaths(g *Graph) {

	g.current.visited = true

	for _, vertex := range g.current.adjacent { // finds all paths that are not intersecting

		if vertex.name == g.end { // check if room neighbor is end
			path := make([]string, len(g.path)) // creates new real path to append
			copy(path, g.path)
			path = path[1:] // deletes start
			if len(path) == 0 {
				g.pathStartEnd = true
			} else {
				g.allPaths = append(g.allPaths, path) // add path to path list
			}
		}
		if !vertex.visited {
			g.current = vertex
			g.path = append(g.path, g.current.name)
			findAllPaths(g)
		}
	}
	if g.current.name != g.start { // clears wrong rooms from the path
		g.current.visited = false       // make deleted rooms unvisited
		g.path = g.path[:len(g.path)-1] // deletes last room from the path
		g.current = g.getVertex(g.path[len(g.path)-1])
	}

}

func sortPaths(g *Graph, s *SortPaths) { // sorts list of all unintersecting paths

	var tmp [][]string
	for i := 0; i < len(g.allPaths); i++ { // here is search from value
		tmp = append(tmp, g.allPaths[i])
		s.search = append(s.search, tmp)
		s.sliceFrom = append(s.sliceFrom, i)
		tmp = nil
	}

	sort.Slice(g.allPaths, func(i, j int) bool {
		return len(g.allPaths[i]) < len(g.allPaths[j])
	})
	sort.Slice(s.search, func(i, j int) bool {
		if len(s.search[i]) != len(s.search[j]) {
			return len(s.search[i]) < len(s.search[j])
		}
		countI, countJ := 0, 0
		for _, subArr := range s.search[i] {
			countI += len(subArr)
		}
		for _, subArr := range s.search[j] {
			countJ += len(subArr)
		}
		return countI < countJ
	})
}

func bestPathCombinations(g *Graph, s *SortPaths) {

	var roomSum int
	var bestSum int

	for _, searchlist := range s.search { // slice of slices of all search elements // here is the problem
		if (len(s.search[0]) == len(s.search[len(s.search)-1]) && len(s.search[0]) == s.counter) || (len(s.search) == 1 && len(s.search[0]) == s.counter) { // condition to search and add shortest path to result
			for _, searchlist := range s.search { // finds summ of one search element
			there:
				for i, room := range searchlist {
					roomSum = roomSum + len(room)
					if i < len(searchlist)-1 {
						continue there
					} else {
						if bestSum == 0 || bestSum > roomSum { // finds shortest path from search elements
							bestSum = roomSum // current best path
						}
					}
					roomSum = 0
				}
			}
			s.counter++

			bestPath := make([][]string, len(searchlist)) // adds end to the result
			copy(bestPath, searchlist)
			for i := range bestPath {
				bestPath[i] = append(bestPath[i], g.end)
			}
			if g.pathStartEnd == true {
				bestPath = append([][]string{{g.end}}, bestPath...)
			}
			s.result = append(s.result, bestPath) // adds tmp to result
		}
		if len(searchlist) > 1 { // combines two path slices into one slice for searching
			var tmp []string
			for _, slice := range searchlist {
				tmp = append(tmp, slice...)
			}
			s.searchSlice = nil
			s.searchSlice = append(s.searchSlice, tmp)

			g.path = nil
		} else {
			s.searchSlice = searchlist // if searchlist has only one path
		}
		for _, path1 := range s.searchSlice { // searching element can be multiple paths in one slice
			s.from = s.sliceFrom[0]
		here:
			for _, path2 := range g.allPaths[s.from:] { // list of all single paths found in dfs sorted in assending order
				for _, room1 := range path1 { // comparing rooms of path1 and path2
					for _, room2 := range path2 {
						if room1 == room2 { // if paths intersect
							if s.from >= len(g.allPaths)-1 { // if all dfs paths where checked >>> n = 0 and change search
								s.search = s.search[1:] // delete current element from search
								s.sliceFrom = s.sliceFrom[1:]
								if len(s.search) > 0 {
									bestPathCombinations(g, s) // start search with new search element
								}
								if len(s.search) == 0 { // if no paths are in search element
									return
								}
							}
							s.from++
							continue here // if rooms are not the same continue without changing path1 element
						}
					}
				} // if paths doesn't have intersecting rooms
				s.from++ // count this element also
				if s.sliceFrom[0] < s.from {
					s.sliceFrom = append(s.sliceFrom, s.from)
				} else {
					s.sliceFrom = append(s.sliceFrom, s.sliceFrom[0])
				}
				var tmp [][]string
				tmp = append(searchlist, path2)  // appending to all search elements
				s.search = append(s.search, tmp) // tmp is used avoid overwriting of elements
				tmp = nil                        // clearing tmp
			}
		}
	}
}
