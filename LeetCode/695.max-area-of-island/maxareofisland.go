package maxareaofisland

import "fmt"

type queue struct {
	x int
	y int
}

func maxAreaOfIsland(grid [][]int) int {
	maxArea := 0
	m := len(grid)
	n := len(grid[0])
	var visited [][]bool
	visited = make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	dx := []int{1, -1, 0, 0}
	dy := []int{0, 0, 1, 0 - 1}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 && !visited[i][j] {
				area := 0
				visited[i][j] = true
				q := []queue{}
				q = append(q, queue{i, j})
				for len(q) != 0 {
					front := q[0]
					area++
					for k := 0; k < 4; k++ {
						px := front.x + dx[k]
						py := front.y + dy[k]
						if px >= 0 && px < m && py >= 0 && py < n && grid[px][py] == 1 && !visited[px][py] {
							q = append(q, queue{px, py})
							visited[px][py] = true
						}
					}
					q = q[1:]
				}
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}

	return maxArea
}

func Main() {
	// 7
	// area := [][]int{
	// 	{0, 1, 1, 0, 1, 1},
	// 	{0, 1, 0, 0, 1, 1},
	// 	{0, 0, 1, 0, 0, 1},
	// 	{1, 0, 1, 0, 1, 1},
	// }
	area := [][]int{{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}}
	fmt.Println(maxAreaOfIsland(area))
}
