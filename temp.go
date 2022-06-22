package main

type Point struct {
	X int
	Y int
}

func pacificAtlantic(heights [][]int) [][]int {
	rows := len(heights)
	cols := len(heights[0])
	pacific := []*Point{}
	atlantic := []*Point{}
	pVisited := map[*Point]bool{}
	aVisited := map[*Point]bool{}
	for i := 0; i < cols; i++ {
		point1 := &Point{X: 0, Y: i}
		point2 := &Point{X: rows - 1, Y: i}
		pacific = append(pacific, point1)
		atlantic = append(atlantic, point2)
		pVisited[point1] = true
		aVisited[point2] = true
	}
	for i := 1; i < rows; i++ {
		point1 := &Point{X: i, Y: 0}
		point2 := &Point{X: i, Y: cols - 1}
		pacific = append(pacific, point1)
		atlantic = append(atlantic, point2)
		pVisited[point1] = true
		aVisited[point2] = true
	}
	totalPacific := runBFSAndGetPoints(pacific, rows, cols, pVisited, heights)
	totalAtlantic := runBFSAndGetPoints(atlantic, rows, cols, aVisited, heights)
	result := [][]int{}
	for p, _ := range totalPacific {
		if _, ok := totalAtlantic[p]; ok {
			result = append(result, []int{p.X, p.Y})
		}
	}
	return result
}

func runBFSAndGetPoints(base []*Point, rows, cols int, visited map[*Point]bool, heights [][]int) map[*Point]bool {
	totalPacific := map[*Point]bool{}
	adjacents := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for len(base) > 0 {
		aux := []*Point{}
		for _, p := range base {
			totalPacific[p] = true
			for _, adj := range adjacents {
				if p.X+adj[0] >= 0 && p.X+adj[0] < rows && p.Y+adj[1] >= 0 && p.Y+adj[1] < cols && heights[p.X+adj[0]][p.Y+adj[1]] >= heights[p.X][p.Y] {
					newPoint := &Point{
						X: p.X + adj[0],
						Y: p.Y + adj[1],
					}
					if _, ok := visited[newPoint]; ok {
						continue
					}
					visited[newPoint] = true
					aux = append(aux, newPoint)
				}
			}
		}
		base = aux
	}
	return totalPacific
}
