package main

import (
	"fmt"
	"log"
	"os"
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

	strReports := strings.Split(input, "\n")
	reports := make([][]int, 0, len(strReports))

	for _, report := range strReports {
		if len(report) < 1 {
			break
		}

		report = strings.TrimSpace(report)

		strLevels := strings.Split(report, " ")
		levels := make([]int, 0, len(strLevels))

		for _, strLevel := range strLevels {
			level, err := strconv.Atoi(strLevel)
			if err != nil {
				log.Println(err)
				os.Exit(1)
			}

			levels = append(levels, level)
		}

		reports = append(reports, levels)
	}

	part1(reports)
	part2(reports)
}

func part1(reports [][]int) {
	totalSafeReports := 0

	for _, levels := range reports {
		if checkReport(levels) == true {
			totalSafeReports++
		}
	}

	fmt.Printf("Total number of safe reports is: %d\n", totalSafeReports)
}

func part2(reports [][]int) {
	dampenedSafeReports := 0

	for _, levels := range reports {
		if len(levels) < 1 {
			continue
		}

		for i := 0; i < len(levels); i++ {
			if checkReport(levels) == true {
				dampenedSafeReports++
				break
			}

			cutLevels := make([]int, len(levels))
			copy(cutLevels, levels)

			if i == len(levels)-1 {
				cutLevels = cutLevels[:i]
			} else {
				cutLevels = append(cutLevels[:i], cutLevels[i+1:]...)
			}

			if checkReport(cutLevels) == true {
				dampenedSafeReports++
				break
			}
		}
	}

	fmt.Printf("Total number of safe dampened reports is: %d\n", dampenedSafeReports)
}

func checkReport(levels []int) bool {
	recordIsSafe := true
	levelsIncreasing := true
	lastIncreased := true
	for i := 1; i < len(levels); i++ {
		level := levels[i]
		prevLevel := levels[i-1]

		diff := intDiff(level, prevLevel)
		if diff < 1 || diff > 3 {
			recordIsSafe = false
			break
		}

		change := level - prevLevel
		lastIncreased = change > 0
		if i == 1 {
			levelsIncreasing = lastIncreased
		} else if lastIncreased != levelsIncreasing {
			recordIsSafe = false
			break
		}
	}

	fmt.Println(levels, recordIsSafe)

	return recordIsSafe
}
