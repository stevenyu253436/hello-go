package main

import "fmt"

// DFS function to find the result of the division query
func dfs(graph map[string]map[string]float64, start string, end string, visited map[string]bool) (float64, bool) {
	if _, ok := graph[start]; !ok {
		return -1.0, false
	}

	if start == end {
		return 1.0, true
	}

	visited[start] = true
	for neighbor, value := range graph[start] {
		if visited[neighbor] {
			continue
		}
		if product, found := dfs(graph, neighbor, end, visited); found {
			return product * value, true
		}
	}
	return -1.0, false
}

// Main function to calculate the results of the queries
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	graph := make(map[string]map[string]float64)

	// Build the graph
	for i, equation := range equations {
		a, b := equation[0], equation[1]
		if _, ok := graph[a]; !ok {
			graph[a] = make(map[string]float64)
		}
		if _, ok := graph[b]; !ok {
			graph[b] = make(map[string]float64)
		}
		graph[a][b] = values[i]
		graph[b][a] = 1.0 / values[i]
	}

	// Process each query
	result := make([]float64, len(queries))
	for i, query := range queries {
		a, b := query[0], query[1]
		if _, ok := graph[a]; !ok {
			result[i] = -1.0
		} else if _, ok := graph[b]; !ok {
			result[i] = -1.0
		} else {
			visited := make(map[string]bool)
			res, found := dfs(graph, a, b, visited)
			if found {
				result[i] = res
			} else {
				result[i] = -1.0
			}
		}
	}

	return result
}

func main() {
	equations := [][]string{{"a", "b"}, {"b", "c"}}
	values := []float64{2.0, 3.0}
	queries := [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}}
	result := calcEquation(equations, values, queries)
	fmt.Println(result)
}
