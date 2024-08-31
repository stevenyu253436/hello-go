package main

func solve(board [][]byte) {
	if len(board) == 0 {
		return
	}

	rows, cols := len(board), len(board[0])

	// 辅助函数，用来标记从边界开始连接的'O'
	var dfs func(int, int)
	dfs = func(r, c int) {
		if r < 0 || c < 0 || r >= rows || c >= cols || board[r][c] != 'O' {
			return
		}
		board[r][c] = '#' // 标记为特殊字符
		dfs(r+1, c)
		dfs(r-1, c)
		dfs(r, c+1)
		dfs(r, c-1)
	}

	// 遍历边界，将所有边界上的'O'及其相连的'O'标记为'#'
	for i := 0; i < rows; i++ {
		if board[i][0] == 'O' {
			dfs(i, 0)
		}
		if board[i][cols-1] == 'O' {
			dfs(i, cols-1)
		}
	}
	for j := 0; j < cols; j++ {
		if board[0][j] == 'O' {
			dfs(0, j)
		}
		if board[rows-1][j] == 'O' {
			dfs(rows-1, j)
		}
	}

	// 遍历整个矩阵，将所有未被标记的'O'变为'X'，将'#'变回'O'
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == '#' {
				board[i][j] = 'O'
			}
		}
	}
}

func main() {
	board := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
	}
	solve(board)
	// 打印结果
	for _, row := range board {
		println(string(row))
	}
}
