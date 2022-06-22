package main

import (
	"fmt"
	"math"
)

type Pair struct {
	X int
	Y int
}

func main() {
	fmt.Println(minimumEffortPath([][]int{
		{4, 3, 4, 10, 5, 5, 9, 2},
		{10, 8, 2, 10, 9, 7, 5, 6},
		{5, 8, 10, 10, 10, 7, 4, 2},
		{5, 1, 3, 1, 1, 3, 1, 9},
		{6, 4, 10, 6, 10, 9, 4, 6},
	}))
}

func minimumEffortPath(heights [][]int) int {
	result := math.MaxInt32
	visited := map[Pair]bool{}
	start := Pair{X: 0, Y: 0}
	visited[start] = true
	path := []Pair{start}
	dfs(heights, 0, 0, visited, path, &result)
	return result
}

func dfs(heights [][]int, row int, col int, visited map[Pair]bool, path []Pair, result *int) {
	rows := len(heights)
	cols := len(heights[0])
	if rows == 1 && cols == 1 {
		*result = 0
		return
	}
	if row == rows-1 && col == cols-1 {
		fmt.Println(path)
		maxAbsoluteDiff := math.MinInt32
		for i := 0; i < len(path)-1; i++ {
			currV := heights[path[i].X][path[i].Y]
			nextV := heights[path[i+1].X][path[i+1].Y]
			maxAbsoluteDiff = max(maxAbsoluteDiff, abs(nextV-currV))
		}
		*result = min(*result, maxAbsoluteDiff)
		return
	}
	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for _, direction := range directions {
		newRow := row + direction[0]
		newCol := col + direction[1]
		if newRow >= 0 && newRow < rows && newCol >= 0 && newCol < cols {
			p := Pair{X: newRow, Y: newCol}
			if _, ok := visited[p]; !ok {
				visited[p] = true
				newPath := make([]Pair, len(path))
				copy(newPath, path)
				newPath = append(newPath, p)
				dfs(heights, newRow, newCol, visited, newPath, result)
				delete(visited, p)
			}
		}
	}
}

func abs(i int) int {
	if i < 0 {
		return -1 * i
	}
	return i
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
