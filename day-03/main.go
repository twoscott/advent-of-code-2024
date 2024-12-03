package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	input := string(bytes)

	part1(input)
	part2(input)
}

func part1(input string) {
	mulRegex := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	matches := mulRegex.FindAllStringSubmatch(input, -1)

	multSum := 0
	for _, match := range matches {
		mult1, err := strconv.Atoi(match[1])
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}

		mult2, err := strconv.Atoi(match[2])
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}

		multSum += mult1 * mult2
	}

	fmt.Printf("The sum of all uncorrupted multiplications is: %d\n", multSum)
}

func part2(input string) {
	mulRegex := regexp.MustCompile(`(?:(do(?:n't)?)\(\)|(mul)\((\d{1,3}),(\d{1,3})\))`)

	matches := mulRegex.FindAllStringSubmatch(input, -1)

	multSum := 0
	multEnabled := true
	for _, match := range matches {
		dodont := match[1]
		mul := match[2]
		mult1Match := match[3]
		mult2Match := match[4]

		if dodont == "do" {
			multEnabled = true
		} else if dodont == "don't" {
			multEnabled = false
		} else if mul == "mul" && multEnabled {
			mult1, err := strconv.Atoi(mult1Match)
			if err != nil {
				log.Println(err)
				os.Exit(1)
			}

			mult2, err := strconv.Atoi(mult2Match)
			if err != nil {
				log.Println(err)
				os.Exit(1)
			}

			multSum += mult1 * mult2
		}
	}

	fmt.Printf("The sum of all uncorrupted multiplications with do & don't instructions is: %d\n", multSum)
}
