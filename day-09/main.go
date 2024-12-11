package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	input := strings.TrimSpace(string(bytes))
	input = strings.ReplaceAll(input, "\r\n", "\n")
	diskMap := []byte(input)

	disk1 := mapOutDisk(diskMap)
	disk2 := make([]int, len(disk1))
	copy(disk2, disk1)

	part1(disk1)
	part2(disk2)
}

func part1(disk []int) {
	disk = compactBlocks(disk)
	checksum := generateChecksum(disk)

	fmt.Printf("The checksum of the disk after compacting the blocks is: %d\n", checksum)
}

func part2(disk []int) {
	disk = compactFiles(disk)
	checksum := generateChecksum(disk)

	fmt.Printf("The checksum of the disk after compacting the files is: %d\n", checksum)
}

func mapOutDisk(diskMap []byte) []int {
	disk := make([]int, 0, len(diskMap))
	for i, val := range diskMap {
		size, err := strconv.Atoi(string(val))
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}

		if isEven(i) {
			disk = append(disk, slices.Repeat([]int{i / 2}, size)...)
		} else {
			disk = append(disk, slices.Repeat([]int{-1}, size)...)
		}
	}

	return disk
}

func compactBlocks(disk []int) []int {
	lastEmpty := -1
	for i := len(disk) - 1; i >= 0; i-- {
		if disk[i] == -1 {
			continue
		}

		for j := lastEmpty + 1; j < i; j++ {
			if disk[j] == -1 {
				disk[j] = disk[i]

				lastEmpty = j
				break
			}
		}
	}

	return disk[:slices.Index(disk, -1)]
}

func compactFiles(disk []int) []int {
	currentID := -1
	fileSize := 0
	lastIDProcessed := -1

	for i := len(disk) - 1; i >= 0; i-- {
		id := disk[i]

		if currentID == -1 || lastIDProcessed >= 0 && currentID >= lastIDProcessed {
			currentID = id
			fileSize = 1
			continue
		}

		if id == currentID {
			fileSize++
			continue
		}

		disk = moveToSpace(disk, fileSize, i+1)

		lastIDProcessed = currentID
		currentID = id
		fileSize = 1
	}

	return disk
}

func moveToSpace(disk []int, fileSize, fileStart int) []int {
	space := 0
	for j := 0; j <= fileStart; j++ {
		if space >= fileSize {
			disk = swapFiles(disk, fileStart, j-fileSize, fileSize)
			break
		}

		if disk[j] == -1 {
			space++
		} else {
			space = 0
		}
	}

	return disk
}

func swapFiles(disk []int, src, dest, size int) []int {
	file := slices.Repeat([]int{disk[src]}, size)
	empty := slices.Repeat([]int{-1}, size)

	disk = append(disk[:src], append(empty, disk[src+size:]...)...)
	disk = append(disk[:dest], append(file, disk[dest+size:]...)...)

	return disk
}

func generateChecksum(disk []int) int64 {
	var checksum int64 = 0
	for i, id := range disk {
		if id == -1 {
			continue
		}

		checksum += int64(i) * int64(id)
	}

	return checksum
}

func isEven(num int) bool {
	return num%2 == 0
}
