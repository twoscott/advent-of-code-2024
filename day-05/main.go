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

	parts := strings.Split(input, "\n\n")
	rules := parseInputs(parts[0], "|")
	updates := parseInputs(parts[1], ",")

	part1(rules, updates)
	part2(rules, updates)
}

func parseInputs(input, sep string) [][]int {
	lines := strings.Split(input, "\n")

	entries := make([][]int, 0, len(lines))
	for _, line := range lines {
		elems := strings.Split(line, sep)

		entry := make([]int, 0, len(elems))
		for _, p := range elems {
			num, err := strconv.Atoi(p)
			if err != nil {
				log.Println(err)
				os.Exit(1)
			}

			entry = append(entry, num)
		}

		entries = append(entries, entry)
	}

	return entries
}

func part1(rules, updates [][]int) {
	middlePageSum := 0

	for _, update := range updates {
		inOrder := true

		for pageIdx := range update {
			inOrder = isPageInOrder(update, pageIdx, rules)
			if !inOrder {
				break
			}
		}

		if inOrder {
			middleIndex := len(update) / 2
			middlePageSum += update[middleIndex]
		}
	}

	fmt.Printf("The sum of middle pages in correctly-ordered updates is %d\n", middlePageSum)
}

func isPageInOrder(update []int, pageIdx int, rules [][]int) bool {
	page := update[pageIdx]
	for _, rule := range rules {
		if rule[1] == page && slices.Contains(update[pageIdx+1:], rule[0]) {
			return false
		}
	}

	return true
}

func part2(rules, updates [][]int) {
	middlePageSum := 0

	for _, update := range updates {
		sortedUpdate := make([]int, len(update))
		copy(sortedUpdate, update)

		slices.SortStableFunc(sortedUpdate, func(a, b int) int {
			for _, r := range rules {
				if r[0] == a && r[1] == b {
					return -1
				} else if r[0] == b && r[1] == a {
					return 1
				}
			}
			return 0
		})

		if slices.Compare(sortedUpdate, update) != 0 {
			middleIndex := len(sortedUpdate) / 2
			middlePageSum += sortedUpdate[middleIndex]
		}
	}

	fmt.Printf("The sum of middle pages in corrected updates is %d\n", middlePageSum)
}
