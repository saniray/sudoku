package main

import (
	"fmt"
	"os"
)

const size = 9

var sudoku [size][size]int

func print() {
	for _, row := range sudoku {
		fmt.Println(row)
	}
}

func isPossible(sudoku [size][size]int, row, col, num int) bool {
	for i := 0; i < size; i++ {
		if sudoku[row][i] == num || sudoku[i][col] == num {
			return false
		}
	}

	firstRow, firstColumn := row-row%3, col-col%3
	for i := 0; i < size/3; i++ {
		for j := 0; j < size/3; j++ {
			if sudoku[i+firstRow][j+firstColumn] == num {
				return false
			}
		}
	}
	return true
}

func solver() bool {
	for row := 0; row < size; row++ {
		for col := 0; col < size; col++ {
			if sudoku[row][col] == 0 {
				for num := 1; num <= size; num++ {
					if isPossible(sudoku, row, col, num) {
						sudoku[row][col] = num
						if solver() {
							return true
						}
						sudoku[row][col] = 0
					}
				}
				return false
			}
		}
	}
	return true
}

func main() {
	if len(os.Args) != 10 {
		fmt.Println("Error")
		return
	}

	for i := 0; i < 9; i++ {
		if len(os.Args[i+1]) != 9 {
			fmt.Println("Error")
			return
		}
		for j := 0; j < 9; j++ {
			if os.Args[i+1][j] == '.' {
				sudoku[i][j] = 0
			} else if os.Args[i+1][j] >= '1' && os.Args[i+1][j] <= '9' {
				sudoku[i][j] = int(os.Args[i+1][j] - '0')
			} else {
				fmt.Println("Error")
				return
			}
		}
	}

	if solver() {
		print()
	} else {
		fmt.Println("Error")
	}
}
