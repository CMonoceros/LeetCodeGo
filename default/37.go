package main

import (
	"strconv"
)

// Hash Table
// Backtracking

// todo:
// 12ms
func solveSudoku(board [][]byte) {
	var empty []int
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == '.' {
				empty = append(empty, i*9+j)
			}
		}
	}
	emptyLen := len(empty)
	dfs(&board, empty, 0, emptyLen)
}

func dfs(board *[][]byte, empty []int, index, size int) bool {
	if index == size {
		return true
	}
	now := empty[index]
	row := now / 9
	col := now % 9
	for i := 1; i <= 9; i++ {
		if isValid(*board, row, col, i) {
			(*board)[row][col] = []byte(strconv.Itoa(i))[0]
			if dfs(board, empty, index+1, size) {
				return true
			}
			(*board)[row][col] = '.'
		}
	}
	return false
}

func isValid(board [][]byte, row, col, num int) bool {
	now := []byte(strconv.Itoa(num))[0]
	for i := 0; i < 9; i++ {
		if board[row][i] == now {
			return false
		}
		if board[i][col] == now {
			return false
		}
		cellRow := (row/3)*3 + i/3
		cellCol := (col/3)*3 + i%3
		if board[cellRow][cellCol] == now {
			return false
		}
	}
	return true
}
