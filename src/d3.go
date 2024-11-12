package main

import (
	"strings"
)

func Excavate(input string) int {
	grid := makeGrid(input)
	m, n := len(grid), len(grid[0])

	result := 0
	change := true
	depth := 0
	for change {
		change = false
		newGrid := make([][]int, 0, m)
		for i := 0; i < m; i++ {
			newRow := make([]int, 0, n)
			for j := 0; j < n; j++ {
				if grid[i][j] == depth {
					result += 1
				}

				if canDig(i, j, depth, grid) {
					newRow = append(newRow, depth+1)
					change = true
				} else {
					newRow = append(newRow, depth)
				}

			}
			newGrid = append(newGrid, newRow)
		}
		grid = newGrid

		depth += 1
	}

	return result
}

func ExcavateFancy(input string) int {
	grid := makeGridFancy(input)
	m, n := len(grid), len(grid[0])

	result := 0
	change := true
	depth := 0
	for change {
		change = false
		newGrid := make([][]int, 0, m)
		for i := 0; i < m; i++ {
			newRow := make([]int, 0, n)
			for j := 0; j < n; j++ {
				if grid[i][j] == depth {
					result += 1
				}

				if canDigDiag(i, j, depth, grid) {
					newRow = append(newRow, depth+1)
					change = true
				} else {
					newRow = append(newRow, depth)
				}

			}
			newGrid = append(newGrid, newRow)
		}
		grid = newGrid

		depth += 1
	}

	return result
}

func canDig(i int, j int, depth int, grid [][]int) bool {
	m, n := len(grid), len(grid[0])
	if i == 0 || j == 0 || i == m-1 || j == n-1 {
		return false
	}

	return grid[i][j] == depth && grid[i-1][j] == depth && grid[i+1][j] == depth && grid[i][j-1] == depth && grid[i][j+1] == depth
}

func canDigDiag(i int, j int, depth int, grid [][]int) bool {
	m, n := len(grid), len(grid[0])
	if i == 0 || j == 0 || i == m-1 || j == n-1 {
		return false
	}

	return grid[i][j] == depth && grid[i-1][j] == depth && grid[i+1][j] == depth && grid[i][j-1] == depth && grid[i][j+1] == depth && grid[i-1][j-1] == depth && grid[i+1][j-1] == depth && grid[i-1][j+1] == depth && grid[i+1][j+1] == depth
}

func makeGrid(input string) [][]int {
	lines := strings.Split(input, "\n")

	grid := make([][]int, 0, len(lines))
	for _, line := range lines {
		row := make([]int, 0, len(line))
		for _, tile := range line {
			char := -1
			if tile == '#' {
				char = 0
			}

			row = append(row, char)
		}
		grid = append(grid, row)
	}

	return grid
}

func makeGridFancy(input string) [][]int {
	lines := strings.Split(input, "\n")
	m := len(lines) + 2
	n := len(lines[0]) + 2

	grid := make([][]int, 0, m)
	row := make([]int, 0, n)
	for _ = range n {
		row = append(row, -1)
	}
	grid = append(grid, row)

	for _, line := range lines {
		row := make([]int, 0, m)
		row = append(row, -1)
		for _, tile := range line {
			char := -1
			if tile == '#' {
				char = 0
			}

			row = append(row, char)
		}
		row = append(row, -1)
		grid = append(grid, row)
	}

	row = make([]int, 0, n)
	for _ = range n {
		row = append(row, -1)
	}
	grid = append(grid, row)

	return grid
}
