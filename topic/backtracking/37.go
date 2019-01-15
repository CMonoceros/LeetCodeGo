package main

import "strconv"

// Hash Table
// Backtracking

// todo:1
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

	var isValid func(board [][]byte, num, row, col int) bool
	isValid = func(board [][]byte, num, row, col int) bool {
		numByte := []byte(strconv.Itoa(num))[0]
		for i := 0; i < 9; i++ {
			if board[row][i] == numByte {
				return false
			}
			if board[i][col] == numByte {
				return false
			}
			cellRow := (row/3)*3 + i/3
			cellCol := (col/3)*3 + i%3
			if board[cellRow][cellCol] == numByte {
				return false
			}
		}
		return true
	}

	var dfs func(board *[][]byte, empty []int, index, size int) bool
	dfs = func(board *[][]byte, empty []int, index, size int) bool {
		if index == size {
			return true
		}
		now := empty[index]
		row := now / 9
		col := now % 9
		for num := 1; num <= 9; num++ {
			if isValid(*board, num, row, col) {
				(*board)[row][col] = []byte(strconv.Itoa(num))[0]
				if dfs(board, empty, index+1, size) {
					return true
				}
				(*board)[row][col] = '*'
			}
		}
		return false
	}

	dfs(&board, empty, 0, len(empty))
}
