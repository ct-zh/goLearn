package _200

// floodfill
func numIslands(grid [][]byte) int {
	s := solution{
		m: len(grid),
		n: len(grid[0]),
		d: [4][2]int{
			{0, 1},
			{1, 0},
			{0, -1},
			{-1, 0},
		},
	}
	if s.m <= 0 || s.n <= 0 {
		return 0
	}

	s.visited = make([][]bool, s.m)
	for i := 0; i < len(s.visited); i++ {
		s.visited[i] = make([]bool, s.n)
	}

	res := 0
	for i := 0; i < s.m; i++ {
		for j := 0; j < s.n; j++ {
			if grid[i][j] == '1' && !s.visited[i][j] {
				res++
				s.dfs(grid, i, j)
			}
		}
	}

	return res
}

func (s solution) dfs(grid [][]byte, x int, y int) {
	s.visited[x][y] = true
	for i := 0; i < 4; i++ {
		newX := x + s.d[i][0]
		newY := y + s.d[i][1]
		if newX < s.m && newY < s.n &&
			s.inArea(newX, newY) &&
			!s.visited[newX][newY] &&
			grid[newX][newY] == '1' {
			s.dfs(grid, newX, newY)
		}
	}

	return
}

type solution struct {
	d       [4][2]int
	m, n    int
	visited [][]bool
}

func (s solution) inArea(x, y int) bool {
	return x >= 0 && x < s.m && y >= 0 && y < s.n
}
