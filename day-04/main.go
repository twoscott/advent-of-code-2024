package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var width int
var height int

var searchWord = "XMAS"

func main() {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	input := strings.TrimSpace(string(bytes))
	input = strings.ReplaceAll(input, "\r\n", "\n")

	matrix := strings.Split(input, "\n")
	width = len(matrix[0])
	height = len(matrix)

	xmasFound := 0
	x_MasFound := 0
	for x, row := range matrix {
		for y := range row {
			xmasFound += findXmas(x, y, matrix)
			if findX_Mas(x, y, matrix) {
				x_MasFound++
			}
		}
	}

	fmt.Printf("XMAS appears %d times in the input\n", xmasFound)
	fmt.Printf("X-MAS appears %d times in the input\n", x_MasFound)
}

func findXmas(x int, y int, matrix []string) int {
	return searchUp(x, y, matrix) +
		searchUpRight(x, y, matrix) +
		searchRight(x, y, matrix) +
		searchDownRight(x, y, matrix) +
		searchDown(x, y, matrix) +
		searchDownLeft(x, y, matrix) +
		searchLeft(x, y, matrix) +
		searchUpLeft(x, y, matrix)

}

func searchUp(x int, y int, matrix []string) int {
	return searchXmas(x, y, matrix, [2]int{0, -1}, 0)
}

func searchUpRight(x int, y int, matrix []string) int {
	return searchXmas(x, y, matrix, [2]int{1, -1}, 0)
}

func searchRight(x int, y int, matrix []string) int {
	return searchXmas(x, y, matrix, [2]int{1, 0}, 0)
}

func searchDownRight(x int, y int, matrix []string) int {
	return searchXmas(x, y, matrix, [2]int{1, 1}, 0)
}

func searchDown(x int, y int, matrix []string) int {
	return searchXmas(x, y, matrix, [2]int{0, 1}, 0)
}

func searchDownLeft(x int, y int, matrix []string) int {
	return searchXmas(x, y, matrix, [2]int{-1, 1}, 0)
}

func searchLeft(x int, y int, matrix []string) int {
	return searchXmas(x, y, matrix, [2]int{-1, 0}, 0)
}

func searchUpLeft(x int, y int, matrix []string) int {
	return searchXmas(x, y, matrix, [2]int{-1, -1}, 0)
}

func searchXmas(x int, y int, matrix []string, transform [2]int, progress int) int {
	if x < 0 || x >= width || y < 0 || y >= height {
		return 0
	}

	if matrix[x][y] == searchWord[progress] {
		if progress == len(searchWord)-1 {
			return 1
		}
		return searchXmas(x+transform[0], y+transform[1], matrix, transform, progress+1)
	}
	return 0
}

func findX_Mas(x int, y int, matrix []string) bool {
	if x < 1 || x >= width-1 || y < 1 || y >= height-1 {
		return false
	}

	if matrix[x][y] != 'A' {
		return false
	}

	backDiagMas := matrix[x-1][y-1] == 'M' && matrix[x+1][y+1] == 'S' ||
		matrix[x-1][y-1] == 'S' && matrix[x+1][y+1] == 'M'
	forwardDiagMas := matrix[x+1][y-1] == 'M' && matrix[x-1][y+1] == 'S' ||
		matrix[x+1][y-1] == 'S' && matrix[x-1][y+1] == 'M'

	return backDiagMas && forwardDiagMas
}
