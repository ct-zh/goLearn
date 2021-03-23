package _51

// 回溯法:
// 任何两个皇后不能处于同一行/列/斜线上
func solveNQueens(n int) [][]string {
	s := solution{
		res:  make([][]string, n),
		col:  make([]bool, n),
		dia1: make([]bool, 2*n-1),
		dia2: make([]bool, 2*n-1),
	}
	for i := 0; i < n; i++ {
		s.res[i] = make([]string, n)
	}

	var rows []int
	s.putQueen(n, 0, &rows)
	return s.res
}

type solution struct {
	res  [][]string
	col  []bool
	dia1 []bool // 右上到左下的对角线，特点：两点之和相等
	dia2 []bool // 左上到右下的对角线，特点：两点之差相等
}

// n: 处理的N皇后的问题； index：当前第几个皇后；
// row：每一行皇后摆放在第几列 例如row[2]=k 代表第二个皇后摆在第k列
func (s solution) putQueen(n int, index int, row *[]int) {
	if index == n {
		s.res = append(s.res, s.generateBoard(n, row))
	}

	for i := 0; i < n; i++ {
		// 尝试将第index行的皇后摆放在第i列
		if !s.col[i] && !s.dia1[index+i] &&
			!s.dia2[index-i+n-1] {
			*row = append(*row, i)
			s.col[i] = true
			s.dia1[index+1] = true
			s.dia2[index-i+n-1] = true
			s.putQueen(n, index+1, row)

			// 回溯
			s.col[i] = false
			s.dia1[index+1] = false
			s.dia2[index-i+n-1] = false
			*row = (*row)[0 : len(*row)-1]
		}
	}
	return
}

func (s solution) generateBoard(n int, row *[]int) []string {
	board := make([]string, n)
	for i := 0; i < n; i++ {
		board[i] = "."
	}
	return board
}
