package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func intDiff(x, y int) int {
	if x < y {
		return y - x
	} else {
		return x - y
	}
}

func main() {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	input := string(bytes)

	leftList := []int{}
	rightList := []int{}
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		if len(line) < 1 {
			break
		}

		entries := strings.Split(line, "   ")

		leftNum, err := strconv.Atoi(strings.TrimSpace(entries[0]))
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}

		rightNum, err := strconv.Atoi(strings.TrimSpace(entries[1]))
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}

		leftList = append(leftList, leftNum)
		rightList = append(rightList, rightNum)
	}

	slices.Sort(leftList)
	slices.Sort(rightList)

	totalDistance := 0
	totalSimilarity := 0

	for i := 0; i < len(leftList); i++ {
		leftNum := leftList[i]
		rightNum := rightList[i]

		totalDistance += intDiff(leftNum, rightNum)

		matches := 0
		for _, r := range rightList {
			if r == leftNum {
				matches++
			}
		}

		totalSimilarity += matches * leftNum
	}

	fmt.Printf("The total distance between the lists is: %d\n", totalDistance)
	fmt.Printf("The total similarity between the lists is: %d\n", totalSimilarity)
}
