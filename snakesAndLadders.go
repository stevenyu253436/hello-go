func snakesAndLadders(board [][]int) int {
	n := len(board)
	// 將棋盤轉換為一維數組
	flatten := make([]int, n*n+1)
	index, row, col, inc := 1, n-1, 0, 1

	for row >= 0 && index <= n*n {
		flatten[index] = board[row][col]
		index++
		col += inc
		if col == n || col == -1 {
			row--
			col -= inc
			inc *= -1
		}
	}

	// BFS 開始
	queue := []int{1}
	visited := make([]bool, n*n+1)
	visited[1] = true
	moves := 0

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			curr := queue[0]
			queue = queue[1:]

			if curr == n*n {
				return moves
			}

			for next := curr + 1; next <= curr+6 && next <= n*n; next++ {
				dest := next
				if flatten[next] != -1 {
					dest = flatten[next]
				}

				if !visited[dest] {
					visited[dest] = true
					queue = append(queue, dest)
				}
			}
		}
		moves++
	}

	return -1 // 無法到達終點
}
