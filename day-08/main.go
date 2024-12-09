package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

type coord [2]int

func main() {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	input := strings.TrimSpace(string(bytes))
	input = strings.ReplaceAll(input, "\r\n", "\n")

	lines := strings.Split(input, "\n")
	grid := [][]byte{}
	for _, line := range lines {
		grid = append(grid, []byte(line))
	}

	part1(grid)
	part2(grid)
}

func part1(grid [][]byte) {
	nodes := getAllAntinodes(grid, 1)

	fmt.Printf("There are %d unique locations that contain antinodes for part 1\n", len(nodes))
}

func part2(grid [][]byte) {
	nodes := getAllAntinodes(grid, -1)

	fmt.Printf("There are %d unique locations that contain antinodes for part 2\n", len(nodes))
}

func getAllAntinodes(grid [][]byte, depth int) []coord {
	antinodes := make([]coord, 0)
	checkedFreqs := make([]byte, 0)

	for _, row := range grid {
		for _, freq := range row {
			if freq == '.' {
				continue
			}

			if slices.Contains(checkedFreqs, freq) {
				continue
			}

			antennas := findAntennas(grid, freq)
			antinodes = findValidAntinodes(grid, antinodes, antennas, depth)

			checkedFreqs = append(checkedFreqs, freq)
		}
	}

	return antinodes
}

func findAntennas(grid [][]byte, freq byte) []coord {
	antennas := make([]coord, 0)

	for y, row := range grid {
		for x, cell := range row {
			if cell == freq {
				antennas = append(antennas, coord{x, y})
			}
		}
	}

	return antennas
}

func findValidAntinodes(grid [][]byte, nodes []coord, antennas []coord, depth int) []coord {
	for i, a1 := range antennas {
		for j, a2 := range antennas {
			if i == j {
				continue
			}

			newNodes := getAntinodes(grid, a1, a2, depth)
			for _, n := range newNodes {
				if !coordPresent(nodes, n) {
					nodes = append(nodes, n)
				}
			}
		}
	}

	return nodes
}

func getAntinodes(grid [][]byte, antenna1 coord, antenna2 coord, depth int) []coord {
	xDiff := antenna2[0] - antenna1[0]
	yDiff := antenna2[1] - antenna1[1]

	nodes := make([]coord, 0)

	// part 2 includes the antennas themselves as antinodes
	if depth < 0 {
		nodes = append(nodes, antenna1, antenna2)
	}

	for i := 0; depth < 0 || i < depth; i++ {
		antenna1[0] -= xDiff
		antenna1[1] -= yDiff
		antenna2[0] += xDiff
		antenna2[1] += yDiff

		if !isInBounds(grid, antenna1) && !isInBounds(grid, antenna2) {
			break
		}

		if isInBounds(grid, antenna1) {
			nodes = append(nodes, antenna1)
		}
		if isInBounds(grid, antenna2) {
			nodes = append(nodes, antenna2)
		}
	}

	return nodes
}

func isInBounds(grid [][]byte, pos coord) bool {
	return pos[0] >= 0 && pos[0] < len(grid[0]) && pos[1] >= 0 && pos[1] < len(grid)
}

func coordPresent(coords []coord, check coord) bool {
	for _, c := range coords {
		if coordsEqual(c, check) {
			return true
		}
	}

	return false
}

func coordsEqual(c1, c2 coord) bool {
	return c1[0] == c2[0] && c1[1] == c2[1]
}
