package _79

// 回溯法
func exist(board [][]byte, word string) bool {
	byteWords := []byte(word)
	if len(board) == 0 || len(byteWords) == 0 {
		return false
	}

	length := len(board)
	width := len(board[0])

	// 先找到起点
	var startCoordinates []struct {
		x int
		y int
	}
	for m := 0; m < length; m++ {
		for n := 0; n < width; n++ {
			if board[m][n] == byteWords[0] {
				startCoordinates = append(startCoordinates, struct {
					x int
					y int
				}{x: m, y: n})
			}
		}
	}

	// 再从起点开始寻路

	return true
}

func find(board [][]byte, item byte) {

}

func exist2(board [][]byte, word string) bool {
	s := search{
		m: len(board),
		n: len(board[0]),
		d: [4][2]int{
			{-1, 0},
			{0, 1},
			{1, 0},
			{0, -1},
		},
	}
	s.visited = make([][]bool, s.m)
	for i := 0; i < len(s.visited); i++ {
		s.visited[i] = make([]bool, s.n)
	}
	if s.m <= 0 {
		return false
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if s.searchWord(board, word, 0, i, j) {
				return true
			}
		}
	}
	return false
}

type search struct {
	m       int
	n       int
	d       [4][2]int
	visited [][]bool
}

// 从board[startx][starty]开始，寻找word[index...len(word)]
func (s search) searchWord(
	board [][]byte,
	word string,
	index int,
	startx int,
	starty int) bool {

	if index == len(word)-1 {
		return board[startx][starty] == word[index]
	}
	if board[startx][starty] == word[index] {
		s.visited[startx][starty] = true
		// 从startx，starty出发，向四个方向寻找
		for i := 0; i < 4; i++ {
			newX := startx + s.d[i][0]
			newY := starty + s.d[i][1]
			if s.inArea(newX, newY) && !s.visited[newX][newY] {
				if s.searchWord(board, word, index+1, newX, newY) {
					return true
				}
			}
		}

		// 回溯
		s.visited[startx][starty] = false
	}
	return false
}

func (s search) inArea(x, y int) bool {
	return x >= 0 && x < s.m && y >= 0 && y < s.n
}
