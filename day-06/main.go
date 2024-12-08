package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type coord [2]int

type direction int

const (
	Left direction = iota
	Up
	Right
	Down
)

type history struct {
	Pos coord
	Dir direction
}

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

	pos := findGuardPos(grid)
	if pos == nil {
		log.Println("guard not found")
		os.Exit(1)
	}

	path := getGuardPath(grid, *pos)
	fmt.Printf("The guard visited %d unique positions\n", len(path))

	blockLoopCount := getBlockLoopCount(grid, path, *pos)
	fmt.Printf("Loops can be made by blocking %d possible locations\n", blockLoopCount)
}

func findGuardPos(grid [][]byte) *coord {
	for y, row := range grid {
		for x, cell := range row {
			if cell == '^' {
				return &coord{x, y}
			}
		}
	}

	return nil
}

func getGuardPath(grid [][]byte, startPos coord) []coord {
	cellHistory := []history{{Pos: startPos, Dir: Up}}
	uniqueVisited := []coord{startPos}

	for {
		last := cellHistory[len(cellHistory)-1]

		nextPos, newDir := findNextCell(grid, last.Pos, last.Dir)
		if !cellIsInBounds(grid, nextPos) {
			break
		}

		cellHistory = append(cellHistory, history{Pos: nextPos, Dir: newDir})

		cellAlreadyVisited := false
		for _, cell := range uniqueVisited {
			if coordsEqual(nextPos, cell) {
				cellAlreadyVisited = true
				break
			}
		}

		if !cellAlreadyVisited {
			uniqueVisited = append(uniqueVisited, nextPos)
		}
	}

	return uniqueVisited
}

func getBlockLoopCount(grid [][]byte, path []coord, startPos coord) int {
	loopBlockLocations := []coord{}
	loopGrid := make([][]byte, len(grid))

	for _, step := range path {
		for i := range loopGrid {
			loopGrid[i] = make([]byte, len(grid[0]))
			copy(loopGrid[i], grid[i])
		}

		loopGrid[step[1]][step[0]] = '#'

		doesLoop := doesGridLoop(loopGrid, startPos, Up)
		if doesLoop {
			loopBlockLocations = append(loopBlockLocations, step)
		}
	}

	return len(loopBlockLocations)
}

func doesGridLoop(grid [][]byte, startPos coord, startDir direction) bool {
	cellHistory := []history{{Pos: startPos, Dir: startDir}}

	for {
		last := cellHistory[len(cellHistory)-1]

		for _, h := range cellHistory[:len(cellHistory)-1] {
			if coordsEqual(h.Pos, last.Pos) && h.Dir == last.Dir {
				return true
			}
		}

		nextPos, newDir := findNextCell(grid, last.Pos, last.Dir)
		if !cellIsInBounds(grid, nextPos) {
			break
		}

		if coordsEqual(nextPos, last.Pos) {
			return true
		}

		cellHistory = append(cellHistory, history{Pos: nextPos, Dir: newDir})
	}

	return false
}

func findNextCell(grid [][]byte, pos coord, dir direction) (coord, direction) {
	rCount := 0

	for {
		if rCount >= 4 {
			return pos, dir
		}

		nextCell := getNextCell(pos, dir)
		if !cellIsInBounds(grid, nextCell) {
			return nextCell, dir
		}

		if grid[nextCell[1]][nextCell[0]] != '#' {
			return nextCell, dir
		}

		dir = rotateDirClockwise(dir)
		rCount++
	}
}

func cellIsInBounds(grid [][]byte, pos coord) bool {
	return pos[0] >= 0 && pos[0] < len(grid[0]) && pos[1] >= 0 && pos[1] < len(grid)
}

func getNextCell(pos coord, dir direction) coord {
	switch dir {
	case Right:
		pos[0]++
	case Down:
		pos[1]++
	case Left:
		pos[0]--
	case Up:
		pos[1]--
	}

	return pos
}

func rotateDirClockwise(dir direction) direction {
	if dir >= Down {
		return Left
	}

	return dir + 1
}

func coordsEqual(c1, c2 coord) bool {
	return c1[0] == c2[0] && c1[1] == c2[1]
}
