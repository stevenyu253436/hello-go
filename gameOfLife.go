package main

import "fmt"

func gameOfLife(board [][]int) {
	rows := len(board)
	cols := len(board[0])

	// Directions for the 8 neighbors (top-left, top, top-right, left, right, bottom-left, bottom, bottom-right)
	directions := [][]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	// Iterate through each cell
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			liveNeighbors := 0

			// Count live neighbors
			for _, dir := range directions {
				ni, nj := i+dir[0], j+dir[1]

				if ni >= 0 && ni < rows && nj >= 0 && nj < cols {
					if board[ni][nj] == 1 || board[ni][nj] == -1 {
						liveNeighbors++
					}
				}
			}

			// Apply the rules
			if board[i][j] == 1 && (liveNeighbors < 2 || liveNeighbors > 3) {
				// Live cell dies
				board[i][j] = -1
			}
			if board[i][j] == 0 && liveNeighbors == 3 {
				// Dead cell becomes live
				board[i][j] = 2
			}
		}
	}

	// Final update of the board
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == -1 {
				board[i][j] = 0
			}
			if board[i][j] == 2 {
				board[i][j] = 1
			}
		}
	}
}

func main() {
	board := [][]int{
		{0, 1, 0},
		{0, 0, 1},
		{1, 1, 1},
		{0, 0, 0},
	}

	fmt.Println("Initial Board:")
	for _, row := range board {
		fmt.Println(row)
	}

	gameOfLife(board)

	fmt.Println("\nNext State Board:")
	for _, row := range board {
		fmt.Println(row)
	}
}
