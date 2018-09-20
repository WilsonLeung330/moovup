package main

import (
	"fmt"
)
func main() {
	graph := make(map[string][]string)

	graph["A"] = []string{"B", "D", "H"}
	graph["B"] = []string{"A", "C", "D"}
	graph["C"] = []string{"B", "D", "F"}
	graph["D"] = []string{"A", "B", "C", "E"}
	graph["E"] = []string{"D", "F", "H"}
	graph["F"] = []string{"C", "E", "G"}
	graph["G"] = []string{"F", "H"}
	graph["H"] = []string{"A", "E", "G"}
	var path []string

	fmt.Println("AllPossiblePath:")
	AllPossiblePath(graph, "A", "H", path)
	fmt.Println("ShortestPath:")
	fmt.Println(ShortestPath(graph, "A", "H", path))

}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}


func AllPossiblePath(graph map[string][]string, start string, end string, path []string) {
	path = append(path, start)

	if start == end {
		fmt.Println(path)
	}

	for _, node := range graph[start]{
		if !stringInSlice(node, path) {
			AllPossiblePath(graph, node, end, path)
		}
	}
}

func ShortestPath(graph map[string][]string, start string, end string, path []string) ([]string) {
	var shortest []string
	path = append(path, start)

	if start == end {
		return path
	}

	for _, node := range graph[start]{
		if !stringInSlice(node, path) {
			var newPath  []string
			newPath = ShortestPath(graph, node, end, path)
			if len(shortest) == 0 || len(newPath) < len(shortest){
				shortest = newPath
			}
		}
	}

	return shortest
}


