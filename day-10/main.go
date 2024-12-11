package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type coord struct {
	X int
	Y int
}

func main() {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	input := strings.TrimSpace(string(bytes))
	input = strings.ReplaceAll(input, "\r\n", "\n")
	grid := parseMap(input)

	trailheadScore := 0
	distinctScore := 0
	for y, row := range grid {
		for x, height := range row {
			if height == 0 {
				trailheadScore += getTrailheadScore(grid, coord{X: x, Y: y}, height)
				distinctScore += getDistinctScore(grid, coord{X: x, Y: y}, height)
			}
		}
	}

	fmt.Printf("The total score of all trailheads is: %d\n", trailheadScore)
	fmt.Printf("The distinct score of all trails is: %d\n", distinctScore)
}

func parseMap(input string) [][]int {
	lines := strings.Split(input, "\n")

	grid := [][]int{}
	for _, line := range lines {
		row := make([]int, 0, len(line))

		for _, c := range line {
			height, err := strconv.Atoi(string(c))
			if err != nil {
				log.Println(err)
				os.Exit(1)
			}

			row = append(row, height)
		}

		grid = append(grid, row)
	}

	return grid
}

func getTrailheadScore(grid [][]int, pos coord, height int) int {
	peaks := make([]coord, 0)
	peaks = getTrailPeaks(grid, pos, height, peaks)
	peaks = getUniquePeaks(peaks)

	return len(peaks)
}

func getTrailPeaks(grid [][]int, pos coord, height int, peaks []coord) []coord {
	if height == 9 {
		return append(peaks, pos)
	}

	for _, adjPos := range getAdjacentPositions(grid, pos) {
		adjHeight := getHeight(grid, adjPos)

		if adjHeight == height+1 {
			newPeaks := make([]coord, 0)
			peaks = append(peaks, getTrailPeaks(grid, adjPos, adjHeight, newPeaks)...)
		}
	}

	return peaks
}

func getDistinctScore(grid [][]int, pos coord, height int) int {
	if height == 9 {
		return 1
	}

	score := 0
	for _, adjPos := range getAdjacentPositions(grid, pos) {
		adjHeight := getHeight(grid, adjPos)

		if adjHeight == height+1 {
			score += getDistinctScore(grid, adjPos, adjHeight)
		}
	}

	return score
}

func getAdjacentPositions(grid [][]int, pos coord) []coord {
	adjacent := make([]coord, 0, 4)

	if pos.X > 0 {
		adjacent = append(adjacent, coord{X: pos.X - 1, Y: pos.Y})
	}
	if pos.Y > 0 {
		adjacent = append(adjacent, coord{X: pos.X, Y: pos.Y - 1})
	}
	if pos.X < len(grid[0])-1 {
		adjacent = append(adjacent, coord{X: pos.X + 1, Y: pos.Y})
	}
	if pos.Y < len(grid)-1 {
		adjacent = append(adjacent, coord{X: pos.X, Y: pos.Y + 1})
	}

	return adjacent
}

func getUniquePeaks(peaks []coord) []coord {
	filtered := make([]coord, 0, len(peaks))
	for _, p := range peaks {
		added := false
		for _, f := range filtered {
			if coordsEqual(f, p) {
				added = true
				break
			}
		}

		if !added {
			filtered = append(filtered, p)
		}
	}

	return filtered
}

func getHeight(grid [][]int, pos coord) int {
	return grid[pos.Y][pos.X]
}

func coordsEqual(p1, p2 coord) bool {
	return p1.X == p2.X && p1.Y == p2.Y
}
